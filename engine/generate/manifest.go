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
	"time"

	"github.com/megakuul/opensail/engine/structure/output"
	"github.com/megakuul/opensail/engine/version"
)

// generateManifest generates the manifest file.
func generateManifest() ([]byte, error) {
	manifestConfig := &output.ManifestConfig{
		EngineVersion: version.Version(),
		Timestamp:     time.Now().Unix(),
	}

	manifestConfigRaw, err := json.Marshal(manifestConfig)
	if err != nil {
		return nil, err
	}

	return manifestConfigRaw, nil
}
