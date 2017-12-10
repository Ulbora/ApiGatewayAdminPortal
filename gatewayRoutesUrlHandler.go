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
	"sync"

	"github.com/gorilla/mux"
)

type gwRouteURLDisplay struct {
	ID           int64
	URI          string
	ClientID     int64
	AssignedRole int64
}

func handleRouteURLs(w http.ResponseWriter, r *http.Request) {
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

			var gr services.GatewayRouteService
			gr.ClientID = getAuthCodeClient()
			gr.Host = getGatewayHost()
			gr.Token = token.AccessToken
			grr := gr.GetRouteList(clientID)

			var page gwPage
			page.GwActive = "active"
			page.Client = res
			page.GatewayClient = gres
			page.GatewayRoutes = grr
			var sm gwSideMenu
			sm.RouteURLsActive = "active teal"
			page.GwSideMenu = &sm
			//fmt.Println(page)
			templates.ExecuteTemplate(w, "gatewayRouteUrls.html", &page)
		}
	}
}

func handleRouteURLsByRoute(w http.ResponseWriter, r *http.Request) {
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

		if clientID != "" && ID != "" {
			var wg sync.WaitGroup

			var c services.ClientService
			token := getToken(w, r)
			c.ClientID = getAuthCodeClient()
			c.Host = getOauthHost()
			c.Token = token.AccessToken
			wg.Add(1)
			var res *services.Client
			go func(clientID string) {
				res = c.GetClient(clientID)
				defer wg.Done()
			}(clientID)

			var g services.GatewayClientService
			//token := getToken(w, r)
			g.ClientID = getAuthCodeClient()
			g.Host = getGatewayHost()
			g.Token = token.AccessToken
			var gres *services.GatewayClient
			wg.Add(1)
			go func(clientID string) {
				gres = g.GetClient(clientID)
				defer wg.Done()
			}(clientID)

			fmt.Println(gres)

			var gr services.GatewayRouteService
			gr.ClientID = getAuthCodeClient()
			gr.Host = getGatewayHost()
			gr.Token = token.AccessToken
			var grr *services.GatewayRoute
			wg.Add(1)
			go func(routeID string, clientID string) {
				grr = gr.GetRoute(routeID, clientID)
				defer wg.Done()
			}(ID, clientID)

			var gu services.GatewayRouteURLService
			gu.ClientID = getAuthCodeClient()
			gu.Host = getGatewayHost()
			gu.Token = token.AccessToken
			wg.Add(1)
			var grus *[]services.GatewayRouteURL
			go func(routeID string, clientID string) {
				grus = gu.GetRouteURLList(ID, clientID)
				defer wg.Done()
			}(ID, clientID)
			wg.Wait()

			var page gwPage
			page.GwActive = "active"
			page.Client = res
			page.GatewayClient = gres
			page.GatewayRoute = grr
			page.GatewayRouteURIs = grus
			var sm gwSideMenu
			sm.EditRoute = "active teal"
			page.GwSideMenu = &sm
			//fmt.Println(page)
			templates.ExecuteTemplate(w, "gatewayRouteUrlsByRoute.html", &page)
		}
	}
}

func handleRouteURLAdd(w http.ResponseWriter, r *http.Request) {
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

		routeIDStr := r.FormValue("routeId")
		routeID, _ := strconv.ParseInt(routeIDStr, 10, 0)
		fmt.Println(routeID)

		clientIDStr := r.FormValue("clientId")
		clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)
		fmt.Print("clientId: ")
		fmt.Println(clientID)

		name := r.FormValue("name")
		fmt.Print("name: ")
		fmt.Println(name)

		gwURL := r.FormValue("gwUrl")
		fmt.Print("gwUrl: ")
		fmt.Println(gwURL)

		if routeIDStr != "" && clientIDStr != "" {
			var gu services.GatewayRouteURLService
			gu.ClientID = getAuthCodeClient()
			gu.Host = getGatewayHost()
			gu.Token = token.AccessToken

			var guu services.GatewayRouteURL
			guu.ClientID = clientID
			guu.RouteID = routeID
			guu.Name = name
			guu.URL = gwURL
			guRes := gu.AddRouteURL(&guu)
			if guRes.Success != true {
				fmt.Println(guRes)
			}
		}
		http.Redirect(w, r, "/gatewayRouteUrlsByRoute/"+routeIDStr+"/"+clientIDStr, http.StatusFound)
	}
}

