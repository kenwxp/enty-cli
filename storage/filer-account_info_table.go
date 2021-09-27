package storage

import (
	"context"
	"database/sql"
	"entysquare/enty-cli/storage/types"
	"entysquare/enty-cli/util"
)

//CREATE SEQUENCE IF NOT EXISTS user_seq START 1;
//goland:noinspection SqlNoDataSourceInspection
const filerUserSchema = `
 CREATE TABLE IF NOT EXISTS filer_account_info  (
        filer_id      TEXT PRIMARY key ,
        filer_name    TEXT,
        reg_time      TEXT,
        mobile        TEXT,                      
        email         TEXT,
        is_valid      TEXT                          
    );
`

const selectAccountByNameSQL = "" +
	"SELECT filer_id , filer_name , reg_time, mobile, email, is_valid  FROM filer_account_info WHERE filer_name = $1 and is_valid = '0' "

type filerAccountsStatements struct {
	selectAccountByNameStmt *sql.Stmt
}

func (s *filerAccountsStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(filerUserSchema)
	return err
}

func (s *filerAccountsStatements) prepare(db *sql.DB) (err error) {
	if s.selectAccountByNameStmt, err = db.Prepare(selectAccountByNameSQL); err != nil {
		return
	}
	return
}

func (s *filerAccountsStatements) selectAccountByName(ctx context.Context, txn *sql.Tx, name string) (account *types.FilerAccountInfo, err error) {
	row := util.TxStmt(txn, s.selectAccountByNameStmt).QueryRowContext(ctx, name)
	account = &types.FilerAccountInfo{}
	if err = row.Scan(
		&account.FilerId,
		&account.FilerName,
		&account.RegTime,
		&account.Mobile,
		&account.Email,
		&account.IsValid,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
	}
	return account, nil
}
