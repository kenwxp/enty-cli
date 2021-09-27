package main

import (
	"entysquare/enty-cli/storage"
	"fmt"
	_ "github.com/lib/pq"
	"testing"
)

func TestSelectAllInProgressOrder(t *testing.T) {
	db, err := storage.NewDatabase(true)
	fmt.Println(db)
	fmt.Println(err)

	//list, err := db.SelectAllInProgressOrder("2021-09-02")
	//fmt.Println("return:")
	//fmt.Println("list:", list)
}
