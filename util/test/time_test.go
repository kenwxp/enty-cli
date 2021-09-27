package main

import (
	"entysquare/filer-backend/util"
	"fmt"
	"testing"
)

func TestTime(t *testing.T) {
	//没生效
	data := util.HandleXP("2021-09-02", 3, 365, "2021-09-02")
	fmt.Println("没生效", data)
	fmt.Println()

	//生效 过去1个月
	data2 := util.HandleXP("2021-08-02", 3, 365, "2021-09-02")
	fmt.Println("生效 过去1个月", data2)
	fmt.Println()

	//生效过去 1年
	data3 := util.HandleXP("2020-09-02", 3, 365, "2021-09-02")
	fmt.Println("生效过去 1年", data3)
	fmt.Println()

	//生效 过去 2年
	data4 := util.HandleXP("2019-09-02", 3, 365, "2021-09-02")
	fmt.Println("生效 过去 2年:", data4)
	fmt.Println()

}
