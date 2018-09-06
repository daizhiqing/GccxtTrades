package huobi

import (
	"testing"
)

func TestHttp(t *testing.T) {

	// t.Error(GetKLine("btcusdt", "1day", 1))
	// t.Error(GetTickerMerged("btcusdt"))
	t.Error(GetDepth("btcusdt", "step1"))
}
