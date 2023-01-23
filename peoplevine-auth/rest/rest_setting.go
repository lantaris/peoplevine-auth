package rest

import "peoplevine-auth/rest_func"

var (
	RestServerConf = TRestServerConf{
		Name:   "REST server",
		Listen: "127.0.0.1:8088",
		Endpoints: []TRestServerEndpoints{
			{Name: "AuthFull", Endpoint: "/auth", Type: "POST", Callback: rest_func.RestFunc},
			{Name: "AuthSimple", Endpoint: "/simple", Type: "POST", Callback: rest_func.RestFunc},
		},
	}
)
