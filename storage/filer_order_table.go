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
const filerOrderSchema = `
 CREATE TABLE IF NOT EXISTS filer_order
(
	order_id    TEXT PRIMARY KEY,
	filer_id    TEXT,
	pay_flow    TEXT,
	product_id  TEXT,
	hold_power  TEXT,
	pay_amount  TEXT,
	order_time  TEXT,
	update_time	TEXT,
	valid_time	TEXT,
	end_time    TEXT,
	order_state TEXT
);
comment on column filer_order.order_id is '订单id';
comment on column filer_order.filer_id is 'filer id';
comment on column filer_order.pay_flow is '交易流水（pay产生）';
comment on column filer_order.product_id is '产品id';
comment on column filer_order.hold_power is '持有算力';
comment on column filer_order.pay_amount is '支付金额';
comment on column filer_order.order_time is '下单时间';
comment on column filer_order.update_time is '更新时间';
comment on column filer_order.valid_time is '生效时间';
comment on column filer_order.end_time is '结束时间';
comment on column filer_order.order_state is '持仓状态 0-待生效 1-已生效 2-持仓已结束 9-收益结算完成';                 
`

const selectAllInProgressOrderSQL = "select o.order_id, " +
	" o.filer_id, " +
	" o.pay_flow, " +
	" o.product_id, " +
	" p.node_id, " +
	" p.period, " +
	" p.valid_plan, " +
	" p.service_rate, " +
	" o.hold_power, " +
	" o.pay_amount, " +
	" o.order_time, " +
	" o.update_time," +
	" o.end_time, " +
	" o.order_state" +
	" from filer_order o, " +
	" filer_product p " +
	" where o.product_id = p.product_id " +
	" and order_state != '9' " +
	" and order_time < $1" +
	" and p.node_id = $2"

const updateFilerOrderStateSQL = "" +
	"update filer_order" +
	" set order_state = $1," +
	"     update_time = $2" +
	" where order_id = $3"

const selectOrderByFilerIdSQL = "select o.order_id, " +
	" o.filer_id, " +
	" o.pay_flow, " +
	" o.product_id, " +
	" o.hold_power, " +
	" o.pay_amount, " +
	" o.order_time, " +
	" o.update_time," +
	" o.end_time, " +
	" o.order_state " +
	" from filer_order o " +
	" where  order_state = $1  and filer_id = $2 "
const insertFilerOrderSQL = "INSERT INTO filer_order" +
	" (order_id,filer_id,pay_flow,product_id,hold_power,pay_amount,order_time,update_time,valid_time,end_time,order_state)" +
	" VALUES" +
	" (gen_random_uuid(), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

const selectInvalidOrderListByFilerIdSQL = "" +
	" select o.order_id," +
	"        o.filer_id," +
	"        o.pay_flow," +
	"        o.hold_power," +
	"        o.pay_amount," +
	"        o.order_time," +
	"        o.valid_time," +
	"        o.order_state," +
	"        p.product_name," +
	"        p.period" +
	" from filer_product p," +
	"      filer_order o" +
	" where p.product_id = o.product_id" +
	"   and o.filer_id = $1" +
	"   and o.order_state='0' "

type filerOrderStatements struct {
	selectAllInProgressOrderStmt        *sql.Stmt
	updateFilerOrderStateStmt           *sql.Stmt
	selectOrderByFilerIdStmt            *sql.Stmt
	insertOrderStmt                     *sql.Stmt
	selectInvalidOrderListByFilerIdStmt *sql.Stmt
	//updatePasswordByPhoneStmt           *sql.Stmt
}

func (s *filerOrderStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(filerOrderSchema)
	return err
}

func (s *filerOrderStatements) prepare(db *sql.DB) (err error) {
	if s.selectAllInProgressOrderStmt, err = db.Prepare(selectAllInProgressOrderSQL); err != nil {
		return
	}
	if s.updateFilerOrderStateStmt, err = db.Prepare(updateFilerOrderStateSQL); err != nil {
		return
	}
	if s.selectOrderByFilerIdStmt, err = db.Prepare(selectOrderByFilerIdSQL); err != nil {
		return
	}
	if s.insertOrderStmt, err = db.Prepare(insertFilerOrderSQL); err != nil {
		return
	}
	if s.selectInvalidOrderListByFilerIdStmt, err = db.Prepare(selectInvalidOrderListByFilerIdSQL); err != nil {
		return
	}
	return
}

