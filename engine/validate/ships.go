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

package validate

import (
	"fmt"
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/go-playground/validator/v10"
	"github.com/megakuul/opensail/engine/adapter/orc"
	"github.com/megakuul/opensail/engine/structure/input"
)

// validateShips performs checks and validations on updated ship register entries.
func validateShips(repoPath string, updatedShips map[string]struct{}, shipStruct input.ShipStructure) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	shipsPath := path.Join(repoPath, shipStruct.BasePath)
	shipsPathInfo, err := os.Stat(shipsPath)
	if err != nil {
		return err
	}
	if !shipsPathInfo.IsDir() {
		return fmt.Errorf("expected ship directory at: %s", shipsPath)
	}

	for ship := range updatedShips {
		shipPath := path.Join(shipsPath, ship)

		shipConfigRaw, err := os.ReadFile(path.Join(shipPath, shipStruct.ConfigFile))
		if err != nil {
			return err
		}

		shipConfig := &input.ShipConfig{}
		err = toml.Unmarshal(shipConfigRaw, shipConfig)
		if err != nil {
			return err
		}
		err = validate.Struct(shipConfig)
		if err != nil {
			return err
		}

		err = validateShipInfo(shipConfig.Info, path.Join(shipPath, shipStruct.InfoFile))
		if err != nil {
			return err
		}

		err = validateShipSpec(shipConfig.Spec, path.Join(shipPath, shipStruct.SpecFile))
		if err != nil {
			return err
		}
	}

	return nil
}

func validateShipInfo(info input.ShipConfigInfo, infoPath string) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	switch info.Source {
	case input.SHIP_INFO_MANUAL:
		shipInfoRaw, err := os.ReadFile(infoPath)
		if err != nil {
			return err
		}
		shipInfo := &input.ShipInfo{}
		err = toml.Unmarshal(shipInfoRaw, shipInfo)
		if err != nil {
			return err
		}
		err = validate.Struct(shipInfo)
		if err != nil {
			return err
		}
	case input.SHIP_INFO_ORC:
		downBoatRms, err := orc.GetDownBoatRMS(info.ORCSailingNo)
		if err != nil {
			return err
		}

		if len(downBoatRms.Rms) < 1 {
			return fmt.Errorf("ship with SailNo. '%s' was not found on orc database", info.ORCSailingNo)
		}
	default:
		return fmt.Errorf("invalid ship info source '%s'", info.Source)
	}
	return nil
}

func validateShipSpec(spec input.ShipConfigSpec, specPath string) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	switch spec.Source {
	case input.SHIP_SPEC_MANUAL:
		shipSpecRaw, err := os.ReadFile(specPath)
		if err != nil {
			return err
		}
		shipSpec := &input.ShipSpec{}
		err = toml.Unmarshal(shipSpecRaw, shipSpec)
		if err != nil {
			return err
		}
		err = validate.Struct(shipSpec)
		if err != nil {
			return err
		}
	case input.SHIP_SPEC_ORC:
		downBoatRms, err := orc.GetDownBoatRMS(spec.ORCSailingNo)
		if err != nil {
			return err
		}

		if len(downBoatRms.Rms) < 1 {
			return fmt.Errorf("ship with SailNo. '%s' was not found on orc database", spec.ORCSailingNo)
		}
	default:
		return fmt.Errorf("invalid ship spec source '%s'", spec.Source)
	}
	return nil
}
