package main

import (
	"context"
	"entysquare/filer-backend/storage"
	"fmt"
	_ "github.com/lib/pq"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	db, err := storage.NewDatabase(true)
	fmt.Println(err)

	list, err := db.SelectFilerBlockAndTimeByTime(context.TODO(), time.Now().AddDate(0, 0, -30), time.Now())
	for k, v := range list {
		fmt.Println(k, v)
	}

}
