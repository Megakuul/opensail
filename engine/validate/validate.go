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
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v67/github"
	"github.com/megakuul/opensail/engine/structure/input"
	"github.com/megakuul/opensail/engine/structure/output"
	"github.com/spf13/cobra"
)

type validateFlags struct {
	inputPath      string
	githubOwner    string
	githubRepo     string
	githubPrNumber int
	githubToken    string
}

func NewValidateCmd(inputStruct *input.Structure, outputStruct *output.Structure) *cobra.Command {
	flags := &validateFlags{}

	cmd := &cobra.Command{
		Use:          "validate",
		Short:        "validate pull requests to the opensail register",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return Run(flags, inputStruct, outputStruct)
		},
	}

	cmd.Flags().SortFlags = false
	cmd.Flags().StringVarP(&flags.inputPath, "input-path", "i",
		".", "specify the repository base path",
	)
	cmd.Flags().StringVar(&flags.githubOwner, "github-owner",
		"", "specify the github repo owner",
	)
	cmd.Flags().StringVar(&flags.githubRepo, "github-repo",
		"", "specify the github repo name",
	)
	cmd.Flags().IntVar(&flags.githubPrNumber, "github-pr-number",
		0, "specify the github pull request number",
	)
	cmd.Flags().StringVar(&flags.githubToken, "github-token",
		"", "specify the github api token",
	)

	return cmd
}

func Run(flags *validateFlags, inputStruct *input.Structure, outputStruct *output.Structure) error {
	client := github.NewClient(nil)
	files, _, err := client.PullRequests.ListFiles(
		context.TODO(),
		flags.githubOwner,
		flags.githubRepo,
		flags.githubPrNumber,
		&github.ListOptions{},
	)
	if err != nil {
		return err
	}

	updatedTeams, updatedShips := map[string]struct{}{}, map[string]struct{}{}
	for _, file := range files {
		if team := findComponentEntry(file.GetFilename(), inputStruct.Team.BasePath); team != "" {
			updatedTeams[team] = struct{}{}
		}
		if ship := findComponentEntry(file.GetFilename(), inputStruct.Ship.BasePath); ship != "" {
			updatedShips[ship] = struct{}{}
		}
	}

	err = validateTeams(flags.inputPath, updatedTeams, inputStruct.Team)
	if err != nil {
		return fmt.Errorf("failure while validating teams: %w", err)
	}

	err = validateShips(flags.inputPath, updatedShips, inputStruct.Ship)
	if err != nil {
		return fmt.Errorf("failure while validating ships: %w", err)
	}

	return nil
}

// findComponent returns the name of the component entry of the file.
// e.g. 'register/teams/example/some/file.toml' & 'register/teams/' returns -> 'example'
func findComponentEntry(filePath, componentPath string) string {
	if strings.HasPrefix(filePath, componentPath) {
		teamPrefix := strings.TrimPrefix(filePath, componentPath)
		fileSegments := strings.Split(teamPrefix, "/")
		if len(fileSegments) > 0 {
			return fileSegments[0]
		}
	}
	return ""
}