func (s *filerOrderStatements) selectAllInProgressOrderByNodeId(ctx context.Context, txn *sql.Tx, statTimeStr string, nodeId string) ([]types.FilerOrder, error) {
	var list []types.FilerOrder
	statTime := util.TimeStringToTime(statTimeStr, "00:00:00", "")
	row, err := txn.Stmt(s.selectAllInProgressOrderStmt).QueryContext(ctx, statTime.AddDate(0, 0, 1).Unix(), nodeId)
	defer row.Close()
	if err != nil {
		fmt.Print("selectStatisticBlockInfo error:", err)
		return nil, err
	}
	for row.Next() {
		var item types.FilerOrder
		err := row.Scan(
			&item.OrderId,
			&item.FilerId,
			&item.PayFlow,
			&item.ProductId,
			&item.NodeId,
			&item.Period,
			&item.ValidPlan,
			&item.ServiceRate,
			&item.HoldPower,
			&item.PayAmount,
			&item.OrderTime,
			&item.UpdateTime,
			&item.EndTime,
			&item.OrderState,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

func (s *filerOrderStatements) updateFilerOrderState(ctx context.Context, txn *sql.Tx, orderState string, orderId string) (err error) {
	updateTime := strconv.FormatInt(util.TimeNow().Unix(), 10)
	stmt := util.TxStmt(txn, s.updateFilerOrderStateStmt)
	r, err := stmt.ExecContext(ctx, orderState, updateTime, orderId)
	if err != nil {
		return err
	}
	if i, _ := r.RowsAffected(); i < 1 {
		return fmt.Errorf("db No data was modified")
	}
	return
}
func (s *filerOrderStatements) selectOrderListByFilerId(ctx context.Context, txn *sql.Tx, state int, filerId int64) (list map[string]types.FilerOrder, err error) {
	rows, err := util.TxStmt(txn, s.selectOrderByFilerIdStmt).QueryContext(ctx, state, filerId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	maps := make(map[string]types.FilerOrder)
	for rows.Next() {
		order := types.FilerOrder{}
		err = rows.Scan(
			&order.OrderId,
			&order.FilerId,
			&order.PayFlow,
			&order.ProductId,
			&order.HoldPower,
			&order.PayAmount,
			&order.OrderTime,
			&order.UpdateTime,
			&order.EndTime,
			&order.OrderState,
		)
		if err != nil {
			return nil, err
		}
		maps[order.OrderId] = order
	}
	list = maps
	return
}
func (s *filerOrderStatements) insertFilerOrder(ctx context.Context, txn *sql.Tx, item *types.FilerOrder) (err error) {
	item.UpdateTime = strconv.FormatInt(util.TimeNow().Unix(), 10)
	_, err = txn.Stmt(s.insertOrderStmt).
		ExecContext(ctx, //uuid
			&item.FilerId,
			&item.PayFlow,
			&item.ProductId,
			&item.HoldPower,
			&item.PayAmount,
			&item.OrderTime,
			&item.UpdateTime,
			&item.ValidTime,
			&item.EndTime,
			&item.OrderState,
		)
	return
}

func (s *filerOrderStatements) selectInvalidOrderListByFilerId(ctx context.Context, txn *sql.Tx, filerId string) ([]types.FilerOrderShow, error) {
	var list []types.FilerOrderShow
	row, err := util.TxStmt(txn, s.selectInvalidOrderListByFilerIdStmt).QueryContext(ctx, filerId)
	defer row.Close()
	if err != nil {
		fmt.Print("selectInvalidOrderListByFilerId error:", err)
		return nil, err
	}
	for row.Next() {
		var item types.FilerOrderShow
		err := row.Scan(
			&item.OrderId,
			&item.PayFlow,
			&item.HoldPower,
			&item.PayAmount,
			&item.OrderTime,
			&item.ValidTime,
			&item.OrderState,
			&item.ProductName,
			&item.Period,
		)
		if err != nil {
			return nil, err
		}
		item.ValidDays = util.GetDurationDaysForTimestamp(item.ValidTime, strconv.FormatInt(util.TimeNow().Unix(), 10))
		list = append(list, item)
	}
	return list, nil
}
