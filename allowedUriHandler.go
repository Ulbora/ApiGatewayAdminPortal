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

type allowedURIDisplay struct {
	ID           int64
	URI          string
	ClientID     int64
	AssignedRole int64
}

func handleAllowedUris(w http.ResponseWriter, r *http.Request) {
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

			var r services.ClientRoleService
			r.ClientID = getAuthCodeClient()
			r.Host = getOauthHost()
			r.Token = token.AccessToken
			rr := r.GetClientRoleList(clientID)
			page.ClientRoles = rr
			//fmt.Println(rr)
			rMap := make(map[int64]int64)
			var ru services.RoleURIService
			ru.ClientID = getAuthCodeClient()
			ru.Host = getOauthHost()
			ru.Token = token.AccessToken
			for _, rrr := range *rr {
				//fmt.Print("start")
				ruu := ru.GetRoleURIList(strconv.FormatInt(rrr.ID, 10))
				//fmt.Print("ruu: ")
				//fmt.Println(ruu)
				for _, rui := range *ruu {
					//fmt.Print("roleUrl: ")
					//fmt.Println(rui)
					//rMap[rui.ClientAllowedURIID] = append(rMap[rui.ClientAllowedURIID], rui.ClientRoleID)
					rMap[rui.ClientAllowedURIID] = rui.ClientRoleID
				}
			}

			//fmt.Println(rMap)
			var a services.AllowedURIService
			a.ClientID = getAuthCodeClient()
			a.Host = getOauthHost()
			a.Token = token.AccessToken
			ares := a.GetAllowedURIList(clientID)
			var newAres []allowedURIDisplay
			for _, ar := range *ares {
				var ard allowedURIDisplay
				ard.AssignedRole = rMap[ar.ID]
				ard.ClientID = ar.ClientID
				ard.ID = ar.ID
				ard.URI = ar.URI
				newAres = append(newAres, ard)
				//fmt.Print("role: ")
				//fmt.Println(ar)
			}
			//fmt.Print("roles: ")
			//fmt.Println(newAres)
			page.AllowedURIs = &newAres

			//fmt.Println(page)
			templates.ExecuteTemplate(w, "allowedUris.html", &page)
		}
	}
}

func handleAllowedUrisAdd(w http.ResponseWriter, r *http.Request) {
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
		uri := r.FormValue("uri")
		fmt.Println(uri)

		roleIDStr := r.FormValue("roleId")
		fmt.Println(roleIDStr)
		roleID, _ := strconv.ParseInt(roleIDStr, 10, 0)

		clientIDStr := r.FormValue("clientId")
		clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)
		fmt.Print("clientId: ")
		fmt.Println(clientID)
		if roleIDStr != "" && clientIDStr != "" {
			var au services.AllowedURIService
			au.ClientID = getAuthCodeClient()
			au.Host = getOauthHost()
			au.Token = token.AccessToken
			var auu services.AllowedURI
			auu.ClientID = clientID
			auu.URI = uri
			aures := au.AddAllowedURI(&auu)
			if aures.Success == true {
				var cr services.RoleURIService
				cr.ClientID = getAuthCodeClient()
				cr.Host = getOauthHost()
				cr.Token = token.AccessToken

				var crr services.RoleURI
				crr.ClientRoleID = roleID
				crr.ClientAllowedURIID = aures.ID
				crres := cr.AddRoleURI(&crr)
				if crres.Success != true {
					fmt.Println(crres)
				}
			}
			//fmt.Println(aures)
		}
		http.Redirect(w, r, "/clientAllowedUris/"+clientIDStr, http.StatusFound)
	}
}

func handleAllowedUrisUpdate(w http.ResponseWriter, r *http.Request) {
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

		IDStr := r.FormValue("id")
		fmt.Println(IDStr)
		ID, _ := strconv.ParseInt(IDStr, 10, 0)
		fmt.Println(ID)

		uri := r.FormValue("uri")
		fmt.Println(uri)

		roleIDStr := r.FormValue("roleId")
		fmt.Println(roleIDStr)
		roleID, _ := strconv.ParseInt(roleIDStr, 10, 0)
		fmt.Println(roleID)

		originalRoleIDStr := r.FormValue("originalRoleId")
		fmt.Println(originalRoleIDStr)
		originalRoleID, _ := strconv.ParseInt(originalRoleIDStr, 10, 0)
		fmt.Println(originalRoleID)
		var updateRole = false
		if roleID != originalRoleID {
			updateRole = true
		}
		fmt.Println(updateRole)

		clientIDStr := r.FormValue("clientId")
		clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)
		fmt.Print("clientId: ")
		fmt.Println(clientID)

		var au services.AllowedURIService
		au.ClientID = getAuthCodeClient()
		au.Host = getOauthHost()
		au.Token = token.AccessToken

		var auu services.AllowedURI
		auu.ID = ID
		auu.ClientID = clientID
		auu.URI = uri
		aures := au.UpdateAllowedURI(&auu)
		if aures.Success == true {
			if updateRole == true {
				var cr services.RoleURIService
				cr.ClientID = getAuthCodeClient()
				cr.Host = getOauthHost()
				cr.Token = token.AccessToken

				var crr services.RoleURI
				crr.ClientRoleID = originalRoleID
				crr.ClientAllowedURIID = ID
				cr.DeleteRoleURI(&crr)

				//var crr services.RoleURI
				crr.ClientRoleID = roleID
				crr.ClientAllowedURIID = ID
				crres := cr.AddRoleURI(&crr)
				if crres.Success != true {
					fmt.Println(crres)
				}
			}
		}
		fmt.Println(aures)
		http.Redirect(w, r, "/clientAllowedUris/"+clientIDStr, http.StatusFound)
	}
}

func handleAllowedUrisDelete(w http.ResponseWriter, r *http.Request) {
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

		IDStr := vars["id"]
		fmt.Println(IDStr)
		ID, _ := strconv.ParseInt(IDStr, 10, 0)
		fmt.Println(ID)

		//clientID := vars["clientId"]

		clientIDStr := vars["clientId"]
		clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)
		fmt.Print("clientId: ")
		fmt.Println(clientID)

		roleIDStr := vars["roleId"]
		fmt.Println(roleIDStr)
		roleID, _ := strconv.ParseInt(roleIDStr, 10, 0)
		fmt.Println(roleID)

		if IDStr != "" && clientIDStr != "" && roleIDStr != "" {
			token := getToken(w, r)

			var cr services.RoleURIService
			cr.ClientID = getAuthCodeClient()
			cr.Host = getOauthHost()
			cr.Token = token.AccessToken

			var crr services.RoleURI
			crr.ClientRoleID = roleID
			crr.ClientAllowedURIID = ID
			cr.DeleteRoleURI(&crr)

			var au services.AllowedURIService
			au.ClientID = getAuthCodeClient()
			au.Host = getOauthHost()
			au.Token = token.AccessToken

			aures := au.DeleteAllowedURI(IDStr)

			fmt.Println(aures)
			http.Redirect(w, r, "/clientAllowedUris/"+clientIDStr, http.StatusFound)
		}
	}
}
