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

		outputShipBaseSpec, err := generateShipBaseSpec(shipConfig.BaseSpec, path.Join(shipPath, shipStruct.BaseSpecFile))
		if err != nil {
			return nil, fmt.Errorf("failed to generate ship spec (ship '%s'): %w", ship, err)
		}

		outputShipExtraSpec, err := generateShipExtraSpec(shipConfig.ExtraSpec, path.Join(shipPath, shipStruct.ExtraSpecFile))
		if err != nil {
			return nil, fmt.Errorf("failed to generate ship spec (ship '%s'): %w", ship, err)
		}

		outputShipRating, err := generateShipRating(outputShipBaseSpec, outputShipExtraSpec)
		if err != nil {
			return nil, fmt.Errorf("failed to generate ship rating (ship '%s'): %w", ship, err)
		}

		shipMap[ship] = output.ShipConfig{
			Team:          shipConfig.Team,
			ShipInfo:      *outputShipInfo,
			ShipBaseSpec:  *outputShipBaseSpec,
			ShipExtraSpec: *outputShipExtraSpec,
			ShipRating:    *outputShipRating,
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

func generateShipBaseSpec(spec input.ShipConfigBaseSpec, specPath string) (*output.ShipConfigBaseSpec, error) {
	switch spec.Source {
	case input.SHIP_BASE_SPEC_MANUAL:
		shipSpecRaw, err := os.ReadFile(specPath)
		if err != nil {
			return nil, err
		}
		shipSpec := &input.ShipBaseSpec{}
		err = toml.Unmarshal(shipSpecRaw, shipSpec)
		if err != nil {
			return nil, err
		}

		return &output.ShipConfigBaseSpec{
			Source: output.SHIP_BASE_SPEC_MANUAL,
			Dimension: output.ShipConfigBaseSpecDimension{
				LengthOverAll:     shipSpec.Dimension.LengthOverAll,
				Draft:             shipSpec.Dimension.Draft,
				Beam:              shipSpec.Dimension.Beam,
				ForestayHeight:    shipSpec.Dimension.ForestayHeight,
				WettedSurfaceArea: shipSpec.Dimension.WettedSurfaceArea,
				Displacement:      shipSpec.Dimension.SailingDisplacement,
				CrewWeight:        shipSpec.Dimension.MaxCrewWeight,
			},
			SailArea: output.ShipConfigBaseSpecSailArea{
				Main:                shipSpec.SailArea.Main,
				Jib:                 shipSpec.SailArea.Jib,
				AsymmetricSpinnaker: shipSpec.SailArea.AsymmetricSpinnaker,
				SymmetricSpinnaker:  shipSpec.SailArea.SymmetricSpinnaker,
			},
		}, nil
	case input.SHIP_BASE_SPEC_ORC:
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

		return &output.ShipConfigBaseSpec{
			Source: output.SHIP_BASE_SPEC_ORC,
			Dimension: output.ShipConfigBaseSpecDimension{
				LengthOverAll:     orcShip.LOA,
				Draft:             orcShip.Draft,
				Beam:              orcShip.MB,
				ForestayHeight:    orcShip.IMSL,
				WettedSurfaceArea: orcShip.WSS,
				Displacement:      orcShip.DsplSailing,
				CrewWeight:        orcShip.CrewWT,
			},
			SailArea: output.ShipConfigBaseSpecSailArea{
				Main:                orcShip.AreaMain,
				Jib:                 orcShip.AreaJib,
				AsymmetricSpinnaker: orcShip.AreaAsym,
				SymmetricSpinnaker:  orcShip.AreaSym,
			},
		}, nil
	default:
		return nil, fmt.Errorf("invalid ship base spec source '%s'", spec.Source)
	}
}

func generateShipExtraSpec(spec input.ShipConfigExtraSpec, specPath string) (*output.ShipConfigExtraSpec, error) {
	switch spec.Source {
	case input.SHIP_EXTRA_SPEC_MANUAL:
		shipSpecRaw, err := os.ReadFile(specPath)
		if err != nil {
			return nil, err
		}
		shipSpec := &input.ShipExtraSpec{}
		err = toml.Unmarshal(shipSpecRaw, shipSpec)
		if err != nil {
			return nil, err
		}

		return &output.ShipConfigExtraSpec{
			Source: output.SHIP_EXTRA_SPEC_MANUAL,
			Design: output.ShipConfigExtraSpecDesign{
				Mode:          output.SHIP_EXTRA_SPEC_DESIGN_MODE(shipSpec.Design.Mode),
				Stabilization: output.SHIP_EXTRA_SPEC_DESIGN_STABILIZATION(shipSpec.Design.Stabilization),
				Hull:          output.SHIP_EXTRA_SPEC_DESIGN_HULL(shipSpec.Design.Hull),
			},
			Composition: output.ShipConfigExtraSpecComposition{
				BallastPercentage: shipSpec.Composition.BallastPercentage,
				CfkPercentage:     shipSpec.Composition.CfkPercentage,
				AluPercentage:     shipSpec.Composition.AluPercentage,
				GfkPercentage:     shipSpec.Composition.GfkPercentage,
				WoodPercentage:    shipSpec.Composition.WoodPercentage,
				EnginePercentage:  shipSpec.Composition.EnginePercentage,
				AmenityPercentage: shipSpec.Composition.AmenityPercentage,
			},
		}, nil
	default:
		return nil, fmt.Errorf("invalid ship extra spec source '%s'", spec.Source)
	}
}

func generateShipRating(baseSpec *output.ShipConfigBaseSpec, extraSpec *output.ShipConfigExtraSpec) (*output.ShipConfigRating, error) {
	mode := openfactor.MODE_DEFAULT
	switch extraSpec.Design.Mode {
	case output.SHIP_EXTRA_SPEC_DESIGN_HYDROFOIL:
		mode = openfactor.MODE_HYDROFOIL
	case output.SHIP_EXTRA_SPEC_DESIGN_PLANING:
		mode = openfactor.MODE_PLANING
	case output.SHIP_EXTRA_SPEC_DESIGN_SEMI:
		mode = openfactor.MODE_SEMI
	case output.SHIP_EXTRA_SPEC_DESIGN_DISPLACE:
		mode = openfactor.MODE_DISPLACE
	}

	stabilization := openfactor.STABILIZATION_DEFAULT
	switch extraSpec.Design.Stabilization {
	case output.SHIP_EXTRA_SPEC_DESIGN_FULLKEEL:
		stabilization = openfactor.STABILIZATION_FULLKEEL
	case output.SHIP_EXTRA_SPEC_DESIGN_BULBKEEL:
		stabilization = openfactor.STABILIZATION_BULBKEEL
	case output.SHIP_EXTRA_SPEC_DESIGN_FINKEEL:
		stabilization = openfactor.STABILIZATION_FINKEEL
	case output.SHIP_EXTRA_SPEC_DESIGN_CENTREBOARD:
		stabilization = openfactor.STABILIZATION_CENTREBOARD
	case output.SHIP_EXTRA_SPEC_DESIGN_DAGGERBOARD:
		stabilization = openfactor.STABILIZATION_DAGGERBOARD
	case output.SHIP_EXTRA_SPEC_DESIGN_FOILS:
		stabilization = openfactor.STABILIZATION_FOILS
	}

	hull := openfactor.HULL_DEFAULT
	switch extraSpec.Design.Hull {
	case output.SHIP_EXTRA_SPEC_DESIGN_MONO:
		hull = openfactor.HULL_MONO
	case output.SHIP_EXTRA_SPEC_DESIGN_MULTI:
		hull = openfactor.HULL_MULTI
	}

	var defaultComposition float64 = 100.0
	defaultComposition -= extraSpec.Composition.BallastPercentage
	defaultComposition -= extraSpec.Composition.CfkPercentage
	defaultComposition -= extraSpec.Composition.AluPercentage
	defaultComposition -= extraSpec.Composition.GfkPercentage
	defaultComposition -= extraSpec.Composition.WoodPercentage
	defaultComposition -= extraSpec.Composition.EnginePercentage
	defaultComposition -= extraSpec.Composition.AmenityPercentage
	if defaultComposition < 0 {
		return nil, fmt.Errorf("invalid ship composition; exceeded 100%%")
	}

	factorOutput, err := openfactor.EvaluateFactor(&openfactor.EvaluationInput{
		LOA:                     baseSpec.Dimension.LengthOverAll,
		MaxDraft:                baseSpec.Dimension.Draft,
		MaxBeam:                 baseSpec.Dimension.Beam,
		IMSL:                    baseSpec.Dimension.ForestayHeight,
		WSA:                     baseSpec.Dimension.WettedSurfaceArea,
		MainSailArea:            baseSpec.SailArea.Main,
		JibSailArea:             baseSpec.SailArea.Jib,
		AsymmetricSpinnakerArea: baseSpec.SailArea.AsymmetricSpinnaker,
		SymmetricSpinnakerArea:  baseSpec.SailArea.SymmetricSpinnaker,
		Displacement:            baseSpec.Dimension.Displacement,
		CrewWeight:              baseSpec.Dimension.CrewWeight,
		Mode:                    mode,
		Stabilization:           stabilization,
		Hull:                    hull,
		Composition: map[openfactor.MATERIAL]float64{
			openfactor.MATERIAL_DEFAULT: defaultComposition,
			openfactor.MATERIAL_BALLAST: extraSpec.Composition.BallastPercentage,
			openfactor.MATERIAL_CFK:     extraSpec.Composition.CfkPercentage,
			openfactor.MATERIAL_ALU:     extraSpec.Composition.AluPercentage,
			openfactor.MATERIAL_GFK:     extraSpec.Composition.GfkPercentage,
			openfactor.MATERIAL_WOOD:    extraSpec.Composition.WoodPercentage,
			openfactor.MATERIAL_ENGINE:  extraSpec.Composition.EnginePercentage,
			openfactor.MATERIAL_AMENITY: extraSpec.Composition.AmenityPercentage,
		},
	})
	if err != nil {
		return nil, err
	}

	return &output.ShipConfigRating{
		Version:             openfactor.Version(),
		TCC:                 factorOutput.TCC,
		SpeedFactor:         factorOutput.SpeedFactor,
		SpeedInfluence:      factorOutput.SpeedInfluence,
		SpeedDragPoints:     factorOutput.SpeedDragPoints,
		SpeedUpwindPoints:   factorOutput.SpeedUpwindPoints,
		SpeedDownwindPoints: factorOutput.SpeedDownwindPoints,

		StabilizationFactor:    factorOutput.StabilizationFactor,
		StabilizationInfluence: factorOutput.StabilizationInfluence,
		StabilizationPoints:    factorOutput.StabilizationPoints,

		AgilityFactor:    factorOutput.AgilityFactor,
		AgilityInfluence: factorOutput.AgilityInfluence,
		AgilityPoints:    factorOutput.AgilityPoints,
	}, nil
}
