package huobi

//具体K线数据
type Data struct {
	Id     int64   `json:"id"`     //K线id
	Open   float64 `json:"open"`   //开盘价
	Close  float64 `json:"close"`  //收盘价,当K线为最晚的一根时，是最新成交价
	Low    float64 `json:"low"`    //最低价
	High   float64 `json:"high"`   // 最高价
	Amount float64 `json:"amount"` //成交量
	Vol    float64 `json:"vol"`    //成交额, 即 sum(每一笔成交价 * 该笔的成交量)
	Count  int64   `json:"count"`  //成交笔数,"count": 103004
}

//响应请求实体
type KLine struct {
	Status string `json:"status"` //请求处理结果	"ok" , "error"
	Ch     string `json:"ch"`     //数据所属的 channel，格式： market.$symbol.kline.$period
	Ts     int64  `json:"ts"`     //响应生成时间点，单位：毫秒
	Data   []Data `json:"data"`   //K线数据
}
