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

package ulborauris

import (
	"fmt"
	"testing"
)

func TestGetUlboraURIs_oauth2_super(t *testing.T) {
	var us UlboraSelection

	us.Oauth2Super = true

	res := GetUlboraURIs(&us)
	fmt.Print("uris: ")
	fmt.Println(res)

	fmt.Print("uris len: ")
	fmt.Println(len(*res))

	if len(*res) != 12 {
		t.Fail()
	}
}

func TestGetUlboraURIs_apiGateway_super(t *testing.T) {
	var us UlboraSelection

	us.APIGatewaySuper = true

	res := GetUlboraURIs(&us)
	fmt.Print("uris: ")
	fmt.Println(res)

	fmt.Print("uris len: ")
	fmt.Println(len(*res))

	if len(*res) != 16 {
		t.Fail()
	}
}

func TestGetUlboraURIs_oauth2(t *testing.T) {
	var us UlboraSelection

	us.Oauth2 = true

	res := GetUlboraURIs(&us)
	fmt.Print("uris: ")
	fmt.Println(res)

	fmt.Print("uris len: ")
	fmt.Println(len(*res))

	if len(*res) != 23 {
		t.Fail()
	}
}

func TestGetUlboraURIs_apiGateway(t *testing.T) {
	var us UlboraSelection

	us.APIGateway = true

	res := GetUlboraURIs(&us)
	fmt.Print("uris: ")
	fmt.Println(res)

	fmt.Print("uris len: ")
	fmt.Println(len(*res))

	if len(*res) != 11 {
		t.Fail()
	}
}

func TestGetUlboraURIs_content(t *testing.T) {
	var us UlboraSelection

	us.Content = true

	res := GetUlboraURIs(&us)
	fmt.Print("uris: ")
	fmt.Println(res)

	fmt.Print("uris len: ")
	fmt.Println(len(*res))

	if len(*res) != 4 {
		t.Fail()
	}
}

func TestGetUlboraURIs_customerAdmin(t *testing.T) {
	var us UlboraSelection

	us.CustomerAdmin = true

	res := GetUlboraURIs(&us)
	fmt.Print("uris: ")
	fmt.Println(res)

	fmt.Print("uris len: ")
	fmt.Println(len(*res))

	if len(*res) != 10 {
		t.Fail()
	}
}

func TestGetUlboraURIs_customerUser(t *testing.T) {
	var us UlboraSelection

	us.CustomerUser = true

	res := GetUlboraURIs(&us)
	fmt.Print("uris: ")
	fmt.Println(res)

	fmt.Print("uris len: ")
	fmt.Println(len(*res))

	if len(*res) != 4 {
		t.Fail()
	}
}

func TestGetUlboraURIs_imageAdmin(t *testing.T) {
	var us UlboraSelection

	us.ImageAdmin = true

	res := GetUlboraURIs(&us)
	fmt.Print("uris: ")
	fmt.Println(res)

	fmt.Print("uris len: ")
	fmt.Println(len(*res))

	if len(*res) != 6 {
		t.Fail()
	}
}

func TestGetUlboraURIs_imageUser(t *testing.T) {
	var us UlboraSelection

	us.ImageUser = true

	res := GetUlboraURIs(&us)
	fmt.Print("uris: ")
	fmt.Println(res)

	fmt.Print("uris len: ")
	fmt.Println(len(*res))

	if len(*res) != 3 {
		t.Fail()
	}
}

func TestGetUlboraURIs_mail(t *testing.T) {
	var us UlboraSelection

	us.Mail = true

	res := GetUlboraURIs(&us)
	fmt.Print("uris: ")
	fmt.Println(res)

	fmt.Print("uris len: ")
	fmt.Println(len(*res))

	if len(*res) != 4 {
		t.Fail()
	}
}

func TestGetUlboraURIs_orderAdmin(t *testing.T) {
	var us UlboraSelection

	us.OrderAdmin = true

	res := GetUlboraURIs(&us)
	fmt.Print("uris: ")
	fmt.Println(res)

	fmt.Print("uris len: ")
	fmt.Println(len(*res))

	if len(*res) != 13 {
		t.Fail()
	}
}

func TestGetUlboraURIs_orderUser(t *testing.T) {
	var us UlboraSelection

	us.OrderUser = true

	res := GetUlboraURIs(&us)
	fmt.Print("uris: ")
	fmt.Println(res)

	fmt.Print("uris len: ")
	fmt.Println(len(*res))

	if len(*res) != 4 {
		t.Fail()
	}
}

func TestGetUlboraURIs_productAdmin(t *testing.T) {
	var us UlboraSelection

	us.ProductAdmin = true

	res := GetUlboraURIs(&us)
	fmt.Print("uris: ")
	fmt.Println(res)

	fmt.Print("uris len: ")
	fmt.Println(len(*res))

	if len(*res) != 22 {
		t.Fail()
	}
}

func TestGetUlboraURIs_productUser(t *testing.T) {
	var us UlboraSelection

	us.ProductUser = true

	res := GetUlboraURIs(&us)
	fmt.Print("uris: ")
	fmt.Println(res)

	fmt.Print("uris len: ")
	fmt.Println(len(*res))

	if len(*res) != 10 {
		t.Fail()
	}
}

func TestGetUlboraURIs_template(t *testing.T) {
	var us UlboraSelection

	us.Template = true

	res := GetUlboraURIs(&us)
	fmt.Print("uris: ")
	fmt.Println(res)

	fmt.Print("uris len: ")
	fmt.Println(len(*res))

	if len(*res) != 3 {
		t.Fail()
	}
}

func TestGetUlboraURIs_combo(t *testing.T) {
	var us UlboraSelection
	us.Template = true     //3
	us.ProductAdmin = true //22
	us.Mail = true         //4
	res := GetUlboraURIs(&us)
	fmt.Print("uris: ")
	fmt.Println(res)

	fmt.Print("uris len: ")
	fmt.Println(len(*res))

	if len(*res) != 29 {
		t.Fail()
	}
}

func TestGetUlboraURIs_combo2(t *testing.T) {
	var us UlboraSelection
	us.APIGatewaySuper = true //16
	us.Oauth2Super = true     //12
	us.Mail = true            //4
	us.ImageAdmin = true      //6
	res := GetUlboraURIs(&us)
	fmt.Print("uris: ")
	fmt.Println(res)

	fmt.Print("uris len: ")
	fmt.Println(len(*res))

	if len(*res) != 38 {
		t.Fail()
	}
}
