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

package ulborauris

//UlborURIs UlborURIs
type UlborURIs struct {
	URI  string
	Role string
}

//UlboraSelection params
type UlboraSelection struct {
	Oauth2          bool
	Oauth2Super     bool
	APIGatewaySuper bool
	APIGateway      bool
	Content         bool
	CustomerAdmin   bool
	CustomerUser    bool
	ImageAdmin      bool
	ImageUser       bool
	Mail            bool
	OrderAdmin      bool
	OrderUser       bool
	ProductAdmin    bool
	ProductUser     bool
	Template        bool
}

//GetUlboraURIs GetUlboraURIs
func GetUlboraURIs(s *UlboraSelection) *[]UlborURIs {
	var rtn = make([]UlborURIs, 0)
	if s != nil && s.Oauth2Super == true {
		for i := range oauthSuperURIs {
			rtn = append(rtn, oauthSuperURIs[i])
		}
	}

	if s != nil && s.APIGatewaySuper == true {
		for i := range apiGatewaySuperURIs {
			rtn = append(rtn, apiGatewaySuperURIs[i])
		}
	}

	if s != nil && s.Oauth2 == true {
		for i := range oauthURIs {
			rtn = append(rtn, oauthURIs[i])
		}
	}

	if s != nil && s.APIGateway == true {
		for i := range apiGatewayURIs {
			rtn = append(rtn, apiGatewayURIs[i])
		}
	}

	if s != nil && s.Content == true {
		for i := range contentURIs {
			rtn = append(rtn, contentURIs[i])
		}
	}

	if s != nil && s.CustomerAdmin == true {
		for i := range customerURIsAdmin {
			rtn = append(rtn, customerURIsAdmin[i])
		}
	}

	if s != nil && s.CustomerUser == true {
		for i := range customerURIsUser {
			rtn = append(rtn, customerURIsUser[i])
		}
	}

	if s != nil && s.ImageAdmin == true {
		for i := range imageURIsAdmin {
			rtn = append(rtn, imageURIsAdmin[i])
		}
	}

	if s != nil && s.ImageUser == true {
		for i := range imageURIsUser {
			rtn = append(rtn, imageURIsUser[i])
		}
	}

	if s != nil && s.Mail == true {
		for i := range mailURIs {
			rtn = append(rtn, mailURIs[i])
		}
	}

	if s != nil && s.OrderAdmin == true {
		for i := range orderURIsAdmin {
			rtn = append(rtn, orderURIsAdmin[i])
		}
	}

	if s != nil && s.OrderUser == true {
		for i := range orderURIsUser {
			rtn = append(rtn, orderURIsUser[i])
		}
	}

	if s != nil && s.ProductAdmin == true {
		for i := range productURIsAdmin {
			rtn = append(rtn, productURIsAdmin[i])
		}
	}

	if s != nil && s.ProductUser == true {
		for i := range productURIsUser {
			rtn = append(rtn, productURIsUser[i])
		}
	}

	if s != nil && s.Template == true {
		for i := range templateURIs {
			rtn = append(rtn, templateURIs[i])
		}
	}

	return &rtn
}
