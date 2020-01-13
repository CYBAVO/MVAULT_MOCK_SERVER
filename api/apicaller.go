// Copyright (c) 2018-2020 The CYBAVO developers
// All Rights Reserved.
// NOTICE: All information contained herein is, and remains
// the property of CYBAVO and its suppliers,
// if any. The intellectual and technical concepts contained
// herein are proprietary to CYBAVO
// Dissemination of this information or reproduction of this materia
// is strictly forbidden unless prior written permission is obtained
// from CYBAVO.

package api

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var baseURL = beego.AppConfig.DefaultString("api_server_url", "")
var apiCode = beego.AppConfig.DefaultString("api_code", "")
var apiSecret = beego.AppConfig.DefaultString("api_secret", "")

func buildChecksum(params []string, secret string, time int64) string {
	params = append(params, fmt.Sprintf("t=%d", time))
	sort.Strings(params)
	params = append(params, fmt.Sprintf("secret=%s", secret))
	return fmt.Sprintf("%x", sha256.Sum256([]byte(strings.Join(params, "&"))))
}

func makeRequest(method string, api string, params []string, postBody []byte) ([]byte, error) {
	if method == "" || api == "" {
		return nil, errors.New("invalid parameters")
	}

	client := &http.Client{}
	t := time.Now().Unix()
	url := fmt.Sprintf("%s%s?t=%d", baseURL, api, t)
	if len(params) > 0 {
		url += fmt.Sprintf("&%s", strings.Join(params, "&"))
	}
	var req *http.Request
	var err error
	if postBody == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest("POST", url, bytes.NewReader(postBody))
		params = append(params, string(postBody))
	}
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-CODE", apiCode)
	req.Header.Set("X-CHECKSUM", buildChecksum(params, apiSecret, t))
	if postBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	logs.Debug("Request URL:", req.URL.String())
	logs.Debug("\tX-CHECKSUM:\t", req.Header.Get("X-CHECKSUM"))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return body, errors.New(res.Status)
	}
	return body, nil
}
