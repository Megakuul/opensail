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

import (
	"math"
)

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

	DRAG_SPEED_POINT_PATCHER     = 60  // empirical value to patch the drag speed points
	UPWIND_SPEED_POINT_PATCHER   = 60  // empirical value to patch the upwind speed points
	DOWNWIND_SPEED_POINT_PATCHER = 60  // empirical value to patch the downwind speed points
	STABILIZATION_POINT_PATCHER  = 0.5 // empirical value to patch the stabilization points
	AGILITY_POINT_PATCHER        = 0.7 // empirical value to patch the agility points
)

type MODE int64

const (
	MODE_DEFAULT MODE = iota
	MODE_HYDROFOIL
	MODE_PLANING
	MODE_SEMI
	MODE_DISPLACE
)

var MODE_DRAG_FACTOR = map[MODE]float64{
	MODE_DEFAULT:   0,    // default mode is considered 0 drag - boat does not touch the water.
	MODE_HYDROFOIL: 0.05, // hydrofoil mode provides optimal drag (nearly 0).
	MODE_PLANING:   0.5,  // assuming the boats has 50% wsa while planing
	MODE_SEMI:      0.7,  // assuming the boats has 50% wsa but is not always planing
	MODE_DISPLACE:  1,    // boat has 100% wsa as it is not planing
}

type STABILIZATION int64

const (
	STABILIZATION_DEFAULT STABILIZATION = iota
	STABILIZATION_BULBKEEL
	STABILIZATION_FINKEEL
	STABILIZATION_FULLKEEL
	STABILIZATION_CENTREBOARD
	STABILIZATION_DAGGERBOARD
	STABILIZATION_FOILS
)

// sorry for this retarded variable name but it fits into scheme...
var STABILIZATION_STABILIZATION_FACTOR = map[STABILIZATION]float64{
	STABILIZATION_DEFAULT:     1,   // default has perfect stabilization
	STABILIZATION_FULLKEEL:    0.9, // fullkeel provides the best stability due to ballast + righting force
	STABILIZATION_BULBKEEL:    0.8, // bulbkeel provides decent stability mainly due to well distributed ballast
	STABILIZATION_FINKEEL:     0.7, // finkeel provides decent stability with decent distributed ballast
	STABILIZATION_CENTREBOARD: 0.2, // centreboard provides some righting force but no distributed ballast
	STABILIZATION_DAGGERBOARD: 0.2, // daggerboard provides some righting force but no distributed ballast
	STABILIZATION_FOILS:       0.1, // foils provide decent righting force but are generally harder to operate
}

var STABILIZATION_AGILITY_FACTOR = map[STABILIZATION]float64{
	STABILIZATION_DEFAULT:     1,   // default has perfect stabilization
	STABILIZATION_BULBKEEL:    0.9, // bulbkeel provides well agility and enables more risky maneuvers due to distributed ballast
	STABILIZATION_FINKEEL:     0.8, // finkeel provides well agility and enables more risky maneuvers due to distributed the ballast
	STABILIZATION_CENTREBOARD: 0.7, // centreboard provides decent agility due to the lack of stabilization
	STABILIZATION_DAGGERBOARD: 0.7, // daggerboard provides decent agility due to the lack of stabilization
	STABILIZATION_FULLKEEL:    0.4, // fullkeel provides rather limited agility due to the strong straight righting force
	STABILIZATION_FOILS:       0.2, // foils provide rather bad agility and are generally harder to operate
}

type HULL int64

const (
	HULL_DEFAULT HULL = iota
	HULL_MULTI
	HULL_MONO
)

var HULL_STABILIZATION_FACTOR = map[HULL]float64{
	HULL_DEFAULT: 1,   // default has perfect stabilization
	HULL_MULTI:   0.9, // multi hull provides almost optimal stabilization
	HULL_MONO:    0.7, // mono hull provides decent stabilization
}

var HULL_AGILITY_FACTOR = map[HULL]float64{
	HULL_DEFAULT: 1,   // default has perfect agility
	HULL_MONO:    0.9, // mono hull has nearly perfect agility
	HULL_MULTI:   0.3, // multi hull has limited agility
}

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
	StabilizationPoints    float64
	StabilizationInfluence float64

	// AgilityFactor specifies the factor the boat retrieved in category "Agility".
	AgilityFactor    float64
	AgilityPoints    float64
	AgilityInfluence float64
}

func EvaluateFactor(input *EvaluationInput) (*EvaluationOutput, error) {
	speedDragPoints := math.Round(evaluateDragSpeedPoints(
		input.Mode,
		input.Displacement,
		input.WSA,
	))
	speedUpwindPoints := math.Round(evaluateUpwindSpeedPoints(
		input.Displacement,
		input.MainSailArea,
		input.AsymmetricSpinnakerArea,
		input.IMSL,
	))
	speedDownwindPoints := math.Round(evaluateDownwindSpeedPoints(
		input.Displacement,
		input.AsymmetricSpinnakerArea,
		input.SymmetricSpinnakerArea,
	))
	speedPoints := (speedDragPoints + speedUpwindPoints + speedDownwindPoints) / 3
	speedFactor := POINT_ANCHOR - (speedPoints / POINT_DIVIDOR)

	stabilizationPoints := math.Round(evaluateStabilizationPoints(
		input.WSA,
		input.MaxDraft,
		input.MaxBeam,
		input.LOA,
		input.Displacement,
		input.MainSailArea,
		input.Composition,
		input.Stabilization,
		input.Hull,
	))
	stabilizationFactor := POINT_ANCHOR - (stabilizationPoints / POINT_DIVIDOR)

	agilityPoints := math.Round(evaluateAgilityPoints(
		input.MaxBeam,
		input.LOA,
		input.CrewWeight,
		input.Displacement,
		input.Stabilization,
		input.Hull,
	))
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
		StabilizationPoints:    stabilizationPoints,
		StabilizationInfluence: STABILIZATION_FACTOR_INFLUENCE,

		AgilityFactor:    agilityFactor,
		AgilityPoints:    agilityPoints,
		AgilityInfluence: AGILITY_FACTOR_INFLUENCE,
	}, nil
}

