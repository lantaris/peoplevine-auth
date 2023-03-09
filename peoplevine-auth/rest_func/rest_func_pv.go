package rest_func

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"peoplevine-auth/logger"
	"peoplevine-auth/pv_types"
	"strconv"
)

//********************************************************
//*** Response
//********************************************************
func restPVAuth(w http.ResponseWriter, r *http.Request, detail bool) {
	var (
		InputData        pv_types.TPVRequest
		ResponseData     pv_types.TRetData
		err              error = nil
		respBody               = []byte("{}")
		DataRecv               = []byte("{}")
		PVDData          pv_types.TPVDData
		PVDataHeader     pv_types.TPVDataHeader
		PVDataWithDetail pv_types.TPVDataWithDetail
	)

	ResponseData.Status.Errcode = pv_types.RET_OK
	ResponseData.Status.Errmsg = "OK"
	ResponseData.Data = nil

	// Parse input data
	_, err = parseInputStructParams(r, &InputData)
	if err != nil {
		ResponseData.Status.Errmsg = "Error parse JSON parameters"
		ResponseData.Status.Errcode = pv_types.RET_PARAM_PARSE_ERROR
		goto send_response
	}

	respBody, err = json.Marshal(InputData)
	if err != nil {
		ResponseData.Status.Errmsg = "Error encode JSON for PV request"
		ResponseData.Status.Errcode = pv_types.RET_PARAM_PARSE_ERROR
		goto send_response
	}

	// Call to PV
	DataRecv, err = restCallPOST(pv_types.PV_URL, respBody)
	if err != nil {
		ResponseData.Status.Errmsg = "Rest call error: " + err.Error()
		ResponseData.Status.Errcode = pv_types.RET_CALL_ERROR
		goto send_response
	}
	// Parse response ( D - field )
	err = json.Unmarshal(DataRecv, &PVDData)
	if err != nil {
		ResponseData.Status.Errmsg = "Error parse data ( D - field ) from PeopleVine response or authentication error "
		ResponseData.Status.Errcode = pv_types.RET_ERROR
		goto send_response
	}

	// Decode response header
	err = json.Unmarshal([]byte(PVDData.D), &PVDataHeader)
	if err != nil {
		ResponseData.Status.Errmsg = "Error parse data ( Header ) from PeopleVine response or authentication error"
		ResponseData.Status.Errcode = pv_types.RET_ERROR
		goto send_response
	}

	ResponseData.Data = PVDataHeader

	// If PV response error
	if PVDataHeader.IsError == true {
		ResponseData.Status.Errmsg = PVDataHeader.Message
		ResponseData.Status.Errcode = pv_types.RET_ERROR
		logger.Logger.Infoln(fmt.Sprintf("User: [%s] not authenticated. IsError = %t, ErrMsg: [%s]  ",
			InputData.Credentials.Email, PVDataHeader.IsError, PVDataHeader.Message))
		goto send_response
	}

	// Decode response detail
	err = json.Unmarshal([]byte(PVDData.D), &PVDataWithDetail)
	if err != nil {
		ResponseData.Status.Errmsg = "Error parse data ( Header with detail )  authentication error"
		ResponseData.Status.Errcode = pv_types.RET_ERROR
		goto send_response
	}

	ResponseData.Data = PVDataWithDetail

	if PVDataWithDetail.ReturnObject.IsMember != true {
		ResponseData.Status.Errcode = pv_types.RET_ERROR
		ResponseData.Status.Errmsg = fmt.Sprintf("Error user: [%s] not authenticated. IsMember = [%t]  ",
			InputData.Credentials.Email, PVDataWithDetail.ReturnObject.IsMember)
		logger.Logger.Infoln(ResponseData.Status.Errmsg)
	}

send_response:
	if detail == false {
		ResponseData.Data = nil
	}

	if ResponseData.Status.Errcode == pv_types.RET_OK {
		w.WriteHeader(http.StatusOK)
	} else {
		logger.Logger.Errorln(ResponseData.Status.Errmsg)
		w.WriteHeader(http.StatusForbidden)
	}

	respBody, _ = json.Marshal(ResponseData)
	w.Write(respBody)
}

//********************************************************
func str2int(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return -1
	} else {
		return i
	}
}

