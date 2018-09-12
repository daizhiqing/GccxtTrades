package utils

import (
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	SYMBOL_USDT = "USDT"
	SYMBOL_USD  = "USD"
	SYMBOL_BTC  = "BTC"
	SYMBOL_ETH  = "ETH"
	SYMBOL_BNB  = "BNB"
	SYMBOL_EUR  = "EUR"
	SYMBOL_HT   = "HT"
	SYMBOL_OKB  = "OKB"
	SYMBOL_ZB   = "ZB"
	SYMBOL_QC   = "QC"
	SYMBOL_KRW  = "KRW"
	SYMBOL_QTUM = "QTUM"
	SYMBOL_KCS  = "KCS"
	SYMBOL_NEO  = "NEO"
	SYMBOL_CNY  = "CNY"
)

func FmtSymbol(s string) (string, string) {
	if len(s) <= 0 {
		return "", ""
	}
	s = strings.Replace(s, "_", "", -1)
	s = strings.Replace(s, "-", "", -1)
	s = strings.Replace(s, "/", "", -1)
	upStr := strings.ToUpper(s)

	var currency, suffix string
	if strings.HasSuffix(upStr, SYMBOL_USDT) {
		suffix = SYMBOL_USDT
	} else if strings.HasSuffix(upStr, SYMBOL_USD) {
		suffix = SYMBOL_USD
	} else if strings.HasSuffix(upStr, SYMBOL_CNY) {
		suffix = SYMBOL_CNY
	} else if strings.HasSuffix(upStr, SYMBOL_BTC) {
		suffix = SYMBOL_BTC
	} else if strings.HasSuffix(upStr, SYMBOL_ETH) {
		suffix = SYMBOL_ETH
	} else if strings.HasSuffix(upStr, SYMBOL_BNB) {
		suffix = SYMBOL_BNB
	} else if strings.HasSuffix(upStr, SYMBOL_EUR) {
		suffix = SYMBOL_EUR
	} else if strings.HasSuffix(upStr, SYMBOL_HT) {
		suffix = SYMBOL_HT
	} else if strings.HasSuffix(upStr, SYMBOL_OKB) {
		suffix = SYMBOL_OKB
	} else if strings.HasSuffix(upStr, SYMBOL_ZB) {
		suffix = SYMBOL_ZB
	} else if strings.HasSuffix(upStr, SYMBOL_QC) {
		suffix = SYMBOL_QC
	} else if strings.HasSuffix(upStr, SYMBOL_KRW) {
		suffix = SYMBOL_KRW
	} else if strings.HasSuffix(upStr, SYMBOL_QTUM) {
		suffix = SYMBOL_QTUM
	} else if strings.HasSuffix(upStr, SYMBOL_KCS) {
		suffix = SYMBOL_KCS
	} else if strings.HasSuffix(upStr, SYMBOL_NEO) {
		suffix = SYMBOL_NEO
	} else {
		logrus.Errorln("未标记的交易对：" + s)
		return "", ""
	}

	currency = string([]rune(upStr)[:(len(upStr) - len(suffix))])

	return currency, suffix
}
