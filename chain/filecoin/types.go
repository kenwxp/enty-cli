package filecoin

type MinerResult struct {
	Code    int
	Message string
	Data    FileData
}
type FileData struct {
	Miner                string
	QualityPower         int64
	QualityPowerStr      string
	QualityPowerPercent  float64
	RawPower             int64
	RawPowerStr          string
	RawPowerPercent      int64
	TotalQualityPower    int64
	TotalQualityPowerStr string
	TotalRawPower        int64
	TotalRawPowerStr     string
	Blocks               int
	WinCount             int
	BlockReward          int64
	blockRewardStr       string
	Owner                string
	Worker               string
	Tag                  string
	IsVerified           int
	PeerId               string
	PowerRank            int
	Sectors              FileSectors
	Balance              FileBalance
	Local                FileLocal
}

type FileSectors struct {
	SectorSize    int64
	SectorSizeStr string
	SectorCount   int
	ActiveCount   int
	FaultCount    int
	RecoveryCount int
}

type FileBalance struct {
	Balance          int64
	BalanceStr       string
	Available        int64
	AvailableStr     string
	SectorsPledge    int64
	SectorsPledgeStr string
	LockedFunds      int64
	LockedFundsStr   string
	FeeDebtStr       string
}

type FileLocal struct {
	Ip       string
	Location string
}

/*
	Block 数据
*/
type FilScoutBlockRet struct {
	Message   string `json:"message"`
	Code      int    `json:"code"`
	PageIndex int    `json:"pageIndex"`
	PageBool  bool   `json:"pageBool"` //是否有下一页
	PageSize  int    `json:"pageSize"`
	Total     int    `json:"total"`

	Data []FilScoutBlockOne `json:"data"`
}

type FilScoutBlockOne struct {
	Cid          string `json:"cid"`
	ExactReward  string `json:"exactReward"`
	Height       int    `json:"height"`
	IsVerified   int    `json:"isVerified"`
	MessageCount int    `json:"messageCount"`
	MineTime     string `json:"mineTime"`
	Miner        string `json:"miner"`
	MinerTag     string `json:"minerTag"`
	Reward       string `json:"reward"`
	Size         int    `json:"size"`
}

type MinerDatePower struct {
	Unix     int64
	Date     string
	Power    int64
	PowerStr string
}
