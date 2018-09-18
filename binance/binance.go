package binance

import "GccxtTrades/config"

const (
	Name = "binance"

	BinanceErrorLimit = config.DEFAULT

	BinanceBufferSize = config.DEFAULT_BUFFER_SIZE

	//交易对
	BinanceSymbole = "https://api.binance.com/api/v1/exchangeInfo"

	BinanceOrigin = "https://api.binance.com"
	BinanceWsUrl  = "wss://stream.binance.com:9443/stream?streams="
)
