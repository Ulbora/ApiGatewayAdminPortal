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

type ulboraAllowedURI struct {
	URI    string
	RoleID int64
}

func handleUlboraUris(w http.ResponseWriter, r *http.Request) {
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
			// //fmt.Println(res)
			var page oauthPage
			page.OauthActive = "active"
			page.Client = res

			// var r services.ClientRoleService
			// r.ClientID = getAuthCodeClient()
			// r.Host = getOauthHost()
			// r.Token = token.AccessToken
			// rr := r.GetClientRoleList(clientID)
			// page.ClientRoles = rr
			// //fmt.Println(rr)
			// rMap := make(map[int64]int64)
			// var ru services.RoleURIService
			// ru.ClientID = getAuthCodeClient()
			// ru.Host = getOauthHost()
			// ru.Token = token.AccessToken
			// for _, rrr := range *rr {
			// 	//fmt.Print("start")
			// 	ruu := ru.GetRoleURIList(strconv.FormatInt(rrr.ID, 10))
			// 	//fmt.Print("ruu: ")
			// 	//fmt.Println(ruu)
			// 	for _, rui := range *ruu {
			// 		//fmt.Print("roleUrl: ")
			// 		//fmt.Println(rui)
			// 		//rMap[rui.ClientAllowedURIID] = append(rMap[rui.ClientAllowedURIID], rui.ClientRoleID)
			// 		rMap[rui.ClientAllowedURIID] = rui.ClientRoleID
			// 	}
			// }

			// //fmt.Println(rMap)
			// var a services.AllowedURIService
			// a.ClientID = getAuthCodeClient()
			// a.Host = getOauthHost()
			// a.Token = token.AccessToken
			// ares := a.GetAllowedURIList(clientID)
			// var newAres []allowedURIDisplay
			// for _, ar := range *ares {
			// 	var ard allowedURIDisplay
			// 	ard.AssignedRole = rMap[ar.ID]
			// 	ard.ClientID = ar.ClientID
			// 	ard.ID = ar.ID
			// 	ard.URI = ar.URI
			// 	newAres = append(newAres, ard)
			// 	//fmt.Print("role: ")
			// 	//fmt.Println(ar)
			// }
			//fmt.Print("roles: ")
			//fmt.Println(newAres)
			//page.AllowedURIs = &newAres
			var sm secSideMenu
			sm.UlboraURIsActive = "active"
			page.SecSideMenu = &sm

			//fmt.Println(page)
			templates.ExecuteTemplate(w, "ulboraUris.html", &page)
		}
	}
}

func handleUlboraUrisAdd(w http.ResponseWriter, r *http.Request) {
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

		oauth2box := r.FormValue("oauth2")
		fmt.Print("oauth2: ")
		fmt.Println(oauth2box)
		if oauth2box == "on" {
			fmt.Println("oauth2 is on")
		}

		apiGateway := r.FormValue("apiGateway")
		fmt.Print("apiGateway: ")
		fmt.Println(apiGateway)

		content := r.FormValue("content")
		fmt.Print("content: ")
		fmt.Println(content)

		customer := r.FormValue("customer")
		fmt.Print("customer: ")
		fmt.Println(customer)

		image := r.FormValue("image")
		fmt.Print("image: ")
		fmt.Println(image)

		mail := r.FormValue("mail")
		fmt.Print("mail: ")
		fmt.Println(mail)

		order := r.FormValue("order")
		fmt.Print("order: ")
		fmt.Println(order)

		product := r.FormValue("product")
		fmt.Print("product: ")
		fmt.Println(product)

		template := r.FormValue("template")
		fmt.Print("template: ")
		fmt.Println(template)

		clientIDStr := r.FormValue("clientId")
		clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)
		fmt.Print("clientId: ")
		fmt.Println(clientID)

		// if roleIDStr != "" && clientIDStr != "" {
		// 	var au services.AllowedURIService
		// 	au.ClientID = getAuthCodeClient()
		// 	au.Host = getOauthHost()
		// 	au.Token = token.AccessToken
		// 	var auu services.AllowedURI
		// 	auu.ClientID = clientID
		// 	auu.URI = uri
		// 	aures := au.AddAllowedURI(&auu)
		// 	if aures.Success == true {
		// 		var cr services.RoleURIService
		// 		cr.ClientID = getAuthCodeClient()
		// 		cr.Host = getOauthHost()
		// 		cr.Token = token.AccessToken

		// 		var crr services.RoleURI
		// 		crr.ClientRoleID = roleID
		// 		crr.ClientAllowedURIID = aures.ID
		// 		crres := cr.AddRoleURI(&crr)
		// 		if crres.Success != true {
		// 			fmt.Println(crres)
		// 		}
		// 	}
		// 	//fmt.Println(aures)
		// }
		http.Redirect(w, r, "/clientAllowedUris/"+clientIDStr, http.StatusFound)
	}
}
