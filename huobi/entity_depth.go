package huobi

type Dtick struct {
	Bids    [][]float32 `json:"bids"`
	Asks    [][]float32 `json:"asks"`
	Ts      int64       `json:"ts"`
	Version int64       `json:"version"`
}

//响应请求实体
type Depth struct {
	Status string `json:"status"` //请求处理结果	"ok" , "error"
	Ch     string `json:"ch"`     //数据所属的 channel，格式： market.$symbol.kline.$period
	Ts     int64  `json:"ts"`     //响应生成时间点，单位：毫秒
	Tick   Dtick  `json:"tick"`   //数据
}
