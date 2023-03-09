package rest

import "peoplevine-auth/rest_func"

var (
	RestServerConf = TRestServerConf{
		Name:   "REST server",
		Listen: "127.0.0.1:8088",
		Endpoints: []TRestServerEndpoints{
			{Name: "AuthFull", Endpoint: "/auth", Type: "POST", Callback: rest_func.RestFunc},
			{Name: "AuthSimple", Endpoint: "/simple", Type: "POST", Callback: rest_func.RestFunc},
			{Name: "GETAuthSimple", Endpoint: "/getsimple/{auth_api_username}/{auth_api_password}/{auth_api_key}/{auth_company_no}/{auth_username}/{auth_password}/{credentials_company_no}/{credentials_email}/{credentials_password}", Type: "GET", Callback: rest_func.RestFunc},
		},
	}
)
