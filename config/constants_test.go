package config

import "testing"

func TestGetExchangeId(t *testing.T) {
	val := GetExchangeId("lbank")
	if val <= 0 {
		t.Error("获取失败")
	}

}
