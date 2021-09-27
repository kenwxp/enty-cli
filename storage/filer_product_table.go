package storage

import (
	"context"
	"database/sql"
	"entysquare/filer-backend/storage/types"
	"entysquare/filer-backend/util"
	"strconv"
)

//goland:noinspection SqlNoDataSourceInspection
const filerProductSchema = `
 	CREATE TABLE IF NOT EXISTS filer_product  (
		product_id    TEXT PRIMARY KEY,
		product_name  TEXT,
		node_id       TEXT,
		cur_id        TEXT,
		period		  TEXT,
		valid_plan	  TEXT,
		price         TEXT,
		pledge_max	  TEXT,
		service_rate  TEXT,
		node1		  TEXT,
		node2		  TEXT,
		shelve_time	  TEXT,
		create_time   TEXT,
		update_time	  TEXT,
		product_state TEXT,
		is_valid	  TEXT
	);

	comment on column filer_product.product_id is '产品ID';
	comment on column filer_product.product_name is '产品名称';
	comment on column filer_product.node_id is '节点';
	comment on column filer_product.cur_id is '币种类型';
	comment on column filer_product.period is '挖矿周期';
	comment on column filer_product.valid_plan is '生效时间';
	comment on column filer_product.price is '每T质押';
	comment on column filer_product.pledge_max is '质押需求额';
	comment on column filer_product.service_rate is '服务费率';
	comment on column filer_product.node1 is '说明（基本规则）';
	comment on column filer_product.node2 is '说明（联合挖矿说明）';
	comment on column filer_product.shelve_time is '上架时间';
	comment on column filer_product.create_time is '创建时间';
	comment on column filer_product.update_time is '更新时间';
	comment on column filer_product.product_state is '产品状态 0-进行中 9-已失效';
	comment on column filer_product.is_valid is '启用标志 0--启用 1-废弃';
	
`
const insertFilerProductSQL = "" +
	" INSERT INTO filer_product " +
	" (product_id,product_name,node_id,cur_id,period,valid_plan,price,pledge_max,service_rate,node1,node2,shelve_time,create_time,update_time,product_state,is_valid) " +
	" VALUES " +
	" ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) "
const updateFilerProductStateSql = "" +
	"UPDATE filer_product" +
	" set product_state = $2," +
	" is_valid = $3 " +
	" where product_id = $1 "
const selectProductInfoByProductIdSql = "" +
	"select product_id,product_name,node_id,cur_id,period,valid_plan,price,pledge_max,service_rate,node1,node2,shelve_time,create_time,update_time,product_state,is_valid" +
	" FROM filer_product WHERE product_id = $1 and is_valid='0'"
const selectProductListByCurIDSql = "" +
	"select product_id,product_name,node_id,cur_id,period,valid_plan,price,pledge_max,service_rate,node1,node2,shelve_time,create_time,update_time,product_state,is_valid" +
	" FROM filer_product WHERE cur_id = $1 and is_valid='0'"
const selectProductListByStateSql = "" +
	"select product_id,product_name,node_id,cur_id,period,valid_plan,price,pledge_max,service_rate,node1,node2,shelve_time,create_time,update_time,product_state,is_valid" +
	" FROM filer_product WHERE product_state = $1 and is_valid='0'"

//const setUserAddressByPhoneNumSQL = "" +
//	"UPDATE user_data SET usdt_address = $1, hsf_address = $2 WHERE phone_num = $3"

type filerProductStatements struct {
	insertFilerProductStmt           *sql.Stmt
	updateFilerProductStateStmt      *sql.Stmt
	selectProductInfoByProductIdStmt *sql.Stmt
	selectProductListByCurIDStmt     *sql.Stmt
	selectProductListByStateStmt     *sql.Stmt
}

func (s *filerProductStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(filerProductSchema)
	return err
}

func (s *filerProductStatements) prepare(db *sql.DB) (err error) {
	if s.insertFilerProductStmt, err = db.Prepare(insertFilerProductSQL); err != nil {
		return
	}
	if s.updateFilerProductStateStmt, err = db.Prepare(updateFilerProductStateSql); err != nil {
		return
	}
	if s.selectProductInfoByProductIdStmt, err = db.Prepare(selectProductInfoByProductIdSql); err != nil {
		return
	}
	if s.selectProductListByCurIDStmt, err = db.Prepare(selectProductListByCurIDSql); err != nil {
		return
	}
	if s.selectProductListByStateStmt, err = db.Prepare(selectProductListByStateSql); err != nil {
		return
	}
	return
}

