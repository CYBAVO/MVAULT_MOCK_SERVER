// Copyright (c) 2018-2020 The CYBAVO developers
// All Rights Reserved.
// NOTICE: All information contained herein is, and remains
// the property of CYBAVO and its suppliers,
// if any. The intellectual and technical concepts contained
// herein are proprietary to CYBAVO
// Dissemination of this information or reproduction of this materia
// is strictly forbidden unless prior written permission is obtained
// from CYBAVO.

package routers

import (
	"github.com/astaxie/beego"
	"github.com/cybavo/MVAULT_MOCK_SERVER/controllers"
)

func init() {
	beego.Router("/v1/mock/transactions/info", &controllers.OuterController{}, "post:GetTransactionInfo")
	beego.Router("/v1/mock/wallets", &controllers.OuterController{}, "get:GetWallets")
	beego.Router("/v1/mock/wallets/balance", &controllers.OuterController{}, "get:GetWalletBalance")
}
