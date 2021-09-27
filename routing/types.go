package routing

type RegisterRequest struct {
	PayID int64 `json:"payId"`
}

type LoginRequest struct {
	PayID int64 `json:"payId"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
type ProductInfoRequest struct {
	ProductID string `json:"productId"`
}
type ProductInfoResponse struct {
	ProductID      string `json:"productId"`
	Currency       string `json:"currency"`
	State          string `json:"state"`
	Name           string `json:"name"`
	PriceService   string `json:"priceService"`
	Note1          string `json:"note1"`
	Note2          string `json:"note2"`
	CycleDay       string `json:"cycleDay"`
	GuessProfitDay string `json:"guessProfitDay"`
	PricePerT      string `json:"pricePerT"` //
}
type ProductListRequest struct {
	State string `json:"state"`
}
type ProductResponse struct {
	ProductID   string `json:"productId"`
	Currency    string `json:"currency"`
	Name        string `json:"name"`
	PricePerT   string `json:"pricePerT"`
	PledgeLimit string `json:"pledgeLimit"`
	CycleDay    string `json:"cycleDay"`
	State       string `json:"state"`
}
type ProductListResponse struct {
	List []ProductResponse `json:"list"`
}
type OrderListRequest struct {
	State int `json:"state"`
}
type OrderResponse struct {
	OrderId              string `json:"orderId"`              //订单id
	PayFlow              string `json:"payFlow"`              //支付流水号
	HoldPower            string `json:"holdPower"`            //持有算力
	PayAmount            string `json:"payAmount"`            //金额质押
	OrderTime            string `json:"orderTime"`            //下单时间 时间戳
	ValidTime            string `json:"validTime"`            //生效时间 时间戳
	OrderState           string `json:"orderState"`           //订单状态 0-待生效 1-已生效 2-持仓已结束 9-收益结算完成
	ProductName          string `json:"productName"`          //产品名
	Period               string `json:"period"`               //周期
	ValidDays            string `json:"validDays"`            //已生效天数
	TotalIncome          string `json:"totalIncome"`          //累计收益
	TotalAvailableIncome string `json:"totalAvailableIncome"` //累计已释放收益
	FreezeIncome         string `json:"freezeIncome"`         //待释放收益
	StatTime             string `json:"statTime"`             //统计日期 yyyy-mm-dd
	DayAvailableIncome   string `json:"dayAvailableIncome"`   //当日释放
	DayRaiseIncome       string `json:"dayRaiseIncome"`       //当日产出
	DayDirectIncome      string `json:"dayDirectIncome"`      //当日直接释放
	DayReleaseIncome     string `json:"dayReleaseIncome"`     //当日线性释放
}
type OrderListResponse struct {
	List []OrderResponse `json:"list"`
}
type ProfitLogResponse struct {
	Uuid                 string `json:"uuid"`
	PledgeSum            string `json:"pledgeSum"`
	HoldPower            string `json:"holdPower"`
	TotalIncome          string `json:"totalIncome"`
	FreezeIncome         string `json:"freezeIncome"`
	TotalAvailableIncome string `json:"totalAvailableIncome"`
	Balance              string `json:"balance"`
	DayAvailableIncome   string `json:"dayAvailableIncome"`
	DayRaiseIncome       string `json:"dayRaiseIncome"`
	DayDirectIncome      string `json:"dayDirectIncome"`
	DayReleaseIncome     string `json:"dayReleaseIncome"`
	StatTime             string `json:"statTime"`
}
type ProfitLogListResponse struct {
	List []ProfitLogResponse `json:"list"`
}

type UserFilResponse struct {
	ProfitAll       string `json:"profitAll"`
	ProfitBalance   string `json:"profitBalance"`
	ProfitAvailable string `json:"profitAvailable"`
	ProfitLock      string `json:"profitLock"`
	ProfitToday     string `json:"profitToday"`
	PledgeAll       string `json:"pledgeAll"`
	Power           string `json:"power"`
}
type HtFilNodeResponse struct {
	PowerAll      string `json:"powerAll"`
	Produce24h    string `json:"produce24h"`
	Produce24hAvg string `json:"produce24hAvg"`
	ProduceAll    string `json:"produceAll"`

	Blocks []BlockResponse `json:"blocks"`
}
type BlockResponse struct {
	Date     string `json:"date"`
	BlockNum string `json:"blockNum"`
}
