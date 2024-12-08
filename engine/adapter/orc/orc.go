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

package orc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const ORC_BASE_ENDPOINT = "https://data.orc.org/public/WPub.dll"

var CERT_FAMILIES = map[string]struct{}{
	"ORC": {},
	"DH":  {},
	"NS":  {},
}

// GetDownBoatRMS executes the DownBoatRMS action on the orc api, querying by orc cert ref number.
func GetDownBoatRMS(refNo string) (*DownBoatRMS, error) {
	orcQuery := url.Values{}
	orcQuery.Add("action", "DownBoatRMS")
	orcQuery.Add("RefNo", refNo)
	orcQuery.Add("ext", "json")

	client := &http.Client{}
	resp, err := client.Get(fmt.Sprintf("%s?%s", ORC_BASE_ENDPOINT, orcQuery.Encode()))
	if err != nil {
		return nil, fmt.Errorf("fetching orc data failed: %w", err)
	}
	defer resp.Body.Close()

	downBoatRmsRaw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading orc data failed: %w", err)
	}

	// Remove retarded BOM header that the orc api is using for whatever reason.
	downBoatRmsRaw = bytes.TrimPrefix(downBoatRmsRaw, []byte("\xef\xbb\xbf"))

	downBoatRms := &DownBoatRMS{}
	err = json.Unmarshal(downBoatRmsRaw, downBoatRms)
	if err != nil {
		return nil, fmt.Errorf("parsing orc data failed: %w", err)
	}

	return downBoatRms, nil
}
