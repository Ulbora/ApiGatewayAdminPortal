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

	"github.com/gorilla/mux"
)

type oauthPage struct {
	ClientActive         string
	OauthActive          string
	GwActive             string
	CanDeleteRedirectURI bool
	ClientList           *[]services.Client
	Client               *services.Client
	RedirectURLs         *[]services.RedirectURI
	GrantTypes           *[]services.GrantType
}

func handleOauth2(w http.ResponseWriter, r *http.Request) {
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
			//fmt.Println(page)
			templates.ExecuteTemplate(w, "oauth2.html", &page)
		}
	}
}

// func handleRedirectURLs(w http.ResponseWriter, r *http.Request) {
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
// 		clientID := vars["clientId"]

// 		if clientID != "" {
// 			var c services.ClientService
// 			token := getToken(w, r)
// 			c.ClientID = getAuthCodeClient()
// 			c.Host = getOauthHost()
// 			c.Token = token.AccessToken

// 			res := c.GetClient(clientID)
// 			//fmt.Println(res)
// 			var page oauthPage
// 			page.OauthActive = "active"
// 			page.Client = res

// 			var r services.RedirectURIService
// 			r.ClientID = getAuthCodeClient()
// 			r.Host = getOauthHost()
// 			r.Token = token.AccessToken
// 			res2 := r.GetRedirectURIList(clientID)
// 			page.RedirectURLs = res2
// 			if len(*res2) > 1 {
// 				page.CanDeleteRedirectURI = true
// 			}
// 			//fmt.Println(page)
// 			templates.ExecuteTemplate(w, "redirectUrls.html", &page)
// 		}
// 	}
// }

// func handleRedirectURLDelete(w http.ResponseWriter, r *http.Request) {
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
// 		ID := vars["id"]
// 		clientID := vars["clientId"]

// 		if ID != "" && clientID != "" {
// 			token := getToken(w, r)
// 			var r services.RedirectURIService
// 			r.ClientID = getAuthCodeClient()
// 			r.Host = getOauthHost()
// 			r.Token = token.AccessToken
// 			dres := r.DeleteRedirectURI(ID)
// 			if dres.Success != true {
// 				fmt.Println(dres)
// 			}

// 			var c services.ClientService

// 			c.ClientID = getAuthCodeClient()
// 			c.Host = getOauthHost()
// 			c.Token = token.AccessToken

// 			res := c.GetClient(clientID)
// 			//fmt.Println(res)
// 			var page oauthPage
// 			page.OauthActive = "active"
// 			page.Client = res

// 			res2 := r.GetRedirectURIList(clientID)
// 			page.RedirectURLs = res2
// 			if len(*res2) > 1 {
// 				page.CanDeleteRedirectURI = true
// 			}
// 			//fmt.Println(page)
// 			templates.ExecuteTemplate(w, "redirectUrls.html", &page)
// 		}
// 	}
// }

// func handleRedirectURLAdd(w http.ResponseWriter, r *http.Request) {
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

// 		redirectURL := r.FormValue("redirectURL")
// 		fmt.Println(redirectURL)

// 		clientIDStr := r.FormValue("clientId")
// 		clientID, _ := strconv.ParseInt(clientIDStr, 10, 0)
// 		fmt.Print("clientId: ")
// 		fmt.Println(clientID)

// 		session.Values["userLoggenIn"] = true

// 		token := getToken(w, r)
// 		var r services.RedirectURIService
// 		r.ClientID = getAuthCodeClient()
// 		r.Host = getOauthHost()
// 		r.Token = token.AccessToken

// 		var rr services.RedirectURI
// 		rr.ClientID = clientID
// 		rr.URI = redirectURL
// 		ares := r.AddRedirectURI(&rr)
// 		if ares.Success != true {
// 			fmt.Println(ares)
// 		}

// 		var c services.ClientService

// 		c.ClientID = getAuthCodeClient()
// 		c.Host = getOauthHost()
// 		c.Token = token.AccessToken

// 		res := c.GetClient(clientIDStr)
// 		//fmt.Println(res)
// 		var page oauthPage
// 		page.OauthActive = "active"
// 		page.Client = res

