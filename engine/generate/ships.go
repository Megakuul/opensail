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

package generate

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/megakuul/opensail/engine/adapter/orc"
	"github.com/megakuul/opensail/engine/structure/input"
	"github.com/megakuul/opensail/engine/structure/output"
	"github.com/megakuul/opensail/openfactor"
)

// generateShips generates the shipMap.
func generateShips(repoPath string, ships map[string]struct{}, shipStruct input.ShipStructure) ([]byte, error) {
	shipMap := output.ShipMap{}

	for ship := range ships {
		shipPath := path.Join(repoPath, shipStruct.BasePath, ship)
		shipConfigRaw, err := os.ReadFile(path.Join(shipPath, shipStruct.ConfigFile))
		if err != nil {
			return nil, fmt.Errorf("failed to read ship config (ship '%s'): %w", ship, err)
		}
		shipConfig := &input.ShipConfig{}
		err = toml.Unmarshal(shipConfigRaw, shipConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to parse ship config (ship '%s'): %w", ship, err)
		}

		outputShipInfo, err := generateShipInfo(shipConfig.Info, path.Join(shipPath, shipStruct.InfoFile))
		if err != nil {
			return nil, fmt.Errorf("failed to generate ship info (ship '%s'): %w", ship, err)
		}

		outputShipSpec, err := generateShipSpec(shipConfig.Spec, path.Join(shipPath, shipStruct.SpecFile))
		if err != nil {
			return nil, fmt.Errorf("failed to generate ship spec (ship '%s'): %w", ship, err)
		}

		outputShipRating, err := generateShipRating(outputShipSpec)
		if err != nil {
			return nil, fmt.Errorf("failed to generate ship rating (ship '%s'): %w", ship, err)
		}

		shipMap[ship] = output.ShipConfig{
			Team:       shipConfig.Team,
			ShipInfo:   *outputShipInfo,
			ShipSpec:   *outputShipSpec,
			ShipRating: *outputShipRating,
		}
	}

	shipMapRaw, err := json.Marshal(shipMap)
	if err != nil {
		return nil, err
	}

	return shipMapRaw, nil
}

func generateShipInfo(info input.ShipConfigInfo, infoPath string) (*output.ShipConfigInfo, error) {
	switch info.Source {
	case input.SHIP_INFO_MANUAL:
		shipInfoRaw, err := os.ReadFile(infoPath)
		if err != nil {
			return nil, err
		}
		shipInfo := &input.ShipInfo{}
		err = toml.Unmarshal(shipInfoRaw, shipInfo)
		if err != nil {
			return nil, err
		}

		return &output.ShipConfigInfo{
			Source:   output.SHIP_INFO_MANUAL,
			Name:     shipInfo.Name,
			Class:    shipInfo.Class,
			Age:      shipInfo.Age,
			Builder:  shipInfo.Builder,
			Designer: shipInfo.Designer,
		}, nil
	case input.SHIP_INFO_ORC:
		if info.ORCRefNo == "" {
			return nil, fmt.Errorf("invalid ship orc RefNo. '%s'", info.ORCRefNo)
		}

		downBoatRms, err := orc.GetDownBoatRMS(info.ORCRefNo)
		if err != nil {
			return nil, err
		}

		if len(downBoatRms.Rms) < 1 {
			return nil, fmt.Errorf("ship with RefNo. '%s' was not found on orc database", info.ORCRefNo)
		}
		orcShip := downBoatRms.Rms[0]

		return &output.ShipConfigInfo{
			Source:   output.SHIP_INFO_ORC,
			Name:     orcShip.YachtName,
			Class:    orcShip.Class,
			Age:      strconv.Itoa(orcShip.AgeYear),
			Builder:  orcShip.Builder,
			Designer: orcShip.Designer,
		}, nil
	default:
		return nil, fmt.Errorf("invalid ship info source '%s'", info.Source)
	}
}

