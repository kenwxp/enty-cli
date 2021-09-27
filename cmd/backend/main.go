package main

import (
	"entysquare/enty-cli/schedule"
	"entysquare/enty-cli/storage"
	"fmt"
)

func main() {
	var localhost bool

	db, err := storage.NewDatabase(localhost)
	if err != nil {
		fmt.Println("err:", err)
		panic("db failed init")
	}
	schedule.StatisticTask(db) // run statistic schedule
}
