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

package input

// RepoStructure holds metainformation used to find the register
// files and directories inside the repository.
type RepoStructure struct {
	Team TeamStructure
	Ship ShipStructure
}

type TeamStructure struct {
	TeamBasePath   string
	TeamConfigFile string
}

type ShipStructure struct {
	ShipBasePath   string
	ShipConfigFile string
	ShipInfoFile   string
	ShipSpecFile   string
}
