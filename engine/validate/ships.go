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
	"github.com/megakuul/opensail/engine/structure/input"
)

// validateShips performs checks and validations on updated ship register entries.
func validateShips(repoPath string, updatedShips map[string]struct{}, shipStruct input.ShipStructure) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	shipsPath := path.Join(repoPath, shipStruct.ShipBasePath)
	shipsPathInfo, err := os.Stat(shipsPath)
	if err != nil {
		return err
	}
	if !shipsPathInfo.IsDir() {
		return fmt.Errorf("expected ship directory at: %s", shipsPath)
	}

	for ship := range updatedShips {
		shipPath := path.Join(shipsPath, ship)

		shipConfigRaw, err := os.ReadFile(path.Join(shipPath, shipStruct.ShipConfigFile))
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
	}

	return nil
}
