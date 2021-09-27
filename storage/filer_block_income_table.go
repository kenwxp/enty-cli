package storage

import (
	"context"
	"database/sql"
	"entysquare/filer-backend/storage/types"
	"entysquare/filer-backend/util"
	"fmt"
	"strconv"
	"time"
)

//goland:noinspection SqlNoDataSourceInspection
const filerBlockIncomeSchema = `
 	CREATE TABLE IF NOT EXISTS filer_block_income  (
		uu_id        TEXT PRIMARY KEY,
		node_id        TEXT,
		block_num      TEXT,
		block_gain     TEXT,
		power			TEXT,
		gain_per_tib	TEXT,
		state          TEXT,
		stat_time    TEXT,
		create_time    TEXT,
		update_time		TEXT
	);

	comment on column filer_block_income.uu_id is 'uuid ';
	comment on column filer_block_income.node_id is '节点id ';
	comment on column filer_block_income.block_num is '爆块数 ';
	comment on column filer_block_income.block_gain is '爆块奖励  FIL ';
	comment on column filer_block_income.power is '总算力 T';
	comment on column filer_block_income.gain_per_tib is '每T收益';
	comment on column filer_block_income.state    is ' 处理标志 0-待处理 1-线性释放 9-已处理 ';
	comment on column filer_block_income.stat_time is '统计时间（yyyy-MM-dd）';
	comment on column filer_block_income.create_time is '创建时间';
	comment on column filer_block_income.update_time is '更新时间';
`
const insertFilerBlockIncomeSQL = "" +
	" INSERT INTO filer_block_income " +
	" (uu_id,node_id,block_num,block_gain,power,gain_per_tib,state,stat_time,create_time,update_time) " +
	" VALUES " +
	" (gen_random_uuid(),$1, $2, $3, $4, $5, $6, $7, $8, $9) "

const selectFilerBlockIncomeByStatRangeSQL = "" +
	"select node_id, block_num, block_gain, power, gain_per_tib,stat_time" +
	" from filer_block_income" +
	" where stat_time >= $1 " +
	"   and stat_time <= $2 " +
	"	and node_id = $3" +
	"   order by create_time"

const selectFilerBlockIncomeByStatTimeAndNodeIdSQL = "" +
	"select uu_id,node_id, block_num, block_gain, power,state, gain_per_tib ,stat_time ,create_time ,update_time " +
	" from filer_block_income" +
	" where stat_time = $1 " +
	"     and node_id = $2 "

const deleteFilerBlockIncomeFromStatTimeSQL = "" +
	"delete from filer_block_income" +
	" where stat_time >= $1" +
	" 	and node_id = $2"

//组装图 出块记录
const selectFilerBlockAndTimeByTimeSQL = "" +
	" select sum(block_gain::bigint),stat_time from filer_block_income " +
	" where " +
	" stat_time > $1 " +
	" and " +
	" stat_time <= $2 " +
	" group by stat_time order by stat_time DESC "

//const setUserAddressByPhoneNumSQL = "" +
//	"UPDATE user_data SET usdt_address = $1, hsf_address = $2 WHERE phone_num = $3"

type filerBlockIncomeStatements struct {
	insertFilerBlockIncomeStmt                    *sql.Stmt
	selectFilerBlockIncomeByStatRangeStmt         *sql.Stmt
	selectFilerBlockIncomeByStatTimeAndNodeIdStmt *sql.Stmt
	deleteFilerBlockIncomeFromStatTimeStmt        *sql.Stmt
	selectFilerBlockAndTimeByTimeStmt             *sql.Stmt
}

func (s *filerBlockIncomeStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(filerBlockIncomeSchema)
	return err
}

