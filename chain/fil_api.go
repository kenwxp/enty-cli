package chain

import (
	"bytes"
	"context"
	"encoding/json"
	"entysquare/enty-cli/chain/filecoin"
	"entysquare/enty-cli/storage"
	"entysquare/enty-cli/storage/types"
	"entysquare/enty-cli/util"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const FileApiUrl = "https://1xR5g3Ym4VwGLmE1BbUJJ1zGJGP:94d7a0fa8dccd26eae3f7346e5e89056@filecoin.infura.io"

/*
	获取区块高度
*/
func GetBlockHeight() (int64, error) {
	req, err := http.NewRequest("POST", FileApiUrl,
		bytes.NewBuffer([]byte(`{ "jsonrpc": "2.0", "method":"Filecoin.ChainHead", "params": [], "id": 3 }`)))
	if err != nil {
		return 0, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return 0, err
	}

	value := gjson.Get(string(body), "result.Height")
	if value.Int() <= 0 {
		return 0, fmt.Errorf("result.Height err 001")
	}

	return value.Int(), nil
}

/*
	获取fil 节点算力
*/
func GetStateMinerPowerByTiB(nodeName string) (int64, error) {
	req, err := http.NewRequest("POST", FileApiUrl,
		bytes.NewBuffer([]byte(`{"id": 0,"jsonrpc": "2.0","method": "Filecoin.StateMinerPower","params": [ "`+nodeName+`",null ]}`)))
	if err != nil {
		return 0, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return 0, err
	}
	power := gjson.Get(string(body), "result.MinerPower.RawBytePower")
	tb := power.Int() / 1024 / 1024 / 1024 / 1024
	return tb, nil
}

/*
	获取需要扫块的高度
	没有则是最新区块-60
*/
func GetBlockScanningHeight(db *storage.Database) (int64, error) {
	var ctx = context.TODO()

	//优先取db记录接着扫
	i, _ := db.SelectFilerBlockTempMaxBlockHeight(ctx)
	if i != 0 {
		return i + 1, nil
	}
	//否则取最新区块数据扫
	num, err := GetBlockHeight()
	return num - 60, err
}

/*
	获得指定区块高度数据
	return json,err
*/
func GetBlockHeightData(height int64) (string, error) {
	req, err := http.NewRequest("POST", FileApiUrl,
		bytes.NewBuffer([]byte(`{"id": 0,"jsonrpc": "2.0","method": "Filecoin.ChainGetTipSetByHeight","params": [ `+strconv.FormatInt(height, 10)+`,null ]}`)))
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}

	return string(body), nil
}

