package filecoin

type FilNodeDataAll struct {
	Name                  string
	QualityPower          float64
	BlockRewardAll        float64
	BlockReward24         float64
	MiningEfficiencyFloat float64
}

var FilNodeUrlData []FilNodeDataAll

//算力数据
var FilNodeUrlPowers []MinerDatePower
