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
const filerOrderIncomeSchema = `
 CREATE TABLE IF NOT EXISTS filer_order_income  (
	uu_id 				TEXT PRIMARY KEY, 
	order_id 			TEXT,
	filer_id			TEXT,
	node_id				TEXT,
	total_income 		TEXT,
	freeze_income 		TEXT,
	total_available_income 	TEXT,
	day_available_income 	TEXT,
	day_raise_income    TEXT,
	day_direct_income	TEXT,
	day_release_income	TEXT,
	stat_time 			TEXT,
	create_time 		TEXT,
	update_time 		TEXT
);
comment on column filer_order_income.uu_id is 'uu_id';
comment on column filer_order_income.order_id is '订单id';
comment on column filer_order_income.filer_id is 'filer ID';
comment on column filer_order_income.node_id is '节点ID';
comment on column filer_order_income.total_income is '总收益';
comment on column filer_order_income.freeze_income is '冻结收益';
comment on column filer_order_income.total_available_income is '累计已释放收益 = 昨日累计收益+当日释放收益 ';
comment on column filer_order_income.day_available_income is '可用收益(当日直接25%+往日75%平账 ';
comment on column filer_order_income.day_raise_income is '日产出收益（当日直接25%+当日冻结75%）';
comment on column filer_order_income.day_direct_income is '当日直接收益（25%';
comment on column filer_order_income.day_release_income is '当日线性释放收益（往日75%平账）';
comment on column filer_order_income.stat_time is '统计时间（yyyy-MM-dd）';
comment on column filer_order_income.create_time is '创建时间 时间戳';
comment on column filer_order_income.update_time is '更新时间 时间戳';
`
const selectFilerOrderIncomeByStatTimeSQL = "" +
	"select uu_id," +
	"       order_id," +
	"       filer_id," +
	"       node_id," +
	"       total_income," +
	"       freeze_income," +
	"       total_available_income," +
	"       day_available_income," +
	"       day_raise_income," +
	"       day_direct_income," +
	"       day_release_income," +
	"       stat_time," +
	"       create_time," +
	"       update_time" +
	" from filer_order_income" +
	" where 1 = 1 " +
	"	and stat_time = $1" +
	" 	and order_id = $2" +
	" 	and node_id = $3"

const insertFilerOrderIncomeSQL = "INSERT INTO filer_order_income" +
	" (uu_id,order_id,filer_id,node_id,total_income,freeze_income,total_available_income,day_available_income,day_raise_income,day_direct_income,day_release_income,stat_time,create_time,update_time)" +
	" VALUES" +
	" (gen_random_uuid(),$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)"

const selectFilerOrderIncomeByOrderIdSQL = "" +
	"select order_id," +
	"       total_income," +
	"       freeze_income," +
	"       total_available_income," +
	"       day_available_income," +
	"       day_raise_income," +
	"       day_direct_income," +
	"       day_release_income," +
	"       stat_time," +
	"       create_time," +
	"       update_time" +
	" from filer_order_income" +
	" where order_id = $1"

const selectStatisticOrderIncomeByNodeIdSQL = "" +
	" select t.filer_id," +
	"        sum(t.hold_power::int8)," +
	"        sum(t.pay_amount::int8)," +
	"        sum(t.total_income::int8)," +
	"        sum(t.freeze_income::int8)," +
	"        sum(t.total_available_income::int8)," +
	"        sum(t.day_available_income::int8)," +
	"        sum(t.day_raise_income::int8)," +
	"        sum(t.day_direct_income::int8)," +
	"        sum(t.day_release_income::int8)" +
	" from (select i.filer_id," +
	"              o.hold_power," +
	"              o.pay_amount," +
	"              i.total_income," +
	"              i.freeze_income," +
	"              i.total_available_income," +
	"              i.day_available_income," +
	"              i.day_raise_income," +
	"              i.day_direct_income," +
	"              i.day_release_income" +
	"       from filer_order o," +
	"            filer_order_income i" +
	"       where o.order_id = i.order_id" +
	"         and i.stat_time = $1" +
	"		  and i.node_id = $2) t" +
	" group by t.filer_id"