//********************************************************
func restPVAuthGet(w http.ResponseWriter, r *http.Request, detail bool) {
	var (
		InputData        pv_types.TPVRequest
		ResponseData     pv_types.TRetData
		err              error = nil
		respBody               = []byte("{}")
		DataRecv               = []byte("{}")
		PVDData          pv_types.TPVDData
		PVDataHeader     pv_types.TPVDataHeader
		PVDataWithDetail pv_types.TPVDataWithDetail
	)

	ResponseData.Status.Errcode = pv_types.RET_OK
	ResponseData.Status.Errmsg = "OK"
	ResponseData.Data = nil

	vars := mux.Vars(r)
	InputData.Auth.ApiUsername = vars["auth_api_username"]
	InputData.Auth.ApiPassword = vars["auth_api_password"]
	InputData.Auth.ApiKey = vars["auth_api_key"]
	InputData.Auth.CompanyNo = str2int(vars["auth_company_no"])
	InputData.Auth.Username = vars["auth_username"]
	InputData.Auth.Password = vars["auth_password"]
	InputData.Credentials.CompanyNo = str2int(vars["credentials_company_no"])
	InputData.Credentials.Email = vars["credentials_email"]
	InputData.Credentials.Password = vars["credentials_password"]

	respBody, err = json.Marshal(InputData)
	if err != nil {
		ResponseData.Status.Errmsg = "Error encode JSON for PV request"
		ResponseData.Status.Errcode = pv_types.RET_PARAM_PARSE_ERROR
		goto send_response
	}

	// Call to PV
	DataRecv, err = restCallPOST(pv_types.PV_URL, respBody)
	if err != nil {
		ResponseData.Status.Errmsg = "Rest call error: " + err.Error()
		ResponseData.Status.Errcode = pv_types.RET_CALL_ERROR
		goto send_response
	}
	// Parse response ( D - field )
	err = json.Unmarshal(DataRecv, &PVDData)
	if err != nil {
		ResponseData.Status.Errmsg = "Error parse data ( D - field ) from PeopleVine response or authentication error "
		ResponseData.Status.Errcode = pv_types.RET_ERROR
		goto send_response
	}

	// Decode response header
	err = json.Unmarshal([]byte(PVDData.D), &PVDataHeader)
	if err != nil {
		ResponseData.Status.Errmsg = "Error parse data ( Header ) from PeopleVine response or authentication error"
		ResponseData.Status.Errcode = pv_types.RET_ERROR
		goto send_response
	}

	ResponseData.Data = PVDataHeader

	// If PV response error
	if PVDataHeader.IsError == true {
		ResponseData.Status.Errmsg = PVDataHeader.Message
		ResponseData.Status.Errcode = pv_types.RET_ERROR
		logger.Logger.Infoln(fmt.Sprintf("User: [%s] not authenticated. IsError = %t, ErrMsg: [%s]  ",
			InputData.Credentials.Email, PVDataHeader.IsError, PVDataHeader.Message))
		goto send_response
	}

	// Decode response detail
	err = json.Unmarshal([]byte(PVDData.D), &PVDataWithDetail)
	if err != nil {
		ResponseData.Status.Errmsg = "Error parse data ( Header with detail )  authentication error"
		ResponseData.Status.Errcode = pv_types.RET_ERROR
		goto send_response
	}

	ResponseData.Data = PVDataWithDetail

	if PVDataWithDetail.ReturnObject.IsMember != true {
		ResponseData.Status.Errcode = pv_types.RET_ERROR
		ResponseData.Status.Errmsg = fmt.Sprintf("Error user: [%s] not authenticated. IsMember = [%t]  ",
			InputData.Credentials.Email, PVDataWithDetail.ReturnObject.IsMember)
		logger.Logger.Infoln(ResponseData.Status.Errmsg)
	}

send_response:
	if detail == false {
		ResponseData.Data = nil
	}

	if ResponseData.Status.Errcode == pv_types.RET_OK {
		w.WriteHeader(http.StatusOK)
	} else {
		logger.Logger.Errorln(ResponseData.Status.Errmsg)
		w.WriteHeader(http.StatusForbidden)
	}

	respBody, _ = json.Marshal(ResponseData)
	w.Write(respBody)
}
