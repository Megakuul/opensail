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
	// Time correction coefficient produced by the algorithm.
	TCC float64
}

func EvaluateFactor(input *EvaluationInput) (*EvaluationOutput, error) {
	return &EvaluationOutput{TCC: 1}, nil
}
