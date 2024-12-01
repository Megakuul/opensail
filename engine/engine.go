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

package main

import (
	"log"
	"os"

	"github.com/megakuul/opensail/engine/generate"
	"github.com/megakuul/opensail/engine/structure/input"
	"github.com/megakuul/opensail/engine/structure/output"
	"github.com/megakuul/opensail/engine/validate"
	"github.com/spf13/cobra"
)

func main() {
	cmd := NewEngineCmd(&input.RepoStructure{
		Team: input.TeamStructure{
			TeamBasePath:   "register/teams/",
			TeamConfigFile: "team.toml",
		},
		Ship: input.ShipStructure{
			ShipBasePath:   "register/ships/",
			ShipConfigFile: "ship.toml",
			ShipInfoFile:   "info.toml",
			ShipSpecFile:   "spec.toml",
		},
	}, &output.DataStructure{
		Team: output.TeamStructure{
			TeamMapFile: "teams.json",
		},
		Ship: output.ShipStructure{
			ShipMapFile: "ships.json",
		},
	})
	if err := cmd.Execute(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}

func NewEngineCmd(inputStruct *input.RepoStructure, outputStruct *output.DataStructure) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "engine",
		Short:         "engine: opensail ci operator",
		Version:       Version(),
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.AddCommand(generate.NewGenerateCmd())
	cmd.AddCommand(validate.NewValidateCmd(inputStruct, outputStruct))

	return cmd
}
