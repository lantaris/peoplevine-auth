package rest_func

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"peoplevine-auth/logger"
	"peoplevine-auth/pv_types"
)

//********************************************************
func RestFunc(w http.ResponseWriter, r *http.Request) {
	var (
		ResponseData pv_types.TRetData
		respBody     = []byte("{}")
		Endpoint     string
	)

	ResponseData.Status.Errcode = pv_types.RET_OK
	ResponseData.Status.Errmsg = "OK"

	Endpoint = mux.CurrentRoute(r).GetName()

	logger.Logger.Infoln("Request: ", Endpoint)
	logger.Logger.Traceln("REST: Request from", r.RemoteAddr, "[", r.RequestURI, "]")

	// Gateway endpoints
	if Endpoint == "AuthFull" {
		restPVAuth(w, r, true)
		return
	} else if Endpoint == "AuthSimple" {
		restPVAuth(w, r, false)
		return
	} else if Endpoint == "GETAuthSimple" {
		restPVAuthGet(w, r, false)
		return
	} else {
		ResponseData.Status.Errcode = pv_types.RET_ERROR
		ResponseData.Status.Errmsg = "Request unknown: " + Endpoint
		logger.Logger.Errorln(ResponseData.Status.Errmsg)
		goto send_response
	}

send_response:
	if ResponseData.Status.Errcode == 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

	respBody, _ = json.Marshal(ResponseData)
	w.Write(respBody)
}
