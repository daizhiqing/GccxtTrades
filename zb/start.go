package zb

import (
	"ccxt/utils"
	"log"
	"strings"
)

func StartWs(proxy string, useProxy bool) {

	if useProxy && proxy != "" {
		utils.UseProxy = useProxy
		utils.ProxyUrl = proxy
	}
	var syListZb []string
	zbSym, _ := utils.HttpGet(ZbSymbols).Map()
	for key, _ := range zbSym {
		syListZb = append(syListZb, strings.Replace(key, "_", "", -1))
	}
	log.Println("ZB:", syListZb)
	go ZbWsConnect(syListZb)
}
