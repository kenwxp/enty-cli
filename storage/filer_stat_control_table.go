package storage

import (
	"context"
	"database/sql"
	"entysquare/filer-backend/storage/types"
	"entysquare/filer-backend/util"
	"fmt"
	"strconv"
)

//goland:noinspection SqlNoDataSourceInspection
const filerStatControlSchema = `
 CREATE TABLE IF NOT EXISTS filer_stat_control
(
	uu_id				TEXT PRIMARY KEY,
	stat_type		TEXT,
	stat_time 	TEXT,
	node_id	    TEXT,
	stat_state	TEXT,
	message     TEXT,
	create_time TEXT,
	update_time TEXT
);
comment on column filer_stat_control.uu_id is 'uuid';
comment on column filer_stat_control.stat_type is '统计类型TASK_BLOCK_STAT,TASK_ORDER_STAT,TASK_ACCOUNT_STAT,TASK_POOL_STAT';
comment on column filer_stat_control.stat_time is '统计时间（yyyy-MM-dd)';
comment on column filer_stat_control.node_id is '统计节点';
comment on column filer_stat_control.stat_state is '统计状态0-新增1-进行中，2-成功3-失败';
comment on column filer_stat_control.message is '信息';      
comment on column filer_stat_control.create_time is '创建时间';
comment on column filer_stat_control.update_time is '更新时间';
`

const selectStatControlInfoSQL = "" +
	"select stat_type,stat_time,node_id,stat_state,message,create_time,update_time" +
	" from filer_stat_control" +
	" where stat_type = $1" +
	"  and node_id = $2" +
	"  and stat_time = $3"

const selectAllNodeStatControlInfoSQL = "" +
	"select stat_type,stat_time,node_id,stat_state,message,create_time,update_time" +
	" from filer_stat_control" +
	" where stat_type = $1" +
	"  and stat_time = $2"

const insertStatControlInfoSQL = "" +
	"insert into filer_stat_control" +
	" (uu_id,stat_type,stat_time,node_id,stat_state,message,create_time,update_time)" +
	" values" +
	" (gen_random_uuid(),$1,$2,$3,$4,$5,$6,$7)"

const updateStatControlInfoSQL = "" +
	"update filer_stat_control" +
	" set stat_state = $1," +
	"     update_time = $2," +
	"     message = $3" +
	" where stat_type = $4" +
	"   and node_id = $5" +
	"   and stat_time = $6"

const deleteStatControlFromStatTimeByNodeIdSQL = "" +
	"delete from filer_stat_control " +
	" where stat_time >=$1" +
	"   and node_id = $2"

const deleteAllStatControlFromStatTimeSQL = "" +
	"delete from filer_stat_control " +
	" where stat_time >=$1"

type filerStatControlStatements struct {
	selectStatControlInfoStmt                 *sql.Stmt
	selectAllNodeStatControlInfoStmt          *sql.Stmt
	insertStatControlInfoStmt                 *sql.Stmt
	updateStatControlInfoStmt                 *sql.Stmt
	deleteStatControlFromStatTimeByNodeIdStmt *sql.Stmt
	deleteAllStatControlFromStatTimeStmt      *sql.Stmt
	deleteControlFromStatTimeForMainStmt      *sql.Stmt
}

func (s *filerStatControlStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(filerStatControlSchema)
	return err
}

func (s *filerStatControlStatements) prepare(db *sql.DB) (err error) {
	if s.selectStatControlInfoStmt, err = db.Prepare(selectStatControlInfoSQL); err != nil {
		return
	}
	if s.selectAllNodeStatControlInfoStmt, err = db.Prepare(selectAllNodeStatControlInfoSQL); err != nil {
		return
	}
	if s.insertStatControlInfoStmt, err = db.Prepare(insertStatControlInfoSQL); err != nil {
		return
	}
	if s.updateStatControlInfoStmt, err = db.Prepare(updateStatControlInfoSQL); err != nil {
		return
	}
	if s.deleteStatControlFromStatTimeByNodeIdStmt, err = db.Prepare(deleteStatControlFromStatTimeByNodeIdSQL); err != nil {
		return
	}
	if s.deleteAllStatControlFromStatTimeStmt, err = db.Prepare(deleteAllStatControlFromStatTimeSQL); err != nil {
		return
	}
	return
}

