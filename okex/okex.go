package okex

import "GccxtTrades/config"

const (

	Name = "okex"

	OkexBufferSize = config.DEFAULT_BUFFER_SIZE

	OKexErrLimit = config.DEFAULT

	OkexOrigin = "https://real.okex.com"
	OkexWsUrl  = "wss://real.okex.com:10441/websocket" //线上
	//OkexWsUrl =  "wss://real.okcoin.com:10440/websocket"  //开发本地

	//获取所有交易对
	OkexSymbols = "https://www.okex.com/v2/spot/markets/products"
)
