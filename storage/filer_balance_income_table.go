package storage

import (
	"context"
	"database/sql"
	"entysquare/enty-cli/storage/types"
	"entysquare/enty-cli/util"
	"fmt"
	"strconv"
)

//goland:noinspection SqlNoDataSourceInspection
const filerBalanceIncomeSchema = `
 CREATE TABLE IF NOT EXISTS filer_balance_income  (
	uu_id 				TEXT PRIMARY KEY,
	filer_id 			TEXT,
	pledge_sum 			TEXT,
	hold_power 			TEXT,
	total_income 		TEXT,
	freeze_income 		TEXT,
	total_available_income 	TEXT,
	balance					TEXT,
	day_available_income 	TEXT,
	day_raise_income    TEXT,
	day_direct_income	TEXT,
	day_release_income	TEXT,            
	stat_time 			TEXT,
	create_time			TEXT,
	update_time			TEXT
);
comment on column filer_balance_income.uu_id is 'uu_id';
comment on column filer_balance_income.filer_id is 'filer id';
comment on column filer_balance_income.pledge_sum is '质押金额';
comment on column filer_balance_income.hold_power is '持有算力';
comment on column filer_balance_income.total_income is '总收益';
comment on column filer_balance_income.freeze_income is '冻结收益';
comment on column filer_balance_income.total_available_income is '累计已释放收益 = 昨日累计收益+当日释放收益 ';
comment on column filer_balance_income.balance is '余额';
comment on column filer_balance_income.day_available_income is '可用收益(当日直接25%+往日75%平账 ';
comment on column filer_balance_income.day_raise_income is '日产出收益（当日直接25%+当日冻结75%）';
comment on column filer_balance_income.day_direct_income is '当日直接收益（25%';
comment on column filer_balance_income.day_release_income is '当日线性释放收益（往日75%平账）';
comment on column filer_balance_income.stat_time is '统计时间（yyyy-MM-dd)';
comment on column filer_balance_income.create_time is '创建时间 时间戳';
comment on column filer_balance_income.update_time is '更新时间 时间戳';
`

const selectFilerBalanceIncomeByStatTimeSQL = "" +
	" select " +
	" 	   uu_id, " +
	" 	   filer_id, " +
	" 	   pledge_sum, " +
	" 	   hold_power, " +
	" 	   total_income, " +
	" 	   freeze_income, " +
	" 	   total_available_income, " +
	" 	   balance, " +
	"      day_available_income, " +
	"      day_raise_income, " +
	"      day_direct_income, " +
	"      day_release_income, " +
	" 	   stat_time, " +
	" 	   create_time, " +
	" 	   update_time " +
	" from filer_balance_income" +
	" where filer_id = $1 " +
	"  and stat_time = $2 "

const insertFilerBalanceIncomeSQL = "insert into filer_balance_income" +
	" (uu_id, filer_id, pledge_sum, hold_power, total_income, freeze_income,total_available_income, balance, day_available_income, day_raise_income, day_direct_income, day_release_income, stat_time, create_time, update_time)" +
	" values (gen_random_uuid(), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)"

const deleteFilerBalanceIncomeFromStatTimeSQL = "" +
	" delete from filer_balance_income" +
	" where stat_time >= $1"

const selectFilerBalanceIncomeByFilerIdSQL = "" +
	" select uu_id, " +
	" 	     filer_id, " +
	" 	     pledge_sum, " +
	" 	     hold_power, " +
	" 	     total_income, " +
	" 	     freeze_income, " +
	" 	     total_available_income, " +
	" 	     balance, " +
	"        day_available_income, " +
	"        day_raise_income, " +
	"        day_direct_income, " +
	"        day_release_income, " +
	" 	     stat_time, " +
	" 	     create_time, " +
	" 	     update_time " +
	"  from filer_balance_income" +
	"  where filer_id = $1 " +
	"  order by stat_time desc " +
	"  limit 1 "

const selectFilerBalanceIncomeListByFilerIdSQL = "" +
	" select uu_id, " +
	" 	     filer_id, " +
	" 	     pledge_sum, " +
	" 	     hold_power, " +
	" 	     total_income, " +
	" 	     freeze_income, " +
	" 	     total_available_income, " +
	" 	     balance, " +
	"        day_available_income, " +
	"        day_raise_income, " +
	"        day_direct_income, " +
	"        day_release_income, " +
	" 	     stat_time, " +
	" 	     create_time, " +
	" 	     update_time " +
	"  from filer_balance_income" +
	"  where filer_id = $1 " +
	"  order by stat_time desc "

type filerBalanceIncomeStatements struct {
	selectFilerBalanceIncomeByStatTimeStmt    *sql.Stmt
	insertFilerBalanceIncomeStmt              *sql.Stmt
	deleteFilerBalanceIncomeFromStatTimeStmt  *sql.Stmt
	selectFilerBalanceIncomeByFilerIdStmt     *sql.Stmt
	selectFilerBalanceIncomeListByFilerIdStmt *sql.Stmt
}

func (s *filerBalanceIncomeStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(filerBalanceIncomeSchema)
	return err
}

