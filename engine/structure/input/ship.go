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

package input

// ShipConfig specifies the toml representation of the ship configuration.
type ShipConfig struct {
	// Team specifies the team identifier currently sailing this boat
	Team string `toml:"team" validate:"required"`
	// Info contains general boat information that doesn't influence rating
	Info ShipConfigInfo `toml:"info" validate:"required"`
	// BaseSpec contains boat dimensions and measurements used for rating
	BaseSpec ShipConfigBaseSpec `toml:"base_spec" validate:"required"`
	// ExtraSpec contains additional boat characteristics used to improve rating
	ExtraSpec ShipConfigExtraSpec `toml:"extra_spec" validate:"required"`
}

type SHIP_INFO_SOURCE string

const (
	SHIP_INFO_MANUAL SHIP_INFO_SOURCE = "manual"
	SHIP_INFO_ORC    SHIP_INFO_SOURCE = "orc"
)

type ShipConfigInfo struct {
	// Source indicates data origin: "manual" or "orc"
	Source SHIP_INFO_SOURCE `toml:"source" validate:"required"`
	// ORCRefNo is the boat certificate identifier in ORC database
	ORCRefNo string `toml:"orc_ref_no"`
}

type SHIP_BASE_SPEC_SOURCE string

const (
	SHIP_BASE_SPEC_MANUAL SHIP_BASE_SPEC_SOURCE = "manual"
	SHIP_BASE_SPEC_ORC    SHIP_BASE_SPEC_SOURCE = "orc"
)

type ShipConfigBaseSpec struct {
	// Source indicates data origin: "manual" or "orc"
	Source SHIP_BASE_SPEC_SOURCE `toml:"source" validate:"required"`
	// ORCRefNo is the boat certificate identifier in ORC database
	ORCRefNo string `toml:"orc_ref_no"`
}

type SHIP_EXTRA_SPEC_SOURCE string

const (
	SHIP_EXTRA_SPEC_MANUAL SHIP_EXTRA_SPEC_SOURCE = "manual"
)

type ShipConfigExtraSpec struct {
	// Source indicates data origin: "manual"
	Source SHIP_EXTRA_SPEC_SOURCE `toml:"source" validate:"required"`
}

// ShipInfo specifies the toml representation of the ship information.
type ShipInfo struct {
	// Name is the friendly name of the boat (e.g. AETHER)
	Name string `toml:"name" validate:"required"`
	// Class specifies the boat class/model (e.g. DEHLER 30 OD)
	Class string `toml:"class"`
	// Age is the vessel's construction year
	Age string `toml:"age"`
	// Builder is the name of the boat manufacturer (e.g. DEHLER)
	Builder string `toml:"builder"`
	// Designer specifies who designed the boat (e.g. JUDEL/VROLIJK)
	Designer string `toml:"designer"`
}

// ShipBaseSpec specifies the toml representation of the ship base specification.
type ShipBaseSpec struct {
	// Dimension contains physical measurements of the boat
	Dimension ShipBaseSpecDimension `toml:"dimension" validate:"required"`
	// SailArea contains measurements for different sail types
	SailArea ShipBaseSpecSailArea `toml:"sail_area" validate:"required"`
}

type ShipBaseSpecDimension struct {
	// LengthOverAll specifies the length in meters from the foremost to the rearmost point of the ship
	LengthOverAll float64 `toml:"length_over_all" validate:"required"`
	// Draft specifies the vertical distance from the bottom of the keel to the waterline in meters
	Draft float64 `toml:"draft" validate:"required"`
	// Beam specifies the largest width of the beam in meters
	Beam float64 `toml:"beam" validate:"required"`
	// ForestayHeight specifies the vertical height of the forestay in meters
	ForestayHeight float64 `toml:"forestay_height" validate:"required"`
	// WettedSurfaceArea specifies the surface area touching the water in square meters
	WettedSurfaceArea float64 `toml:"wetted_surface_area" validate:"required"`
	// SailingDisplacement specifies the displacement while sailing in kg
	SailingDisplacement float64 `toml:"sailing_displacement" validate:"required"`
	// MaxCrewWeight specifies the maximum crew weight in kg
	MaxCrewWeight float64 `toml:"max_crew_weight" validate:"required"`
}

