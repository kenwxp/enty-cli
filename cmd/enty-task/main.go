package main

import (
	"entysquare/enty-cli/schedule"
	"entysquare/enty-cli/storage"
	"fmt"
)

func main() {
	db, err := storage.NewDatabase()
	if err != nil {
		fmt.Println("err:", err)
		panic("db failed init")
	}
	schedule.StatisticTask(db) // run statistic schedule
}