// evaluateDragSpeedPoints calcs the drag speed. more points == fewer drag == good
func evaluateDragSpeedPoints(mode MODE, displ, wsa float64) float64 {
	displVol := displ / 1000 // assuming water is 1000 kg / m3
	// ratio between the edge length of wsa and displArea.
	wsaDisplRatio := math.Pow(math.Pow(wsa, 1/2)/math.Pow(displVol, 1/3), 2)

	impact := (wsa * wsaDisplRatio * MODE_DRAG_FACTOR[mode])
	// normalize (as more drag == less points) and reverse result into a scale 1.0-2.0
	return (2.0 - (impact / 20)) * DRAG_SPEED_POINT_PATCHER
}

// evaluateDownwindSpeedPoints calcs the downwind speed. more points == faster == good
func evaluateDownwindSpeedPoints(displ, asym, sym float64) float64 {
	// asymmetric and symmetric downwindsails are not differentiated, as its considered a "strategic decision".
	// the largest sail is counted, other smaller sails may be used in the race.
	sailArea := math.Max(sym, asym)
	displVol := displ / 1000 // assuming water is 1000 kg / m3

	// ratio between the edge length of sailArea and displVol.
	sailDisplRatio := math.Pow(math.Pow(sailArea, 1/2)/math.Pow(displVol, 1/3), 2)

	impact := (sailArea * sailDisplRatio)
	// normalize result into a scale 1.0-2.0
	return (1.0 + impact/(impact+10)) * DOWNWIND_SPEED_POINT_PATCHER
}

// evaluateUpwindSpeedPoints calcs the upwind speed. more points == faster == good
func evaluateUpwindSpeedPoints(displ, main, jib, forestay float64) float64 {
	// higher forestay means the sails can be trimmed to use higher winds which are generally faster due to surface friction.
	// this is not very influential, so only a small fraction of the jib is added.
	forestayFactor := forestay / 100
	sailArea := main + jib + (math.Pow(math.Pow(jib, 1/2)*forestayFactor, 2))
	displVol := displ / 1000 // assuming water is 1000 kg / m3

	// ratio between the edge length of sailArea and displVol.
	sailDisplRatio := math.Pow(math.Pow(sailArea, 1/2)/math.Pow(displVol, 1/3), 2)

	impact := (sailArea * sailDisplRatio)
	// normalize result into a scale 1.0-2.0
	return (1.0 + impact/(impact+10)) * UPWIND_SPEED_POINT_PATCHER
}

// evaluateStabilizationPoints calcs the boat stabilization. more points == better stabilization == good
func evaluateStabilizationPoints(wsa, draft, beam, loa, displ, main float64, material map[MATERIAL]float64, stabilization STABILIZATION, hull HULL) float64 {
	displVol := displ / 1000 // assuming water is 1000 kg / m3
	// sailDisplRatio is added to take strong heeling forces into account.
	// spinnaker and jib sails are generally a more controllable and minor heeling forces and therefore ignored.
	sailDisplRatio := math.Pow(main, 1/2) / math.Pow(displVol, 1/3)
	// wsaDisplRatio is added to take into account how much relative wsa the boat has.
	// A huge wsa is required for heavier boats, but for light boats this enhances formstability.
	wsaDisplRatio := math.Pow(wsa, 1/2) / math.Pow(displVol, 1/3)
	// beamLoaRatio is added to take into account the beam relative to loa.
	// A wide beam (relative to loa) also provides more formstability to the boat.
	beamLoaRatio := beam / loa

	// draft is taken into account because generally more draft == more ship in water == more stability
	basicStabilizationPoints := wsaDisplRatio * sailDisplRatio * beamLoaRatio * draft

	// ballastFactor is calculated by the amount of special ballast on board
	// it is multiplied by the stabilization factor, this makes ballast in combination with
	// keels weight stronger then ballast in combination with daggerboards.
	ballastPercentage := material[MATERIAL_BALLAST]
	if ballastPercentage <= 0 {
		ballastPercentage = 100 // by default the ballast is 100% of the displacement
	}
	ballastFactor := (ballastPercentage * STABILIZATION_STABILIZATION_FACTOR[stabilization]) / 100

	return basicStabilizationPoints * HULL_STABILIZATION_FACTOR[hull] * ballastFactor * STABILIZATION_POINT_PATCHER
}

// evaluateAgilityPoints calcs the boat agility. more points == better agility == good
func evaluateAgilityPoints(beam, loa, crew, displ float64, stabilization STABILIZATION, hull HULL) float64 {
	// beamLoaDiff is added to take into account the difference between loa and beam.
	// large difference means the ship is compact (and agile), small difference means it's long and thin which makes it less agile.
	beamLoaDiff := math.Max(loa, beam) / (math.Abs(loa-beam) + 0.0000001)
	// crewDisplRatio is added to take into account that more crew weight == more agility.
	// with more crew you can more effective rebalance the weight of the ship and therefore operate more agile.
	crewDisplRatio := crew / displ

	basicAgility := beamLoaDiff * crewDisplRatio

	// loa is mixed in here because in general longer ships have
	return basicAgility * HULL_AGILITY_FACTOR[hull] * STABILIZATION_AGILITY_FACTOR[stabilization] * AGILITY_POINT_PATCHER
}
