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

// TeamConfig specifies the toml representation of the teams configuration.
type TeamConfig struct {
	// Name specifies the team name/identifier
	Name string `toml:"name" validate:"required,lowercase"`
	// Members contains the list of team members
	Members []TeamConfigMember `toml:"members"`
}

type TeamConfigMember struct {
	// Name specifies the member's full name
	Name string `toml:"name" validate:"required"`
	// Roles specifies the member's roles (skipper, trimmer, bowman)
	Roles []string `toml:"roles" validate:"dive,oneof=skipper trimmer bowman"`
}