/*
	www.filscout.com 获取出块数据
	return json,err
*/
func GetFilScoutData(apiUrl, minerId string, pageIndex int) (*filecoin.FilScoutBlockRet, error) {
	time.Sleep(time.Second * 2)
	req, err := http.NewRequest("POST", apiUrl,
		bytes.NewBuffer([]byte(`{
       "miner":"`+minerId+`",
       "pageIndex":`+strconv.Itoa(pageIndex)+`,
       "pageSize":100
		}`)))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var data = filecoin.FilScoutBlockRet{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	if float64(data.PageIndex) < float64(data.Total)/100.0 {
		data.PageBool = true
	}

	return &data, nil
}

/*
	www.filscout.com 算力数据
	dayStr = 30d ，180d
	return json,err
*/
func GetFilScoutPowerStatsData(minerId string, dayStr string) (string, error) {
	time.Sleep(time.Second * 2)
	//https://api.filscout.com/api/v1/miners/f0166425/powerstats
	req, err := http.NewRequest("POST", "https://api.filscout.com/api/v1/miners/"+minerId+"/powerstats",
		bytes.NewBuffer([]byte(`{"statsType":"`+dayStr+`"}`)))
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	return string(body), nil
}

/*
	获取算力记录
	前180天平均记录 + 前30天 精确记录
*/
func GetFilScoutPowerList(minerId string) ([]filecoin.MinerDatePower, error) {
	var list []filecoin.MinerDatePower

	json180d, err := GetFilScoutPowerStatsData(minerId, "180d")
	if err != nil {
		return nil, err
	}
	json30d, err := GetFilScoutPowerStatsData(minerId, "30d")
	if err != nil {
		return nil, err
	}
	arr180 := gjson.Get(json180d, "data.powers").Array()
	arr30 := gjson.Get(json30d, "data.powers").Array()
	if len(arr180) == 0 || len(arr30) == 0 {
		return nil, err
	}
	for _, v := range arr180 {
		if v.Get("heightTime").Int() >= arr30[0].Get("heightTime").Int() {
			break
		}
		list = append(list, filecoin.MinerDatePower{
			Unix:     v.Get("heightTime").Int(),
			Date:     v.Get("heightTimeStr").String()[0:10],
			Power:    v.Get("rawPower").Int(),
			PowerStr: v.Get("rawPowerStr").String(),
		})
	}
	for _, v := range arr30 {
		list = append(list, filecoin.MinerDatePower{
			Unix:     v.Get("heightTime").Int(),
			Date:     v.Get("heightTimeStr").String()[0:10],
			Power:    v.Get("rawPower").Int(),
			PowerStr: v.Get("rawPowerStr").String(),
		})
	}
	return list, nil
}

/*
	从历史数据得出 算力
	历史数据没有则倒退模糊计算
*/
func GetFilPowerByMinerDatePower(minerId string, checkUnix int64, list []filecoin.MinerDatePower) (int64, error) {
	var checkTime = time.Unix(checkUnix, 0)

	for i, v := range list {
		//精确数据
		if checkTime.Format("2006-01-02") == v.Date {
			return v.Power / 1024 / 1024 / 1024 / 1024, nil
		}
		//相对数据
		if v.Unix >= checkUnix {
			if i == 0 {
				return v.Power / 1024 / 1024 / 1024 / 1024, nil
			}

			//精确计算
			var before = list[i-1] //得出之前的数据

			//计算日期差值
			var dataDiffer = (v.Unix - before.Unix) / (60 * 60 * 24)
			var checkDiffer = (checkUnix - before.Unix) / (60 * 60 * 24)
			var timeDiffer = float64(checkDiffer) / float64(dataDiffer)

			//计算算力差值
			powerDiffer := v.Power - before.Power

			//算力差值 * 时间差值 算力振幅数量
			powerIncrease := util.MultiplicationPrec0(powerDiffer, timeDiffer)
			powerIncreaseInt64, err := strconv.ParseInt(powerIncrease, 10, 64)
			if err != nil {
				return 0, err
			}
			return (before.Power + powerIncreaseInt64) / 1024 / 1024 / 1024 / 1024, err
		}
	}
	return GetStateMinerPowerByTiB(minerId)
}

/*
	检查出块者是否包含
	jsonStr 链数据json
	myMiners 需要扫的节点地址
	rewardNum 区块奖励数量 24.14 fil

	是 返回出块记录
	否 空出块记录
	err 异常
*/
func handleBlockMiner(jsonStr string, myMiners []types.FilerPool, rewardNum string) ([]types.FilerBlockTemp, error) {
	var list []types.FilerBlockTemp

	//判断区块高度是否异常
	blockHeight := gjson.Get(jsonStr, "result.Height").Int()
	if blockHeight <= 0 {
		return nil, fmt.Errorf("checkBlockMiner err! height <= 0")
	}
	//开扫
	for _, block := range gjson.Get(jsonStr, "result.Blocks").Array() {

		//出块者
		toMiner := block.Get("Miner").String()
		fmt.Println("爆块者:", toMiner)

		//不是内部节点出块
		if miner := checkMiner(toMiner, myMiners); miner != nil {
			var winCount = block.Get("ElectionProof.WinCount").String()
			var blockGain = ""
			if winCount != "" {
				blockGain = util.Multiplication(rewardNum, winCount)
			}
			tib, _ := GetStateMinerPowerByTiB(miner.NodeName)
			list = append(list, types.FilerBlockTemp{
				Power:       strconv.FormatInt(tib, 10),
				NodeId:      miner.NodeId,
				BlockNum:    winCount,  //出块数量
				BlockGain:   blockGain, //出块收益总数
				BlockHeight: strconv.FormatInt(blockHeight, 10),
				ChainTime:   block.Get("Timestamp").String(),
			})

		}
	}
	return list, nil
}

func checkMiner(toMiner string, myMiners []types.FilerPool) *types.FilerPool {
	for _, v := range myMiners {
		if v.NodeName == toMiner {
			return &v
		}
	}
	return nil
}
