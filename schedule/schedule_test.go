package schedule

import (
	"entysquare/filer-backend/storage"
	"fmt"
	_ "github.com/lib/pq"
	"testing"
)

func TestStatisticMethod(t *testing.T) {
	db, err := storage.NewDatabase(true)
	fmt.Println(db)
	fmt.Println(err)
	//统计时间为当天执行时最早时间 的前一天
	// (当天-2021-09-01)
	//duration := -1
	//statTime := util.TimeStringToTime(util.TimeNow().Format("2006-01-02"), "00:00:00", "").AddDate(0, 0, duration)
	//ctx :=context.TODO()

	//scheduleExecute(db, statTime, nodeId)
	ManualScheduleRun(db, "2021-09-01", "")
	//err = statisticBlockInfo(db, statTime)
	//err = statisticOrderIncome(db, statTime)
	//err = statisticAccountIncome(db, statTime)
	//err = statisticPoolIncome(db, statTime)

	//db.WithTransaction(func(txn *sql.Tx) error {
	//	statisticBalanceIncome(db, context.TODO(), txn, statTime)
	//	return nil
	//})
	//AutoScheduleRun(db)
}