func (s *filerBlockIncomeStatements) prepare(db *sql.DB) (err error) {
	if s.insertFilerBlockIncomeStmt, err = db.Prepare(insertFilerBlockIncomeSQL); err != nil {
		return
	}
	if s.selectFilerBlockIncomeByStatRangeStmt, err = db.Prepare(selectFilerBlockIncomeByStatRangeSQL); err != nil {
		return
	}
	if s.selectFilerBlockIncomeByStatTimeAndNodeIdStmt, err = db.Prepare(selectFilerBlockIncomeByStatTimeAndNodeIdSQL); err != nil {
		return
	}
	if s.deleteFilerBlockIncomeFromStatTimeStmt, err = db.Prepare(deleteFilerBlockIncomeFromStatTimeSQL); err != nil {
		return
	}
	if s.selectFilerBlockAndTimeByTimeStmt, err = db.Prepare(selectFilerBlockAndTimeByTimeSQL); err != nil {
		return
	}

	return
}

//stat_id,node_id,block_num,block_gain,power,gain_per_tib,state,create_time,update_time
func (s *filerBlockIncomeStatements) insertFilerBlockIncome(ctx context.Context, txn *sql.Tx, f *types.FilerBlockIncome) (err error) {
	f.CreateTime = strconv.FormatInt(util.TimeNow().Unix(), 10)
	f.UpdateTime = strconv.FormatInt(util.TimeNow().Unix(), 10)
	f.State = "0"
	_, err = txn.Stmt(s.insertFilerBlockIncomeStmt).
		ExecContext(ctx, //uuid
			f.NodeId,
			f.BlockNum,
			f.BlockGain,
			f.Power,
			f.GainPerTib,
			f.State, //处理标志 0-待处理 1-线性释放 9-已处理
			f.StatTime,
			f.CreateTime,
			f.UpdateTime)
	return
}

func (s *filerBlockIncomeStatements) selectFilerBlockIncomeByStatRange(ctx context.Context, txn *sql.Tx, startTime string, endTime string, nodeId string) ([]types.FilerBlockIncome, error) {
	var list []types.FilerBlockIncome
	row, err := txn.Stmt(s.selectFilerBlockIncomeByStatRangeStmt).QueryContext(ctx, startTime, endTime, nodeId)
	defer row.Close()
	if err != nil {
		fmt.Print("selectStatisticBlockInfo error:", err)
		return nil, err
	}
	//stat_id, node_id, block_num, block_gain, power, gain_per_tib
	for row.Next() {
		var item types.FilerBlockIncome
		err := row.Scan(
			&item.NodeId,
			&item.BlockNum,
			&item.BlockGain,
			&item.Power,
			&item.GainPerTib,
			&item.StatTime,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

func (s *filerBlockIncomeStatements) selectFilerBlockAndTimeByTime(ctx context.Context, start, end time.Time) ([]types.BlockRecord, error) {
	row, err := s.selectFilerBlockAndTimeByTimeStmt.QueryContext(ctx,
		start.Format("2006-01-02"),
		end.Format("2006-01-02"))
	defer row.Close()
	if err != nil {
		fmt.Print("selectFilerBlockAndTimeByTIme error:", err)
		return nil, err
	}
	var list []types.BlockRecord
	for row.Next() {
		var item types.BlockRecord
		err := row.Scan(
			&item.Num,
			&item.Time,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, err
}

//select uu_id,node_id, block_num, block_gain, power,state, gain_per_tib ,stat_time ,create_time ,update_time
func (s *filerBlockIncomeStatements) selectFilerBlockIncomeByStatTimeAndNodeId(ctx context.Context, startTime string, NodeId string) *types.FilerBlockIncome {
	var f types.FilerBlockIncome
	err := s.selectFilerBlockIncomeByStatTimeAndNodeIdStmt.
		QueryRowContext(ctx, startTime, NodeId).
		Scan(&f.UuId,
			&f.NodeId,
			&f.BlockNum,
			&f.BlockGain,
			&f.Power,
			&f.State,
			&f.GainPerTib,
			&f.StatTime,
			&f.CreateTime,
			&f.UpdateTime)
	if err != nil {
		return nil
	}
	return &f
}
func (s *filerBlockIncomeStatements) deleteFilerBlockIncomeFromStatTime(ctx context.Context, txn *sql.Tx, statTime string, nodeId string) (err error) {
	stmt := TxStmt(txn, s.deleteFilerBlockIncomeFromStatTimeStmt)
	_, err = stmt.ExecContext(ctx, statTime, nodeId)
	if err != nil {
		return err
	}
	return
}
