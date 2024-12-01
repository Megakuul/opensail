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
	"io"
	"net/http"
	"net/url"
)

const ORC_BASE_ENDPOINT = "http://data.orc.org/public/WPub.dll"

// GetDownBoatRMS executes the DownBoatRMS action on the orc api, querying by sail number.
func GetDownBoatRMS(sailNo string) (*DownBoatRMS, error) {
	orcUrl, err := url.Parse(ORC_BASE_ENDPOINT)
	if err != nil {
		return nil, err
	}
	orcUrl.Query().Add("action", "DownBoatRMS")
	orcUrl.Query().Add("SailNo", sailNo)
	orcUrl.Query().Add("ext", "json")

	client := &http.Client{}
	resp, err := client.Get(orcUrl.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	downBoatRmsRaw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	downBoatRms := &DownBoatRMS{}
	err = json.Unmarshal(downBoatRmsRaw, downBoatRms)
	if err != nil {
		return nil, err
	}

	return downBoatRms, nil
}
