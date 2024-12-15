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

package output

type ShipMap map[string]ShipConfig

type ShipConfig struct {
	Team          string              `json:"team"`
	ShipInfo      ShipConfigInfo      `json:"boat_info"`
	ShipBaseSpec  ShipConfigBaseSpec  `json:"boat_base_spec"`
	ShipExtraSpec ShipConfigExtraSpec `json:"boat_extra_spec"`
	ShipRating    ShipConfigRating    `json:"boat_rating"`
}

type SHIP_INFO_SOURCE string

const (
	SHIP_INFO_MANUAL SHIP_INFO_SOURCE = "manual"
	SHIP_INFO_ORC    SHIP_INFO_SOURCE = "orc"
)

type ShipConfigInfo struct {
	Source   SHIP_INFO_SOURCE `json:"source"`
	Name     string           `json:"name"`
	Class    string           `json:"class"`
	Age      string           `json:"age"`
	Builder  string           `json:"builder"`
	Designer string           `json:"designer"`
}

type SHIP_BASE_SPEC_SOURCE string

const (
	SHIP_BASE_SPEC_MANUAL SHIP_BASE_SPEC_SOURCE = "manual"
	SHIP_BASE_SPEC_ORC    SHIP_BASE_SPEC_SOURCE = "orc"
)

type ShipConfigBaseSpec struct {
	Source    SHIP_BASE_SPEC_SOURCE       `json:"source"`
	Dimension ShipConfigBaseSpecDimension `json:"dimension"`
	SailArea  ShipConfigBaseSpecSailArea  `json:"sail_area"`
}

type ShipConfigBaseSpecDimension struct {
	LengthOverAll     float64 `json:"length_over_all"`
	Draft             float64 `json:"draft"`
	Beam              float64 `json:"beam"`
	ForestayHeight    float64 `json:"forestay_height"`
	WettedSurfaceArea float64 `json:"wetted_surface_area"`
	Displacement      float64 `json:"displacement"`
	CrewWeight        float64 `json:"crew_weight"`
}

type ShipConfigBaseSpecSailArea struct {
	Main                float64 `json:"main"`
	Jib                 float64 `json:"jib"`
	AsymmetricSpinnaker float64 `json:"asymmetric_spinnaker"`
	SymmetricSpinnaker  float64 `json:"symmetric_spinnaker"`
}

type SHIP_EXTRA_SPEC_SOURCE string

const (
	SHIP_EXTRA_SPEC_MANUAL SHIP_EXTRA_SPEC_SOURCE = "manual"
)

type ShipConfigExtraSpec struct {
	Source      SHIP_EXTRA_SPEC_SOURCE         `json:"source"`
	Design      ShipConfigExtraSpecDesign      `json:"design"`
	Composition ShipConfigExtraSpecComposition `json:"composition"`
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
	SHIP_EXTRA_SPEC_DESIGN_FOILS       SHIP_EXTRA_SPEC_DESIGN_HULL = "foils"
	SHIP_EXTRA_SPEC_DESIGN_DAGGERBOARD SHIP_EXTRA_SPEC_DESIGN_HULL = "daggerboard"
	SHIP_EXTRA_SPEC_DESIGN_CENTREBOARD SHIP_EXTRA_SPEC_DESIGN_HULL = "centreboard"
	SHIP_EXTRA_SPEC_DESIGN_KEEL        SHIP_EXTRA_SPEC_DESIGN_HULL = "keel"
)

type SHIP_EXTRA_SPEC_DESIGN_HULL string

const (
	SHIP_EXTRA_SPEC_DESIGN_MONO  SHIP_EXTRA_SPEC_DESIGN_HULL = "manual"
	SHIP_EXTRA_SPEC_DESIGN_MULTI SHIP_EXTRA_SPEC_DESIGN_HULL = "multi"
)

type ShipConfigExtraSpecDesign struct {
	Mode          SHIP_EXTRA_SPEC_DESIGN_MODE          `toml:"mode"`
	Stabilization SHIP_EXTRA_SPEC_DESIGN_STABILIZATION `toml:"stabilization"`
	Hull          SHIP_EXTRA_SPEC_DESIGN_HULL          `toml:"hull"`
}

type ShipConfigExtraSpecComposition struct {
	BallastPercentage float64 `toml:"ballast_percentage"`
	CfkPercentage     float64 `toml:"cfk_percentage"`
	AluPercentage     float64 `toml:"alu_percentage"`
	GfkPercentage     float64 `toml:"gfk_percentage"`
	WoodPercentage    float64 `toml:"wood_percentage"`
	EnginePercentage  float64 `toml:"engine_percentage"`
	AmenityPercentage float64 `toml:"amenity_percentage"`
}

type ShipConfigRating struct {
	Version             string  `json:"version"`
	Lifter              float64 `json:"lifter"`
	TCC                 float64 `json:"tcc"`
	SpeedFactor         float64 `json:"speed_factor"`
	AgilityFactor       float64 `json:"agility_factor"`
	StabilizationFactor float64 `json:"stabilization_factor"`
}
