package huobi

//
const huobiDomain string = "https://api.huobipro.com"

//获取K线数据
const KLineApi string = huobiDomain + "/market/history/kline?symbol=%s&period=%s&size=%d"

//获取聚合行情(Ticker)
const TickerMergedApi = huobiDomain + "/market/detail/merged?symbol=%s"

//获取tickers 数据
const TickersApi = huobiDomain + "/market/tickers"

//获取深度
const DepthApi = huobiDomain + "/market/depth?symbol=%s&type=%s"

//历史交易单
const TradeApi = huobiDomain + "/market/trade?symbol=%s"
