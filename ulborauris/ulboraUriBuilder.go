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

package ulborauri

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

var oauthSuperURIs = []UlborURIs{
	{"/ulbora/rs/client/add", "superAdmin"},
	{"/ulbora/rs/client/update", "superAdmin"},
	{"/ulbora/rs/client/get", "superAdmin"},
	{"/ulbora/rs/client/list", "superAdmin"},
	{"/ulbora/rs/client/delete", "superAdmin"},
	{"/ulbora/rs/client/search", "superAdmin"},

	{"/ulbora/rs/user/add", "superAdmin"},
	{"/ulbora/rs/user/update", "superAdmin"},
	{"/ulbora/rs/user/get", "superAdmin"},
	{"/ulbora/rs/user/search", "superAdmin"},
	{"/ulbora/rs/user/delete", "superAdmin"},
}

var apiGatewaySuperURIs = []UlborURIs{
	{"/ulbora/rs/gwClient/add", "superAdmin"},
	{"/ulbora/rs/gwClient/update", "superAdmin"},
	{"/ulbora/rs/gwClient/get", "superAdmin"},
	{"/ulbora/rs/gwClient/list", "superAdmin"},
	{"/ulbora/rs/gwClient/delete", "superAdmin"},

	{"/ulbora/rs/gwRouteUrlSuper/add", "superAdmin"},
	{"/ulbora/rs/gwRouteUrlSuper/update", "superAdmin"},
	{"/ulbora/rs/gwRouteUrlSuper/activate", "superAdmin"},
	{"/ulbora/rs/gwRouteUrlSuper/get", "superAdmin"},
	{"/ulbora/rs/gwRouteUrlSuper/delete", "superAdmin"},
	{"/ulbora/rs/gwRouteUrlSuper/list", "superAdmin"},

	{"/ulbora/rs/gwRestRouteSuper/add", "superAdmin"},
	{"/ulbora/rs/gwRestRouteSuper/update", "superAdmin"},
	{"/ulbora/rs/gwRestRouteSuper/get", "superAdmin"},
	{"/ulbora/rs/gwRestRouteSuper/list", "superAdmin"},
	{"/ulbora/rs/gwRestRouteSuper/delete", "superAdmin"},
}

var oauthURIs = []UlborURIs{

	{"/ulbora/rs/clientRoleUri/add", "admin"},
	{"/ulbora/rs/clientRoleUri/list", "admin"},
	{"/ulbora/rs/clientRoleUri/delete", "admin"},

	{"/ulbora/rs/clientRole/add", "admin"},
	{"/ulbora/rs/clientRole/list", "admin"},
	{"/ulbora/rs/clientRole/delete", "admin"},

	{"/ulbora/rs/clientRedirectUri/add", "admin"},
	{"/ulbora/rs/clientRedirectUri/list", "admin"},
	{"/ulbora/rs/clientRedirectUri/delete", "admin"},

	{"/ulbora/rs/clientGrantType/add", "admin"},
	{"/ulbora/rs/clientGrantType/list", "admin"},
	{"/ulbora/rs/clientGrantType/delete", "admin"},

	{"/ulbora/rs/clientAllowedUri/add", "admin"},
	{"/ulbora/rs/clientAllowedUri/update", "admin"},
	{"/ulbora/rs/clientAllowedUri/get", "admin"},
	{"/ulbora/rs/clientAllowedUri/list", "admin"},
	{"/ulbora/rs/clientAllowedUri/delete", "admin"},

	{"/ulbora/rs/client/user/add", "admin"},
	{"/ulbora/rs/client/user/update", "admin"},
	{"/ulbora/rs/client/user/get", "admin"},
	{"/ulbora/rs/client/user/search", "admin"},
	{"/ulbora/rs/client/user/delete", "admin"},
}

