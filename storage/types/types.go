package types

type FilerAccountInfo struct {
	FilerId   string
	FilerName string
	RegTime   string
	Mobile    string
	Email     string
	IsValid   string
}

//爆块信息临时表
type FilerBlockTemp struct {
	UuId        string //uuid
	NodeId      string //节点id
	BlockNum    string //爆块数
	BlockGain   string //爆块奖励 FIL
	BlockHeight string //区块高度
	Power       string //算力
	State       string
	ChainTime   string //链时间
	CreateTime  string //创建时间
	UpdateTime  string //更新时间

	NodeName string //临时使用 节点名称
}

//爆块信息表
type FilerBlockIncome struct {
	UuId       string //uu id
	NodeId     string //节点id
	BlockNum   string //爆块数
	BlockGain  string //爆块奖励 FIL
	Power      string //总算力
	GainPerTib string //每T收益
	State      string //处理标志 0-待处理 1-线性释放 9-已处理
	StatTime   string //统计ID(yyyy-MM-dd)
	CreateTime string //创建时间
	UpdateTime string //更新时间
}

//产品信息表
type FilerProduct struct {
	ProductId    string //产品ID
	ProductName  string //产品名称
	NodeId       string //节点
	CurId        string //币种类型
	Period       string //挖矿周期
	ValidPlan    string //生效时间
	Price        string //每T质押
	PledgeMax    string //质押需求额
	ServiceRate  string //服务费率
	Note1        string
	Note2        string
	ShelveTime   string //上架时间
	CreateTime   string //创建时间
	UpdateTime   string //更新时间
	ProductState string //产品状态 0-进行中 9-已失效
	IsValid      string //启用标志 0--启用 1-废弃
}

//持仓信息表
type FilerOrder struct {
	OrderId     string //	订单id
	FilerId     string //	filer id
	PayFlow     string //	交易流水（pay产生）
	ProductId   string //	产品id
	NodeId      string //   节点 id
	Period      string //	挖矿周期 (冗余）
	ValidPlan   string //	生效时间 (冗余）
	ServiceRate string //	服务费率 (冗余）
	HoldPower   string //	持有算力
	PayAmount   string //	支付金额
	OrderTime   string //	下单时间
	UpdateTime  string //   更新时间
	ValidTime   string //   生效时间
	EndTime     string //	结束时间
	OrderState  string //	持仓状态	0-待生效 1-已生效 9-失败
}

//持仓信息表
type FilerOrderShow struct {
	OrderId              string //订单id
	PayFlow              string //支付流水号
	HoldPower            string //持有算力
	PayAmount            string //金额质押
	OrderTime            string //下单时间 时间戳
	ValidTime            string //生效时间 时间戳
	OrderState           string //订单状态 0-待生效 1-已生效 2-持仓已结束 9-收益结算完成
	ProductName          string //产品名
	Period               string //周期
	ValidDays            string //已生效天数
	TotalIncome          string //累计收益
	TotalAvailableIncome string //累计已释放收益
	FreezeIncome         string //待释放收益
	StatTime             string //统计日期 yyyy-mm-dd
	DayAvailableIncome   string //当日释放
	DayRaiseIncome       string //当日产出
	DayDirectIncome      string //当日直接释放
	DayReleaseIncome     string //当日线性释放
}

//持仓信息平账表
type FilerOrderIncome struct {
	Uuid                 string //	uuid
	OrderId              string //	订单id
	FilerId              string //	FilerId
	NodeId               string //	节点
	TotalIncome          string //	总收益
	FreezeIncome         string //	冻结收益
	TotalAvailableIncome string //	累计已释放收益 = 昨日累计收益+当日释放收益
	DayAvailableIncome   string //	可用收益(当日直接25%+往日75%平账 ）
	DayRaiseIncome       string //日产出收益（当日直接25%+当日冻结75%） |
	DayDirectIncome      string //  当日直接收益（25%）	|
	DayReleaseIncome     string //  当日线性释放收益（往日75%平账）	|
	StatTime             string //	统计id （yyyy-MM-dd）
	CreateTime           string //	创建时间
	UpdateTime           string //	更新时间
}

