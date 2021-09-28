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
const filerAccountIncomeSchema = `
 CREATE TABLE IF NOT EXISTS filer_account_income  (
	uu_id 				TEXT PRIMARY KEY,
	filer_id 			TEXT,
	node_id				TEXT,
	pledge_sum 			TEXT,
	hold_power 			TEXT,
	total_income 		TEXT,
	freeze_income 		TEXT,
	total_available_income 	TEXT,
	day_available_income 	TEXT,
	day_raise_income    TEXT,
	day_direct_income	TEXT,
	day_release_income	TEXT,
	stat_time 			TEXT,
	create_time			TEXT
);
comment on column filer_account_income.uu_id is 'uu_id';
comment on column filer_account_income.filer_id is 'filer id';
comment on column filer_account_income.node_id is '节点id';
comment on column filer_account_income.pledge_sum is '质押金额';
comment on column filer_account_income.hold_power is '持有算力';
comment on column filer_account_income.total_income is '总收益';
comment on column filer_account_income.freeze_income is '冻结收益';
comment on column filer_account_income.total_available_income is '累计已释放收益 = 昨日累计收益+当日释放收益 ';
comment on column filer_account_income.day_available_income is '可用收益(当日直接25%+往日75%平账 ';
comment on column filer_account_income.day_raise_income is '日产出收益（当日直接25%+当日冻结75%）';
comment on column filer_account_income.day_direct_income is '当日直接收益（25%';
comment on column filer_account_income.day_release_income is '当日线性释放收益（往日75%平账）';
comment on column filer_account_income.stat_time is '统计时间（yyyy-MM-dd)';
comment on column filer_account_income.create_time is '创建时间 时间戳';
`
const selectFilerAccountIncomeByFilerIdSQL = "select o.uu_id, " +
	" o.filer_id, " +
	" o.node_id, " +
	" o.pledge_sum, " +
	" o.hold_power, " +
	" o.total_income, " +
	" o.freeze_income, " +
	" o.total_available_income, " +
	" o.day_available_income, " +
	" o.day_raise_income, " +
	" o.day_direct_income, " +
	" o.day_release_income, " +
	" o.stat_time, " +
	" o.create_time " +
	" from filer_account_income o" +
	" where filer_id = $1 order by create_time desc"

const selectFilerAccountIncomeByStatTimeSQL = "select " +
	" uu_id, " +
	" filer_id, " +
	" node_id, " +
	" pledge_sum, " +
	" hold_power, " +
	" total_income, " +
	" freeze_income, " +
	" total_available_income, " +
	" day_available_income, " +
	" day_raise_income, " +
	" day_direct_income, " +
	" day_release_income, " +
	" stat_time, " +
	" create_time " +
	" from filer_account_income" +
	" where filer_id = $1 " +
	"  and stat_time = $2 " +
	"  and node_id = $3"

const insertFilerAccountIncomeSQL = "insert into filer_account_income" +
	" (uu_id, filer_id, node_id, pledge_sum, hold_power, total_income, freeze_income, total_available_income, day_available_income,day_raise_income, day_direct_income, day_release_income, stat_time, create_time)" +
	" values (gen_random_uuid(), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10 ,$11, $12, $13)"

const deleteFilerAccountIncomeFromStatTimeSQL = "" +
	" delete from filer_account_income" +
	" where stat_time >= $1" +
	" and node_id = $2"

const selectStatisticBalanceIncomeSQL = "" +
	" select t.filer_id," +
	"        sum(t.pledge_sum::int8)," +
	"        sum(t.hold_power::float8)," +
	"        sum(t.total_income::int8)," +
	"        sum(t.freeze_income::int8)," +
	"        sum(t.total_available_income::int8)," +
	"        sum(t.day_available_income::int8)," +
	"        sum(t.day_raise_income::int8)," +
	"        sum(t.day_direct_income::int8)," +
	"        sum(t.day_release_income::int8)" +
	" from filer_account_income t" +
	" where t.stat_time = $1" +
	" group by t.filer_id"

type filerAccountIncomeStatements struct {
	selectFilerAccountIncomeByFilerIdStmt    *sql.Stmt
	selectFilerAccountIncomeByStatTimeStmt   *sql.Stmt
	insertFilerAccountIncomeStmt             *sql.Stmt
	deleteFilerAccountIncomeFromStatTimeStmt *sql.Stmt
	selectStatisticBalanceIncomeStmt         *sql.Stmt
}

func (s *filerAccountIncomeStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(filerAccountIncomeSchema)
	return err
}

