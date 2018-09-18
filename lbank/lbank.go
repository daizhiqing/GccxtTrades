package lbank

import "GccxtTrades/config"

const (

	Name = "lbank"

	LBankErrorLimit = config.DEFAULT

	LBankBufferSzie = config.DEFAULT_BUFFER_SIZE

	LBankOrigin = "https://api.lbank.info"

	//LBankWsUrl = "wss://api.lbank.info/ws/V2/"
	LBankWsUrl = "wss://api.lbkex.com/ws/V2/"

	LBankSymbole = "http://api.lbank.info/v1/currencyPairs.do"
)
