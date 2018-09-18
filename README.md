# Gccxt
数字货币交易所WebSocket对接历史交易记录 Go版本实现，主要对接各大交易所的历史交易单数据，提供到消息队列进行统一数据分析

目前已经对接交易所(有时间的话会持续更新。。。):

交易所名称 | WebSocket | 官网
--- | ----- | -----
 huobipro |  wss://api.huobi.br.com/ws           | https://www.hbg.com
 lbank    |  wss://api.lbkex.com/ws/V2/          | https://www.lbank.info/  
 zb       |  wss://api.zb.cn:9999/websocket      | https://www.zb.com/
 okex     |  wss://real.okex.com:10441/websocket | https://www.okex.com/
 bitfinex |  wss://api.bitfinex.com/ws           | https://www.binance.com/
 binance  |  wss://stream.binance.com:9443/stream| https://www.bitfinex.com/
 gateio   |  wss://ws.gateio.io/v3/              | https://www.gate.io/
 fcoin    |  wss://api.fcoin.com/v2/ws           | https://www.fcoin.com/
 hitbtc2  |  wss://api.hitbtc.com/api/2/ws       | https://www.hitbtc.com/
 hadax    |  wss://api.hadax.com/ws              | https://www.hadax.com

每个交易所根据不同包名分割，根目录下启动(部分交易所需要翻墙，建议海外或者香港服务器)
```bash
go run main.go -name="交易所名字" -mq="amqp://user:pwd@host:port/vhost"
```

# 项目规划
* 对接比较出名的几家交易所历史交易WebSocket接口（其实K线、涨幅等数据都可以根据历史交易计算得出）
* 断线重连
* 统一的数据格式输出到RabbitMQ :
```json
{"exchangeId":11,"symbol":"NANO/BTC","trades":[{"symbol":"NANO/BTC","side":"buy","amount":"12.84000000","price":"0.00037230","timestamp":"1536910985284"}],"exchange":"binance"}
```

