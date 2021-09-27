package main

import (
	"context"
	"database/sql"
	"entysquare/enty-cli/storage"
	"entysquare/enty-cli/storage/types"
	"entysquare/enty-cli/util"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
	"testing"
)

func TestInsertFilerBlockTemp(t *testing.T) {
	db, err := storage.NewDatabase(true)
	fmt.Println(db)
	fmt.Println(err)
	err = db.WithTransaction(func(txn *sql.Tx) error {
		err := db.InsertFilerBlockTemp(context.TODO(), txn, &types.FilerBlockTemp{
			NodeId:      "403050",
			BlockNum:    "123",
			BlockGain:   "321",
			BlockHeight: "1234",
			Power:       "567",
			ChainTime:   strconv.FormatInt(util.TimeNow().Unix(), 10),
		})
		fmt.Println(err)
		return err
	})
	//fmt.Println("return:")
	fmt.Println(err)
}

func TestInsertFilerBlockTemp2(t *testing.T) {
	//now := strconv.FormatInt(util.TimeNow().Unix(), 10)
	//fmt.Println("now:", now)
	//
	//now2 := strconv.FormatInt(util.TimeNow().AddDate(0, 0, -1).Unix(), 10)
	//fmt.Println("now2:", now2)
	//
	//t2, err := time.ParseInLocation("2006-01-02 15:04:05",
	//	util.TimeNow().AddDate(0, 0, -1).Format("2006-01-02 ")+"00:00:01", time.Local)
	//fmt.Println("t2", t2, err)
	//t1, err := time.Parse("2006-01-02 15:04:05",
	//	util.TimeNow().AddDate(0, 0, -1).Format("2006-01-02 ")+"00:00:01")
	//fmt.Println("t1", t1, err)
	//str1 := util.TimeNow().AddDate(0, 0, -1).Format("2006-01-02 ")
	//fmt.Println("str1", str1)
	//
	//now3 := util.TimeNow()
	//fmt.Println("now3", now3)
	//
	//now4 := time.Now()
	//fmt.Println("now4", now4)
	//
	//fmt.Println(time.Unix(1630511999, 0).Format("2006-01-02 15:04:05"))
	//
	//fmt.Println("~~~~~~~~~~~~~~")
	//start, end := util.TimeDayUnix(util.TimeNow().AddDate(0, 0, -1))
	//fmt.Println("start", start.Unix())
	//fmt.Println("end", end.Unix())
	time := util.TimeStringToTime("2021-09-02", "00:00:00", "")
	fmt.Println("time", time)
}

func TestSelectStatisticBlockInfo(t *testing.T) {
	db, err := storage.NewDatabase(true)
	fmt.Println(db)
	fmt.Println(err)
	//beginTime, _ := util.TimeDayUnix(util.TimeNow().AddDate(0, 0, 1))
	//list, err := db.SelectStatisticBlockInfo(beginTime)
	//fmt.Println("return:")
	//fmt.Println("list:", list)
}
