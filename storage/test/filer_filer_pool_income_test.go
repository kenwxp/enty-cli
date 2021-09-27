package main

import (
	"context"
	"database/sql"
	"entysquare/filer-backend/storage"
	"entysquare/filer-backend/storage/types"
	"fmt"
	_ "github.com/lib/pq"
	"testing"
)

func TestInsertFilerPoolIncome(t *testing.T) {
	db, err := storage.NewDatabase(true)
	fmt.Println(db)
	fmt.Println(err)
	err = db.WithTransaction(func(txn *sql.Tx) error {
		err := db.InsertFilerPoolIncome(context.TODO(), txn, &types.FilerPoolIncome{
			NodeId:               "031ffdcb-c9e1-4065-891e-411cecf312ce",
			Balance:              "Balance",
			PledgeSum:            "PledgeSum",
			TotalPower:           "TotalPower",
			TotalIncome:          "TotalIncome",
			FreezeIncome:         "FreezeIncome",
			AvailableIncome:      "AvailableIncome",
			TodayIncomeTotal:     "TodayIncomeTotal",
			TodayIncomeFreeze:    "TodayIncomeFreeze",
			TodayIncomeAvailable: "TodayIncomeAvailable",
			StatTime:             "2021-01-01",
			CreateTime:           "CreateTime",
		})
		fmt.Println(err)
		return err
	})
	fmt.Println(err)
}
func TestSelectFilerPoolIncomeByNodeId(t *testing.T) {
	//db, err := storage.NewDatabase(true)
	//fmt.Println(db)
	//fmt.Println(err)
	//data := db.SelectFilerPoolIncomeByNodeId(context.TODO(), "031ffdcb-c9e1-4065-891e-411cecf312ce")
	//fmt.Println(data)
}
