package huobi

import "GccxtTrades/config"

const (
	Name = "huobipro"
	//火币WS的最大EOF连接次数
	HuoBiErroLimit = config.DEFAULT

	HuoBiOrigin = "https://api.huobi.br.com/"
	HuoBiWsUrl  = "wss://api.huobi.br.com/ws"

	HuoBiSymbols = "https://api.huobi.pro/v1/common/symbols"

	HuoBiMsgBufferSize = config.DEFAULT_BUFFER_SIZE

	//订阅的Generated ID
	HuoBiGId = 1000

	//direction成交方向类型：出售
	HuoBiSell = "sell"
	//direction成交方向类型：购买
	HuobiBuy = "buy"
)
