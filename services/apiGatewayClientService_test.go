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

var GwCid int64
var GwCidStr string

func TestGatewayClientService_AddClient(t *testing.T) {
	GwCid = 55445588844444444
	GwCidStr = "55445588844444444"
	var c GatewayClientService
	c.ClientID = "403"
	c.Host = "http://localhost:3011"
	c.Token = tempToken
	var cc GatewayClient
	cc.ClientID = GwCid
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

func TestGatewayClientService_UpdateClient(t *testing.T) {
	var c GatewayClientService
	c.ClientID = "403"
	c.Host = "http://localhost:3011"
	c.Token = tempToken
	var cc GatewayClient
	cc.ClientID = GwCid
	cc.APIKey = "12344444"
	cc.Enabled = false
	cc.Level = "small"
	res := c.UpdateClient(&cc)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}

func TestGatewayClientService_GetClient(t *testing.T) {
	var c GatewayClientService
	c.ClientID = "403"
	c.Host = "http://localhost:3011"
	c.Token = tempToken

	res := c.GetClient(GwCidStr)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Enabled == true || res.ClientID != GwCid {
		t.Fail()
	}
}

func TestGatewayClientService_GetClientList(t *testing.T) {
	var c GatewayClientService
	c.ClientID = "403"
	c.Host = "http://localhost:3011"
	c.Token = tempToken

	res := c.GetClientList()
	fmt.Print("res: ")
	fmt.Println(res)
	if len(*res) == 0 {
		t.Fail()
	}
}

func TestGatewayClientService_DeleteClient(t *testing.T) {
	var c GatewayClientService
	c.ClientID = "403"
	c.Host = "http://localhost:3011"
	c.Token = tempToken

	res := c.DeleteClient(GwCidStr)
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Success != true {
		t.Fail()
	}
}
