package storage

import (
	"context"
	"database/sql"
	"entysquare/enty-cli/storage/types"
	"entysquare/enty-cli/util"
	"fmt"
	"strconv"
	"time"
)

//goland:noinspection SqlNoDataSourceInspection
const filerBlockTempSchema = `
 CREATE TABLE IF NOT EXISTS filer_block_temp  (
        uu_id          TEXT PRIMARY KEY ,
        node_id        TEXT,
        block_num      TEXT,
        block_gain     TEXT,
        block_height   TEXT,
        power   	   TEXT,
        state          TEXT,
        chain_time     TEXT,
        create_time    TEXT,
        update_time    TEXT
    );
`

const insertFilerBlockTempSQL = "" +
	" INSERT INTO filer_block_temp " +
	" (uu_id,node_id,block_num,block_gain,block_height,power,state,chain_time,create_time,update_time) " +
	" VALUES " +
	" (gen_random_uuid(), $1, $2, $3, $4, $5, $6, $7,$8,$9) "

const selectFilerBlockTempMaxBlockHeightSQL = "" +
	" select max(block_height::int8) from filer_block_temp "

const selectFilerBlockTempMaxBlockHeightByNodeIdSQL = "" +
	" select max(block_height::int8) from filer_block_temp where node_id = $1"

const selectStatisticBlockInfoSQL = "" +
	"select node_id," +
	"sum(block_num::int2)," +
	"sum(block_gain::int8)," +
	"max(power::int8)," +
	"sum(block_gain::int8) / max(power::int8)" +
	" from filer_block_temp" +
	" where 1=1 " +
	" and node_id = $1 " +
	" and chain_time::int8 >= $2" +
	" and chain_time::int8 < $3" +
	" group by node_id"

const updateFilerBlockTempStateSQL = "" +
	"update filer_block_temp" +
	" set state      = '9'," +
	" update_time = $1 " +
	" where state = '0' " +
	" and chain_time::int8 >= $2" +
	" and chain_time::int8 < $3" +
	" and node_id = $4"

//const updateGoogleVerifyByIdSQL = "" +
//	"UPDATE user_data SET google_key = $2, google_ver_flag = $3 WHERE user_id = $1"

type filerBlockTempStatements struct {
	insertFilerBlockTempStmt                       *sql.Stmt
	selectFilerBlockTempMaxBlockHeightStmt         *sql.Stmt
	selectFilerBlockTempMaxBlockHeightByNodeIdStmt *sql.Stmt
	selectStatisticBlockInfoStmt                   *sql.Stmt
	updateFilerBlockTempStateStmt                  *sql.Stmt
}

func (s *filerBlockTempStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(filerBlockTempSchema)
	return err
}

func (s *filerBlockTempStatements) prepare(db *sql.DB) (err error) {
	s.insertFilerBlockTempStmt, err = db.Prepare(insertFilerBlockTempSQL)
	s.selectFilerBlockTempMaxBlockHeightStmt, err = db.Prepare(selectFilerBlockTempMaxBlockHeightSQL)
	s.selectFilerBlockTempMaxBlockHeightByNodeIdStmt, err = db.Prepare(selectFilerBlockTempMaxBlockHeightByNodeIdSQL)
	s.selectStatisticBlockInfoStmt, err = db.Prepare(selectStatisticBlockInfoSQL)
	s.updateFilerBlockTempStateStmt, err = db.Prepare(updateFilerBlockTempStateSQL)
	return
}

func (s *filerBlockTempStatements) insertFilerBlockTemp(ctx context.Context, txn *sql.Tx, f *types.FilerBlockTemp) (err error) {
	f.CreateTime = strconv.FormatInt(util.TimeNow().Unix(), 10)
	f.UpdateTime = strconv.FormatInt(util.TimeNow().Unix(), 10)
	_, err = txn.Stmt(s.insertFilerBlockTempStmt).
		ExecContext(ctx, //uuid
			f.NodeId,
			f.BlockNum,
			f.BlockGain,
			f.BlockHeight,
			f.Power,
			0,
			f.ChainTime,
			f.CreateTime,
			f.UpdateTime)
	return
}
func (s *filerBlockTempStatements) selectFilerBlockTempMaxBlockHeight(ctx context.Context) (height int64, err error) {
	result := s.selectFilerBlockTempMaxBlockHeightStmt.QueryRowContext(ctx)
	err = result.Scan(&height)
	return height, err
}
func (s *filerBlockTempStatements) selectFilerBlockTempMaxBlockHeightByNodeId(ctx context.Context, nodeId string) (height int64, err error) {
	result := s.selectFilerBlockTempMaxBlockHeightByNodeIdStmt.QueryRowContext(ctx, nodeId)
	err = result.Scan(&height)
	return height, err
}

func (s *filerBlockTempStatements) selectStatisticBlockInfoByNodeId(ctx context.Context, txn *sql.Tx, statTime time.Time, nodeId string) ([]types.FilerBlockIncome, error) {
	var list []types.FilerBlockIncome
	row, err := txn.Stmt(s.selectStatisticBlockInfoStmt).QueryContext(ctx, nodeId, statTime.Unix(), statTime.AddDate(0, 0, 1).Unix())
	defer row.Close()
	if err != nil {
		fmt.Print("selectStatisticBlockInfo error:", err)
		return nil, err
	}
	for row.Next() {
		var item types.FilerBlockIncome
		err := row.Scan(
			&item.NodeId,
			&item.BlockNum,
			&item.BlockGain,
			&item.Power,
			&item.GainPerTib,
		)
		if err != nil {
			return nil, err
		}
		item.StatTime = statTime.Format("2006-01-02")
		list = append(list, item)
	}
	return list, nil
}

func (s *filerBlockTempStatements) updateFilerBlockTempState(ctx context.Context, txn *sql.Tx, statTime time.Time, nodeId string) error {
	updateTime := strconv.FormatInt(util.TimeNow().Unix(), 10)
	stmt := util.TxStmt(txn, s.updateFilerBlockTempStateStmt)
	r, err := stmt.ExecContext(ctx, updateTime, statTime.Unix(), statTime.AddDate(0, 0, 1).Unix(), nodeId)

	if err != nil {
		return err
	}
	if i, _ := r.RowsAffected(); i < 1 {
		return fmt.Errorf("db No data was modified")
	}
	return nil
}
