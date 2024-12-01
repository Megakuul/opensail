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

package orc

// DownBoatRMS specifies the response object from the offshore racing congress
// api action 'DownBoatRMS' (https://data.orc.org/tools.php?c=pcs).
type DownBoatRMS struct {
	Rms []RMS `json:"rms"`
}

// I have no idea for what orc RMS stands for, but its essentially the ship configuration.
type RMS struct {
	// NatAuth is the National Authority issuing the certificate (e.g., GRE for Greece)
	NatAuth string `json:"NatAuth"`
	// BIN is the Boat Identification Number in the ORC database
	BIN string `json:"BIN"`
	// CertNo is the unique certificate number
	CertNo string `json:"CertNo"`
	// RefNo is the reference number for the certificate
	RefNo string `json:"RefNo"`
	// SailNo is the official sail number of the boat
	SailNo string `json:"SailNo"`
	// YachtName is the registered name of the vessel
	YachtName string `json:"YachtName"`
	// Class represents the boat model/class
	Class string `json:"Class"`
	// Builder is the manufacturer of the boat
	Builder string `json:"Builder"`
	// Designer is the naval architect/design firm
	Designer string `json:"Designer"`
	// C_Type represents the certificate type
	C_Type string `json:"C_Type"`
	// Family indicates the boat family/category
	Family string `json:"Family"`
	// IssueDate is when the certificate was issued
	IssueDate string `json:"IssueDate"`
	// WSS is the Wetted Surface Area (square meters)
	WSS float64 `json:"WSS"`
	// AreaMain is the mainsail area (square meters)
	AreaMain float64 `json:"Area_Main"`
	// AreaJib is the maximum jib area (square meters)
	AreaJib float64 `json:"Area_Jib"`
	// AreaSym is the symmetric spinnaker area if equipped (square meters)
	AreaSym float64 `json:"Area_Sym"`
	// AreaAsym is the asymmetric spinnaker area (square meters)
	AreaAsym float64 `json:"Area_Asym"`
	// AgeYear is the year of initial build
	AgeYear int `json:"Age_Year"`
	// CrewWT is the maximum crew weight allowed (kg)
	CrewWT float64 `json:"CrewWT"`
	// LOA is Length Overall (meters)
	LOA float64 `json:"LOA"`
	// IMSL is the forestay height (meters)
	IMSL float64 `json:"IMSL"`
	// Draft is the maximum depth below waterline (meters)
	Draft float64 `json:"Draft"`
	// MB is the Maximum Beam width (meters)
	MB float64 `json:"MB"`
	// DsplSailing is the displacement in sailing condition (kg)
	DsplSailing float64 `json:"Dspl_Sailing"`
	// DsplMeasurement is the measured displacement (kg)
	DsplMeasurement float64 `json:"Dspl_Measurement"`
	// StabilityIndex represents the boat's stability rating
	StabilityIndex float64 `json:"Stability_Index"`
	// GPH is the General Purpose Handicap (seconds/mile)
	GPH float64 `json:"GPH"`
}
