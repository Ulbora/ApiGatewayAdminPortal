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
	"net/http"

	oauth2 "github.com/Ulbora/go-oauth2-client"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	s.InitSessionStore(w, r)
	authorize(w, r)
}

// login handler
func handleLogout(w http.ResponseWriter, r *http.Request) {
	removeToken(w, r)
	cookie := &http.Cookie{
		Name:   "ugw-user-session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	cookie2 := &http.Cookie{
		Name:   "ulbora_oauth2_server",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie2)
	http.Redirect(w, r, "/", http.StatusFound)
}

func authorize(res http.ResponseWriter, req *http.Request) bool {
	fmt.Println("in authorize")
	fmt.Println(schemeDefault)
	var a oauth2.AuthCodeAuthorize
	a.ClientID = getAuthCodeClient()
	a.OauthHost = getOauthRedirectHost()
	a.RedirectURI = getRedirectURI(req, authCodeRedirectURI)
	a.Scope = "write"
	a.State = authCodeState
	a.Res = res
	a.Req = req
	resp := a.AuthCodeAuthorizeUser()
	if resp != true {
		fmt.Println("Authorize Failed")
	}
	fmt.Print("Resp: ")
	fmt.Println(resp)
	return resp
}

func handleToken(res http.ResponseWriter, req *http.Request) {
	code := req.URL.Query().Get("code")
	state := req.URL.Query().Get("state")
	fmt.Println("handle token")
	if state == authCodeState {
		var tn oauth2.AuthCodeToken
		tn.OauthHost = getOauthHost()
		tn.ClientID = getAuthCodeClient()
		tn.Secret = getAuthCodeSecret()
		tn.Code = code
		tn.RedirectURI = getRedirectURI(req, authCodeRedirectURI)
		fmt.Println("getting token")
		resp := tn.AuthCodeToken()
		fmt.Print("token len: ")
		fmt.Println(len(resp.AccessToken))
		fmt.Println(resp.AccessToken)
		if resp != nil && resp.AccessToken != "" {
			//fmt.Println(resp.AccessToken)
			//token = resp
			session, err := s.GetSession(req)
			if err != nil {
				fmt.Println(err)
				http.Error(res, err.Error(), http.StatusInternalServerError)
			} else {
				session.Values["userLoggenIn"] = true
				var accKey = generateTokenKey()
				session.Values["accessTokenKey"] = accKey
				tokenMap[accKey] = resp
				//fmt.Print("session id: ")
				//fmt.Println(session.ID)
				err := session.Save(req, res)
				fmt.Println(err)
				http.Redirect(res, req, "/clients", http.StatusFound)

				// decode token and get user id
			}
		}
	}
}
