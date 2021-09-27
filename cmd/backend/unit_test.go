package main

import (
	"encoding/json"
	"entysquare/enty-cli/storage/types"
	"entysquare/enty-cli/util"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestAllHandler_ServeHTTP(t *testing.T) {
	fmt.Println("123")
}

func Test123(t *testing.T) {
	fmt.Println("456")
}

func TestTime(t *testing.T) {
	// 1 billion seconds of Unix, three ways.
	now := util.TimeNow()
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
}

func TestInsert(t *testing.T) {
	vi, _ := strconv.Atoi(strconv.FormatInt(100, 10))
	ci, _ := strconv.Atoi(strconv.FormatInt(1, 10))
	uf := types.UserFinance{
		UserId:      1,
		ProjectId:   1,
		InvestValue: vi,
		InvestClass: ci,
		ShareOut:    0,
		Flag:        1,
		OrderState:  1,
		OrderNum:    "1",
		OrderSec:    "",
		CreateTimes: 1,
	}
	str := ""
	marshal, _ := json.Marshal(uf)
	for i := 0; i < len(marshal); i++ {
		str += string(marshal[i])
	}
	println(str)
}
