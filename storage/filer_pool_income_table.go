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
const filerPoolIncomeSchema = `
 CREATE TABLE IF NOT EXISTS filer_pool_income  (
  uu_id				        TEXT PRIMARY KEY,
  node_id				    TEXT,
  balance				    TEXT,
  pledge_sum				TEXT,
  total_power				TEXT,
  total_income				TEXT,
  freeze_income				TEXT,
  available_income			TEXT,
  today_income_total		TEXT,
  today_income_freeze		TEXT,
  today_income_available	TEXT,
  stat_time					TEXT,
  create_time				TEXT
);

comment on column filer_pool_income.uu_id	is 'pk';
comment on column filer_pool_income.node_id	is '节点id';
comment on column filer_pool_income.balance   is '矿池余额 ';
comment on column filer_pool_income.pledge_sum	is '质押金额';
comment on column filer_pool_income.total_power is '总算力';
comment on column filer_pool_income.total_income	is '总收益';
comment on column filer_pool_income.freeze_income	is '冻结收益';
comment on column filer_pool_income.available_income	is '可用收益';
comment on column filer_pool_income.today_income_total	is '今日产出收益';
comment on column filer_pool_income.today_income_freeze	is '今日冻结收益';
comment on column filer_pool_income.today_income_available	is '今日可用收益';
comment on column filer_pool_income.stat_time	is '统计时间（yyyy-MM-dd）';
comment on column filer_pool_income.create_time	is '创建时间';
`

const insertFilerPoolIncomeSQL = "" +
	" INSERT INTO filer_pool_income " +
	" (uu_id,node_id,balance,pledge_sum,total_power,total_income,freeze_income,available_income,today_income_total,today_income_freeze,today_income_available,stat_time,create_time) " +
	" VALUES " +
	" (gen_random_uuid(), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) "

const selectFilerPoolIncomeByNodeIdSQL = "" +
	" SELECT " +
	" uu_id,node_id,balance,pledge_sum,total_power,total_income,freeze_income,available_income,today_income_total,today_income_freeze,today_income_available,stat_time,create_time " +
	" FROM filer_pool_income " +
	" WHERE " +
	" node_id = $1 and stat_time = $2 "

const selectFilerPoolIncomeAllSQL = "" +
	" SELECT " +
	" uu_id,node_id,balance,pledge_sum,total_power,total_income,freeze_income,available_income,today_income_total,today_income_freeze,today_income_available,stat_time,create_time " +
	" FROM filer_pool_income "
const selectFilerPoolIncomeByNodeIdLastSQL = "" +
	" SELECT " +
	" uu_id,node_id,balance,pledge_sum,total_power,total_income,freeze_income,available_income,today_income_total,today_income_freeze,today_income_available,stat_time,create_time " +
	" FROM filer_pool_income " +
	" WHERE " +
	" node_id = $1 order by create_time desc limit 1  "

//改收益部分sql
const updateFilerPoolIncomeProfitByNodeIdSQL = "" +
	" UPDATE filer_pool_income " +
	" SET " +
	" total_power = $1 ,			" + //总算力
	" total_income = $2 ,			" + //总收益
	" freeze_income = $3 ,			" + //冻结收益
	" available_income = $4 ,		" + //可用收益
	" today_income_total = $5 ,		" + //今日产出收益
	" today_income_freeze = $6 ,	" + //今日冻结收益
	" today_income_available = $7 ,	" + //今日可用收益
	" stat_time = $8 				" + //今日可用收益
	" WHERE " +
	" node_id = $9 "

const deleteFilerPoolIncomeFromStatTimeSQL = "" +
	" delete from filer_pool_income " +
	" where stat_time >= $1" +
	"	and node_id = $2"

type filerPoolIncomeStatements struct {
	insertFilerPoolIncomeStmt               *sql.Stmt
	selectFilerPoolIncomeByNodeIdStmt       *sql.Stmt
	updateFilerPoolIncomeProfitByNodeIdStmt *sql.Stmt
	selectFilerPoolIncomeAllStmt            *sql.Stmt
	selectFilerPoolIncomeByNodeIdLastStmt   *sql.Stmt
	deleteFilerPoolIncomeFromStatTimeStmt   *sql.Stmt
}

func (s *filerPoolIncomeStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(filerPoolIncomeSchema)
	return err
}

func (s *filerPoolIncomeStatements) prepare(db *sql.DB) (err error) {
	if s.insertFilerPoolIncomeStmt, err = db.Prepare(insertFilerPoolIncomeSQL); err != nil {
		return
	}
	if s.selectFilerPoolIncomeByNodeIdStmt, err = db.Prepare(selectFilerPoolIncomeByNodeIdSQL); err != nil {
		return
	}
	if s.updateFilerPoolIncomeProfitByNodeIdStmt, err = db.Prepare(updateFilerPoolIncomeProfitByNodeIdSQL); err != nil {
		return
	}
	if s.selectFilerPoolIncomeAllStmt, err = db.Prepare(selectFilerPoolIncomeAllSQL); err != nil {
		return
	}
	if s.selectFilerPoolIncomeByNodeIdLastStmt, err = db.Prepare(selectFilerPoolIncomeByNodeIdLastSQL); err != nil {
		return
	}
	if s.deleteFilerPoolIncomeFromStatTimeStmt, err = db.Prepare(deleteFilerPoolIncomeFromStatTimeSQL); err != nil {
		return
	}
	return
}

