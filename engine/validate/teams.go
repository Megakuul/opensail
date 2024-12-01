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
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/go-playground/validator/v10"
	"github.com/megakuul/opensail/engine/structure/input"
)

// validateTeams performs checks and validations on updated team register entries.
func validateTeams(repoPath string, teams map[string]struct{}, teamStruct input.TeamStructure) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	teamsPath := path.Join(repoPath, teamStruct.BasePath)
	teamsPathInfo, err := os.Stat(teamsPath)
	if err != nil {
		return err
	}
	if !teamsPathInfo.IsDir() {
		return fmt.Errorf("expected team directory at: %s", teamsPath)
	}

	for team := range teams {
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

		err = validateTeamMembers(teamConfig.Members)
		if err != nil {
			return err
		}
	}

	return nil
}

func validateTeamMembers(members []input.TeamConfigMember) error {
	if len(members) < 1 {
		return fmt.Errorf("expected at least 1 team member")
	}
	for _, member := range members {
		if member.Name == "" {
			return fmt.Errorf("invalid team member name: %s", member.Name)
		}

		for _, role := range member.Roles {
			if _, ok := input.TEAM_MEMBER_ROLES[strings.ToLower(role)]; !ok {
				return fmt.Errorf("invalid team member role: '%s'; expected one of %v", role, input.TEAM_MEMBER_ROLES)
			}
		}
	}
	return nil
}