const deleteFilerOrderIncomeFromStatTimeSQL = "" +
	" delete from filer_order_income" +
	" where stat_time >= $1" +
	" 	and node_id = $2"

const selectFilerOrderListWithIncomeInfoByFilerIdAndOrderStateSQL = "" +
	"select o.order_id," +
	"        o.pay_flow," +
	"        o.hold_power," +
	"        o.pay_amount," +
	"        o.order_time," +
	"        o.valid_time," +
	"        o.order_state," +
	"        p.product_name," +
	"        p.period," +
	"        i.total_income," +
	"        i.total_available_income," +
	"        i.freeze_income," +
	"        i.stat_time," +
	"        i.day_available_income," +
	"        i.day_raise_income," +
	"        i.day_direct_income," +
	"        i.day_release_income" +
	" from filer_product p," +
	"      filer_order o," +
	"      (select order_id," +
	"              total_income," +
	"              total_available_income," +
	"              freeze_income," +
	"              stat_time," +
	"              day_available_income," +
	"              day_raise_income," +
	"              day_direct_income," +
	"              day_release_income," +
	"              row_number() over (partition by order_id order by stat_time desc) rownum" +
	"       from filer_order_income) i" +
	" where p.product_id = o.product_id" +
	"   and o.order_id = i.order_id " +
	"   and o.filer_id = $1" +
	"   and o.order_state = $2" +
	"   and i.rownum =1"

type filerOrderIncomeStatements struct {
	selectFilerOrderIncomeByStatTimeStmt                         *sql.Stmt
	insertFilerOrderIncomeStmt                                   *sql.Stmt
	selectFilerOrderIncomeByOrderIdStmt                          *sql.Stmt
	selectStatisticOrderIncomeByNodeIdStmt                       *sql.Stmt
	deleteFilerOrderIncomeFromStatTimeStmt                       *sql.Stmt
	selectFilerOrderListWithIncomeInfoByFilerIdAndOrderStateStmt *sql.Stmt
}

func (s *filerOrderIncomeStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(filerOrderIncomeSchema)
	return err
}

