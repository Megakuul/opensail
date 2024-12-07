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
	// Spec contains boat dimensions and measurements used for rating
	Spec ShipConfigSpec `toml:"spec" validate:"required"`
}

type SHIP_INFO_SOURCE string

const (
	SHIP_INFO_MANUAL SHIP_INFO_SOURCE = "manual"
	SHIP_INFO_ORC    SHIP_INFO_SOURCE = "orc"
)

type ShipConfigInfo struct {
	// Source indicates data origin: "manual" or "orc"
	Source SHIP_INFO_SOURCE `toml:"source" validate:"required"`
	// ORCSailingNo is the boat identifier in ORC database
	ORCSailingNo string `toml:"orc_sailing_no"`
}

type SHIP_SPEC_SOURCE string

const (
	SHIP_SPEC_MANUAL SHIP_SPEC_SOURCE = "manual"
	SHIP_SPEC_ORC    SHIP_SPEC_SOURCE = "orc"
)

type ShipConfigSpec struct {
	// Source indicates data origin: "manual" or "orc"
	Source SHIP_SPEC_SOURCE `toml:"source" validate:"required"`
	// ORCSailingNo is the boat identifier in ORC database
	ORCSailingNo string `toml:"orc_sailing_no"`
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

// ShipSpec specifies the toml representation of the ship specification.
type ShipSpec struct {
	// Dimension contains physical measurements of the boat
	Dimension ShipSpecDimension `toml:"dimension" validate:"required"`
	// SailArea contains measurements for different sail types
	SailArea ShipSpecSailArea `toml:"sail_area" validate:"required"`
	// Misc contains various other specifications
	Misc ShipSpecMisc `toml:"misc" validate:"required"`
}

type ShipSpecDimension struct {
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
}

type ShipSpecSailArea struct {
	// Main specifies the area of the main sail in square meters
	Main float64 `toml:"main" validate:"required"`
	// Jib specifies the area of the largest onboard jib sail in square meters
	Jib float64 `toml:"jib" validate:"required"`
	// AsymmetricSpinnaker specifies the area of the largest onboard downwind sail in square meters
	AsymmetricSpinnaker float64 `toml:"asymmetric_spinnaker" validate:"required"`
	// SymmetricSpinnaker specifies the area of the largest onboard symmetric spinnaker in square meters
	SymmetricSpinnaker float64 `toml:"symmetric_spinnaker" validate:"required"`
}

type ShipSpecMisc struct {
	// StabilityIndex specifies the stability index of the ship
	StabilityIndex float64 `toml:"stability_index" validate:"required"`
	// SailingDisplacement specifies the displacement while sailing in kg
	SailingDisplacement float64 `toml:"sailing_displacement" validate:"required"`
	// MeasuredDisplacement specifies the measured displacement in kg
	MeasuredDisplacement float64 `toml:"measured_displacement" validate:"required"`
	// MaxCrewWeight specifies the maximum crew weight in kg
	MaxCrewWeight float64 `toml:"max_crew_weight" validate:"required"`
}
