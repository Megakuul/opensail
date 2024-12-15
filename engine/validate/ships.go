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
	"regexp"

	"github.com/BurntSushi/toml"
	"github.com/go-playground/validator/v10"
	"github.com/megakuul/opensail/engine/adapter/orc"
	"github.com/megakuul/opensail/engine/structure/input"
)

// validateShips performs checks and validations on updated ship register entries.
func validateShips(repoPath string, ships map[string]struct{}, shipStruct input.ShipStructure) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	shipsPath := path.Join(repoPath, shipStruct.BasePath)
	shipsPathInfo, err := os.Stat(shipsPath)
	if err != nil {
		return err
	}
	if !shipsPathInfo.IsDir() {
		return fmt.Errorf("expected ship directory at: %s", shipsPath)
	}

	for ship := range ships {
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

		ok, err := regexp.MatchString("^[a-z]{3}_[a-z]{1,10}$", ship)
		if err != nil || !ok {
			return fmt.Errorf("ship identifier does not match the required format (e.g., 'sui_example')")
		}

		err = validateShipInfo(shipConfig.Info, path.Join(shipPath, shipStruct.InfoFile))
		if err != nil {
			return err
		}

		err = validateShipBaseSpec(shipConfig.BaseSpec, path.Join(shipPath, shipStruct.BaseSpecFile))
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
		if info.ORCRefNo == "" {
			return fmt.Errorf("invalid ship orc RefNo. '%s'", info.ORCRefNo)
		}
		downBoatRms, err := orc.GetDownBoatRMS(info.ORCRefNo)
		if err != nil {
			return err
		}

		if len(downBoatRms.Rms) < 1 {
			return fmt.Errorf("ship with RefNo. '%s' was not found on orc database", info.ORCRefNo)
		}
	default:
		return fmt.Errorf("invalid ship info source '%s'", info.Source)
	}
	return nil
}

func validateShipBaseSpec(spec input.ShipConfigBaseSpec, specPath string) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	switch spec.Source {
	case input.SHIP_BASE_SPEC_MANUAL:
		shipSpecRaw, err := os.ReadFile(specPath)
		if err != nil {
			return err
		}
		shipSpec := &input.ShipBaseSpec{}
		err = toml.Unmarshal(shipSpecRaw, shipSpec)
		if err != nil {
			return err
		}
		err = validate.Struct(shipSpec)
		if err != nil {
			return err
		}
	case input.SHIP_BASE_SPEC_ORC:
		if spec.ORCRefNo == "" {
			return fmt.Errorf("invalid ship orc RefNo. '%s'", spec.ORCRefNo)
		}
		downBoatRms, err := orc.GetDownBoatRMS(spec.ORCRefNo)
		if err != nil {
			return err
		}

		if len(downBoatRms.Rms) < 1 {
			return fmt.Errorf("ship with RefNo. '%s' was not found on orc database", spec.ORCRefNo)
		}
	default:
		return fmt.Errorf("invalid ship base spec source '%s'", spec.Source)
	}
	return nil
}

func validateShipExtraSpec(spec input.ShipConfigExtraSpec, specPath string) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	switch spec.Source {
	case input.SHIP_EXTRA_SPEC_MANUAL:
		shipSpecRaw, err := os.ReadFile(specPath)
		if err != nil {
			return err
		}
		shipSpec := &input.ShipExtraSpec{}
		err = toml.Unmarshal(shipSpecRaw, shipSpec)
		if err != nil {
			return err
		}
		err = validate.Struct(shipSpec)
		if err != nil {
			return err
		}

		switch shipSpec.Design.Mode {
		case input.SHIP_EXTRA_SPEC_DESIGN_DISPLACE:
		case input.SHIP_EXTRA_SPEC_DESIGN_SEMI:
		case input.SHIP_EXTRA_SPEC_DESIGN_PLANING:
		case input.SHIP_EXTRA_SPEC_DESIGN_HYDROFOIL:
			break
		default:
			return fmt.Errorf("invalid ship extra spec mode '%s'", shipSpec.Design.Mode)
		}

		switch shipSpec.Design.Stabilization {
		case input.SHIP_EXTRA_SPEC_DESIGN_FOILS:
		case input.SHIP_EXTRA_SPEC_DESIGN_CENTREBOARD:
		case input.SHIP_EXTRA_SPEC_DESIGN_DAGGERBOARD:
		case input.SHIP_EXTRA_SPEC_DESIGN_KEEL:
			break
		default:
			return fmt.Errorf("invalid ship extra spec stabilization '%s'", shipSpec.Design.Stabilization)
		}

		switch shipSpec.Design.Hull {
		case input.SHIP_EXTRA_SPEC_DESIGN_MONO:
		case input.SHIP_EXTRA_SPEC_DESIGN_MULTI:
			break
		default:
			return fmt.Errorf("invalid ship extra spec hull '%s'", shipSpec.Design.Hull)
		}
	default:
		return fmt.Errorf("invalid ship extra spec source '%s'", spec.Source)
	}
	return nil
}
