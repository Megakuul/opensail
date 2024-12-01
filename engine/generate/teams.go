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

	"github.com/BurntSushi/toml"
	"github.com/megakuul/opensail/engine/structure/input"
	"github.com/megakuul/opensail/engine/structure/output"
)

// generateTeams generates the teamMap.
func generateTeams(repoPath string, teams map[string]struct{}, teamStruct input.TeamStructure) ([]byte, error) {
	teamMap := output.TeamMap{}

	for team := range teams {
		teamPath := path.Join(repoPath, teamStruct.BasePath, team)
		teamConfigRaw, err := os.ReadFile(path.Join(teamPath, teamStruct.ConfigFile))
		if err != nil {
			return nil, fmt.Errorf("failed to read team config (team '%s'): %w", team, err)
		}
		teamConfig := &input.TeamConfig{}
		err = toml.Unmarshal(teamConfigRaw, teamConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to parse team config (team '%s'): %w", team, err)
		}

		outputTeamConfig := output.TeamConfig{
			Name: teamConfig.Name,
		}

		for _, member := range teamConfig.Members {
			outputTeamConfig.Members = append(outputTeamConfig.Members, output.TeamConfigMember{
				Name:  member.Name,
				Roles: member.Roles,
			})
		}

		teamMap[team] = outputTeamConfig
	}

	teamMapRaw, err := json.Marshal(teamMap)
	if err != nil {
		return nil, err
	}

	return teamMapRaw, nil
}
