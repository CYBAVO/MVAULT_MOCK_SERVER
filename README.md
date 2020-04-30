# Transaction API

#### Query transaction info by given TXIDs
- **POST** /v1/api/transactions/info
	- Request
		-	Params
			- 	`currency`: Specified the currency
				-	BTC: 0
				-	LTC: 2
				-	ETH: 60
				-	XRP: 144
				-	BCH: 145
				-	EOS: 194
			-  `txids`: Specify the TX ID of transactions
				-	**NOTE: Only those transactions you made can be queried by this API**
		-  Sample:
			
			```
			{
			  "currency": 60,
			  "txids": [
			    "0x806058ca93570df0fd75439e1a15ada125f0f72ff52af1e3b8c5b1b1e7cd8b6e"
			  ]
			}
			```

	- Response
		-	Params
			-	`blockNumber`: block height of this transaction
			-	`confirm_blocks`: already confirmed block numbers
			-	`blockTimeStamp`: time to pack into block
			-	`contract_address`: token contract address
			-	`fee`: transaction fee
			-	`id`: 
			-	`result`: message for transaction fail
			-	`success`: transaction is success or not
			-	`data`: trasaction data
		-	Sample:
	
		```
		{
		  "transactions_info": {
		    "0x806058ca93570df0fd75439e1a15ada125f0f72ff52af1e3b8c5b1b1e7cd8b6e": {
		      "blockNumber": 6362254,
		      "confirm_blocks": 6320,
		      "blockTimeStamp": 1568098193,
		      "contract_address": "",
		      "fee": "0.00021",
		      "id": "",
		      "result": "",
		      "success": true,
		      "data": "0x0xa9059cbb000000000000000000000000926adfba641fb58d1741fe6ce525d0075886ca5a00000000000000000000000000000000000000000000000029a2241af62c0000"
		    }
		  }
		}
		```
	**NOTE: how to adjust a successful transaction, you need to check if success = true and block number is not 0 
	
#### Query wallet list by given cryptocurrency

- **GET** /v1/api/wallets
	- Request
		- Query
			- `currency`: Specified the currency
				-	BTC: 0
				-	LTC: 2
				-	ETH: 60
				-	XRP: 144
				-	BCH: 145
				-	EOS: 194
			- `token_address`: Specified the token address
			- `start_index`: Index of starting wallets
			- `request_number`: Count of returning wallets
		- Sample:
		```
		/v1/api/wallets?currency=60&start_index=0&token_address=0xd1d8d3fd8bc9e88c4767e46be7ce970683f92811&request_number=1000
		```
	- Response
		- Params
		    - `total_count`: total wallet count of currecy/token. Due to performance concern, this field only shows when `start_index` is 0 (first page)
			- `user_id`: user unique id
			- `user_register_account`: user register method, 可以用以下prefix檢查該用戶用什麼方式註冊
				- `GOOG-`: Google
				- `AXCEL-`: AXCEL
				- `FB-`: Facebook
				- `FORYOU-`: Foryou
				- `LINE-`: LINE
				- `WECH-`: Wechat
				- `MOCK-`: Mock
			- `user_name`: user name if have
			- `user_email`: user Email if have
			- `currency`
			- `token_address`
			- `wallet_id`
			- `wallet_address`: Address of wallet
			- `balance_invalidated`: if balance is invalid
			- `balance`: balance of wallet
			- `balance_time`: last time to get balance
		- Sample:
		```
		{
		  "addresses": [
		    {
		      "user_id": 393,
		      "user_register_account": "MOCK-VVVToken203543",
		      "user_name": "UserVVVToken203543",
		      "user_email": "",
		      "currency": 60,
		      "token_address": "0xd1d8d3fd8bc9e88c4767e46be7ce970683f92811",
		      "wallet_id": 637254071,
		      "wallet_address": "0x24cd8c0A6F8c48E70a9Ce1CEC6D3E1fAd0913089",
		      "balance_invalidated": false,
		      "balance": "0",
		      "balance_time": 1577250748
		    }
		  ],
		  "total_count": 1
		}
		```

#### Query wallet balance

- **GET** /v1/api/wallets/balance

	Query wallet balance
	
	- Request
		- Query
			- `currency`: Wallet currency
			- `token_address`: Wallet token address
			- `address`: Wallet address
		- Sample:
		```
		/v1/api/wallets/balance?currency=60&token_address=0x1be7cfd61aa8daaa9ff2f3b8820888f09462d037&address=0x4a7c66d3E7C7d2b0A4FDE97F75975433Ace9c643
		```
		
	- Response
		- Params
		    - `balance`: wallet balance
		- Sample:
		```
		{
		  "balance": "0.1"
		}
		```

# Mock Server

### How to build
-	glide install
-	go build ./mockserver.go

### Set API server URL
-	mockserver.app.conf

```
api_server_url="API_SERVER_URL"
```

### Query API token

-	Operate on web console
	-	API-CODE, API-SECRET
- 	Fill **api_code** and **api_secret** with CODE and SECRET in mockserver.app.conf

```
api_code="API-CODE"
api_secret="API-SECRET"
```

### Query transaction info

```
curl -X POST -d '{"currency":60,"txids":["0x806058ca93570df0fd75439e1a15ada125f0f72ff52af1e3b8c5b1b1e7cd8b6e"]}' \
http://localhost:8889/v1/mock/transactions/info
```

### Query wallet list

```
curl http://localhost:8889/v1/mock/wallets?currency=60&start_index=0&token_address=0xd1d8d3fd8bc9e88c4767e46be7ce970683f92811&request_number=1000
```

### Query wallet balance

```
curl http://localhost:8889/v1/mock/wallets/balance?currency=60&token_address=0x1be7cfd61aa8daaa9ff2f3b8820888f09462d037&address=0x4a7c66d3E7C7d2b0A4FDE97F75975433Ace9c643
```