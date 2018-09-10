# Gccxt
数字货币交易所WebSocket对接历史交易记录 Go版本实现

目前已经对接交易所(持续更新。。。):

* HuobiPro： wss://api.huobi.br.com/ws
* LBank：    wss://api.lbkex.com/ws/V2/
* ZB :       wss://api.zb.cn:9999/websocket
* Okex :     wss://real.okex.com:10441/websocket
* Bitfinex   wss://api.bitfinex.com/ws
* Bitfinex   wss://api.bitfinex.com/ws
* Gateio     wss://ws.gateio.io/v3/

每个交易所根据不同包名分割，根目录下启动(部分交易所需要翻墙，建议海外或者香港服务器)
```bash
go run main.go
```
```go
//开启对应交易所WebSocket连接 
//针对http做了代理之前主要用于本地测试，http接口主要请求交易所支持的所有交易对
huobi.StartWs(""  , false)
zb.StartWs("http代理地址"  , true)
...
```

# 项目规划
* 对接比较出名的几家交易所历史交易WebSocket接口（其实K线、涨幅等数据都可以根据历史交易计算得出）

* 统一的数据格式输出（待开发）
* 添加数据存储模块Hbase或redis（待开发）
