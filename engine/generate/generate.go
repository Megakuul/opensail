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
	"os"
	"path"

	"github.com/megakuul/opensail/engine/structure/input"
	"github.com/megakuul/opensail/engine/structure/output"
	"github.com/spf13/cobra"
)

type generateFlags struct {
	inputPath  string
	outputPath string
}

func NewGenerateCmd(inputStruct *input.Structure, outputStruct *output.Structure) *cobra.Command {
	flags := &generateFlags{}

	cmd := &cobra.Command{
		Use:          "generate",
		Short:        "generate opensail api data from register",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return Run(flags, inputStruct, outputStruct)
		},
	}

	cmd.Flags().SortFlags = false
	cmd.Flags().StringVarP(&flags.inputPath, "input-path", "i",
		".", "specify the repository base path",
	)
	cmd.Flags().StringVarP(&flags.outputPath, "output-path", "o",
		"./out", "specify the data output path",
	)

	return cmd
}

func Run(flags *generateFlags, inputStruct *input.Structure, outputStruct *output.Structure) error {

	teamsDirectory, err := os.ReadDir(path.Join(flags.inputPath, inputStruct.Team.BasePath))
	if err != nil {
		return err
	}
	teams := map[string]struct{}{}
	for _, entry := range teamsDirectory {
		if entry.IsDir() {
			teams[entry.Name()] = struct{}{}
		}
	}

	shipsDirectory, err := os.ReadDir(path.Join(flags.inputPath, inputStruct.Ship.BasePath))
	if err != nil {
		return err
	}
	ships := map[string]struct{}{}
	for _, entry := range shipsDirectory {
		if entry.IsDir() {
			ships[entry.Name()] = struct{}{}
		}
	}

	manifestData, err := generateManifest()
	if err != nil {
		return err
	}
	err = os.WriteFile(outputStruct.Manifest.ConfigFile, manifestData, 0644)
	if err != nil {
		return err
	}

	teamsData, err := generateTeams(teams, inputStruct.Team)
	if err != nil {
		return err
	}
	err = os.WriteFile(outputStruct.Team.MapFile, teamsData, 0644)
	if err != nil {
		return err
	}

	shipData, err := generateShips(ships, inputStruct.Ship)
	if err != nil {
		return err
	}
	err = os.WriteFile(outputStruct.Ship.MapFile, shipData, 0644)
	if err != nil {
		return err
	}

	return nil
}