type ShipBaseSpecSailArea struct {
	// Main specifies the area of the main sail in square meters
	Main float64 `toml:"main" validate:"required"`
	// Jib specifies the area of the largest onboard jib sail in square meters
	Jib float64 `toml:"jib" validate:"required"`
	// AsymmetricSpinnaker specifies the area of the largest onboard downwind sail in square meters
	AsymmetricSpinnaker float64 `toml:"asymmetric_spinnaker" validate:"required"`
	// SymmetricSpinnaker specifies the area of the largest onboard symmetric spinnaker in square meters
	SymmetricSpinnaker float64 `toml:"symmetric_spinnaker" validate:"required"`
}

// ShipExtraSpec specifies the toml representation of the ship extra specification.
type ShipExtraSpec struct {
	// Design holds information about the ships design characteristics
	Design ShipExtraSpecDesign `toml:"design" validate:"required"`
	// Composition specifies how the ships weight is composed
	Composition ShipExtraSpecComposition `toml:"composition" validate:"required"`
}

type SHIP_EXTRA_SPEC_DESIGN_MODE string

const (
	SHIP_EXTRA_SPEC_DESIGN_DISPLACE  SHIP_EXTRA_SPEC_DESIGN_MODE = "displace"
	SHIP_EXTRA_SPEC_DESIGN_SEMI      SHIP_EXTRA_SPEC_DESIGN_MODE = "semi"
	SHIP_EXTRA_SPEC_DESIGN_PLANING   SHIP_EXTRA_SPEC_DESIGN_MODE = "planing"
	SHIP_EXTRA_SPEC_DESIGN_HYDROFOIL SHIP_EXTRA_SPEC_DESIGN_MODE = "hydrofoil"
)

type SHIP_EXTRA_SPEC_DESIGN_STABILIZATION string

const (
	SHIP_EXTRA_SPEC_DESIGN_FULLKEEL    SHIP_EXTRA_SPEC_DESIGN_STABILIZATION = "fullkeel"
	SHIP_EXTRA_SPEC_DESIGN_BULBKEEL    SHIP_EXTRA_SPEC_DESIGN_STABILIZATION = "bulbkeel"
	SHIP_EXTRA_SPEC_DESIGN_FINKEEL     SHIP_EXTRA_SPEC_DESIGN_STABILIZATION = "finkeel"
	SHIP_EXTRA_SPEC_DESIGN_DAGGERBOARD SHIP_EXTRA_SPEC_DESIGN_STABILIZATION = "daggerboard"
	SHIP_EXTRA_SPEC_DESIGN_CENTREBOARD SHIP_EXTRA_SPEC_DESIGN_STABILIZATION = "centreboard"
	SHIP_EXTRA_SPEC_DESIGN_FOILS       SHIP_EXTRA_SPEC_DESIGN_STABILIZATION = "foils"
)

type SHIP_EXTRA_SPEC_DESIGN_HULL string

const (
	SHIP_EXTRA_SPEC_DESIGN_MONO  SHIP_EXTRA_SPEC_DESIGN_HULL = "mono"
	SHIP_EXTRA_SPEC_DESIGN_MULTI SHIP_EXTRA_SPEC_DESIGN_HULL = "multi"
)

type ShipExtraSpecDesign struct {
	// Mode specifies the type of the hull design/mode [displace; semi; planing; hydrofoil;]
	Mode SHIP_EXTRA_SPEC_DESIGN_MODE `toml:"mode"`
	// Stabilization specifies the method used to stabilize the ship [foils; centreboard; daggerboard; keel;]
	Stabilization SHIP_EXTRA_SPEC_DESIGN_STABILIZATION `toml:"stabilization"`
	// Hull specifies the hull type [mono; multi;]
	Hull SHIP_EXTRA_SPEC_DESIGN_HULL `toml:"hull"`
}

type ShipExtraSpecComposition struct {
	// BallastPercentage specifies how much weight of the ship is pure ballast (usually lead)
	BallastPercentage float64 `toml:"ballast_percentage"`
	// CfkPercentage specifies how much weight of the ship is cfk
	CfkPercentage float64 `toml:"cfk_percentage"`
	// AluPercentage specifies how much weight of the ship is aluminium
	AluPercentage float64 `toml:"alu_percentage"`
	// GfkPercentage specifies how much weight of the ship is gfk
	GfkPercentage float64 `toml:"gfk_percentage"`
	// WoodPercentage specifies how much weight of the ship is wood
	WoodPercentage float64 `toml:"wood_percentage"`
	// EnginePercentage specifies how much weight of the ship is engine
	EnginePercentage float64 `toml:"engine_percentage"`
	// AmenityPercentage specifies how much weight of the ship is amenities
	AmenityPercentage float64 `toml:"amenity_percentage"`
}
