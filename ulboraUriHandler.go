package main

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

import (
	"fmt"
	services "github.com/Ulbora/ApiGatewayAdminPortal/services"
	"github.com/Ulbora/ApiGatewayAdminPortal/ulborauris"
	"net/http"
	"strconv"
	"sync"

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
			sm.UlboraURIsActive = "active teal"
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
		var uAdded = false

		var usel ulborauris.UlboraSelection

		oauth2 := r.FormValue("oauth2")
		if oauth2 == "superUser" {
			usel.Oauth2Super = true
			uAdded = true
		} else if oauth2 == "admin" {
			usel.Oauth2 = true
			uAdded = true
		}
		fmt.Print("oauth2: ")
		fmt.Println(oauth2)

		apiGateway := r.FormValue("apiGateway")
		if apiGateway == "superUser" {
			usel.APIGatewaySuper = true
			uAdded = true
		} else if apiGateway == "admin" {
			usel.APIGateway = true
			uAdded = true
		}
		fmt.Print("apiGateway: ")
		fmt.Println(apiGateway)

		content := r.FormValue("content")
		if content == "on" {
			usel.Content = true
			uAdded = true
		}
		fmt.Print("content: ")
		fmt.Println(content)

		customer := r.FormValue("customer")
		if customer == "admin" {
			usel.CustomerAdmin = true
			uAdded = true
		} else if customer == "user" {
			usel.CustomerUser = true
			uAdded = true
		}
		fmt.Print("customer: ")
		fmt.Println(customer)

		image := r.FormValue("image")
		if image == "admin" {
			usel.ImageAdmin = true
			uAdded = true
		} else if image == "user" {
			usel.ImageUser = true
			uAdded = true
		}
		fmt.Print("image: ")
		fmt.Println(image)

		mail := r.FormValue("mail")
		if mail == "on" {
			usel.Mail = true
			uAdded = true
		}
		fmt.Print("mail: ")
		fmt.Println(mail)

		order := r.FormValue("order")
		if order == "admin" {
			usel.OrderAdmin = true
			uAdded = true
		} else if order == "user" {
			usel.OrderUser = true
			uAdded = true
		}
		fmt.Print("order: ")
		fmt.Println(order)

		product := r.FormValue("product")
		if product == "admin" {
			usel.ProductAdmin = true
			uAdded = true
		} else if product == "user" {
			usel.ProductUser = true
			uAdded = true
		}
		fmt.Print("product: ")
		fmt.Println(product)

		template := r.FormValue("template")
		if template == "on" {
			usel.Template = true
			uAdded = true
		}
		fmt.Print("template: ")
		fmt.Println(template)

		clientIDStr := r.FormValue("clientId")
		clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)
		fmt.Print("clientId: ")
		fmt.Println(clientID)

		if uAdded == true {

			uuris := ulborauris.GetUlboraURIs(&usel)
			fmt.Println(uuris)

			var rs services.ClientRoleService
			rs.ClientID = getAuthCodeClient()
			rs.Host = getOauthHost()
			rs.Token = token.AccessToken
			rr := rs.GetClientRoleList(clientIDStr)

			rMap := make(map[string]int64)

			for _, rrr := range *rr {
				rMap[rrr.Role] = rrr.ID
			}
			fmt.Println(rMap)

			var au services.AllowedURIService
			au.ClientID = getAuthCodeClient()
			au.Host = getOauthHost()
			au.Token = token.AccessToken
			var insCnt = 0
			var tmpUris = make([]ulborauris.UlborURIs, 0)
			for _, uuri := range *uuris {
				tmpUris = append(tmpUris, uuri)
			}
			var wg sync.WaitGroup
			for i := range tmpUris {
				wg.Add(1)
				//fmt.Println(tmpUris[i].URI)
				//var tu = tmpUris[i].URI
				go func(val ulborauris.UlborURIs) {
					//fmt.Print("in thread: ")
					defer wg.Done()
					//fmt.Println(val.URI)
					insCnt++
					if rMap[val.Role] != 0 {
						var auu services.AllowedURI
						auu.ClientID = clientID
						auu.URI = val.URI
						aures := au.AddAllowedURI(&auu)
						if aures.Success != true {
							fmt.Print("error inserting record:")
							fmt.Println(val.URI)
							fmt.Println(aures)
						} else {
							var crr services.RoleURI
							crr.ClientRoleID = rMap[val.Role]
							crr.ClientAllowedURIID = aures.ID
							var cr services.RoleURIService
							cr.ClientID = getAuthCodeClient()
							cr.Host = getOauthHost()
							cr.Token = token.AccessToken

							crres := cr.AddRoleURI(&crr)
							if crres.Success != true {
								fmt.Println(crres)
							}
						}
					}
				}(tmpUris[i])
			}
			wg.Wait()

			fmt.Print("cnt: ")
			fmt.Println(insCnt)

		}

		http.Redirect(w, r, "/clientAllowedUris/"+clientIDStr, http.StatusFound)
	}
}