func (s *filerStatControlStatements) selectStatControlInfo(ctx context.Context, txn *sql.Tx, statType string, statTimeStr string, nodeId string) (*types.FilerStatControl, error) {
	var statControl types.FilerStatControl
	err := TxStmt(txn, s.selectStatControlInfoStmt).
		QueryRowContext(ctx, statType, statTimeStr, nodeId).
		Scan(
			&statControl.StatType,
			&statControl.StatTime,
			&statControl.NodeId,
			&statControl.StatState,
			&statControl.CreateTime,
			&statControl.UpdateTime,
			&statControl.Message,
		)
	if err != nil && err.Error() == "sql: no rows in result set" {
		err = nil
	}
	return &statControl, err
}

func (s *filerStatControlStatements) selectAllNodeStatControlInfo(ctx context.Context, txn *sql.Tx, statType string, statTime string) ([]types.FilerStatControl, error) {
	var list []types.FilerStatControl
	row, err := txn.Stmt(s.selectAllNodeStatControlInfoStmt).QueryContext(ctx, statType, statTime)
	defer row.Close()
	if err != nil {
		fmt.Print("selectStatisticBlockInfo error:", err)
		return nil, err
	}
	for row.Next() {
		var statControl types.FilerStatControl
		err := row.Scan(
			&statControl.StatType,
			&statControl.StatTime,
			&statControl.NodeId,
			&statControl.StatState,
			&statControl.CreateTime,
			&statControl.UpdateTime,
			&statControl.Message,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, statControl)
	}
	return list, nil
}

func (s *filerStatControlStatements) insertStatControlInfo(ctx context.Context, txn *sql.Tx, statControl *types.FilerStatControl) (err error) {
	statControl.CreateTime = strconv.FormatInt(util.TimeNow().Unix(), 10)
	statControl.UpdateTime = strconv.FormatInt(util.TimeNow().Unix(), 10)
	statControl.StatState = "0"
	r, err := TxStmt(txn, s.insertStatControlInfoStmt).
		ExecContext(ctx, //uuid
			statControl.StatType,
			statControl.StatTime,
			statControl.NodeId,
			statControl.StatState,
			statControl.Message,
			statControl.CreateTime,
			statControl.UpdateTime)
	if err != nil {
		return err
	}
	if i, _ := r.RowsAffected(); i < 1 {
		return fmt.Errorf("db No data was inserted")
	}
	return
}

func (s *filerStatControlStatements) updateStatControlInfo(ctx context.Context, txn *sql.Tx, statControl *types.FilerStatControl) (err error) {
	statControl.UpdateTime = strconv.FormatInt(util.TimeNow().Unix(), 10)
	stmt := TxStmt(txn, s.updateStatControlInfoStmt)
	r, err := stmt.ExecContext(ctx,
		statControl.StatState,
		statControl.UpdateTime,
		statControl.Message,
		statControl.StatType,
		statControl.NodeId,
		statControl.StatTime,
	)
	if err != nil {
		return err
	}
	if i, _ := r.RowsAffected(); i < 1 {
		return fmt.Errorf("db No data was modified")
	}
	return
}

func (s *filerStatControlStatements) deleteStatControlFromStatTimeByNodeId(ctx context.Context, txn *sql.Tx, statTime string, nodeId string) (err error) {
	stmt := TxStmt(txn, s.deleteStatControlFromStatTimeByNodeIdStmt)
	_, err = stmt.ExecContext(ctx, statTime, nodeId)
	if err != nil {
		return err
	}
	return
}
func (s *filerStatControlStatements) deleteAllStatControlFromStatTime(ctx context.Context, txn *sql.Tx, statTime string) (err error) {
	stmt := TxStmt(txn, s.deleteAllStatControlFromStatTimeStmt)
	_, err = stmt.ExecContext(ctx, statTime)
	if err != nil {
		return err
	}
	return
}
