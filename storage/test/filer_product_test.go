package main

import (
	"context"
	"database/sql"
	"entysquare/enty-cli/storage"
	"entysquare/enty-cli/storage/types"
	"fmt"
	_ "github.com/lib/pq"
	"testing"
)

func TestInsertFilerProduct(t *testing.T) {
	db, err := storage.NewDatabase(true)
	fmt.Println(db)
	fmt.Println(err)
	err = db.WithTransaction(func(txn *sql.Tx) error {
		err := db.InsertFilerProduct(context.TODO(), txn, &types.FilerProduct{
			ProductId:    "222222", //产品名称
			ProductName:  "13134",  //产品名称
			NodeId:       "13134",  //节点
			CurId:        "13134",  //币种类型
			Period:       "13134",  //挖矿周期
			ValidPlan:    "13134",  //生效时间
			Price:        "13134",  //每T质押
			PledgeMax:    "13134",  //质押需求额
			ServiceRate:  "13134",  //服务费率
			Note1:        "13134",  //说明
			Note2:        "13134",  //说明
			ShelveTime:   "13134",  //上架时间
			CreateTime:   "13134",  //创建时间
			UpdateTime:   "13134",  //更新时间
			ProductState: "13134",  //产品状态 0-进行中 9-已失效
			IsValid:      "13134",  //启用标志 0--启用 1-废弃
		})
		fmt.Println(err)
		return err
	})
	fmt.Println(err)
}
func TestSelectProductInfoById(t *testing.T) {
	db, err := storage.NewDatabase(true)
	fmt.Println(db)
	fmt.Println(err)

	err = db.WithTransaction(func(txn *sql.Tx) error {
		data, err := db.SelectProductInfoById(context.TODO(), txn, "001")
		fmt.Println(data)
		return err
	})

}
