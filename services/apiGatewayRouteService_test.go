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

package services

import (
	"fmt"
	"strconv"
	"testing"
)

var GwCid2 int64
var GwCidStr2 string
var RTID int64

func TestGatewayRouteService_AddClient(t *testing.T) {
	GwCid2 = 5544558884444466666
	GwCidStr2 = "5544558884444466666"
	var c GatewayClientService
	c.ClientID = "403"
	c.Host = "http://localhost:3011"
	c.Token = tempToken
	var cc GatewayClient
	cc.ClientID = GwCid2
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

func TestGatewayRouteService_AddRoute(t *testing.T) {
	var r GatewayRouteService
	r.ClientID = "403"
	r.Host = "http://localhost:3011"
	r.Token = tempToken

	var rr GatewayRoute
	rr.ClientID = GwCid2
	rr.Route = "anewroute"
	res := r.AddRoute(&rr)
	RTID = res.ID
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayRouteService_UpdateRoute(t *testing.T) {
	var r GatewayRouteService
	r.ClientID = "403"
	r.Host = "http://localhost:3011"
	r.Token = tempToken

	var rr GatewayRoute
	rr.ID = RTID
	rr.ClientID = GwCid2
	rr.Route = "anewmailroute"
	res := r.UpdateRoute(&rr)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayRouteService_GetRoute(t *testing.T) {
	var r GatewayRouteService
	r.ClientID = "403"
	r.Host = "http://localhost:3011"
	r.Token = tempToken

	res := r.GetRoute(strconv.FormatInt(RTID, 10), GwCidStr2)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Route != "anewmailroute" {
		t.Fail()
	}
}

func TestGatewayRouteService_GetRouteList(t *testing.T) {
	var r GatewayRouteService
	r.ClientID = "403"
	r.Host = "http://localhost:3011"
	r.Token = tempToken

	res := r.GetRouteList(GwCidStr2)
	fmt.Print("res: ")
	fmt.Println(res)
	if len(*res) == 0 {
		t.Fail()
	}
}

func TestGatewayRouteService_DeleteRoute(t *testing.T) {
	var r GatewayRouteService
	r.ClientID = "403"
	r.Host = "http://localhost:3011"
	r.Token = tempToken

	res := r.DeleteRoute(strconv.FormatInt(RTID, 10), GwCidStr2)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayRouteService_DeleteClient(t *testing.T) {
	var c GatewayClientService
	c.ClientID = "403"
	c.Host = "http://localhost:3011"
	c.Token = tempToken

	res := c.DeleteClient(GwCidStr2)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}
