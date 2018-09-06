package config

const (
	//火币WS的最大EOF连接次数
	HuoBiErroLimit  = DEFAULT

	HuoBiOrigin = "https://api.huobi.br.com/"
	HuoBiWsUrl = "wss://api.huobi.br.com/ws"

	HuoBiMsgBufferSize = 512

	//订阅的Generated ID
	HuoBiGId = 1000

	//direction成交方向类型：出售
	HuoBiSell = "sell";
	//direction成交方向类型：购买
	HuobiBuy = "buy"
)
