package util

import (
	"fmt"
	"testing"
	"time"
)

const TIME_NOW = "2021-09-06 14:24:30"
const TIME_NOW_LONDON = "2021-09-06 5:58:30"

func TestTimeZone(t *testing.T) {
	//now := time.Now().Unix()
	//fmt.Println(now)
	//ts := TimeStrToUnix(TIME_NOW)
	//fmt.Println(now, ts, (now-ts)/60/60)
	//zoneTs := TimeStrToUnix(TIME_NOW)

	now := time.Now()
	nowTs := now.Unix()
	fmt.Println(nowTs)
	t2, err := time.Parse("2006/01/02 15:04:05.000000", now.Format("2006/01/02 15:04:05.000000"))
	if err != nil {
		panic("")
	}
	fmt.Println(t2.Unix())

}
