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

import "github.com/spf13/cobra"

type generateFlags struct {
	inputPath      string
	outputPath     string
	orcApiEndpoint string
}

func NewGenerateCmd() *cobra.Command {
	flags := &generateFlags{}

	cmd := &cobra.Command{
		Use:          "generate",
		Short:        "generate opensail api data from register",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return Run(flags)
		},
	}

	cmd.Flags().SortFlags = false
	cmd.Flags().StringVarP(&flags.inputPath, "input-path", "i",
		".", "specify the repository base path",
	)
	cmd.Flags().StringVarP(&flags.outputPath, "output-path", "o",
		"./out", "specify the data output path",
	)
	cmd.Flags().StringVar(&flags.orcApiEndpoint, "orc-api-endpoint",
		"https://data.orc.org/public/WPub.dll", "specify the base url for the orc api",
	)

	return cmd
}

func Run(flags *generateFlags) error {
	return nil
}
