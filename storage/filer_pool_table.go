package storage

import (
	"context"
	"database/sql"
	"entysquare/enty-cli/storage/types"
)

//goland:noinspection SqlNoDataSourceInspection
const filerPoolSchema = `
 	CREATE TABLE IF NOT EXISTS filer_pool (
	node_id    		TEXT PRIMARY KEY,
	node_name   	TEXT,
	location    	TEXT,
	mobile  		TEXT,
	email   		TEXT,
	create_time		TEXT,
	update_time		TEXT,
	is_valid  		TEXT
    );
	comment on column filer_pool.node_id is '节点id';
	comment on column filer_pool.node_name is '节点名';
	comment on column filer_pool.location is '位置';
	comment on column filer_pool.mobile is '手机';
	comment on column filer_pool.email is '邮箱';
	comment on column filer_pool.create_time is '创建时间';
	comment on column filer_pool.update_time is '更新时间';
	comment on column filer_pool.is_valid is '启用标志';
`

const insertFilerPoolSQL = "" +
	" INSERT INTO filer_pool " +
	" (node_id,node_name,location,mobile,email,create_time,update_time,is_valid) " +
	" VALUES " +
	" ($1, $2, $3, $4, $5, $6,$7,$8) "

const selectFilerPoolSQL = "" +
	" SELECT node_id,node_name,location,mobile,email,create_time,update_time,is_valid FROM filer_pool where is_valid='0'"
const selectHtFilerPoolSQL = "" +
	" SELECT node_id,node_name,location,mobile,email,create_time,update_time,is_valid FROM filer_pool where is_valid='0' and node_name = 'ht' "

type filerPoolStatements struct {
	insertFilerPoolStmt   *sql.Stmt
	selectFilerPoolStmt   *sql.Stmt
	selectHtFilerPoolStmt *sql.Stmt
}

func (s *filerPoolStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(filerPoolSchema)
	return err
}

func (s *filerPoolStatements) prepare(db *sql.DB) (err error) {
	if s.insertFilerPoolStmt, err = db.Prepare(insertFilerPoolSQL); err != nil {
		return
	}
	if s.selectFilerPoolStmt, err = db.Prepare(selectFilerPoolSQL); err != nil {
		return
	}
	if s.selectHtFilerPoolStmt, err = db.Prepare(selectHtFilerPoolSQL); err != nil {
		return
	}
	return
}

//node_id,node_name,location,mobile,email,valid_create,is_valid
func (s *filerPoolStatements) insertFilerPool(ctx context.Context, txn *sql.Tx, f *types.FilerPool) (err error) {
	_, err = txn.Stmt(s.insertFilerPoolStmt).
		ExecContext(ctx, //uuid
			f.NodeId,
			f.NodeName,
			f.Location,
			f.Mobile,
			f.Email,
			f.CreateTime,
			f.UpdateTime,
			f.IsValid)
	return
}

//node_id,node_name,location,mobile,email,valid_create,is_valid
func (s *filerPoolStatements) selectFilerPool() ([]types.FilerPool, error) {
	var list []types.FilerPool
	row, err := s.selectFilerPoolStmt.Query()
	defer row.Close()
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var f types.FilerPool
		err := row.Scan(
			&f.NodeId,
			&f.NodeName,
			&f.Location,
			&f.Mobile,
			&f.Email,
			&f.CreateTime,
			&f.UpdateTime,
			&f.IsValid)
		if err != nil {
			return nil, err
		}
		list = append(list, f)
	}
	return list, nil
}

func (s *filerPoolStatements) selectAllFilerPool(ctx context.Context, txn *sql.Tx) ([]types.FilerPool, error) {
	var list []types.FilerPool
	row, err := TxStmt(txn, s.selectFilerPoolStmt).QueryContext(ctx)
	defer row.Close()
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var f types.FilerPool
		err := row.Scan(
			&f.NodeId,
			&f.NodeName,
			&f.Location,
			&f.Mobile,
			&f.Email,
			&f.CreateTime,
			&f.UpdateTime,
			&f.IsValid)
		if err != nil {
			return nil, err
		}
		list = append(list, f)
	}
	return list, nil
}
func (s *filerPoolStatements) selectHtFilerPool(ctx context.Context, txn *sql.Tx) (htPool types.FilerPool, err error) {
	rows, err := TxStmt(txn, s.selectHtFilerPoolStmt).QueryContext(ctx)
	if err != nil {
		return types.FilerPool{}, err

	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(
			&htPool.NodeId,
			&htPool.NodeName,
			&htPool.Location,
			&htPool.Mobile,
			&htPool.Email,
			&htPool.CreateTime,
			&htPool.UpdateTime,
			&htPool.IsValid); err != nil {
			if err == sql.ErrNoRows {
				return
			}
		}
	}
	return htPool, rows.Err()
}
