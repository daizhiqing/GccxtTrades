package bitfinex

import "GccxtTrades/config"

const (
	Name = "bitfinex"

	BitfinexErrorLimit = config.DEFAULT

	BitfinexBufferSize = config.DEFAULT_BUFFER_SIZE

	//交易对
	BitfinexSymbole = "https://api.bitfinex.com/v1/symbols"

	BitfinexOrigin = "https://api.bitfinex.com"
	BitfinexWsUrl  = "wss://api.bitfinex.com/ws"
)