var apiGatewayURIs = []UlborURIs{
	{"/ulbora/rs/gwRouteUrl/add", "admin"},
	{"/ulbora/rs/gwRouteUrl/update", "admin"},
	{"/ulbora/rs/gwRouteUrl/activate", "admin"},
	{"/ulbora/rs/gwRouteUrl/get", "admin"},
	{"/ulbora/rs/gwRouteUrl/delete", "admin"},
	{"/ulbora/rs/gwRouteUrl/list", "admin"},

	{"/ulbora/rs/gwRestRoute/add", "admin"},
	{"/ulbora/rs/gwRestRoute/update", "admin"},
	{"/ulbora/rs/gwRestRoute/get", "admin"},
	{"/ulbora/rs/gwRestRoute/list", "admin"},
	{"/ulbora/rs/gwRestRoute/delete", "admin"},
}

var contentURIs = []UlborURIs{
	{"/ulbora/rs/content/add", "admin"},
	{"/ulbora/rs/content/update", "admin"},
	{"/ulbora/rs/content/hits", "admin"},
	{"/ulbora/rs/content/delete", "admin"},
}

var customerURIsAdmin = []UlborURIs{
	{"/ulbora/rs/customer/add", "admin"},
	{"/ulbora/rs/customer/update", "admin"},
	{"/ulbora/rs/customer/get", "user"},
	{"/ulbora/rs/customer/list", "user"},
	{"/ulbora/rs/customer/delete", "admin"},

	{"/ulbora/rs/address/add", "admin"},
	{"/ulbora/rs/address/update", "admin"},
	{"/ulbora/rs/address/get", "user"},
	{"/ulbora/rs/address/list", "user"},
	{"/ulbora/rs/address/delete", "admin"},
}

var customerURIsUser = []UlborURIs{
	{"/ulbora/rs/customer/get", "user"},
	{"/ulbora/rs/customer/list", "user"},

	{"/ulbora/rs/address/get", "user"},
	{"/ulbora/rs/address/list", "user"},
}

var imageURIsAdmin = []UlborURIs{
	{"/ulbora/rs/image/add", "admin"},
	{"/ulbora/rs/image/update", "admin"},
	{"/ulbora/rs/image/details", "user"},
	{"/ulbora/rs/image/page/count", "user"},
	{"/ulbora/rs/image/list", "user"},
	{"/ulbora/rs/image/delete", "admin"},
}

var imageURIsUser = []UlborURIs{
	{"/ulbora/rs/image/details", "user"},
	{"/ulbora/rs/image/page/count", "user"},
	{"/ulbora/rs/image/list", "user"},
}

var mailURIs = []UlborURIs{
	{"/ulbora/rs/mail/send", "admin"},
	{"/ulbora/rs/mailServer/add", "admin"},
	{"/ulbora/rs/mailServer/update", "admin"},
	{"/ulbora/rs/mailServer/get", "admin"},
}

var orderURIsAdmin = []UlborURIs{
	{"/ulbora/rs/order/add", "admin"},
	{"/ulbora/rs/order/update", "admin"},
	{"/ulbora/rs/order/get", "user"},
	{"/ulbora/rs/order/list", "user"},
	{"/ulbora/rs/order/customer/list", "user"},
	{"/ulbora/rs/order/delete", "admin"},

	{"/ulbora/rs/order/item/add", "admin"},
	{"/ulbora/rs/order/item/update", "admin"},
	{"/ulbora/rs/order/item/delete", "admin"},

	{"/ulbora/rs/order/package/add", "admin"},
	{"/ulbora/rs/order/package/update", "admin"},
	{"/ulbora/rs/order/package/get", "user"},
	{"/ulbora/rs/order/package/delete", "admin"},
}

var orderURIsUser = []UlborURIs{
	{"/ulbora/rs/order/get", "user"},
	{"/ulbora/rs/order/list", "user"},
	{"/ulbora/rs/order/customer/list", "user"},

	{"/ulbora/rs/order/package/get", "user"},
}