// 		res2 := r.GetRedirectURIList(clientIDStr)
// 		page.RedirectURLs = res2
// 		if len(*res2) > 1 {
// 			page.CanDeleteRedirectURI = true
// 		}
// 		//fmt.Println(page)
// 		templates.ExecuteTemplate(w, "redirectUrls.html", &page)

// 	}
// }

// func handleGrantType(w http.ResponseWriter, r *http.Request) {
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
// 		clientID := vars["clientId"]

// 		if clientID != "" {
// 			var c services.ClientService
// 			token := getToken(w, r)
// 			c.ClientID = getAuthCodeClient()
// 			c.Host = getOauthHost()
// 			c.Token = token.AccessToken

// 			res := c.GetClient(clientID)
// 			//fmt.Println(res)
// 			var page oauthPage
// 			page.OauthActive = "active"
// 			page.Client = res

// 			var r services.RedirectURIService
// 			r.ClientID = getAuthCodeClient()
// 			r.Host = getOauthHost()
// 			r.Token = token.AccessToken
// 			res2 := r.GetRedirectURIList(clientID)
// 			page.RedirectURLs = res2
// 			if len(*res2) > 1 {
// 				page.CanDeleteRedirectURI = true
// 			}
// 			//fmt.Println(page)
// 			templates.ExecuteTemplate(w, "redirectUrls.html", &page)
// 		}
// 	}
// }

// // user handlers-----------------------------------------------------
// func handleClients(w http.ResponseWriter, r *http.Request) {
// 	fmt.Print("url: ")
// 	fmt.Println(r.URL)
// 	fmt.Println(r.Host)
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
// 	var res *[]services.Client
// 	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
// 		authorize(w, r)
// 	} else {
// 		session.Values["userLoggenIn"] = true
// 		clientName := r.FormValue("clientName")
// 		fmt.Print("clientName: ")
// 		fmt.Println(clientName)

// 		if clientName != "" {
// 			var c services.ClientService
// 			token := getToken(w, r)
// 			c.ClientID = getAuthCodeClient()
// 			c.Host = getOauthHost()
// 			c.Token = token.AccessToken
// 			var cc services.Client
// 			cc.Name = clientName
// 			res = c.SearchClient(&cc)
// 		}
// 		var page clientPage
// 		page.ClientActive = "active"
// 		page.ClientList = res
// 		templates.ExecuteTemplate(w, "clients.html", &page)
// 	}

// }

// func handleAddClient(w http.ResponseWriter, r *http.Request) {
// 	s.InitSessionStore(w, r)
// 	session, err := s.GetSession(r)
// 	if err != nil {
// 		fmt.Println(err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// 	loggedIn := session.Values["userLoggenIn"]
// 	token := getToken(w, r)
// 	fmt.Print("Logged in: ")
// 	fmt.Println(loggedIn)
// 	//fmt.Println(token.AccessToken)
// 	//var res *[]services.Client
// 	if loggedIn == nil || loggedIn.(bool) == false || token == nil {
// 		authorize(w, r)
// 	} else {
// 		templates.ExecuteTemplate(w, "addClient.html", nil)
// 	}

// }

// func handleEditClient(w http.ResponseWriter, r *http.Request) {
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
// 		clientID := vars["clientId"]

// 		if clientID != "" {
// 			var c services.ClientService
// 			token := getToken(w, r)
// 			c.ClientID = getAuthCodeClient()
// 			c.Host = getOauthHost()
// 			c.Token = token.AccessToken

// 			res := c.GetClient(clientID)
// 			templates.ExecuteTemplate(w, "editClient.html", &res)
// 		}
// 	}
// }

// func handleNewClient(w http.ResponseWriter, r *http.Request) {
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
// 		clientName := r.FormValue("clientName")
// 		fmt.Print("clientName: ")
// 		fmt.Println(clientName)

// 		webSite := r.FormValue("webSite")
// 		fmt.Print("webSite: ")
// 		fmt.Println(webSite)

// 		emailAddress := r.FormValue("emailAddress")
// 		fmt.Print("emailAddress: ")
// 		fmt.Println(emailAddress)