func (s *filerBalanceIncomeStatements) prepare(db *sql.DB) (err error) {

	if s.selectFilerBalanceIncomeByStatTimeStmt, err = db.Prepare(selectFilerBalanceIncomeByStatTimeSQL); err != nil {
		return
	}
	if s.insertFilerBalanceIncomeStmt, err = db.Prepare(insertFilerBalanceIncomeSQL); err != nil {
		return
	}
	if s.deleteFilerBalanceIncomeFromStatTimeStmt, err = db.Prepare(deleteFilerBalanceIncomeFromStatTimeSQL); err != nil {
		return
	}
	if s.selectFilerBalanceIncomeByFilerIdStmt, err = db.Prepare(selectFilerBalanceIncomeByFilerIdSQL); err != nil {
		return
	}
	if s.selectFilerBalanceIncomeListByFilerIdStmt, err = db.Prepare(selectFilerBalanceIncomeListByFilerIdSQL); err != nil {
		return
	}
	return
}

func (s *filerBalanceIncomeStatements) selectFilerBalanceIncomeByStatTime(ctx context.Context, txn *sql.Tx, statTime string, filerId string) ([]types.FilerBalanceIncome, error) {
	var list []types.FilerBalanceIncome
	row, err := txn.Stmt(s.selectFilerBalanceIncomeByStatTimeStmt).QueryContext(ctx, filerId, statTime)
	defer row.Close()
	if err != nil {
		fmt.Print("selectStatisticBlockInfo error:", err)
		return nil, err
	}
	for row.Next() {
		var item types.FilerBalanceIncome
		err := row.Scan(
			&item.Uuid,
			&item.FilerId,
			&item.PledgeSum,
			&item.HoldPower,
			&item.TotalIncome,
			&item.FreezeIncome,
			&item.TotalAvailableIncome,
			&item.Balance,
			&item.DayAvailableIncome,
			&item.DayRaiseIncome,
			&item.DayDirectIncome,
			&item.DayReleaseIncome,
			&item.StatTime,
			&item.CreateTime,
			&item.UpdateTime,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

func (s *filerBalanceIncomeStatements) selectFilerBalanceIncomeByFilerId(ctx context.Context, txn *sql.Tx, filerId string) ([]types.FilerBalanceIncome, error) {
	var list []types.FilerBalanceIncome
	row, err := TxStmt(txn, s.selectFilerBalanceIncomeByFilerIdStmt).QueryContext(ctx, filerId)
	defer row.Close()
	if err != nil {
		fmt.Print("selectFilerBalanceIncomeByFilerId error:", err)
		return nil, err
	}
	for row.Next() {
		var item types.FilerBalanceIncome
		err := row.Scan(
			&item.Uuid,
			&item.FilerId,
			&item.PledgeSum,
			&item.HoldPower,
			&item.TotalIncome,
			&item.FreezeIncome,
			&item.TotalAvailableIncome,
			&item.Balance,
			&item.DayAvailableIncome,
			&item.DayRaiseIncome,
			&item.DayDirectIncome,
			&item.DayReleaseIncome,
			&item.StatTime,
			&item.CreateTime,
			&item.UpdateTime,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

func (s *filerBalanceIncomeStatements) selectFilerBalanceIncomeListByFilerId(ctx context.Context, txn *sql.Tx, filerId string) ([]types.FilerBalanceIncome, error) {
	var list []types.FilerBalanceIncome
	row, err := TxStmt(txn, s.selectFilerBalanceIncomeListByFilerIdStmt).QueryContext(ctx, filerId)
	defer row.Close()
	if err != nil {
		fmt.Print("selectFilerBalanceIncomeByFilerId error:", err)
		return nil, err
	}
	for row.Next() {
		var item types.FilerBalanceIncome
		err := row.Scan(
			&item.Uuid,
			&item.FilerId,
			&item.PledgeSum,
			&item.HoldPower,
			&item.TotalIncome,
			&item.FreezeIncome,
			&item.TotalAvailableIncome,
			&item.Balance,
			&item.DayAvailableIncome,
			&item.DayRaiseIncome,
			&item.DayDirectIncome,
			&item.DayReleaseIncome,
			&item.StatTime,
			&item.CreateTime,
			&item.UpdateTime,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

func (s *filerBalanceIncomeStatements) insertFilerBalanceIncome(ctx context.Context, txn *sql.Tx, f *types.FilerBalanceIncome) (err error) {
	f.CreateTime = strconv.FormatInt(util.TimeNow().Unix(), 10)
	_, err = txn.Stmt(s.insertFilerBalanceIncomeStmt).
		ExecContext(ctx, //uuid
			f.FilerId,
			f.PledgeSum,
			f.HoldPower,
			f.TotalIncome,
			f.FreezeIncome,
			f.TotalAvailableIncome,
			f.Balance,
			f.DayAvailableIncome,
			f.DayRaiseIncome,
			f.DayDirectIncome,
			f.DayReleaseIncome,
			f.StatTime,
			f.CreateTime,
			f.UpdateTime,
		)
	return
}

func (s *filerBalanceIncomeStatements) deleteFilerBalanceIncomeFromStatTime(ctx context.Context, txn *sql.Tx, statTime string) (err error) {
	stmt := TxStmt(txn, s.deleteFilerBalanceIncomeFromStatTimeStmt)
	_, err = stmt.ExecContext(ctx, statTime)
	if err != nil {
		return err
	}
	return
}