var productURIsAdmin = []UlborURIs{
	{"/ulbora/rs/product/add", "admin"},
	{"/ulbora/rs/product/update", "admin"},
	{"/ulbora/rs/product/get", "user"},
	{"/ulbora/rs/product/delete", "admin"},

	{"/ulbora/rs/options/add", "admin"},
	{"/ulbora/rs/options/update", "admin"},
	{"/ulbora/rs/options/get", "user"},
	{"/ulbora/rs/options/getByDetails", "user"},
	{"/ulbora/rs/options/searchByOption", "user"},
	{"/ulbora/rs/options/delete", "admin"},

	{"/ulbora/rs/details/add", "admin"},
	{"/ulbora/rs/details/update", "admin"},
	{"/ulbora/rs/details/get", "user"},
	{"/ulbora/rs/details/getByProduct", "user"},
	{"/ulbora/rs/details/getBySku", "user"},
	{"/ulbora/rs/details/getByBarCode", "user"},
	{"/ulbora/rs/details/delete", "admin"},

	{"/ulbora/rs/barCode/add", "admin"},
	{"/ulbora/rs/barCode/update", "admin"},
	{"/ulbora/rs/barCode/get", "user"},
	{"/ulbora/rs/barCode/getByDetails", "user"},
	{"/ulbora/rs/barCode/delete", "admin"},
}

var productURIsUser = []UlborURIs{
	{"/ulbora/rs/product/get", "user"},

	{"/ulbora/rs/options/get", "user"},
	{"/ulbora/rs/options/getByDetails", "user"},
	{"/ulbora/rs/options/searchByOption", "user"},

	{"/ulbora/rs/details/get", "user"},
	{"/ulbora/rs/details/getByProduct", "user"},
	{"/ulbora/rs/details/getBySku", "user"},
	{"/ulbora/rs/details/getByBarCode", "user"},

	{"/ulbora/rs/barCode/get", "user"},
	{"/ulbora/rs/barCode/getByDetails", "user"},
}

var templateURIs = []UlborURIs{
	{"/ulbora/rs/template/add", "admin"},
	{"/ulbora/rs/template/update", "admin"},
	{"/ulbora/rs/template/delete", "admin"},
}

//GetUlboraURIs GetUlboraURIs
func GetUlboraURIs(s *UlboraSelection) *[]UlborURIs {
	var rtn = make([]UlborURIs, 0)
	if s.Oauth2Super == true {
		for i := range oauthSuperURIs {
			rtn = append(rtn, oauthSuperURIs[i])
		}
	}

	if s.Oauth2 == true {
		for i := range oauthURIs {
			rtn = append(rtn, oauthURIs[i])
		}
	}

	if s.APIGatewaySuper == true {
		for i := range apiGatewaySuperURIs {
			rtn = append(rtn, apiGatewaySuperURIs[i])
		}
	}

	if s.APIGateway == true {
		for i := range apiGatewayURIs {
			rtn = append(rtn, apiGatewayURIs[i])
		}
	}

	if s.Content == true {
		for i := range contentURIs {
			rtn = append(rtn, contentURIs[i])
		}
	}

	if s.CustomerAdmin == true {
		for i := range customerURIsAdmin {
			rtn = append(rtn, customerURIsAdmin[i])
		}
	}

	if s.CustomerUser == true {
		for i := range customerURIsUser {
			rtn = append(rtn, customerURIsUser[i])
		}
	}

	if s.ImageAdmin == true {
		for i := range imageURIsAdmin {
			rtn = append(rtn, imageURIsAdmin[i])
		}
	}

	if s.ImageUser == true {
		for i := range imageURIsUser {
			rtn = append(rtn, imageURIsUser[i])
		}
	}

	if s.Mail == true {
		for i := range mailURIs {
			rtn = append(rtn, mailURIs[i])
		}
	}

	if s.OrderAdmin == true {
		for i := range orderURIsAdmin {
			rtn = append(rtn, orderURIsAdmin[i])
		}
	}

	if s.OrderUser == true {
		for i := range orderURIsUser {
			rtn = append(rtn, orderURIsUser[i])
		}
	}

	if s.ProductAdmin == true {
		for i := range productURIsAdmin {
			rtn = append(rtn, productURIsAdmin[i])
		}
	}

	if s.ProductUser == true {
		for i := range productURIsUser {
			rtn = append(rtn, productURIsUser[i])
		}
	}

	if s.Template == true {
		for i := range templateURIs {
			rtn = append(rtn, templateURIs[i])
		}
	}

	return &rtn
}
