/*
 Copyright (C) 2017 Ulbora Labs Inc. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2017 Ken Williamson
 All rights reserved.

 Certain inventions and disclosures in this file may be claimed within
 patents owned or patent applications filed by Ulbora Labs Inc., or third
 parties.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as published
 by the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.

 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//ClientService service
type ClientService struct {
	Token    string
	ClientID string
	APIKey   string
	UserID   string
	Hashed   string
	Host     string
}

//Client client
type Client struct {
	ClientID     int64         `json:"clientId"`
	Secret       string        `json:"secret"`
	Name         string        `json:"name"`
	WebSite      string        `json:"webSite"`
	Email        string        `json:"email"`
	Enabled      bool          `json:"enabled"`
	RedirectURIs []RedirectURI `json:"redirectUrls"`
}

//ClientResponse resp
type ClientResponse struct {
	Success bool  `json:"success"`
	ID      int64 `json:"id"`
	Code    int   `json:"code"`
}

//AddClient add template
func (c *ClientService) AddClient(client *Client) *ClientResponse {
	var rtn = new(ClientResponse)
	var addURL = c.Host + "/rs/client/add"
	aJSON, err := json.Marshal(client)

	if err != nil {
		fmt.Println(err)
	} else {
		req, rErr := http.NewRequest("POST", addURL, bytes.NewBuffer(aJSON))
		if rErr != nil {
			fmt.Print("request err: ")
			fmt.Println(rErr)
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+c.Token)
			req.Header.Set("clientId", c.ClientID)
			//req.Header.Set("userId", c.UserID)
			//req.Header.Set("hashed", c.Hashed)
			req.Header.Set("apiKey", c.APIKey)
			client := &http.Client{}
			resp, cErr := client.Do(req)
			if cErr != nil {
				fmt.Print("Template Add err: ")
				fmt.Println(cErr)
			} else {
				defer resp.Body.Close()
				decoder := json.NewDecoder(resp.Body)
				error := decoder.Decode(&rtn)
				if error != nil {
					log.Println(error.Error())
				}
				rtn.Code = resp.StatusCode
			}
		}
	}
	return rtn
}
