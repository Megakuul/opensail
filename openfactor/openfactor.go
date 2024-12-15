/**
 * Opensail System
 *
 * Copyright (C) 2024 Linus Ilian Moser <linus.moser@megakuul.ch>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package openfactor

import "math"

// Version returns the algorithm version.
func Version() string {
	return "v0.0.1"
}

const (
	POINT_DIVIDOR = 100 // number used to convert points to decimal
	POINT_ANCHOR  = 2   // anchor used to convert points to factor

	SPEED_FACTOR_INFLUENCE         = 1.5  // empirical value specifying how much TCC influence the speed has
	STABILIZATION_FACTOR_INFLUENCE = 0.75 // empirical value specifying how much TCC influence the stabilization has
	AGILITY_FACTOR_INFLUENCE       = 0.75 // empirical value specifying how much TCC influence the agility has

	DRAG_SPEED_POINT_PATCHER     = 0.6 // empirical value to patch the drag speed points
	UPWIND_SPEED_POINT_PATCHER   = 0.8 // empirical value to patch the upwind speed points
	DOWNWIND_SPEED_POINT_PATCHER = 0.5 // empirical value to patch the downwind speed points

)

type MODE int64

const (
	MODE_DEFAULT MODE = iota // default==unspecified provides 0 extra rating
	MODE_HYDROFOIL
	MODE_PLANING
	MODE_SEMI
	MODE_DISPLACE
)

var MODE_FACTOR = map[MODE]float64{
	MODE_DEFAULT:   0,    // default mode is considered 0 drag - boat does not touch the water.
	MODE_HYDROFOIL: 0.05, // hydrofoil mode provides optimal drag (nearly 0).
	MODE_PLANING:   0.5,  // assuming the boats has 50% wsa while planing
	MODE_SEMI:      0.7,  // assuming the boats has 50% wsa but is not always planing
	MODE_DISPLACE:  1,    // boat has 100% wsa as it is not planing
}

type STABILIZATION int64

const (
	STABILIZATION_DEFAULT STABILIZATION = iota
	STABILIZATION_KEEL
	STABILIZATION_CENTREBOARD
	STABILIZATION_DAGGERBOARD
	STABILIZATION_FOILS
)

type HULL int64

const (
	HULL_DEFAULT HULL = iota
	HULL_MULTI
	HULL_MONO
)

type MATERIAL int64

const (
	MATERIAL_DEFAULT MATERIAL = iota
	MATERIAL_BALLAST
	MATERIAL_CFK
	MATERIAL_ALU
	MATERIAL_GFK
	MATERIAL_WOOD
	MATERIAL_ENGINE
	MATERIAL_AMENITY
)

var MATERIAL_FACTOR = map[MATERIAL]float64{
	MATERIAL_DEFAULT: 1,    // default material is considered perfect for its function
	MATERIAL_BALLAST: 1,    // no matter what, ballast is always 100% perfect in its function
	MATERIAL_CFK:     0.9,  // cfk is nearly perfect in its function
	MATERIAL_ALU:     0.5,  // assuming it takes ~1.8 times more alu then cfk to get the same structural strength
	MATERIAL_GFK:     0.4,  // assuming it takes ~2.2 times more gfk then cfk to get the same structural strength
	MATERIAL_WOOD:    0.23, // assuming it takes ~4 times more wood then cfk to get the same structural strength
	MATERIAL_ENGINE:  0,    // the engine has no function for the sailor and is therefore ignored
	MATERIAL_AMENITY: 0,    // the amenities have no function for the sailor and are therefore ignored
}

// EvaluationInput specifies data that is used to derive the rating.
type EvaluationInput struct {
	// LOA specifies the ship length in meters
	// from the foremost to the rearmost point.
	LOA float64
	// MaxDraft specifies the maximum vertical distance
	// from the bottom of the keel to the waterline in meters.
	MaxDraft float64
	// MaxBeam specifies the maximum width of the beam in meters.
	MaxBeam float64
	// IMSL is a weird term for the vertical forestray height in meters.
	IMSL float64
	// WSA specifies the wetted surface area in square meters.
	WSA float64

	// MainSailArea specifies the size of the
	// largest mainsail in square meters.
	MainSailArea float64
	// JibSailArea specifies the size of the
	// largest jibsail in square meters.
	JibSailArea float64
	// AsymmetricSpinnakerArea specifies the size of the
	// largest asymmetric downwind sail in square meters.
	AsymmetricSpinnakerArea float64
	// SymmetricSpinnakerArea specifies the size of the
	// largest symmetric downwind sail in square meters.
	SymmetricSpinnakerArea float64

	// Displacement specifies the ship displacement in KG (this is the same as the weight of the ship)
	Displacement float64
	// CrewWeight specifies the weight of the crew on the ship.
	CrewWeight float64

	// Mode specifies the hull operation mode.
	Mode MODE
	// Stabilization specifies how the ship is stabilized.
	Stabilization STABILIZATION
	// Hull specifies the hull design of the ship.
	Hull HULL

	// Composition specifies how the ship is composed.
	Composition map[MATERIAL]float64
}

type EvaluationOutput struct {
	// Time correction coefficient produced by the algorithm.
	TCC float64
	// SpeedFactor specifies the factor the boat retrieved in category "Speed".
	SpeedFactor         float64
	SpeedInfluence      float64
	SpeedDragPoints     float64
	SpeedUpwindPoints   float64
	SpeedDownwindPoints float64

	// StabilizationFactor specifies the factor the boat retrieved in category "Stability".
	StabilizationFactor    float64
	StabilizationInfluence float64

	// AgilityFactor specifies the factor the boat retrieved in category "Agility".
	AgilityFactor    float64
	AgilityInfluence float64
}

func EvaluateFactor(input *EvaluationInput) (*EvaluationOutput, error) {
	speedDragPoints := evaluateDragSpeedPoints(
		input.Mode,
		input.Displacement,
		input.WSA,
	)
	speedUpwindPoints := evaluateUpwindSpeedPoints(
		input.Displacement,
		input.MainSailArea,
		input.AsymmetricSpinnakerArea,
		input.IMSL,
	)
	speedDownwindPoints := evaluateDownwindSpeedPoints(
		input.Displacement,
		input.AsymmetricSpinnakerArea,
		input.SymmetricSpinnakerArea,
	)
	speedPoints := speedDragPoints + speedUpwindPoints + speedDownwindPoints
	speedFactor := POINT_ANCHOR - (speedPoints / POINT_DIVIDOR)

	stabilizationFactor := POINT_ANCHOR - (stabilizationPoints / POINT_DIVIDOR)

	agilityFactor := POINT_ANCHOR - (agilityPoints / POINT_DIVIDOR)

	tcc := ((speedFactor * SPEED_FACTOR_INFLUENCE) +
		(stabilizationFactor * STABILIZATION_FACTOR_INFLUENCE) +
		(agilityFactor * AGILITY_FACTOR_INFLUENCE)) / 3

	return &EvaluationOutput{
		TCC:                 tcc,
		SpeedFactor:         speedFactor,
		SpeedInfluence:      SPEED_FACTOR_INFLUENCE,
		SpeedDragPoints:     speedDragPoints,
		SpeedUpwindPoints:   speedUpwindPoints,
		SpeedDownwindPoints: speedDownwindPoints,

		StabilizationFactor:    stabilizationFactor,
		StabilizationInfluence: STABILIZATION_FACTOR_INFLUENCE,

		AgilityFactor:    agilityFactor,
		AgilityInfluence: AGILITY_FACTOR_INFLUENCE,
	}, nil
}

func evaluateDragSpeedPoints(mode MODE, displ, wsa float64) float64 {
	displArea := displ / 1000 // assuming water is 1000 kg / m3
	// ratio between the edge length of wsa and displArea.
	wsaDisplRatio := math.Pow(math.Pow(wsa, 1/2)/math.Pow(displArea, 1/3), 2)
	// multiply wsa by the mode factor to adjust its influence based on the mode.
	// this means: low factors (foils) reduces wsa impact, while high factors (displacement) increase the wsa impact.
	// the reference value 100 inverts the scale, so higher wsa means higher drag and lower points.
	return (500 - (wsa * wsaDisplRatio * MODE_FACTOR[mode])) * DRAG_SPEED_POINT_PATCHER
}

func evaluateDownwindSpeedPoints(displ, asym, sym float64) float64 {
	// asymmetric and symmetric downwindsails are not differentiated, as its considered a "strategic decision".
	// the largest sail is counted, other smaller sails may be used in the race.
	sailArea := math.Max(sym, asym)
	displArea := displ / 1000 // assuming water is 1000 kg / m3

	// ratio between the edge length of sailArea and displArea.
	sailDisplRatio := math.Pow(math.Pow(sailArea, 1/2)/math.Pow(displArea, 1/3), 2)

	return (sailArea * sailDisplRatio) * DOWNWIND_SPEED_POINT_PATCHER
}

func evaluateUpwindSpeedPoints(displ, main, jib, forestay float64) float64 {
	// higher forestay means the sails can be trimmed to use higher winds which are generally faster due to surface friction.
	// this is not very influential, so only a small fraction of the jib is added.
	forestayFactor := forestay / 100
	sailArea := main + jib + (math.Pow(math.Pow(jib, 1/2)*forestayFactor, 2))
	displArea := displ / 1000 // assuming water is 1000 kg / m3

	// ratio between the edge length of sailArea and displArea.
	sailDisplRatio := math.Pow(math.Pow(sailArea, 1/2)/math.Pow(displArea, 1/3), 2)

	return (sailArea * sailDisplRatio) * UPWIND_SPEED_POINT_PATCHER
}

func evaluateAgilityFactor(loa, beam, draft, wsa, main, jib, asym, sym float64) float64 {

}

func evaluateStabilizationFactor(loa, beam, draft, wsa, main, jib float64, hull HULL, stabilization STABILIZATION) float64 {

}