func handleRouteURLEdit(w http.ResponseWriter, r *http.Request) {
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
		//ID, _ := strconv.ParseInt(IDStr, 10, 0)
		fmt.Println(IDStr)

		routeIDStr := vars["routeId"]
		//routeID, _ := strconv.ParseInt(routeIDStr, 10, 0)
		fmt.Println(routeIDStr)

		clientIDStr := vars["clientId"]
		//clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)

		fmt.Print("clientId: ")
		fmt.Println(clientIDStr)

		if IDStr != "" && routeIDStr != "" && clientIDStr != "" {
			var wg sync.WaitGroup

			var c services.ClientService
			token := getToken(w, r)
			c.ClientID = getAuthCodeClient()
			c.Host = getOauthHost()
			c.Token = token.AccessToken
			wg.Add(1)
			var res *services.Client
			go func(clientID string) {
				res = c.GetClient(clientIDStr)
				defer wg.Done()
			}(clientIDStr)

			var g services.GatewayClientService
			//token := getToken(w, r)
			g.ClientID = getAuthCodeClient()
			g.Host = getGatewayHost()
			g.Token = token.AccessToken
			var gres *services.GatewayClient
			wg.Add(1)
			go func(clientID string) {
				gres = g.GetClient(clientIDStr)
				defer wg.Done()
			}(clientIDStr)

			fmt.Println(gres)

			var gr services.GatewayRouteService
			gr.ClientID = getAuthCodeClient()
			gr.Host = getGatewayHost()
			gr.Token = token.AccessToken
			var grr *services.GatewayRoute
			wg.Add(1)
			go func(routeID string, clientID string) {
				grr = gr.GetRoute(routeIDStr, clientIDStr)
				defer wg.Done()
			}(routeIDStr, clientIDStr)

			var gu services.GatewayRouteURLService
			gu.ClientID = getAuthCodeClient()
			gu.Host = getGatewayHost()
			gu.Token = token.AccessToken
			var guRes *services.GatewayRouteURL
			wg.Add(1)
			go func(id string, routeID string, clientID string) {
				guRes = gu.GetRouteURL(IDStr, routeIDStr, clientIDStr)
				defer wg.Done()
			}(IDStr, routeIDStr, clientIDStr)

			var cu services.GatewayBreakerService
			cu.ClientID = getAuthCodeClient()
			cu.Host = getGatewayHost()
			cu.Token = token.AccessToken
			var cRes *services.GatewayBreaker
			wg.Add(1)
			go func(urlID string, routeID string, clientID string) {
				cRes = cu.GetBreaker(IDStr, routeIDStr, clientIDStr)
				defer wg.Done()
			}(IDStr, routeIDStr, clientIDStr)

			wg.Wait()
			var page gwPage
			page.GwActive = "active"
			page.Client = res
			page.GatewayRoute = grr
			page.GatewayClient = gres
			page.GatewayRouteURI = guRes
			page.CircuitBreaker = cRes
			if cRes.ID != 0 {
				page.CircuitBreakerEnabled = true
			}
			var sm gwSideMenu
			sm.EditURL = "active teal"
			page.GwSideMenu = &sm
			//fmt.Println(page)
			templates.ExecuteTemplate(w, "gatewayRouteUrl.html", &page)
		}
	}
}

func handleRouteURLActivate(w http.ResponseWriter, r *http.Request) {
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
		ID, _ := strconv.ParseInt(IDStr, 10, 0)
		fmt.Println(IDStr)

		routeIDStr := vars["routeId"]
		routeID, _ := strconv.ParseInt(routeIDStr, 10, 0)
		fmt.Println(routeIDStr)

		clientIDStr := vars["clientId"]
		clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)

		fmt.Print("clientId: ")
		fmt.Println(clientIDStr)

		if IDStr != "" && routeIDStr != "" && clientIDStr != "" {
			token := getToken(w, r)

			var gu services.GatewayRouteURLService
			gu.ClientID = getAuthCodeClient()
			gu.Host = getGatewayHost()
			gu.Token = token.AccessToken

			var guu services.GatewayRouteURL
			guu.ClientID = clientID
			guu.RouteID = routeID
			guu.ID = ID
			//guu.Name = name
			//guu.URL = gwURL

			guRes := gu.ActivateRouteURL(&guu)
			if guRes.Success != true {
				fmt.Println(guRes)
			}
			fmt.Println(guRes)
			http.Redirect(w, r, "/gatewayRouteUrlsByRoute/"+routeIDStr+"/"+clientIDStr, http.StatusFound)
		}
	}
}

