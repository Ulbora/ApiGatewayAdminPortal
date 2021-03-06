/*
 Copyright (C) 2017 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2017 Ken Williamson
 All rights reserved.

 Certain inventions and disclosures in this file may be claimed within
 patents owned or patent applications filed by Ulbora Labs LLC., or third
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
	"fmt"
	services "github.com/Ulbora/ApiGatewayAdminPortal/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func handleGrantType(w http.ResponseWriter, r *http.Request) {
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
			//fmt.Println(res)
			var page oauthPage
			page.OauthActive = "active"
			page.Client = res

			var g services.GrantTypeService
			g.ClientID = getAuthCodeClient()
			g.Host = getOauthHost()
			g.Token = token.AccessToken
			res2 := g.GetGrantTypeList(clientID)
			page.GrantTypes = res2
			var sm secSideMenu
			sm.GrantTypeActive = "active teal"
			page.SecSideMenu = &sm

			//fmt.Println(page)
			templates.ExecuteTemplate(w, "grantTypes.html", &page)
		}
	}
}

func handleGrantTypeAdd(w http.ResponseWriter, r *http.Request) {
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

		grantType := r.FormValue("grantType")
		fmt.Println(grantType)

		clientIDStr := r.FormValue("clientId")
		clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)
		fmt.Print("clientId: ")
		fmt.Println(clientID)

		session.Values["userLoggenIn"] = true

		token := getToken(w, r)

		var g services.GrantTypeService
		g.ClientID = getAuthCodeClient()
		g.Host = getOauthHost()
		g.Token = token.AccessToken
		if grantType != "" {
			var gg services.GrantType
			gg.ClientID = clientID
			gg.GrantType = grantType
			gres := g.AddGrantType(&gg)
			if gres.Success != true {
				fmt.Println(gres)
			}
		}
		http.Redirect(w, r, "/clientGrantTypes/"+clientIDStr, http.StatusFound)
	}
}

func handleGrantTypeDelete(w http.ResponseWriter, r *http.Request) {
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
		ID := vars["id"]
		clientID := vars["clientId"]

		if ID != "" && clientID != "" {
			token := getToken(w, r)
			var g services.GrantTypeService
			g.ClientID = getAuthCodeClient()
			g.Host = getOauthHost()
			g.Token = token.AccessToken
			gres := g.DeleteGrantType(ID)
			if gres.Success != true {
				fmt.Println(gres)
			}
		}
		http.Redirect(w, r, "/clientGrantTypes/"+clientID, http.StatusFound)
	}
}