func generateShipSpec(spec input.ShipConfigSpec, specPath string) (*output.ShipConfigSpec, error) {
	switch spec.Source {
	case input.SHIP_SPEC_MANUAL:
		shipSpecRaw, err := os.ReadFile(specPath)
		if err != nil {
			return nil, err
		}
		shipSpec := &input.ShipSpec{}
		err = toml.Unmarshal(shipSpecRaw, shipSpec)
		if err != nil {
			return nil, err
		}

		return &output.ShipConfigSpec{
			Source: output.SHIP_SPEC_MANUAL,
			Dimension: output.ShipConfigDimension{
				LengthOverAll:     shipSpec.Dimension.LengthOverAll,
				Draft:             shipSpec.Dimension.Draft,
				Beam:              shipSpec.Dimension.Beam,
				ForestayHeight:    shipSpec.Dimension.ForestayHeight,
				WettedSurfaceArea: shipSpec.Dimension.WettedSurfaceArea,
			},
			SailArea: output.ShipConfigSailArea{
				Main:                shipSpec.SailArea.Main,
				Jib:                 shipSpec.SailArea.Jib,
				AsymmetricSpinnaker: shipSpec.SailArea.AsymmetricSpinnaker,
				SymmetricSpinnaker:  shipSpec.SailArea.SymmetricSpinnaker,
			},
			Misc: output.ShipConfigMisc{
				SailingDisplacement: shipSpec.Misc.SailingDisplacement,
				MaxCrewWeight:       shipSpec.Misc.MaxCrewWeight,
			},
		}, nil
	case input.SHIP_SPEC_ORC:
		if spec.ORCRefNo == "" {
			return nil, fmt.Errorf("invalid ship orc RefNo. '%s'", spec.ORCRefNo)
		}

		downBoatRms, err := orc.GetDownBoatRMS(spec.ORCRefNo)
		if err != nil {
			return nil, err
		}

		if len(downBoatRms.Rms) < 1 {
			return nil, fmt.Errorf("ship with RefNo. '%s' was not found on orc database", spec.ORCRefNo)
		}
		orcShip := downBoatRms.Rms[0]

		return &output.ShipConfigSpec{
			Source: output.SHIP_SPEC_ORC,
			Dimension: output.ShipConfigDimension{
				LengthOverAll:     orcShip.LOA,
				Draft:             orcShip.Draft,
				Beam:              orcShip.MB,
				ForestayHeight:    orcShip.IMSL,
				WettedSurfaceArea: orcShip.WSS,
			},
			SailArea: output.ShipConfigSailArea{
				Main:                orcShip.AreaMain,
				Jib:                 orcShip.AreaJib,
				AsymmetricSpinnaker: orcShip.AreaAsym,
				SymmetricSpinnaker:  orcShip.AreaSym,
			},
			Misc: output.ShipConfigMisc{
				StabilityIndex:       orcShip.StabilityIndex,
				MeasuredDisplacement: orcShip.DsplMeasurement,
				SailingDisplacement:  orcShip.DsplSailing,
				MaxCrewWeight:        orcShip.CrewWT,
			},
		}, nil
	default:
		return nil, fmt.Errorf("invalid ship spec source '%s'", spec.Source)
	}
}

func generateShipRating(spec *output.ShipConfigSpec) (*output.ShipConfigRating, error) {
	factorOutput, err := openfactor.EvaluateFactor(&openfactor.EvaluationInput{
		LOA:                     spec.Dimension.LengthOverAll,
		MaxDraft:                spec.Dimension.Draft,
		MaxBeam:                 spec.Dimension.Beam,
		IMSL:                    spec.Dimension.ForestayHeight,
		WSS:                     spec.Dimension.WettedSurfaceArea,
		MainSailArea:            spec.SailArea.Main,
		JibSailArea:             spec.SailArea.Jib,
		AsymmetricSpinnakerArea: spec.SailArea.AsymmetricSpinnaker,
		SymmetricSpinnakerArea:  spec.SailArea.SymmetricSpinnaker,
		StabilityIndex:          spec.Misc.StabilityIndex,
		Displacement:            spec.Misc.SailingDisplacement,
		CrewWeight:              spec.Misc.MaxCrewWeight,
	})
	if err != nil {
		return nil, err
	}

	return &output.ShipConfigRating{
		Version: openfactor.Version(),
		TCC:     factorOutput.TCC,
	}, nil
}
