package huobi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//请求Huobi K线数据
// symbol 交易对 btcusdt, bchbtc, rcneth ...
// period 时间维度 1min, 5min, 15min, 30min, 60min, 1day, 1mon, 1week, 1year
// size 获取数量 [1,1000]
func GetKLine(symbol, period string, size int) KLine {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	if len(symbol) <= 0 {
		panic("huobi:GetKLine symbol is nil")
	}

	if len(symbol) <= 0 {
		panic("huobi:GetKLine period is nil")
	}

	if size <= 0 {
		size = 150
	}
	if size > 1000 {
		size = 1000
	}

	klineUrl := fmt.Sprintf(KLineApi, symbol, period, size)
	resp, err := http.Get(klineUrl)
	log.Println(klineUrl)
	if err != nil {
		log.Println("接口请求失败", klineUrl, symbol, period, size)
		panic(err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var result string = string(body)
	log.Println(result)
	var klineObject KLine
	json.Unmarshal(body, &klineObject)
	log.Println(klineObject)
	return klineObject
}

// 获取聚合行情(Ticker)
//symbol 交易对 btcusdt, bchbtc, rcneth ...
func GetTickerMerged(symbol string) Merged {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	if len(symbol) <= 0 {
		panic("huobi:GetTickerMerged symbol is nil")
	}

	tikerMergerUrl := fmt.Sprintf(TickerMergedApi, symbol)
	resp, err := http.Get(tikerMergerUrl)
	if err != nil {
		log.Println("接口请求失败", tikerMergerUrl, symbol)
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result string = string(body)
	log.Println(result)

	var merged Merged
	json.Unmarshal(body, &merged)
	log.Println(merged)
	return merged
}

func GetDepth(symbol, _type string) Depth {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	if len(symbol) <= 0 {
		panic("huobi:GetDepth symbol is nil")
	}
	if len(_type) <= 0 {
		panic("huobi:GetDepth type is nil")
	}

	depthUrl := fmt.Sprintf(DepthApi, symbol, _type)

	resp, err := http.Get(depthUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var result string = string(body)
	log.Println(result)
	var depth Depth
	json.Unmarshal(body, &depth)
	log.Println(depth)
	return depth
}
