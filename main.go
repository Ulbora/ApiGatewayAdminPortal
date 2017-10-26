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
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	//services "UlboraCmsV3/services"

	usession "github.com/Ulbora/go-better-sessions"
	oauth2 "github.com/Ulbora/go-oauth2-client"
	"github.com/gorilla/mux"
)

var s usession.Session

//var token *oauth2.Token
var tokenMap map[string]*oauth2.Token
var credentialToken *oauth2.Token

var templates = template.Must(template.ParseFiles("./static/index.html", "./static/header.html",
	"./static/footer.html", "./static/navbar.html", "./static/clients.html", "./static/addClient.html",
	"./static/editClient.html", "./static/oauth2.html", "./static/redirectUrls.html", "./static/grantTypes.html",
	"./static/roles.html", "./static/allowedUris.html"))

//var username string

func main() {
	//gob.Register(oauth2.Token)
	s.MaxAge = sessingTimeToLive
	s.Name = userSession
	if os.Getenv("SESSION_SECRET_KEY") != "" {
		s.SessionKey = os.Getenv("SESSION_SECRET_KEY")
	} else {
		s.SessionKey = "115722gggg14ddfg4567"
	}

	tokenMap = make(map[string]*oauth2.Token)

	router := mux.NewRouter()

	router.HandleFunc("/", handleIndex)
	router.HandleFunc("/clients", handleClients)
	router.HandleFunc("/addClient", handleAddClient)
	router.HandleFunc("/editClient/{clientId}", handleEditClient)
	router.HandleFunc("/newClient", handleNewClient)
	router.HandleFunc("/updateClient", handleUpdateClient)

	router.HandleFunc("/oauth2/{clientId}", handleOauth2)

	router.HandleFunc("/clientRedirectUrls/{clientId}", handleRedirectURLs)
	router.HandleFunc("/addRedirectUrl", handleRedirectURLAdd)
	router.HandleFunc("/deleteRedirectUri/{id}/{clientId}", handleRedirectURLDelete)

	router.HandleFunc("/clientGrantTypes/{clientId}", handleGrantType)
	router.HandleFunc("/addGrantType", handleGrantTypeAdd)
	router.HandleFunc("/deleteGrantType/{id}/{clientId}", handleGrantTypeDelete)

	router.HandleFunc("/clientRoles/{clientId}", handleRoles)
	router.HandleFunc("/addClientRole", handleRoleAdd)
	router.HandleFunc("/deleteClientRoles/{id}/{clientId}", handleRoleDelete)

	router.HandleFunc("/clientAllowedUris/{clientId}", handleAllowedUris)
	router.HandleFunc("/addAllowedUri", handleAllowedUrisAdd)
	router.HandleFunc("/deleteAllowedUri/{id}/{clientId}", handleAllowedUrisDelete)

	//router.HandleFunc("/gateway/{clientId}", handleGateway)

	router.HandleFunc("/tokenHandler", handleToken)
	router.HandleFunc("/login", handleLogin)
	router.HandleFunc("/logout", handleLogout)

	// admin resources
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	//http.Handle("/js", fs)

	fmt.Println("API Gateway Admin Portal running!")
	log.Println("Listening on :8091...")
	http.ListenAndServe(":8091", router)
}
