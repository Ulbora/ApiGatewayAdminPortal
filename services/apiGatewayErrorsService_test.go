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

package services

import (
	"fmt"
	"testing"
)

var GwCid3335 int64
var GwCidStr3335 string
var RTID2225 int64
var RUID22225 int64
var RUID222225 int64

func TestGatewayErrorsService_AddClient(t *testing.T) {
	GwCid3335 = 5677733211111
	GwCidStr3335 = "5677733211111"
	var c GatewayClientService
	c.ClientID = "403"
	c.Host = "http://localhost:3011"
	c.Token = tempToken
	var cc GatewayClient
	cc.ClientID = GwCid3335
	cc.APIKey = "55511111155"
	cc.Enabled = true
	cc.Level = "small"
	res := c.AddClient(&cc)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayErrorsService_AddRoute(t *testing.T) {
	var r GatewayRouteService
	r.ClientID = "403"
	r.Host = "http://localhost:3011"
	r.Token = tempToken

	var rr GatewayRoute
	rr.ClientID = GwCid3335
	rr.Route = "anewroute"
	res := r.AddRoute(&rr)
	RTID2225 = res.ID
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayErrorsService_AddRouteURL(t *testing.T) {
	var r GatewayRouteURLService
	r.ClientID = "403"
	r.Host = "http://localhost:3011"
	r.Token = tempToken

	var rr GatewayRouteURL
	rr.ClientID = GwCid3335
	rr.RouteID = RTID2225
	rr.Name = "blue"
	rr.URL = "http://www.google.com/test"
	res := r.AddRouteURL(&rr)
	RUID22225 = res.ID
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayErrorsService_GetErrors(t *testing.T) {
	var e GatewayErrorsService
	e.ClientID = "403"
	e.Host = "http://localhost:3011"
	e.Token = tempToken

	var ee GatewayError
	ee.RouteURIID = RUID22225
	ee.RestRouteID = RTID2225
	ee.ClientID = GwCid3335
	res := e.GetRouteErrors(&ee)
	fmt.Print("get performance res: ")
	fmt.Println(res)
	if res == nil || len(*res) != 0 {
		t.Fail()
	}
}

func TestGatewayErrorsService_DeleteClient(t *testing.T) {
	var c GatewayClientService
	c.ClientID = "403"
	c.Host = "http://localhost:3011"
	c.Token = tempToken

	res := c.DeleteClient(GwCidStr3335)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		//can't delete active route url
		t.Fail()
	}
}
