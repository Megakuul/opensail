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
	Owner      string           `json:"owner" validate:"required"`
	Team       string           `json:"team" validate:"required"`
	BoatInfo   ShipConfigInfo   `json:"boat_info" validate:"required"`
	BoatSpec   ShipConfigSpec   `json:"boat_spec" validate:"required"`
	BoatRating ShipConfigRating `json:"boat_rating" validate:"required"`
}

type ShipConfigInfo struct {
	Source               string `json:"source"`
	Name                 string `json:"name"`
	Class                string `json:"class"`
	Family               string `json:"family"`
	Age                  string `json:"age"`
	Builder              string `json:"builder"`
	Designer             string `json:"designer"`
	CertificateType      string `json:"certificate_type"`
	CertificateIssueDate string `json:"certificate_issue_date"`
}

type ShipConfigSpec struct {
	Source    string              `json:"source"`
	Dimension ShipConfigDimension `json:"dimension"`
	SailArea  ShipConfigSailArea  `json:"sail_area"`
	Misc      ShipConfigMisc      `json:"misc"`
}

type ShipConfigDimension struct {
	LengthOverAll     float64 `json:"length_over_all"`
	Draft             float64 `json:"draft"`
	Beam              float64 `json:"beam"`
	ForestrayHeight   float64 `json:"forestray_height"`
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
	OS             ShipConfigRatingSystem `json:"os" `
	SwissYardstick ShipConfigRatingSystem `json:"swiss_yardstick" `
	ORC            ShipConfigRatingSystem `json:"orc"`
	IRC            ShipConfigRatingSystem `json:"irc"`
}

type ShipConfigRatingSystem struct {
	Source string  `json:"source"`
	Value  string  `json:"value"`
	TCC    float64 `json:"tcc"`
}
