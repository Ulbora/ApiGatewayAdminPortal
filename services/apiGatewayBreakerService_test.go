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

var GwCid33 int64
var GwCidStr33 string
var RTID22 int64
var RUID22 int64
var RUID222 int64
var BKID int64

func TestGatewayBreakerService_AddClient(t *testing.T) {
	GwCid33 = 559994444466567666
	GwCidStr33 = "559994444466567666"
	var c GatewayClientService
	c.ClientID = "403"
	c.Host = "http://localhost:3011"
	c.Token = tempToken
	var cc GatewayClient
	cc.ClientID = GwCid33
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

func TestGatewayBreakerService_AddRoute(t *testing.T) {
	var r GatewayRouteService
	r.ClientID = "403"
	r.Host = "http://localhost:3011"
	r.Token = tempToken

	var rr GatewayRoute
	rr.ClientID = GwCid33
	rr.Route = "anewroute"
	res := r.AddRoute(&rr)
	RTID22 = res.ID
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayBreakerService_AddRouteURL(t *testing.T) {
	var r GatewayRouteURLService
	r.ClientID = "403"
	r.Host = "http://localhost:3011"
	r.Token = tempToken

	var rr GatewayRouteURL
	rr.ClientID = GwCid33
	rr.RouteID = RTID22
	rr.Name = "blue"
	rr.URL = "http://www.google.com/test"
	res := r.AddRouteURL(&rr)
	RUID22 = res.ID
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayBreakerService_InsertBreaker(t *testing.T) {
	var b GatewayBreakerService
	b.ClientID = "403"
	b.Host = "http://localhost:3011"
	b.Token = tempToken

	var bb GatewayBreaker
	bb.ClientID = GwCid33
	bb.FailoverRouteName = "red"
	bb.FailureThreshold = 3
	bb.HealthCheckTimeSeconds = 60
	bb.RestRouteID = RTID22
	bb.RouteURIID = RUID22
	res := b.InsertBreaker(&bb)
	fmt.Print("breaker res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayBreakerService_GetBreaker(t *testing.T) {
	var b GatewayBreakerService
	b.ClientID = "403"
	b.Host = "http://localhost:3011"
	b.Token = tempToken

	res := b.GetBreaker(strconv.FormatInt(RUID22, 10), strconv.FormatInt(RTID22, 10), strconv.FormatInt(GwCid33, 10))
	fmt.Print("get breaker res: ")
	fmt.Println(res)
	if res == nil {
		t.Fail()
	} else {
		BKID = res.ID
	}
}

func TestGatewayBreakerService_UpdateBreaker(t *testing.T) {
	var b GatewayBreakerService
	b.ClientID = "403"
	b.Host = "http://localhost:3011"
	b.Token = tempToken

	var bb GatewayBreaker
	bb.ID = BKID
	bb.ClientID = GwCid33
	bb.FailoverRouteName = "green"
	bb.FailureThreshold = 1
	bb.HealthCheckTimeSeconds = 120
	bb.RestRouteID = RTID22
	bb.RouteURIID = RUID22
	res := b.UpdateBreaker(&bb)
	fmt.Print("update breaker res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayBreakerService_GetBreaker2(t *testing.T) {
	var b GatewayBreakerService
	b.ClientID = "403"
	b.Host = "http://localhost:3011"
	b.Token = tempToken

	res := b.GetBreaker(strconv.FormatInt(RUID22, 10), strconv.FormatInt(RTID22, 10), strconv.FormatInt(GwCid33, 10))
	fmt.Print("get breaker res 2: ")
	fmt.Println(res)
	if res == nil {
		t.Fail()
	}
}

func TestGatewayBreakerService_GetBreakerStatus(t *testing.T) {
	var b GatewayBreakerService
	b.ClientID = "403"
	b.Host = "http://localhost:3011"
	b.Token = tempToken

	res := b.GetBreakerStatus(strconv.FormatInt(RUID22, 10), strconv.FormatInt(GwCid33, 10))
	fmt.Print("get breaker status: ")
	fmt.Println(res)
	if res == nil {
		t.Fail()
	}
}

func TestGatewayBreakerService_ResetBreakerStatus(t *testing.T) {
	var b GatewayBreakerService
	b.ClientID = "403"
	b.Host = "http://localhost:3011"
	b.Token = tempToken

	var bb GatewayBreaker
	bb.ClientID = GwCid33
	bb.RestRouteID = RTID22
	bb.RouteURIID = RUID22
	res := b.ResetBreaker(&bb)
	fmt.Print("reset breaker res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayBreakerService_DeleteBreakerStatus(t *testing.T) {
	var b GatewayBreakerService
	b.ClientID = "403"
	b.Host = "http://localhost:3011"
	b.Token = tempToken

	res := b.DeleteBreaker(strconv.FormatInt(RUID22, 10), strconv.FormatInt(RTID22, 10), strconv.FormatInt(GwCid33, 10))
	fmt.Print("delete breaker res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayBreakerService_DeleteClient(t *testing.T) {
	var c GatewayClientService
	c.ClientID = "403"
	c.Host = "http://localhost:3011"
	c.Token = tempToken

	res := c.DeleteClient(GwCidStr33)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		//can't delete active route url
		t.Fail()
	}
}
