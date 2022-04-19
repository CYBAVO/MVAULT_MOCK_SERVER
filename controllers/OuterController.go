// Copyright (c) 2018-2020 The CYBAVO developers
// All Rights Reserved.
// NOTICE: All information contained herein is, and remains
// the property of CYBAVO and its suppliers,
// if any. The intellectual and technical concepts contained
// herein are proprietary to CYBAVO
// Dissemination of this information or reproduction of this materia
// is strictly forbidden unless prior written permission is obtained
// from CYBAVO.

package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/cybavo/MVAULT_MOCK_SERVER/api"
)

type OuterController struct {
	beego.Controller
}

func (c *OuterController) AbortWithError(status int, err error) {
	resp := api.ErrorCodeResponse{
		ErrMsg:  err.Error(),
		ErrCode: status,
	}
	c.Data["json"] = resp
	c.Abort(strconv.Itoa(status))
}

func getQueryString(ctx *context.Context) []string {
	var qs []string
	tokens := strings.Split(ctx.Request.URL.RawQuery, "&")
	for _, token := range tokens {
		qs = append(qs, token)
	}
	return qs
}

// @Title Get detail transaction info by given TXIDs
// @router /transactions/info [post]
func (c *OuterController) GetTransactionInfo() {
	defer c.ServeJSON()

	var request api.GetTransactionsInfoRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	resp, err := api.GetTransactionInfo(&request)
	if err != nil {
		logs.Error("GetTransactionInfo failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @Title Get wallets by given currency
// @router /wallets [get]
func (c *OuterController) GetWallets() {
	defer c.ServeJSON()

	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.GetWallets(qs)
	if err != nil {
		logs.Error("GetWallets failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @Title Get wallet balance
// @router /wallets/balance [get]
func (c *OuterController) GetWalletBalance() {
	defer c.ServeJSON()

	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.GetWalletBalance(qs)
	if err != nil {
		logs.Error("GetWalletBalance failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @Title Get wallet balance
// @router /user/wallets [post]
func (c *OuterController) GetUserWalletList() {
	defer c.ServeJSON()

	qs := getQueryString(c.Ctx)
	if qs == nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("no required info"))
	}

	resp, err := api.GetUserWalletList(qs)
	if err != nil {
		logs.Error("api.GetUserWalletList failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}

// @Title Get wallet balance
// @router /user/wallets [post]
func (c *OuterController) GetSupportedCurrencies() {
	defer c.ServeJSON()

	resp, err := api.GetSupportedCurrencies()
	if err != nil {
		logs.Error("api.GetSupportedCurrencies failed", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Data["json"] = resp
}
