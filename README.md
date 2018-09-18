# Gccxt
数字货币交易所WebSocket对接历史交易记录 Go版本实现，主要对接各大交易所的历史交易单数据，提供到消息队列进行统一数据分析

目前已经对接交易所(持续更新。。。):

* HuobiPro :  wss://api.huobi.br.com/ws
* LBank    :  wss://api.lbkex.com/ws/V2/
* ZB       :  wss://api.zb.cn:9999/websocket
* Okex     :  wss://real.okex.com:10441/websocket
* Bitfinex :  wss://api.bitfinex.com/ws
* Binance  :  wss://stream.binance.com:9443/stream
* Gateio   :  wss://ws.gateio.io/v3/
* FCoin    :  wss://api.fcoin.com/v2/ws
* Hitbtc   :  wss://api.hitbtc.com/api/2/ws
* Hadax    :  wss://api.hadax.com/ws

每个交易所根据不同包名分割，根目录下启动(部分交易所需要翻墙，建议海外或者香港服务器)
```bash
go run main.go -name="交易所名字" -mq="amqp://user:pwd@host:port/vhost"
```
```go
//开启对应交易所WebSocket连接 
//针对http做了代理之前主要用于本地测试，http接口主要请求交易所支持的所有交易对
huobi.StartWs(""  , false)
zb.StartWs("http代理地址"  , true)
...
```
交易所名称列表：

交易所名称 | 官网
---- | ---
binance | https://www.binance.com/
bitfinex | https://www.bitfinex.com/
fcoin | https://www.fcoin.com/
gateio | https://www.gate.io/
hadax | https://www.hadax.com
hitbtc2 | https://www.hitbtc.com/
huobipro | https://www.hbg.com
lbank | https://www.lbank.info/
okex | https://www.okex.com/
zb | https://www.zb.com/
```

# 项目规划
* 对接比较出名的几家交易所历史交易WebSocket接口（其实K线、涨幅等数据都可以根据历史交易计算得出）
* 断线重连
* 统一的数据格式输出到RabbitMQ :
```json
{"exchangeId":11,"symbol":"NANO/BTC","trades":[{"symbol":"NANO/BTC","side":"buy","amount":"12.84000000","price":"0.00037230","timestamp":"1536910985284"}],"exchange":"binance"}
```