// 		redirectURLStr := r.FormValue("redirectURLs")
// 		redirectURLStr = strings.Replace(redirectURLStr, " ", "", -1)
// 		fmt.Print("redirectURLStr: ")
// 		fmt.Println(redirectURLStr)
// 		redirectURLs := strings.Split(redirectURLStr, ",")
// 		fmt.Println(redirectURLs)
// 		enabled := r.FormValue("enabled")
// 		fmt.Print("enabled: ")
// 		fmt.Println(enabled)

// 		var c services.ClientService
// 		token := getToken(w, r)
// 		c.ClientID = getAuthCodeClient()
// 		c.Host = getOauthHost()
// 		c.Token = token.AccessToken

// 		var cc services.Client
// 		cc.Name = clientName
// 		cc.Email = emailAddress
// 		if enabled == "yes" {
// 			cc.Enabled = true
// 		} else {
// 			cc.Enabled = false
// 		}
// 		cc.WebSite = webSite

// 		var uris []services.RedirectURI
// 		for i := range redirectURLs {
// 			var uri services.RedirectURI
// 			uri.URI = redirectURLs[i]
// 			uris = append(uris, uri)
// 		}
// 		cc.RedirectURIs = uris
// 		//cc.Secret = generateClientSecret()
// 		res := c.AddClient(&cc)
// 		if res.Success == true {
// 			http.Redirect(w, r, "/clients", http.StatusFound)
// 		} else {
// 			http.Redirect(w, r, "/addClient", http.StatusFound)
// 		}
// 	}
// }

// func handleUpdateClient(w http.ResponseWriter, r *http.Request) {
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
// 		clientIDStr := r.FormValue("clientId")
// 		clientID, errID := strconv.ParseInt(clientIDStr, 10, 0)
// 		if errID != nil {
// 			fmt.Print(errID)
// 		}
// 		fmt.Print("clientId: ")
// 		fmt.Println(clientID)

// 		clientName := r.FormValue("clientName")
// 		fmt.Print("clientName: ")
// 		fmt.Println(clientName)

// 		secret := r.FormValue("secret")
// 		fmt.Print("secret: ")
// 		fmt.Println(secret)

// 		webSite := r.FormValue("webSite")
// 		fmt.Print("webSite: ")
// 		fmt.Println(webSite)

// 		emailAddress := r.FormValue("emailAddress")
// 		fmt.Print("emailAddress: ")
// 		fmt.Println(emailAddress)

// 		// redirectURLStr := r.FormValue("redirectURLs")
// 		// fmt.Print("redirectURLStr: ")
// 		// fmt.Println(redirectURLStr)
// 		// redirectURLs := strings.Split(redirectURLStr, ",")
// 		//fmt.Println(redirectURLs)
// 		enabled := r.FormValue("enabled")
// 		fmt.Print("enabled: ")
// 		fmt.Println(enabled)

// 		var c services.ClientService
// 		token := getToken(w, r)
// 		c.ClientID = getAuthCodeClient()
// 		c.Host = getOauthHost()
// 		c.Token = token.AccessToken

// 		var cc services.Client
// 		cc.ClientID = clientID
// 		cc.Secret = secret
// 		cc.Name = clientName
// 		cc.Email = emailAddress
// 		if enabled == "yes" {
// 			cc.Enabled = true
// 		} else {
// 			cc.Enabled = false
// 		}
// 		cc.WebSite = webSite

// 		// var uris []services.RedirectURI
// 		// for i := range redirectURLs {
// 		// 	var uri services.RedirectURI
// 		// 	uri.URI = redirectURLs[i]
// 		// 	uris = append(uris, uri)
// 		// }
// 		// cc.RedirectURIs = uris
// 		//cc.Secret = generateClientSecret()
// 		res := c.UpdateClient(&cc)
// 		if res.Success == true {
// 			http.Redirect(w, r, "/clients", http.StatusFound)
// 		} else {
// 			fmt.Println(res)
// 			http.Redirect(w, r, "/editClient/"+clientIDStr, http.StatusFound)
// 		}
// 	}
// 	//if res.Success == true {
// 	//http.Redirect(w, r, "/clients", http.StatusFound)
// 	//} else {
// 	//http.Redirect(w, r, "/admin/addContent", http.StatusFound)
// 	//}
// }
