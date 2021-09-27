package filecoin

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"entysquare/enty-cli/storage"
	"entysquare/enty-cli/util"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"time"
	"unsafe"
)

func ScanMiner(db *storage.Database) {
	for {
		fmt.Println("处理 node all 数据")
		var list []FilNodeDataAll
		err := util.WithTransaction(db.Db, func(txn *sql.Tx) error {
			filerPool, err := db.SelectFilerPool()
			if err != nil {
				return err
			}
			for _, v := range filerPool {
				time.Sleep(time.Second * 2)
				json1, err := MinerInfo(v.NodeName)
				if err != nil {

					return err
				}
				json24h, err := MinerInfoBlock(v.NodeName)
				if err != nil {
					return err
				}
				list = append(list, FilNodeDataAll{
					Name:                  v.NodeName,
					QualityPower:          gjson.Get(json1, "data.qualityPower").Float(),
					BlockRewardAll:        gjson.Get(json1, "data.blockReward").Float(),
					BlockReward24:         gjson.Get(json24h, "data.blockReward").Float(),
					MiningEfficiencyFloat: gjson.Get(json24h, "data.miningEfficiencyFloat").Float(),
				})
			}
			return err
		})
		if err == nil {
			FilNodeUrlData = list
		}
		time.Sleep(time.Second * 60 * 2)
	}
}
func MinerInfo(miner string) (jsonstr string, err error) {
	song := make(map[string]interface{})
	bytesData, err := json.Marshal(song)
	if err != nil {
		return "", err
	}
	reader := bytes.NewReader(bytesData)
	url := "https://api.filscout.com/api/v1/miners/"
	url += miner
	url += "/info"
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return "", err
	}
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	str := (*string)(unsafe.Pointer(&respBytes))
	return *str, nil
}

//https://api.filscout.com/api/v1/miners/f01159979/miningstats
func MinerInfoBlock(miner string) (jsonstr string, err error) {
	url := "https://api.filscout.com/api/v1/miners/"
	url += miner
	url += "/miningstats"

	req, err := http.NewRequest("POST", url,
		bytes.NewBuffer([]byte(`{"statsType":"24h"}`)))
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