//stat_id,node_id,block_num,block_gain,power,gain_per_tib,state,create_time,update_time
func (s *filerProductStatements) insertFilerProduct(ctx context.Context, txn *sql.Tx, f *types.FilerProduct) (err error) {
	f.CreateTime = strconv.FormatInt(util.TimeNow().Unix(), 10)
	f.UpdateTime = strconv.FormatInt(util.TimeNow().Unix(), 10)
	f.ProductState = "0"
	f.IsValid = "0"
	_, err = txn.Stmt(s.insertFilerProductStmt).
		ExecContext(ctx, //uuid
			f.ProductId,
			f.ProductName,
			f.NodeId,
			f.CurId,
			f.Period,
			f.ValidPlan,
			f.Price,
			f.PledgeMax,
			f.ServiceRate,
			f.Note1,
			f.Note2,
			f.ShelveTime,
			f.CreateTime,
			f.UpdateTime,
			f.ProductState,
			f.IsValid) //处理标志 0-待处理 1-线性释放 9-已处理)
	return
}
func (s *filerProductStatements) updateProductState(ctx context.Context, txn *sql.Tx, productId string, productState string, isValid string) (err error) {
	stmt := TxStmt(txn, s.updateFilerProductStateStmt)
	_, err = stmt.ExecContext(ctx, productId, productState, isValid)
	if err != nil {
		return err
	}
	return
}
func (s *filerProductStatements) selectProductInfoByProductID(ctx context.Context, txn *sql.Tx, productId string) (product types.FilerProduct, err error) {
	rows, err := TxStmt(txn, s.selectProductInfoByProductIdStmt).QueryContext(ctx, productId)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(
			&product.ProductId,
			&product.ProductName,
			&product.NodeId,
			&product.CurId,
			&product.Period,
			&product.ValidPlan,
			&product.Price,
			&product.PledgeMax,
			&product.ServiceRate,
			&product.Note1,
			&product.Note2,
			&product.ShelveTime,
			&product.CreateTime,
			&product.UpdateTime,
			&product.ProductState,
			&product.IsValid,
		); err != nil {
			if err == sql.ErrNoRows {
				return
			}
		}
	}
	return product, rows.Err()
}
func (s *filerProductStatements) selectProductListByCurId(ctx context.Context, txn *sql.Tx, curId string) (list map[string]types.FilerProduct, err error) {
	rows, err := TxStmt(txn, s.selectProductListByCurIDStmt).QueryContext(ctx, curId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	maps := make(map[string]types.FilerProduct)
	for rows.Next() {
		product := types.FilerProduct{}
		err = rows.Scan(
			&product.ProductId,
			&product.ProductName,
			&product.NodeId,
			&product.CurId,
			&product.Period,
			&product.ValidPlan,
			&product.Price,
			&product.PledgeMax,
			&product.ServiceRate,
			&product.Note1,
			&product.Note2,
			&product.ShelveTime,
			&product.CreateTime,
			&product.UpdateTime,
			&product.ProductState,
			&product.IsValid)
		if err != nil {
			return nil, err
		}
		maps[product.ProductId] = product
	}
	list = maps
	return
}
func (s *filerProductStatements) selectProductListByState(ctx context.Context, txn *sql.Tx, state string) (list map[string]types.FilerProduct, err error) {
	rows, err := TxStmt(txn, s.selectProductListByStateStmt).QueryContext(ctx, state)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	maps := make(map[string]types.FilerProduct)
	for rows.Next() {
		product := types.FilerProduct{}
		err = rows.Scan(
			&product.ProductId,
			&product.ProductName,
			&product.NodeId,
			&product.CurId,
			&product.Period,
			&product.ValidPlan,
			&product.Price,
			&product.PledgeMax,
			&product.ServiceRate,
			&product.Note1,
			&product.Note2,
			&product.ShelveTime,
			&product.CreateTime,
			&product.UpdateTime,
			&product.ProductState,
			&product.IsValid)
		if err != nil {
			return nil, err
		}
		maps[product.ProductId] = product
	}
	list = maps
	return
}
