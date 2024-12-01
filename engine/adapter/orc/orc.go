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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const ORC_BASE_ENDPOINT = "http://data.orc.org/public/WPub.dll"

var CERT_FAMILIES = map[string]struct{}{
	"ORC": {},
	"DH":  {},
	"NS":  {},
}

// GetDownBoatRMS executes the DownBoatRMS action on the orc api, querying by sail number.
func GetDownBoatRMS(sailNo, certFamily string) (*DownBoatRMS, error) {
	if _, ok := CERT_FAMILIES[strings.ToUpper(certFamily)]; !ok {
		return nil, fmt.Errorf(
			"preparing orc data fetch failed: invalid certificate family: '%s'; expected one of %v",
			certFamily, CERT_FAMILIES,
		)
	}

	orcUrl, _ := url.Parse(ORC_BASE_ENDPOINT)
	orcUrl.Query().Add("action", "DownBoatRMS")
	orcUrl.Query().Add("SailNo", sailNo)
	orcUrl.Query().Add("Family", strings.ToUpper(certFamily))
	orcUrl.Query().Add("ext", "json")

	client := &http.Client{}
	resp, err := client.Get(orcUrl.String())
	if err != nil {
		return nil, fmt.Errorf("fetching orc data failed: %w", err)
	}
	defer resp.Body.Close()

	downBoatRmsRaw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading orc data failed: %w", err)
	}

	println(string(downBoatRmsRaw))

	downBoatRms := &DownBoatRMS{}
	err = json.Unmarshal(downBoatRmsRaw, downBoatRms)
	if err != nil {
		return nil, fmt.Errorf("parsing orc data failed: %w", err)
	}

	return downBoatRms, nil
}
