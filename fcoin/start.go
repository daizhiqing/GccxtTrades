package fcoin

import (
	"GccxtTrades/utils"

	"github.com/sirupsen/logrus"
)

func StartWs(proxy string, useProxy bool) {
	if useProxy {
		utils.UseProxy = useProxy
		if proxy != "" {
			utils.ProxyUrl = proxy
		}
	}

	resp, err := utils.HttpGet(FCoinSymbol).Get("data").Array()
	if err != nil {
		logrus.Panic(err)
	}
	var symboleList []string
	for _, m := range resp {
		symboleList = append(symboleList, m.(map[string]interface{})["name"].(string))
	}
	logrus.Warnln(symboleList)
	FcoinWsConnect(symboleList)
}
