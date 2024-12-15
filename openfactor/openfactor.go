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

type MODE int64

const (
	MODE_DEFAULT MODE = iota // default==unspecified provides 0 extra rating
	MODE_HYDROFOIL
	MODE_PLANING
	MODE_SEMI
	MODE_DISPLACE
)

var MODE_FACTOR = map[MODE]float64{
	MODE_DEFAULT:   1,
	MODE_HYDROFOIL: 0.9,
	MODE_PLANING:   0.7,
	MODE_SEMI:      0.5,
	MODE_DISPLACE:  0.3,
}

type STABILIZATION int64

const (
	STABILIZATION_DEFAULT STABILIZATION = iota // default==unspecified provides 0 extra rating
	STABILIZATION_KEEL
	STABILIZATION_CENTREBOARD
	STABILIZATION_DAGGERBOARD
	STABILIZATION_FOILS
)

type HULL int64

const (
	HULL_DEFAULT HULL = iota // default==unspecified provides 0 extra rating
	HULL_MULTI
	HULL_MONO
)

type MATERIAL int64

const (
	MATERIAL_DEFAULT MATERIAL = iota // default==unspecified provides 0 extra rating
	MATERIAL_KEEL
	MATERIAL_CFK
	MATERIAL_GFK
	MATERIAL_WOOD
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
	// WSS specifies the wetted surface area in square meters.
	WSS float64

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
	// Openfactor is calculated flat from 0.0 on. To create a more human friendly coefficient
	// the resulting factors are lifted by this number. Therefore the Lifter is the difference between
	// the reference ship factor and 1.0 (e.g. RefShipTCC = 0.83; Lifter = 0.17; results in RefShipTCC == 1.00).
	// The Lifter is already applied to all results and is just returned for informational purposes.
	Lifter float64

	// Time correction coefficient produced by the algorithm.
	TCC float64
	// SpeedFactor specifies the factor the boat retrieved in category "Speed".
	SpeedFactor float64
	// AgilityFactor specifies the factor the boat retrieved in category "Agility".
	AgilityFactor float64
	// StabilizationFactor specifies the factor the boat retrieved in category "Stability".
	StabilizationFactor float64
}

func EvaluateFactor(input *EvaluationInput) (*EvaluationOutput, error) {
	return &EvaluationOutput{TCC: 1}, nil
}

func evaluateHullFactor(displacement, crew float64, composition map[MATERIAL]float64) float64 {

}

func evaluateSailFactor(displacement, crew float64, composition map[MATERIAL]float64) float64 {

}

func evaluateWeightFactor(displacement, crew float64, composition map[MATERIAL]float64) float64 {

}

func evaluateDownwindSpeedFactor(displ, asym, sym float64) float64 {
	sailArea := math.Max(sym, asym)
	displArea := displ / 1000 // assuming water is 1000 kg / m3

	// ratio between the edge length of sailArea and displArea.
	sailDisplRatio := math.Pow(sailArea, 1/2) / math.Pow(displArea, 1/3)

	return (sailArea * sailDisplRatio) / DOWNWIND_SPEED_FACTOR
}

func evaluateAgilityFactor(loa, beam, draft, wss, main, jib, asym, sym float64) float64 {

}

func evaluateStabilizationFactor(loa, beam, draft, wss, main, jib float64, hull HULL, stabilization STABILIZATION) float64 {

}