func (s *filerAccountIncomeStatements) prepare(db *sql.DB) (err error) {
	if s.selectFilerAccountIncomeByFilerIdStmt, err = db.Prepare(selectFilerAccountIncomeByFilerIdSQL); err != nil {
		return
	}
	if s.selectFilerAccountIncomeByStatTimeStmt, err = db.Prepare(selectFilerAccountIncomeByStatTimeSQL); err != nil {
		return
	}
	if s.insertFilerAccountIncomeStmt, err = db.Prepare(insertFilerAccountIncomeSQL); err != nil {
		return
	}
	if s.deleteFilerAccountIncomeFromStatTimeStmt, err = db.Prepare(deleteFilerAccountIncomeFromStatTimeSQL); err != nil {
		return
	}
	if s.selectStatisticBalanceIncomeStmt, err = db.Prepare(selectStatisticBalanceIncomeSQL); err != nil {
		return
	}
	return
}
func (s *filerAccountIncomeStatements) selectFilerAccountIncomeByFilerId(ctx context.Context, txn *sql.Tx, filerId int64) (list map[string]types.FilerAccountIncome, err error) {
	rows, err := util.TxStmt(txn, s.selectFilerAccountIncomeByFilerIdStmt).QueryContext(ctx, filerId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	maps := make(map[string]types.FilerAccountIncome)
	for rows.Next() {
		filerAccountIncome := types.FilerAccountIncome{}
		err = rows.Scan(
			&filerAccountIncome.Uuid,
			&filerAccountIncome.FilerId,
			&filerAccountIncome.NodeId,
			&filerAccountIncome.PledgeSum,
			&filerAccountIncome.HoldPower,
			&filerAccountIncome.TotalIncome,
			&filerAccountIncome.FreezeIncome,
			&filerAccountIncome.TotalAvailableIncome,
			&filerAccountIncome.DayAvailableIncome,
			&filerAccountIncome.DayRaiseIncome,
			&filerAccountIncome.DayDirectIncome,
			&filerAccountIncome.DayReleaseIncome,
			&filerAccountIncome.StatTime,
			&filerAccountIncome.CreateTime,
		)
		if err != nil {
			return nil, err
		}
		maps[filerAccountIncome.Uuid] = filerAccountIncome
	}
	list = maps
	return
}

func (s *filerAccountIncomeStatements) selectFilerAccountIncomeByStatTime(ctx context.Context, txn *sql.Tx, statTime string, filerId string, nodeId string) ([]types.FilerAccountIncome, error) {
	var list []types.FilerAccountIncome
	row, err := txn.Stmt(s.selectFilerAccountIncomeByStatTimeStmt).QueryContext(ctx, filerId, statTime, nodeId)
	defer row.Close()
	if err != nil {
		fmt.Print("selectStatisticBlockInfo error:", err)
		return nil, err
	}
	for row.Next() {
		var item types.FilerAccountIncome
		err := row.Scan(
			&item.Uuid,
			&item.FilerId,
			&item.NodeId,
			&item.PledgeSum,
			&item.HoldPower,
			&item.TotalIncome,
			&item.FreezeIncome,
			&item.TotalAvailableIncome,
			&item.DayAvailableIncome,
			&item.DayRaiseIncome,
			&item.DayDirectIncome,
			&item.DayReleaseIncome,
			&item.StatTime,
			&item.CreateTime,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

func (s *filerAccountIncomeStatements) insertFilerAccountIncome(ctx context.Context, txn *sql.Tx, f *types.FilerAccountIncome) (err error) {
	f.CreateTime = strconv.FormatInt(util.TimeNow().Unix(), 10)
	_, err = txn.Stmt(s.insertFilerAccountIncomeStmt).
		ExecContext(ctx, //uuid
			f.FilerId,
			f.NodeId,
			f.PledgeSum,
			f.HoldPower,
			f.TotalIncome,
			f.FreezeIncome,
			f.TotalAvailableIncome,
			f.DayAvailableIncome,
			f.DayRaiseIncome,
			f.DayDirectIncome,
			f.DayReleaseIncome,
			f.StatTime,
			f.CreateTime,
		)
	return
}

func (s *filerAccountIncomeStatements) deleteFilerAccountIncomeFromStatTime(ctx context.Context, txn *sql.Tx, statTime string, nodeId string) (err error) {
	stmt := util.TxStmt(txn, s.deleteFilerAccountIncomeFromStatTimeStmt)
	_, err = stmt.ExecContext(ctx, statTime, nodeId)
	if err != nil {
		return err
	}
	return
}

func (s *filerAccountIncomeStatements) selectStatisticBalanceIncome(ctx context.Context, txn *sql.Tx, statTime string) ([]types.FilerBalanceIncome, error) {
	var list []types.FilerBalanceIncome
	row, err := txn.Stmt(s.selectStatisticBalanceIncomeStmt).QueryContext(ctx, statTime)
	defer row.Close()
	if err != nil {
		fmt.Print("selectStatisticBalanceIncome error:", err)
		return nil, err
	}
	for row.Next() {
		var item types.FilerBalanceIncome
		err := row.Scan(
			&item.FilerId,
			&item.PledgeSum,
			&item.HoldPower,
			&item.TotalIncome,
			&item.FreezeIncome,
			&item.TotalAvailableIncome,
			&item.DayAvailableIncome,
			&item.DayRaiseIncome,
			&item.DayDirectIncome,
			&item.DayReleaseIncome,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}
