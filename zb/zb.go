package zb

import "GccxtTrades/config"

const (

	Name = "zb"

	ZbErrorLimit = config.DEFAULT

	ZbMsgBufferSize = config.DEFAULT_BUFFER_SIZE * 10

	ZbSymbols = "http://api.zb.cn/data/v1/markets"

	ZbOrigi = "https://api.zb.cn/"
	ZbWsUrl = "wss://api.zb.cn:9999/websocket"
)