//uu_id,node_id,balance,pledge_sum,total_power,total_income,freeze_income,available_income,allocated_income,
//today_income_total,today_income_freeze,today_income_available,state,stat_time,create_time
func (s *filerPoolIncomeStatements) insertFilerPoolIncome(ctx context.Context, txn *sql.Tx, f *types.FilerPoolIncome) (err error) {
	f.CreateTime = strconv.FormatInt(util.TimeNow().Unix(), 10)
	_, err = txn.Stmt(s.insertFilerPoolIncomeStmt).
		ExecContext(ctx, //uuid
			f.NodeId,
			f.Balance,
			f.PledgeSum,
			f.TotalPower,
			f.TotalIncome,
			f.FreezeIncome,
			f.AvailableIncome,
			f.TodayIncomeTotal,
			f.TodayIncomeFreeze,
			f.TodayIncomeAvailable,
			f.StatTime,
			f.CreateTime)
	return
}

//select uu_id,node_id, block_num, block_gain, power,state, gain_per_tib ,stat_time ,create_time ,update_time
func (s *filerPoolIncomeStatements) selectFilerPoolIncomeByNodeId(ctx context.Context, txn *sql.Tx, nodeId, statTime string) (*types.FilerPoolIncome, error) {
	var f types.FilerPoolIncome
	err := txn.Stmt(s.selectFilerPoolIncomeByNodeIdStmt).
		QueryRowContext(ctx, nodeId, statTime).
		Scan(
			&f.UuId,
			&f.NodeId,
			&f.Balance,
			&f.PledgeSum,
			&f.TotalPower,
			&f.TotalIncome,
			&f.FreezeIncome,
			&f.AvailableIncome,
			&f.TodayIncomeTotal,
			&f.TodayIncomeFreeze,
			&f.TodayIncomeAvailable,
			&f.StatTime,
			&f.CreateTime)
	if err != nil && err.Error() == "sql: no rows in result set" {
		err = nil
	}
	return &f, err
}

func (s *filerPoolIncomeStatements) selectFilerPoolIncomeAll() ([]types.FilerPoolIncome, error) {
	var list []types.FilerPoolIncome
	row, err := s.selectFilerPoolIncomeAllStmt.Query()
	defer row.Close()
	if err != nil {
		fmt.Print("selectFilerPoolIncomeAll error:", err)
		return nil, err
	}
	for row.Next() {
		var f types.FilerPoolIncome
		err := row.Scan(
			&f.UuId,
			&f.NodeId,
			&f.Balance,
			&f.PledgeSum,
			&f.TotalPower,
			&f.TotalIncome,
			&f.FreezeIncome,
			&f.AvailableIncome,
			&f.TodayIncomeTotal,
			&f.TodayIncomeFreeze,
			&f.TodayIncomeAvailable,
			&f.StatTime,
			&f.CreateTime,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, f)
	}
	return list, nil
}

func (s *filerPoolIncomeStatements) updateFilerPoolIncomeProfitByNodeId(ctx context.Context, txn *sql.Tx, f *types.FilerPoolIncome) (err error) {
	if f.NodeId == "" {
		return fmt.Errorf("node id is null")
	}
	r, err := txn.Stmt(s.updateFilerPoolIncomeProfitByNodeIdStmt).ExecContext(ctx,
		f.TotalPower,           //总算力
		f.TotalIncome,          //总收益
		f.FreezeIncome,         //冻结收益
		f.AvailableIncome,      //可用收益
		f.TodayIncomeTotal,     //今日产出收益
		f.TodayIncomeFreeze,    //今日冻结收益
		f.TodayIncomeAvailable, //今日可用收益
		f.StatTime,             //今日可用收益
		f.NodeId)               //节点id
	if err != nil {
		return err
	}
	if i, _ := r.RowsAffected(); i < 1 {
		return fmt.Errorf("db No data was modified")
	}
	return nil
}

//select uu_id,node_id, block_num, block_gain, power,state, gain_per_tib ,stat_time ,create_time ,update_time
func (s *filerPoolIncomeStatements) selectFilerPoolIncomeByNodeIdLast(ctx context.Context, nodeId string) *types.FilerPoolIncome {
	var f types.FilerPoolIncome
	err := s.selectFilerPoolIncomeByNodeIdLastStmt.
		QueryRowContext(ctx, nodeId).
		Scan(
			&f.UuId,
			&f.NodeId,
			&f.Balance,
			&f.PledgeSum,
			&f.TotalPower,
			&f.TotalIncome,
			&f.FreezeIncome,
			&f.AvailableIncome,
			&f.TodayIncomeTotal,
			&f.TodayIncomeFreeze,
			&f.TodayIncomeAvailable,
			&f.StatTime,
			&f.CreateTime)
	if err != nil {
		return nil
	}
	return &f
}

func (s *filerPoolIncomeStatements) deleteFilerPoolIncomeFromStatTime(ctx context.Context, txn *sql.Tx, statTime string, nodeId string) (err error) {
	stmt := TxStmt(txn, s.deleteFilerPoolIncomeFromStatTimeStmt)
	_, err = stmt.ExecContext(ctx, statTime, nodeId)
	if err != nil {
		return err
	}
	return
}
