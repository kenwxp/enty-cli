package main

import (
	"context"
	"database/sql"
	"entysquare/filer-backend/routing"
	"entysquare/filer-backend/schedule"
	"entysquare/filer-backend/storage"
	"entysquare/filer-backend/storage/types"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	var localhost bool

	//var cli *payment.PayClient
	routers := mux.NewRouter()

	db, err := storage.NewDatabase(localhost)
	if err != nil {
		fmt.Println("err:", err)
		panic("db failed init")
	}
	env, bool := os.LookupEnv("MODE")
	if !bool {
		fmt.Println("environment MODE not set")
	}
	if env == "dev" {
		err = injectMockDB(db)
		if err != nil {
			fmt.Println("err:", err)
			panic("db can not inject")
		}

	}
	schedule.StatisticTask(db) // run statistic schedule

	routing.Setup(routers, db)
	err = http.ListenAndServe("0.0.0.0:9000", routers)
	if err != nil {
		panic("error")
	}
}

func injectMockDB(db *storage.Database) (err error) {
	ctx := context.TODO()
	err = db.WithTransaction(func(txn *sql.Tx) error {

		db.InsertFilerPoolIncome(ctx, txn, &types.FilerPoolIncome{
			NodeId:               "1",
			Balance:              "2000000000",
			PledgeSum:            "1",
			TotalPower:           "1",
			TotalIncome:          "1000000000",
			FreezeIncome:         "500000000",
			AvailableIncome:      "500000000",
			TodayIncomeTotal:     "10000000",
			TodayIncomeFreeze:    "5000000",
			TodayIncomeAvailable: "5000000",
			StatTime:             "2021-09-01",
		})
		db.InsertFilerProduct(ctx, txn, &types.FilerProduct{
			ProductId:   "1",
			ProductName: "test1",
			NodeId:      "1",
			CurId:       "fil",
			Period:      "540",
			ValidPlan:   "1",
			Price:       "1",
			PledgeMax:   "1",
			ServiceRate: "40%",
			Note1:       "1",
			Note2:       "1",
			ShelveTime:  "1",
		})
		db.InsertFilerProduct(ctx, txn, &types.FilerProduct{
			ProductId:   "2",
			ProductName: "test2",
			NodeId:      "1",
			CurId:       "fil",
			Period:      "360",
			ValidPlan:   "1",
			Price:       "1",
			PledgeMax:   "1",
			ServiceRate: "30%",
			Note1:       "1",
			Note2:       "1",
			ShelveTime:  "1",
		})
		if err != nil {
			fmt.Println("err:", err)
			return err
		}
		return nil
	})
	return nil
}
