package rest_func

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"peoplevine-auth/logger"
)

// ********************************************************
func restCallPOST(URL string, data []byte) (rbody []byte, err error) {
	var (
		Retry int = 3
		resp  *http.Response
	)

	err = nil

	for RetryCnt := 1; RetryCnt <= Retry; RetryCnt++ {
		resp, err = http.Post(URL, "application/json", bytes.NewBuffer(data))
		if err != nil {
			logger.Logger.Traceln(fmt.Sprintf("Error call to [%s]. ERR: [%s] Retry...", URL, err.Error()))
			continue
		}

		rbody, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Logger.Traceln(fmt.Sprintf("Error read data from [%s]. ERR: [%s]. Retry...  ", URL, err.Error()))
			continue
		}
		return rbody, nil
	}
	return nil, err
}

//********************************************************
func parseInputStructParams(r *http.Request, JsonParams interface{}) (fields map[string]interface{}, err error) {
	var (
		rbody []byte = nil
	)

	fields = make(map[string]interface{})

	rbody, err = ioutil.ReadAll(r.Body)
	if err != nil {
		err = errors.New("Error get request body:" + err.Error())
		logger.Logger.Errorln(err.Error())
		return fields, err
	}

	err = json.Unmarshal(rbody, &fields)
	if err != nil {
		err = errors.New("Error parse request body to fields")
		logger.Logger.Errorln(err.Error())
		return fields, err
	}

	err = json.Unmarshal(rbody, JsonParams)
	if err != nil {
		err = errors.New("Error parse request body")
		logger.Logger.Errorln(err.Error())
		return fields, err
	}
	return fields, nil
}

//********************************************************
func checkFieldExist(FieldsList map[string]interface{}, Field string) bool {
	if _, ok := FieldsList[Field]; ok {
		return true
	}
	return false
}
