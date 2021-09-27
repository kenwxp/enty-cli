package filecoin

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"strconv"
	"testing"
)

func Test1(t *testing.T) {
	str, _ := MinerInfo("f024089")
	println(str)
	minerResult := MinerResult{}
	err := json.Unmarshal([]byte(str), &minerResult)
	if err != nil {
		println(err)
	}
	println(minerResult.Code)
	balanceStr := strconv.FormatInt(minerResult.Data.Balance.Balance, 10)
	println(minerResult.Data.Balance.BalanceStr)
	println(balanceStr)
}

func Test2(t *testing.T) {
	jsonStr, err := MinerInfo("f01159979")
	fmt.Println(err)
	fmt.Println(jsonStr)
	fmt.Println("总算力", gjson.Get(jsonStr, "data.qualityPower"))
	fmt.Println("累计出块：", gjson.Get(jsonStr, "data.blockReward"))
	json2, err := MinerInfoBlock("f01159979")
	fmt.Println("24小时出块", gjson.Get(json2, "data.blockReward"))
	fmt.Println("平均出块：", gjson.Get(json2, "data.miningEfficiencyFloat")) //fil

}
