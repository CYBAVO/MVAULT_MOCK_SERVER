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
	"encoding/json"

	"github.com/astaxie/beego/logs"
)

type CommonResponse struct {
	Result int64 `json:"result"`
}

type ErrorCodeResponse struct {
	ErrMsg  string `json:"error,omitempty"`
	ErrCode int    `json:"error_code,omitempty"`
}

type GetTransactionsInfoRequest struct {
	Currency int64    `json:"currency"`
	TXIDs    []string `json:"txids"`
}

type GetTransactionInfoResponse struct {
	BlockNumber     int64  `json:"blockNumber"`
	ConfirmBlocks   int64  `json:"confirm_blocks"`
	BlockTimeStamp  int64  `json:"blockTimeStamp"`
	ContractAddress string `json:"contract_address"`
	Fee             string `json:"fee"`
	Id              string `json:"id"`
	ResMessage      string `json:"resMessage"`
	Result          string `json:"result"`
	Success         bool   `json:"success"`
	From            string `json:"from"`
	To              string `json:"to"`
	Amount          string `json:"amount"`
	Data            string `json:"data"`
}

type GetTransactionsInfoResponse struct {
	TransactionsInfo map[string]GetTransactionInfoResponse `json:"transactions_info"`
}

type AddressItem struct {
	UserID              int64  `json:"user_id"`
	UserRegisterAccount string `json:"user_register_account"`
	UserName            string `json:"user_name"`
	UserEnmail          string `json:"user_email"`
	Curency             int64  `json:"currency"`
	TokenAddress        string `json:"token_address"`
	WalletID            int64  `json:"wallet_id"`
	WalletAddress       string `json:"wallet_address"`
	BalanceInvalidated  bool   `json:"balance_invalidated"`
	Balance             string `json:"balance"`
	BalanceTime         int64  `json:"balance_time"`
	ChainID             int64  `json:"chain_id"`
	WalletName          string `json:"wallet_name,omitempty"`
}

type GetWalletsResponse struct {
	Address    []*AddressItem `json:"addresses"`
	TotalCount int64          `json:"total_count,omitempty"`
}

type GetWalletBalanceResponse struct {
	Balance string `json:"balance"`
}

type CurrencyInfo struct {
	Currency             int64  `json:"currency"`
	CurrencyName         string `json:"currency_name"`
	TokenName            string `json:"token_name,omitempty"`
	TokenSymbol          string `json:"token_symbol,omitempty"`
	TokenContractAddress string `json:"token_contract_address,omitempty"`
	TokenDecimals        string `json:"token_decimals,omitempty"`
	TokenVersion         int64  `json:"token_version,omitempty"`
	ChainID              int64  `json:"chain_id"`
}

type SupportedCurrenciesResponse struct {
	Currenices []CurrencyInfo `json:"currencies"`
}

func GetTransactionInfo(request *GetTransactionsInfoRequest) (response *GetTransactionsInfoResponse, err error) {
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return
	}
	resp, err := makeRequest("POST", "/v1/api/transactions/info", nil, jsonRequest)
	if err != nil {
		return
	}

	response = &GetTransactionsInfoResponse{}
	err = json.Unmarshal(resp, response)

	logs.Debug("GetTransactionInfo() => ", response)
	return
}

func GetWallets(qs []string) (response *GetWalletsResponse, err error) {
	resp, err := makeRequest("GET", "/v1/api/wallets", qs, nil)
	if err != nil {
		return
	}

	response = &GetWalletsResponse{}
	err = json.Unmarshal(resp, response)

	logs.Debug("GetWallets() => ", response)
	return
}

func GetWalletBalance(qs []string) (response *GetWalletBalanceResponse, err error) {
	resp, err := makeRequest("GET", "/v1/api/wallets/balance", qs, nil)
	if err != nil {
		return
	}

	response = &GetWalletBalanceResponse{}
	err = json.Unmarshal(resp, response)

	logs.Debug("GetWalletBalance() => ", response)
	return
}

func GetUserWalletList(qs []string) (response *GetWalletsResponse, err error) {
	resp, err := makeRequest("GET", "/v1/api/user/wallets", qs, nil)
	if err != nil {
		return
	}

	response = &GetWalletsResponse{}
	err = json.Unmarshal(resp, response)

	logs.Debug("GetUserWalletList() => ", response)
	return
}

func GetSupportedCurrencies() (response *SupportedCurrenciesResponse, err error) {
	resp, err := makeRequest("GET", "/v1/api/wallets/currencies", nil, nil)
	if err != nil {
		return
	}

	response = &SupportedCurrenciesResponse{}
	err = json.Unmarshal(resp, response)

	logs.Debug("GetSupportedCurrencies() => ", response)
	return
}

type GetCurrencyTransactionsRequest struct {
	Currency      int64  `json:"currency"`
	Token         string `json:"token"`
	WalletAddress string `json:"wallet_address"`
	StartIndex    int64  `json:"start_index"`
	Limit         int64  `json:"limit"`
}

type CurrencyTransactionInfo struct {
	BlockNumber     int64  `json:"blockNumber"`
	ConfirmBlocks   int64  `json:"confirm_blocks"`
	BlockTimeStamp  int64  `json:"blockTimeStamp"`
	ContractAddress string `json:"contract_address"`
	Fee             string `json:"fee"`
	Result          string `json:"result"`
	Success         bool   `json:"success"`
	From            string `json:"from"`
	To              string `json:"to"`
	Amount          string `json:"amount"`
	Data            string `json:"data"`
}

type GetCurrencyTransactionsResponse struct {
	Transactions []CurrencyTransactionInfo `json:"transactions"`
	TotalCount   int64                     `json:"total_count"`
}

func GetCurrencyTransactions(request GetCurrencyTransactionsRequest) (response *GetCurrencyTransactionsResponse, err error) {

	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return
	}
	resp, err := makeRequest("POST", "/v1/api/transactions/currency", nil, jsonRequest)
	if err != nil {
		return
	}

	response = &GetCurrencyTransactionsResponse{}
	err = json.Unmarshal(resp, response)

	logs.Debug("GetTransactionsCurrency() => ", response)
	return
}
