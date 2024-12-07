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
	Team       string           `json:"team"`
	ShipInfo   ShipConfigInfo   `json:"boat_info"`
	ShipSpec   ShipConfigSpec   `json:"boat_spec"`
	ShipRating ShipConfigRating `json:"boat_rating"`
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

type SHIP_SPEC_SOURCE string

const (
	SHIP_SPEC_MANUAL SHIP_SPEC_SOURCE = "manual"
	SHIP_SPEC_ORC    SHIP_SPEC_SOURCE = "orc"
)

type ShipConfigSpec struct {
	Source    SHIP_SPEC_SOURCE    `json:"source"`
	Dimension ShipConfigDimension `json:"dimension"`
	SailArea  ShipConfigSailArea  `json:"sail_area"`
	Misc      ShipConfigMisc      `json:"misc"`
}

type ShipConfigDimension struct {
	LengthOverAll     float64 `json:"length_over_all"`
	Draft             float64 `json:"draft"`
	Beam              float64 `json:"beam"`
	ForestayHeight    float64 `json:"forestay_height"`
	WettedSurfaceArea float64 `json:"wetted_surface_area"`
}

type ShipConfigSailArea struct {
	Main                float64 `json:"main"`
	Jib                 float64 `json:"jib"`
	AsymmetricSpinnaker float64 `json:"asymmetric_spinnaker"`
	SymmetricSpinnaker  float64 `json:"symmetric_spinnaker"`
}

type ShipConfigMisc struct {
	StabilityIndex       float64 `json:"stability_index"`
	SailingDisplacement  float64 `json:"sailing_displacement"`
	MeasuredDisplacement float64 `json:"measured_displacement"`
	MaxCrewWeight        float64 `json:"max_crew_weight"`
}

type ShipConfigRating struct {
	Version string  `json:"version"`
	TCC     float64 `json:"tcc"`
}