func handleRouteURLUpdate(w http.ResponseWriter, r *http.Request) {
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
		ID, _ := strconv.ParseInt(IDStr, 10, 0)
		fmt.Println(IDStr)

		routeIDStr := r.FormValue("routeId")
		routeID, _ := strconv.ParseInt(routeIDStr, 10, 0)
		fmt.Println(routeID)

		clientIDStr := r.FormValue("clientId")
		clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)
		fmt.Print("clientId: ")
		fmt.Println(clientID)

		name := r.FormValue("name")
		fmt.Print("name: ")
		fmt.Println(name)

		gwURL := r.FormValue("gwUrl")
		fmt.Print("gwUrl: ")
		fmt.Println(gwURL)

		cbEnabled := r.FormValue("cbEnabled")
		fmt.Print("cbEnabled: ")
		fmt.Println(cbEnabled)

		cbIDStr := r.FormValue("cbId")
		var cbID int64
		if cbIDStr != "" {
			cbID, _ = strconv.ParseInt(cbIDStr, 10, 0)
		}
		fmt.Println(cbID)

		ftStr := r.FormValue("failureThreshold")
		var ft int
		if ftStr != "" {
			ft, _ = strconv.Atoi(ftStr)
		}
		fmt.Println(ft)

		htStr := r.FormValue("healthCheckTimeSeconds")
		var ht int
		if htStr != "" {
			ht, _ = strconv.Atoi(htStr)
		}
		fmt.Println(ht)

		fname := r.FormValue("failoverRouteName")
		fmt.Print("fname: ")
		fmt.Println(fname)

		if IDStr != "" && routeIDStr != "" && clientIDStr != "" {
			token := getToken(w, r)

			var gu services.GatewayRouteURLService
			gu.ClientID = getAuthCodeClient()
			gu.Host = getGatewayHost()
			gu.Token = token.AccessToken

			var guu services.GatewayRouteURL
			guu.ClientID = clientID
			guu.RouteID = routeID
			guu.ID = ID
			guu.Name = name
			guu.URL = gwURL

			guRes := gu.UpdateRouteURL(&guu)
			if guRes.Success != true {
				fmt.Println(guRes)
			}

			var cbs services.GatewayBreakerService
			cbs.ClientID = getAuthCodeClient()
			cbs.Host = getGatewayHost()
			cbs.Token = token.AccessToken
			if cbEnabled == "on" {
				var cb services.GatewayBreaker
				cb.ClientID = clientID
				cb.RestRouteID = routeID
				cb.RouteURIID = ID
				if cbID != 0 {
					cb.ID = cbID
					cb.FailureThreshold = ft
					cb.HealthCheckTimeSeconds = ht
					cb.FailoverRouteName = fname
					cbRes := cbs.UpdateBreaker(&cb)
					if cbRes.Success != true {
						fmt.Println(cbRes)
					}
				} else {
					cb.FailureThreshold = ft
					cb.HealthCheckTimeSeconds = ht
					cb.FailoverRouteName = fname
					cbRes := cbs.InsertBreaker(&cb)
					if cbRes.Success != true {
						fmt.Println(cbRes)
					}
				}

			} else {
				if cbID != 0 {
					cbdRes := cbs.DeleteBreaker(IDStr, routeIDStr, clientIDStr)
					if cbdRes.Success != true {
						fmt.Println(cbdRes)
					}
				}
			}

			fmt.Println(guRes)
			http.Redirect(w, r, "/gatewayRouteUrlsByRoute/"+routeIDStr+"/"+clientIDStr, http.StatusFound)
		}
	}
}

func handleRouteURLDelete(w http.ResponseWriter, r *http.Request) {
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

		routeID := vars["routeId"]
		fmt.Println(routeID)

		clientIDStr := vars["clientId"]

		fmt.Print("clientId: ")
		fmt.Println(clientIDStr)

		if IDStr != "" && clientIDStr != "" {
			token := getToken(w, r)

			var gu services.GatewayRouteURLService
			gu.ClientID = getAuthCodeClient()
			gu.Host = getGatewayHost()
			gu.Token = token.AccessToken

			guRes := gu.DeleteRouteURL(IDStr, routeID, clientIDStr)
			if guRes.Success != true {
				fmt.Println(guRes)
			}
			fmt.Println(guRes)
			http.Redirect(w, r, "/gatewayRouteUrlsByRoute/"+routeID+"/"+clientIDStr, http.StatusFound)
		}
	}
}
