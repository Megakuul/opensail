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

// validateTeams performs checks and validations on updated team register entries.
func validateTeams(repoPath string, updatedTeams map[string]struct{}, teamStruct input.TeamStructure) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	teamsPath := path.Join(repoPath, teamStruct.TeamBasePath)
	teamsPathInfo, err := os.Stat(teamsPath)
	if err != nil {
		return err
	}
	if !teamsPathInfo.IsDir() {
		return fmt.Errorf("expected team directory at: %s", teamsPath)
	}

	for team := range updatedTeams {
		teamPath := path.Join(teamsPath, team)

		teamConfigRaw, err := os.ReadFile(path.Join(teamPath, "team.toml"))
		if err != nil {
			return err
		}

		teamConfig := &input.TeamConfig{}
		err = toml.Unmarshal(teamConfigRaw, teamConfig)
		if err != nil {
			return err
		}
		err = validate.Struct(teamConfig)
		if err != nil {
			return err
		}
	}

	return nil
}
