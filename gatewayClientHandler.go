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

package main

import (
	services "ApiGatewayAdminPortal/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func handleGatewayClient(w http.ResponseWriter, r *http.Request) {
	s.InitSessionStore(w, r)
	session, err := s.GetSession(r)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	token := getToken(w, r)
	fmt.Print("in main page. Logged in: ")
	fmt.Println(loggedIn)
	//fmt.Println(token.AccessToken)
	//var res *[]services.Client
	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
		authorize(w, r)
	} else {
		session.Values["userLoggenIn"] = true
		vars := mux.Vars(r)
		clientID := vars["clientId"]

		if clientID != "" {
			var c services.ClientService
			token := getToken(w, r)
			c.ClientID = getAuthCodeClient()
			c.Host = getOauthHost()
			c.Token = token.AccessToken

			res := c.GetClient(clientID)

			var g services.GatewayClientService
			//token := getToken(w, r)
			g.ClientID = getAuthCodeClient()
			g.Host = getGatewayHost()
			g.Token = token.AccessToken

			gres := g.GetClient(clientID)
			fmt.Println(gres)
			var page gwPage
			page.GwActive = "active"
			page.Client = res
			page.GatewayClient = gres
			var sm gwSideMenu
			sm.GWClientActive = "active teal"
			page.GwSideMenu = &sm
			//fmt.Println(page)
			templates.ExecuteTemplate(w, "gatewayClient.html", &page)
		}
	}
}

func handleAddGatewayClient(w http.ResponseWriter, r *http.Request) {
	s.InitSessionStore(w, r)
	session, err := s.GetSession(r)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	token := getToken(w, r)
	fmt.Print("in main page. Logged in: ")
	fmt.Println(loggedIn)
	//fmt.Println(token.AccessToken)
	//var res *[]services.Client
	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
		authorize(w, r)
	} else {
		clientIDStr := r.FormValue("clientId")
		clientID, errID := strconv.ParseInt(clientIDStr, 10, 0)
		if errID != nil {
			fmt.Print(errID)
		}
		fmt.Print("clientId: ")
		fmt.Println(clientID)

		level := r.FormValue("level")
		fmt.Print("level: ")
		fmt.Println(level)

		var g services.GatewayClientService
		g.ClientID = getAuthCodeClient()
		g.Host = getGatewayHost()
		g.Token = token.AccessToken

		var gc services.GatewayClient
		gc.ClientID = clientID
		gc.Enabled = true
		gc.Level = level
		gc.APIKey = generateAPIKey()
		res := g.AddClient(&gc)
		if res.Success != true {
			fmt.Println("Error adding Gateway Client")
		}
		http.Redirect(w, r, "/gateway/"+clientIDStr, http.StatusFound)
	}
}

func handleUpdateGatewayClient(w http.ResponseWriter, r *http.Request) {
	s.InitSessionStore(w, r)
	session, err := s.GetSession(r)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	loggedIn := session.Values["userLoggenIn"]
	token := getToken(w, r)
	fmt.Print("in main page. Logged in: ")
	fmt.Println(loggedIn)
	//fmt.Println(token.AccessToken)
	//var res *[]services.Client
	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
		authorize(w, r)
	} else {
		clientIDStr := r.FormValue("clientId")
		clientID, errID := strconv.ParseInt(clientIDStr, 10, 0)
		if errID != nil {
			fmt.Print(errID)
		}
		fmt.Print("clientId: ")
		fmt.Println(clientID)

		level := r.FormValue("level")
		fmt.Print("level: ")
		fmt.Println(level)

		apiKey := r.FormValue("apiKey")
		fmt.Print("apiKey: ")
		fmt.Println(apiKey)

		enabled := r.FormValue("enabled")
		fmt.Print("enabled: ")
		fmt.Println(enabled)

		var g services.GatewayClientService
		g.ClientID = getAuthCodeClient()
		g.Host = getGatewayHost()
		g.Token = token.AccessToken

		var gc services.GatewayClient
		gc.ClientID = clientID
		if enabled == "yes" {
			gc.Enabled = true
		} else {
			gc.Enabled = false
		}
		gc.Level = level
		gc.APIKey = apiKey

		res := g.UpdateClient(&gc)
		if res.Success != true {
			fmt.Println("Error adding Gateway Client")
		}
		http.Redirect(w, r, "/gateway/"+clientIDStr, http.StatusFound)
	}
}
