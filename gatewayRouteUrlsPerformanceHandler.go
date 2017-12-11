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
	"sync"

	"github.com/gorilla/mux"
)

func handleRouteURLsPerformance(w http.ResponseWriter, r *http.Request) {
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
		ID := vars["routeId"]
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

			var ps services.GatewayPerformanceService
			ps.ClientID = getAuthCodeClient()
			ps.Host = getGatewayHost()
			ps.Token = token.AccessToken
			//var cRes *services.GatewayBreaker

			var grusDisp = make([]gatewayRouteURLDisp, 0)
			for _, u := range *grus {
				var gudisp gatewayRouteURLDisp
				gudisp.ID = u.ID
				gudisp.Name = u.Name
				gudisp.URL = u.URL
				gudisp.RouteID = u.RouteID
				gudisp.ClientID = u.ClientID
				gudisp.Active = u.Active

				var pss services.GatewayPerformance
				pss.ClientID = u.ClientID
				pss.RouteURIID = u.ID
				pss.RestRouteID = u.RouteID
				pRes := ps.GetRoutePerformance(&pss)
				var lat int64
				var cnt int64
				for _, p := range *pRes {
					lat += p.LatencyMsTotal
					cnt += p.Calls
				}
				if lat > 0 && cnt > 0 {
					aveLat := (lat / cnt)
					gudisp.AverageLatency = aveLat
				}

				// cRes := cu.GetBreakerStatus(strconv.FormatInt(u.ID, 10), clientID)
				// if cRes.Open == true {
				// 	gudisp.Healthy = false
				// 	gudisp.BreakerStatus = "Open"
				// } else if cRes.PartialOpen == true {
				// 	gudisp.Healthy = true
				// 	gudisp.BreakerStatus = "Partially Open"
				// } else if cRes.Warning == true {
				// 	gudisp.Healthy = true
				// 	gudisp.BreakerStatus = "Warning"
				// } else {
				// 	gudisp.Healthy = true
				// 	gudisp.BreakerStatus = "Normal"
				// }

				grusDisp = append(grusDisp, gudisp)
			}

			var page gwPage
			page.GwActive = "active"
			page.Client = res
			page.GatewayClient = gres
			page.GatewayRoute = grr
			page.GatewayRouteURLsDisp = &grusDisp

			var sm gwSideMenu
			sm.EditRoute = "active teal"
			page.GwSideMenu = &sm
			//fmt.Println(page)
			templates.ExecuteTemplate(w, "gatewayRouteUrlsPerformance.html", &page)
		}
	}
}

// func handleResetBreaker(w http.ResponseWriter, r *http.Request) {
// 	s.InitSessionStore(w, r)
// 	session, err := s.GetSession(r)
// 	if err != nil {
// 		fmt.Println(err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// 	loggedIn := session.Values["userLoggenIn"]
// 	token := getToken(w, r)
// 	fmt.Print("in main page. Logged in: ")
// 	fmt.Println(loggedIn)
// 	//fmt.Println(token.AccessToken)
// 	//var res *[]services.Client
// 	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
// 		authorize(w, r)
// 	} else {
// 		session.Values["userLoggenIn"] = true
// 		vars := mux.Vars(r)

// 		IDStr := vars["urlId"]
// 		ID, _ := strconv.ParseInt(IDStr, 10, 0)
// 		fmt.Println(IDStr)

// 		routeIDStr := vars["routeId"]
// 		//routeID, _ := strconv.ParseInt(routeIDStr, 10, 0)
// 		fmt.Println(routeIDStr)

// 		clientIDStr := vars["clientId"]
// 		clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)

// 		fmt.Print("clientId: ")
// 		fmt.Println(clientIDStr)

// 		if IDStr != "" && clientIDStr != "" {
// 			token := getToken(w, r)

// 			var cu services.GatewayBreakerService
// 			cu.ClientID = getAuthCodeClient()
// 			cu.Host = getGatewayHost()
// 			cu.Token = token.AccessToken

// 			var cuu services.GatewayBreaker
// 			cuu.ClientID = clientID
// 			cuu.RouteURIID = ID

// 			//guu.Name = name
// 			//guu.URL = gwURL

// 			cuRes := cu.ResetBreaker(&cuu)
// 			if cuRes.Success != true {
// 				fmt.Println(cuRes)
// 			}
// 			http.Redirect(w, r, "/gatewayRouteUrlsStatus/"+routeIDStr+"/"+clientIDStr, http.StatusFound)
// 		}
// 	}
// }