func (s *filerOrderIncomeStatements) prepare(db *sql.DB) (err error) {
	if s.selectFilerOrderIncomeByStatTimeStmt, err = db.Prepare(selectFilerOrderIncomeByStatTimeSQL); err != nil {
		return
	}
	if s.insertFilerOrderIncomeStmt, err = db.Prepare(insertFilerOrderIncomeSQL); err != nil {
		return
	}
	if s.selectFilerOrderIncomeByOrderIdStmt, err = db.Prepare(selectFilerOrderIncomeByOrderIdSQL); err != nil {
		return
	}
	if s.selectStatisticOrderIncomeByNodeIdStmt, err = db.Prepare(selectStatisticOrderIncomeByNodeIdSQL); err != nil {
		return
	}
	if s.deleteFilerOrderIncomeFromStatTimeStmt, err = db.Prepare(deleteFilerOrderIncomeFromStatTimeSQL); err != nil {
		return
	}
	if s.selectFilerOrderListWithIncomeInfoByFilerIdAndOrderStateStmt, err = db.Prepare(selectFilerOrderListWithIncomeInfoByFilerIdAndOrderStateSQL); err != nil {
		return
	}
	return
}
func (s *filerOrderIncomeStatements) selectFilerOrderIncomeByStatTime(ctx context.Context, txn *sql.Tx, statTime string, orderId string, nodeId string) ([]types.FilerOrderIncome, error) {
	var list []types.FilerOrderIncome
	row, err := txn.Stmt(s.selectFilerOrderIncomeByStatTimeStmt).QueryContext(ctx, statTime, orderId, nodeId)
	defer row.Close()
	if err != nil {
		fmt.Print("selectStatisticBlockInfo error:", err)
		return nil, err
	}
	for row.Next() {
		var item types.FilerOrderIncome
		err := row.Scan(
			&item.Uuid,
			&item.OrderId,
			&item.FilerId,
			&item.NodeId,
			&item.TotalIncome,
			&item.FreezeIncome,
			&item.TotalAvailableIncome,
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

func (s *filerOrderIncomeStatements) insertFilerOrderIncome(ctx context.Context, txn *sql.Tx, f *types.FilerOrderIncome) (err error) {
	f.CreateTime = strconv.FormatInt(util.TimeNow().Unix(), 10)
	f.UpdateTime = strconv.FormatInt(util.TimeNow().Unix(), 10)
	_, err = txn.Stmt(s.insertFilerOrderIncomeStmt).
		ExecContext(ctx, //uuid
			f.OrderId,
			f.FilerId,
			f.NodeId,
			f.TotalIncome,
			f.FreezeIncome,
			f.TotalAvailableIncome,
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
func (s *filerOrderIncomeStatements) selectFilerOrderIncomeByOrderId(ctx context.Context, txn *sql.Tx, orderId string) (orderIncome types.FilerOrderIncome, err error) {
	rows, err := util.TxStmt(txn, s.selectFilerOrderIncomeByOrderIdStmt).QueryContext(ctx, orderId)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(
			&orderIncome.OrderId,
			&orderIncome.TotalIncome,
			&orderIncome.FreezeIncome,
			&orderIncome.TotalAvailableIncome,
			&orderIncome.DayAvailableIncome,
			&orderIncome.DayRaiseIncome,
			&orderIncome.DayDirectIncome,
			&orderIncome.DayReleaseIncome,
			&orderIncome.StatTime,
			&orderIncome.CreateTime,
			&orderIncome.UpdateTime,
		); err != nil {
			if err == sql.ErrNoRows {
				return
			}
		}
	}
	return orderIncome, rows.Err()
}

func (s *filerOrderIncomeStatements) selectStatisticOrderIncomeByNodeId(ctx context.Context, txn *sql.Tx, statTime string, nodeId string) ([]types.FilerAccountIncome, error) {
	var list []types.FilerAccountIncome
	row, err := txn.Stmt(s.selectStatisticOrderIncomeByNodeIdStmt).QueryContext(ctx, statTime, nodeId)
	defer row.Close()
	if err != nil {
		fmt.Print("selectStatisticOrderIncomeInfo error:", err)
		return nil, err
	}
	for row.Next() {
		var item types.FilerAccountIncome
		err := row.Scan(
			&item.FilerId,
			&item.HoldPower,
			&item.PledgeSum,
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

func (s *filerOrderIncomeStatements) deleteFilerOrderIncomeFromStatTime(ctx context.Context, txn *sql.Tx, statTime string, nodeId string) (err error) {
	stmt := util.TxStmt(txn, s.deleteFilerOrderIncomeFromStatTimeStmt)
	_, err = stmt.ExecContext(ctx, statTime, nodeId)
	if err != nil {
		return err
	}
	return
}

func (s *filerOrderIncomeStatements) selectFilerOrderListWithIncomeInfoByFilerIdAndOrderState(ctx context.Context, txn *sql.Tx, filerId string, orderState string) ([]types.FilerOrderShow, error) {
	var list []types.FilerOrderShow
	row, err := util.TxStmt(txn, s.selectFilerOrderListWithIncomeInfoByFilerIdAndOrderStateStmt).QueryContext(ctx, filerId, orderState)
	defer row.Close()
	if err != nil {
		fmt.Print("selectFilerOrderListWithIncomeInfoByFilerIdAndOrderState error:", err)
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
			&item.TotalIncome,
			&item.TotalAvailableIncome,
			&item.FreezeIncome,
			&item.StatTime,
			&item.DayAvailableIncome,
			&item.DayRaiseIncome,
			&item.DayDirectIncome,
			&item.DayReleaseIncome,
		)
		if err != nil {
			return nil, err
		}
		stateTimeTs := util.TimeStringToTime(item.StatTime, "00:00:00", "").Unix()
		item.ValidDays = util.CalculateString(util.GetDurationDaysForTimestamp(item.ValidTime, strconv.FormatInt(stateTimeTs, 10)), "1", "add")
		list = append(list, item)
	}
	return list, nil
}
