package main

import (
	"entysquare/enty-cli/storage"
	"fmt"
	_ "github.com/lib/pq"
	"testing"
)

func TestInsertFilerPool(t *testing.T) {
	//db, err := storage.NewDatabase(true)
	//fmt.Println(db)
	//fmt.Println(err)
	//err = db.WithTransaction(func(txn *sql.Tx) error {
	//	err := db.InsertFilerPool(context.TODO(), txn, &types.FilerPool{
	//		NodeName:    "f0154944",
	//		Location:    "Location1",
	//		Mobile:      "Mobile1",
	//		Email:       "Email1",
	//		ValidCreate: "ValidCreate1",
	//		IsValid:     "IsValid1",
	//	})
	//	err = db.InsertFilerPool(context.TODO(), txn, &types.FilerPool{
	//		NodeName:    "f0109070",
	//		Location:    "Location1",
	//		Mobile:      "Mobile1",
	//		Email:       "Email1",
	//		ValidCreate: "ValidCreate1",
	//		IsValid:     "IsValid1",
	//	})
	//	fmt.Println(err)
	//	return err
	//})
	////fmt.Println("return:")
	//fmt.Println(err)
}

func TestSelectFilerPool(t *testing.T) {
	db, err := storage.NewDatabase(true)
	fmt.Println("db:", err)
	list, err := db.SelectFilerPool()
	fmt.Println("db select:", err)

	for k, v := range list {
		fmt.Println(k, v)
	}
}
