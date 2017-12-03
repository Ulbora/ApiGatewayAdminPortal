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

			// var gu services.GatewayRouteURLService
			// gu.ClientID = getAuthCodeClient()
			// gu.Host = getGatewayHost()
			// gu.Token = token.AccessToken
			// gu.GetRouteURLList

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
			grr := gr.GetRoute(ID, clientID)

			var gu services.GatewayRouteURLService
			gu.ClientID = getAuthCodeClient()
			gu.Host = getGatewayHost()
			gu.Token = token.AccessToken
			grus := gu.GetRouteURLList(ID, clientID)

			var page gwPage
			page.GwActive = "active"
			page.Client = res
			page.GatewayClient = gres
			page.GatewayRoute = grr
			page.GatewayRouteURIs = grus
			var sm gwSideMenu
			//sm.RouteURLsActive = "active teal"
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
			//fmt.Println(aures)
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

		routeID := vars["id"]
		clientID := vars["clientId"]

		if clientID != "" && routeID != "" {
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
			grr := gr.GetRoute(routeID, clientID)
			var page gwPage
			page.GwActive = "active"
			page.Client = res
			page.GatewayClient = gres
			page.GatewayRoute = grr
			var sm gwSideMenu
			//sm.RouteActive = "active teal"
			page.GwSideMenu = &sm
			//fmt.Println(page)
			templates.ExecuteTemplate(w, "editGatewayRoute.html", &page)
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
		fmt.Println(IDStr)
		ID, _ := strconv.ParseInt(IDStr, 10, 0)
		fmt.Println(ID)

		clientIDStr := r.FormValue("clientId")
		clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)
		fmt.Print("clientId: ")
		fmt.Println(clientID)

		gwRoute := r.FormValue("gwRoute")
		fmt.Print("gwRoute: ")
		fmt.Println(gwRoute)

		if IDStr != "" && clientIDStr != "" && gwRoute != "" {
			var gr services.GatewayRouteService
			gr.ClientID = getAuthCodeClient()
			gr.Host = getGatewayHost()
			gr.Token = token.AccessToken

			var grr services.GatewayRoute
			grr.ClientID = clientID
			grr.Route = gwRoute
			grr.ID = ID
			grRes := gr.UpdateRoute(&grr)
			if grRes.Success != true {
				fmt.Println(grRes)
			}
		}
		http.Redirect(w, r, "/gatewayRoutes/"+clientIDStr, http.StatusFound)

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