//矿池信息
type FilerPool struct {
	NodeId     string //	节点id	pk
	NodeName   string //	节点名
	Location   string //   位置
	Mobile     string //	手机
	Email      string //	邮箱
	CreateTime string //	启用时间
	UpdateTime string //	启用时间
	IsValid    string //	启用标志
}

type FilerPoolIncome struct {
	UuId                 string // pk
	NodeId               string // 节点id
	Balance              string // 矿池余额
	PledgeSum            string // 质押金额
	TotalPower           string // 总算力
	TotalIncome          string // 总收益*
	FreezeIncome         string // 冻结收益*
	AvailableIncome      string // 可用收益*
	TodayIncomeTotal     string // 今日产出收益
	TodayIncomeFreeze    string // 今日冻结收益
	TodayIncomeAvailable string // 今日可用收益
	StatTime             string // 统计时间（yyyy-MM-dd）
	CreateTime           string // 创建时间
}

type FilerAccountIncome struct {
	Uuid                 string //uu_id
	FilerId              string //filer id
	NodeId               string //node id
	PledgeSum            string //质押金额
	HoldPower            string //持有算力
	TotalIncome          string //总收益
	FreezeIncome         string //冻结收益
	TotalAvailableIncome string //	累计已释放收益 = 昨日累计收益+当日释放收益
	DayAvailableIncome   string //	可用收益(当日直接25%+往日75%平账 ）
	DayRaiseIncome       string //日产出收益（当日直接25%+当日冻结75%） |
	DayDirectIncome      string //  当日直接收益（25%）	|
	DayReleaseIncome     string //  当日线性释放收益（往日75%平账）	|
	StatTime             string //统计时间（yyyy-MM-dd)
	CreateTime           string //创建时间 时间戳
}

type FilerBalanceIncome struct {
	Uuid                 string //uu_id
	FilerId              string //filer id
	FilerName            string
	PledgeSum            string //质押金额
	HoldPower            string //持有算力
	TotalIncome          string //总收益
	FreezeIncome         string //冻结收益
	TotalAvailableIncome string //	累计已释放收益 = 昨日累计收益+当日释放收益
	Balance              string //余额 =昨日余额+可用收益(当日直接25%+往日75%平账 +手工入账-提币出账
	DayAvailableIncome   string //	可用收益(当日直接25%+往日75%平账 ）
	DayRaiseIncome       string //日产出收益（当日直接25%+当日冻结75%） |
	DayDirectIncome      string //  当日直接收益（25%）	|
	DayReleaseIncome     string //  当日线性释放收益（往日75%平账）	|
	StatTime             string //统计时间（yyyy-MM-dd)
	CreateTime           string //创建时间 时间戳
	UpdateTime           string //更新时间 时间戳
}
type FilerBalanceFlow struct {
	Uuid       string // uuid
	FilerId    string // filer id
	OperType   string // 操作类型 0-收益入账 1-手工入账 2-提现出账
	Amount     string // 金额
	CreateTime string // 创建时间 时间戳
}

type FilerStatControl struct {
	StatType   string //统计类型 TASK_BLOCK_STAT,TASK_ORDER_STAT,TASK_ACCOUNT_STAT,TASK_POOL_STAT
	StatTime   string //统计时间（yyyy-MM-dd)
	NodeId     string //统计节点
	StatState  string //统计状态 0-新增 1-进行中，2-成功 3-失败
	CreateTime string //创建时间
	UpdateTime string //更新时间
	Message    string //信息
}

type BlockRecord struct {
	Num  string
	Time string
}

/*
	因为是字符串类型 初始化结构体赋予默认"0"参数
*/
func NewFilerPoolIncome() *FilerPoolIncome {
	return &FilerPoolIncome{
		UuId:                 "",
		NodeId:               "",
		Balance:              "0",
		PledgeSum:            "0",
		TotalPower:           "0",
		TotalIncome:          "0",
		FreezeIncome:         "0",
		AvailableIncome:      "0",
		TodayIncomeTotal:     "0",
		TodayIncomeFreeze:    "0",
		TodayIncomeAvailable: "0",
		StatTime:             "",
		CreateTime:           "0",
	}
}
