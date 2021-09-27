package chain

import (
	"entysquare/enty-cli/chain/filecoin"
	"entysquare/enty-cli/storage"
	"entysquare/enty-cli/util"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
	"testing"
	"time"
)

func TestFil(t *testing.T) {
	num, err := GetBlockHeight()
	fmt.Println(num)
	fmt.Println(err)
	//return *str,nil

	GetBlockHeightData(1073500)
}

func TestFilProfitScanning(t *testing.T) {
	FilProfitScanning()
}

func TestFilBlockTempStatistics(t *testing.T) {
	db, _ := storage.NewDatabase(true)
	FilProfitScanningByFilScout(db)
}

func TestGetStateMinerPower(t *testing.T) {
	//tib, err := GetStateMinerPowerByTiB("f02528") //算力: 24.04 PiB
	//fmt.Println(tib)
	//fmt.Println(err)
	db, _ := storage.NewDatabase(true)
	SetFilNodeUrlPowersServer(db)

	//for k,v := range list{
	//	fmt.Println(k,v)
	//}
}

func TestGetFilScoutData(t *testing.T) {
	var checkTime int64 = 1619336483
	var checkTime2 = time.Unix(checkTime, 0)

	fmt.Println(checkTime2)
	list, err := GetFilScoutPowerList("f0680529")
	fmt.Println(err)
	var f *filecoin.MinerDatePower
	for k, v := range list {
		//精确数据
		if checkTime2.Format("2006-01-02") == v.Date {
			fmt.Println(k, "^^^", v)
			f = &v
			continue
		}
		//相对数据
		if v.Unix >= checkTime {
			fmt.Println(k, "===", v)
			f = &v
			//精确计算
			var before = list[k-1] //得出之前的数据
			var dataDiffer = (v.Unix - before.Unix) / (60 * 60 * 24)
			var checkDiffer = (checkTime - before.Unix) / (60 * 60 * 24)
			fmt.Println("数据 相差天数：", dataDiffer)
			fmt.Println("查询 相差天数：", checkDiffer)
			fmt.Println("时间相差比例：", float64(checkDiffer)/float64(dataDiffer))
			powerDiffer := v.Power - before.Power
			i1, str := util.FileSizeStr(float64(powerDiffer))
			fmt.Println("相差算力：", i1, str)

			powerIncrease := util.MultiplicationPrec0(powerDiffer, float64(checkDiffer)/float64(dataDiffer))
			powerIncreaseInt64, _ := strconv.ParseInt(powerIncrease, 10, 64)
			fmt.Println("算力振幅：", before.Power+powerIncreaseInt64)

			fileSize, fileCompany := util.FileSizeStr(float64(before.Power + powerIncreaseInt64))
			f = &filecoin.MinerDatePower{
				Unix:     checkTime,
				Date:     checkTime2.Format("2006-01-02"),
				Power:    before.Power + powerIncreaseInt64,
				PowerStr: fmt.Sprintf("%.2f", fileSize) + " " + fileCompany,
			}
			break
		}
		fmt.Println(k, v)
	}

	fmt.Println("ffff")
	fmt.Println(f)
	if f == nil {
		fmt.Println("没数据 取当前算力数据")
	}
}
