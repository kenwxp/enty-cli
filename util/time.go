package util

import (
	"errors"
	"strconv"
	"time"
)

type ValuedBlockRange struct {
	Code  int //Code = 0-收益未开始，1-开始生效 2-有收益，3-有收益且今日到期 9-收益结算完毕 -1 err
	Start string
	End   string
}

//获取中国时区当前时间
func TimeNow() time.Time {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	return time.Now().In(cstSh)
}

/*
获取 日期两端 时间戳
start 2021-12-30 00:00:01
end   2021-12-30 23:59:59
*/
func TimeDayUnix(time2 time.Time) (start time.Time, end time.Time) {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	start, _ = time.ParseInLocation("2006-01-02 15:04:05", time2.Format("2006-01-02 ")+"00:00:00", cstSh)
	end, _ = time.ParseInLocation("2006-01-02 15:04:05", time2.Format("2006-01-02 ")+"23:59:59", cstSh)
	return
}

/*
	"2021-09-03 17:11:00" -> 转 时间搓
*/
func TimeStrToUnix(timeStr string) int64 {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, cstSh)
	if err != nil {
		return -1
	}
	return t2.Unix()
}

/*
 具有特定格式（yyyy-MM-dd）的字符串，转换为Time对象
	timeStr yyyy-MM-dd
	suffix hh:mm:ss
	location 如Asia/Shanghai
*/
func TimeStringToTime(timeStr string, suffix string, location string) time.Time {
	if location == "" {
		location = "Asia/Shanghai"
	}
	var cstSh, _ = time.LoadLocation(location) //上海
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" "+suffix, cstSh)
	return t
}

/*
	处理每条持仓所拥有的爆块收益时间区间
	orderTimeStr = 下单时间 时间戳
	validPlan = 生效时间（天） T+3 = 3天
	period = 周期（天）
	statTimeStr = 今天时间 字符串 "2020-01-02"
*/
func GetValuedBlockRange(orderTimeStr string, validPlanStr string, periodStr string, statTimeStr string) ValuedBlockRange {
	//时区 上海
	var cstSh, _ = time.LoadLocation("Asia/Shanghai")
	//下单时间
	orderTimeInt, _ := strconv.ParseInt(orderTimeStr, 10, 64)
	orderFormatTimeStr := time.Unix(orderTimeInt, 0).Format("2006-01-02")
	orderTime, _ := time.ParseInLocation("2006-01-02 15:04:05", orderFormatTimeStr+" 12:00:00", cstSh)
	//生效时间
	validPlan, _ := strconv.Atoi(validPlanStr)
	valid := orderTime.AddDate(0, 0, validPlan)
	//周期 结束时间
	period, _ := strconv.Atoi(periodStr)
	periodEnd := valid.AddDate(0, 0, period)
	//今天时间
	statTime, _ := time.ParseInLocation("2006-01-02 15:04:05", statTimeStr+" 12:00:00", cstSh)
	//今天前推180天
	start180 := statTime.AddDate(0, 0, -180)

	//fmt.Println("下单时间:", orderTime.Unix(), orderTime.Format("2006-01-02 15:04:05"))
	//fmt.Println("生效时间:", valid.Unix(), valid.Format("2006-01-02 15:04:05"))
	//fmt.Println("周期:", periodEnd.Unix(), periodEnd.Format("2006-01-02 15:04:05"))
	//fmt.Println("今天时间:", statTime.Unix(), statTime.Format("2006-01-02 15:04:05"))
	//fmt.Println("今天前推180天:", start180.Unix(), start180.Format("2006-01-02 15:04:05"))

	// 		    | [---持仓周期---]
	// [--180--]|
	if statTime.Unix() < valid.Unix() {
		return ValuedBlockRange{Code: 0}
	}
	if statTime.Unix() == valid.Unix() {
		return ValuedBlockRange{Code: 1, Start: valid.Format("2006-01-02"), End: statTime.Format("2006-01-02")}
	}
	if statTime.Unix() > valid.Unix() && statTime.Unix() < periodEnd.Unix() {
		// 	   [---|--持仓周期---]
		// [--180--|]
		if start180.Unix() < valid.Unix() {
			return ValuedBlockRange{Code: 2, Start: valid.Format("2006-01-02"), End: statTime.Format("2006-01-02")}
		}
		//  [-----持仓周期---|--]
		//         [--180--|]
		if start180.Unix() >= valid.Unix() {
			return ValuedBlockRange{Code: 2, Start: start180.Format("2006-01-02"), End: statTime.Format("2006-01-02")}
		}
	}
	if statTime.Unix() == periodEnd.Unix() {
		return ValuedBlockRange{Code: 3, Start: start180.Format("2006-01-02"), End: statTime.Format("2006-01-02")}
	}

	if statTime.Unix() > periodEnd.Unix() {
		//[---持仓周期---]    |
		//			[--180--|]
		if start180.Unix() <= periodEnd.Unix() {
			return ValuedBlockRange{Code: 2, Start: start180.Format("2006-01-02"), End: periodEnd.Format("2006-01-02")}
		}
		//[---持仓周期---]	|
		//                	|[---180---]
		if start180.Unix() > periodEnd.Unix() {
			return ValuedBlockRange{Code: 9}
		}
	}

	return ValuedBlockRange{Code: -1}
}

/**
  根据所传时间字符串（格式：yyyy-MM-dd）计算与当前统计时间之间的间隔天数
*/
func GetDuratinDaysFromCurrentStatTime(timeStr string) (int, error) {
	statTimestamp := TimeStringToTime(timeStr, "00:00:00", "").Unix()
	curTimeStr := TimeNow().AddDate(0, 0, -1).Format("2006-01-02")
	curTimestamp := TimeStringToTime(curTimeStr, "00:00:00", "").Unix()
	if statTimestamp > curTimestamp {
		return 0, errors.New("the statTimeStr set beyond the current statistic time:" + timeStr + ">" + curTimeStr)
	}
	statTime, _ := time.Parse("2006-01-02", timeStr)
	curTime, _ := time.Parse("2006-01-02", curTimeStr)
	durationOfDays := curTime.Sub(statTime).Hours() / 24
	return int(durationOfDays), nil
}

func ConvertTimestampToTimeStr(timestamp string, format string) string {
	timeString, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return ""
	}
	return time.Unix(timeString, 0).Format(format)
}

/*
 粗略计算两天间隔，不考虑具体小时，如2021-09-01 23:59:59 和2021-09-02 00:00:00 相差一天
*/
func GetDurationDaysForTimestamp(foreTsStr string, lateTsStr string) string {
	foreTs, err := strconv.ParseInt(foreTsStr, 10, 64)
	if err != nil {
		return ""
	}
	lateTs, err := strconv.ParseInt(lateTsStr, 10, 64)
	if err != nil {
		return ""
	}
	foreTimeTs := TimeStringToTime(time.Unix(foreTs, 0).Format("2006-01-02"), "00:00:00", "")
	lateTimeTs := TimeStringToTime(time.Unix(lateTs, 0).Format("2006-01-02"), "00:00:00", "")
	durationDays := int64(lateTimeTs.Sub(foreTimeTs).Hours() / 24)
	return strconv.FormatInt(durationDays, 10)
}
