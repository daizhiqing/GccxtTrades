package huobi

type Tick struct {
	Id     int64     `json:"id"`     //"id": K线id,
	Amount float64   `json:"amount"` //"amount": 成交量,
	Count  int64     `json:"count"`  //"count": 成交笔数,
	Open   float64   `json:"open"`   //"open": 开盘价,
	Close  float64   `json:"close"`  //"close": 收盘价,当K线为最晚的一根时，是最新成交价
	Low    float64   `json:"low"`    //"low": 最低价,
	High   float64   `json:"high"`   //"high": 最高价,
	Vol    float64   `json:"vol"`    //"vol": 成交额, 即 sum(每一笔成交价 * 该笔的成交量)
	Bid    []float64 `json:"bid"`    //"bid": [买1价,买1量],
	Ask    []float64 `json:"ask"`    //"ask": [卖1价,卖1量]
}

//响应请求实体
type Merged struct {
	Status string `json:"status"` //请求处理结果	"ok" , "error"
	Ch     string `json:"ch"`     //数据所属的 channel，格式： market.$symbol.kline.$period
	Ts     int64  `json:"ts"`     //响应生成时间点，单位：毫秒
	Tick   Tick   `json:"tick"`   //数据
}
