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
const filerBalanceFlowSchema = `
	CREATE TABLE IF NOT EXISTS filer_balance_flow
	(
		uu_id       TEXT PRIMARY KEY,
		filer_id    TEXT,
		oper_type   TEXT,
		amount      TEXT,
		create_time TEXT
	);
	comment on column filer_balance_flow.uu_id is 'uu_id';
	comment on column filer_balance_flow.filer_id is 'filer id';
	comment on column filer_balance_flow.oper_type is '操作类型 0-收益入账 1-手工入账 2-提现出账';
	comment on column filer_balance_flow.amount is '金额';
	comment on column filer_balance_flow.create_time is '创建时间 时间戳';
`
const selectStatisticBalanceFlowListByFilerIdSQL = "" +
	" select filer_id," +
	"        oper_type," +
	"        sum(amount::int8)" +
	" from filer_balance_flow" +
	" where filer_id = $1 " +
	" and create_time::int8 >= $2 " +
	" and create_time::int8 < $3 " +
	" group by filer_id, oper_type"

const insertBalanceFlowSQL = "" +
	" insert into filer_balance_flow" +
	"    (uu_id, filer_id, oper_type, amount, create_time)" +
	" values " +
	"    (gen_random_uuid(), $1, $2, $3, $4)"

const deleteBalanceFlowByOperTypeSQL = "" +
	" delete" +
	" from filer_balance_flow" +
	" where filer_id = $1" +
	"  and create_time::int8 >= $2" +
	"  and create_time::int8 < $3" +
	"  and oper_type = $4"

type filerBalanceFlowStatements struct {
	selectStatisticBalanceFlowListByFilerIdStmt *sql.Stmt
	insertBalanceFlowStmt                       *sql.Stmt
	deleteBalanceFlowByOperTypeStmt             *sql.Stmt
}

func (s *filerBalanceFlowStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(filerBalanceFlowSchema)
	return err
}

func (s *filerBalanceFlowStatements) prepare(db *sql.DB) (err error) {
	if s.selectStatisticBalanceFlowListByFilerIdStmt, err = db.Prepare(selectStatisticBalanceFlowListByFilerIdSQL); err != nil {
		return
	}
	if s.insertBalanceFlowStmt, err = db.Prepare(insertBalanceFlowSQL); err != nil {
		return
	}
	if s.deleteBalanceFlowByOperTypeStmt, err = db.Prepare(deleteBalanceFlowByOperTypeSQL); err != nil {
		return
	}
	return
}
func (s *filerBalanceFlowStatements) selectStatisticBalanceFlowListByFilerId(ctx context.Context, txn *sql.Tx, filerId string, statTimeStr string) ([]types.FilerBalanceFlow, error) {
	var list []types.FilerBalanceFlow
	statTime := util.TimeStringToTime(statTimeStr, "00:00:00", "")
	row, err := util.TxStmt(txn, s.selectStatisticBalanceFlowListByFilerIdStmt).QueryContext(ctx, filerId, statTime.Unix(), statTime.AddDate(0, 0, 1).Unix())
	defer row.Close()
	if err != nil {
		fmt.Print("selectStatisticBalanceFlowListByFilerId error:", err)
		return nil, err
	}
	for row.Next() {
		var item types.FilerBalanceFlow
		err := row.Scan(
			&item.FilerId,
			&item.OperType,
			&item.Amount,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

func (s *filerBalanceFlowStatements) insertBalanceFlow(ctx context.Context, txn *sql.Tx, f *types.FilerBalanceFlow) (err error) {
	if f.OperType != "0" {
		f.CreateTime = strconv.FormatInt(util.TimeNow().Unix(), 10)
	}
	r, err := txn.Stmt(s.insertBalanceFlowStmt).
		ExecContext(ctx, //uuid
			f.FilerId,
			f.OperType,
			f.Amount,
			f.CreateTime,
		)
	if err != nil {
		return err
	}
	if i, _ := r.RowsAffected(); i < 1 {
		return fmt.Errorf("db No data was created")
	}
	return err
}

func (s *filerBalanceFlowStatements) deleteBalanceFlowByOperType(ctx context.Context, txn *sql.Tx, filerId string, statTimeStr string, operType string) (err error) {
	statTime := util.TimeStringToTime(statTimeStr, "00:00:00", "")
	stmt := util.TxStmt(txn, s.deleteBalanceFlowByOperTypeStmt)
	r, err := stmt.ExecContext(ctx, filerId, statTime.Unix(), statTime.AddDate(0, 0, 1).Unix(), operType)
	if err != nil {
		return err
	}
	if i, _ := r.RowsAffected(); i < 1 {
		return fmt.Errorf("db No data was removed")
	}
	return
}
