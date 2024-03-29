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
			- `chain_id`: currency's chain id. In non-evm currencies, -1 is tetenet and 1 is mainnet.
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
		      "balance_time": 1577250748,
			  "chain_id": 3,
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

#### Query wallet list by given user's unique token

- **GET** /v1/api/user/wallets
	- Request
		- Query
			- `unique_token`: Specified the user's unique token
			- `start_index`: Index of starting wallets
			- `request_number`: Count of returning wallets
		- Sample:
		```
		/v1/api/user/wallets?unique_token=GOOG-12345abc&start_index=0&request_number=1000
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
			- `chain_id`: currency's chain id. In non-evm currencies, -1 is tetenet and 1 is mainnet.
			- `wallet_name`: user defined wallet name
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
		      "balance_time": 1577250748,
			  "chain_id": 3,
			  "wallet_name": "Just Test"
		    }
		  ],
		  "total_count": 1
		}
		```
#### Query wallet list by given user's unique token

- **GET** /v1/api/wallets/currencies
	- Request
		- Sample:
		```
		/v1/api/wallets/currencies
		```
	- Response
		- Params
		    - `currency`: currency's BIP44
			- `currency_name`: currency's name
			- `token_name`: token name in blockchain 
			- `token_symbol`: token symbol in blockchain
			- `token_contract_address`: contract address
			- `token_decimals`: token decimal in blockchain
			- `token_version`: token's EIP version
			- `chain_id`: currency's chain id
		- Sample:
		```
		{
		  "currencies": [
        	{
            	"currency": 60,
            	"currency_name": "ETH",
            	"token_name": "etherlands chunk",
            	"token_symbol": "ELC",
            	"token_contract_address": "0x45072d88faea89dd42791808f8b491ab70b279fa",
				"token_version": 20,
            	"token_decimals": "0",
           	 "chain_id": 3
        	}]
		}
		```

#### Query currency transactions by currency, token and wallet address

- **POST** /v1/api/transactions/currency
	- Request
		- Query
			- `currency`: currency's BIP44
			- `token`: contract address
			- `start_index`: Index of starting transaction
			- `limit`: Count of returning transaction
			- `wallet_address`: (Optional) specified wallet address to query
		- Sample:
		```
		/v1/api/transactions/currency
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
			-	`total_count`: when start_index = 0, all item number will present in total_count
		- Sample:
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
		  "total_count": 101
		}
		```

#### Refresh wallet balance to sync the vaule from blokchain

- **POST** /v1/api/wallets/balance/refresh
	- Request
		- Query
			- `currency`: currency's BIP44
			- `wallets`: wallet list
				- `token_address`: empty or contract address 
				- `wallet_address`: user's wallet address in system
		- Sample:
		```
		/v1/api/wallets/balance/refresh
		{
			"currency":60,
			"wallets":[
        		{
            		"token_address": "",
            		"wallet_address": "0x8C83bd97b5521eE8A1955401489184eE775D458A"
        		},
        		{ 
            		"token_address": "0x45072d88faea89dd42791808f8b491ab70b279fa",
            		"wallet_address": "0x8C83bd97b5521eE8A1955401489184eE775D458A"
        		}
    		]
		}
		```
	- Response
		-	Params
			-	`currency`: currency's BIP44
			-	`success_count`: the number of wallets that successfully refreshed the amount
			-	`unmatched_wallets`: unmatched wallet list from request
		- Sample:
		```
		{
    		"currency": 60,
    		"success_count": 1,
    		"unmatched_wallets": [
        		"0x45072d88faea89dd42791808f8b491ab70b279fa:0x8C83bd97b5521eE8A1955401489184eE775D458A"
    		]
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

### Query wallet list by user unique token

```
curl http://localhost:8889/v1/mock/user/wallets?unique_token=GOOG-12345abc&start_index=0&request_number=1000
```

### Query supported currencies

```
curl http://localhost:8889/v1/mock/wallets/currencies
```

### Refresh wallet balance to sync the vaule from blokchain

```
curl http://localhost:8889/v1/mock/wallets/balance/refresh
```