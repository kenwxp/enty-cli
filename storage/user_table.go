package storage

import (
	"context"
	"database/sql"
	"entysquare/enty-cli/storage/types"
)

//CREATE SEQUENCE IF NOT EXISTS user_seq START 1;
//goland:noinspection SqlNoDataSourceInspection
const userSchema = `
 CREATE TABLE IF NOT EXISTS filer_account  (
        filer_id      serial not null,
        pay_id        INTEGER,
        token         text
    );
create unique index IF NOT EXISTS pay_id_index on filer_account (pay_id);
`
const insertAccountSQL = "" +
	"INSERT INTO filer_account(pay_id, token) VALUES ($1, $2)"

const updateFilerAccountTokenSQL = "" +
	"UPDATE filer_account SET token = $2 WHERE pay_id = $1"

const updatePasswordByMailSQL = "" +
	"UPDATE filer_account SET password = $1 WHERE email = $2"

const selectAccountByPayIDSQL = "" +
	"SELECT filer_id , pay_id , token FROM filer_account WHERE pay_id = $1"

const selectAccountByTokenSQL = "" +
	"SELECT filer_id , pay_id , token FROM filer_account WHERE token LIKE $1"

const selectAccountByUserPhoneSQL = "" +
	"SELECT user_id, user_name, password, phone_num, email, gender, age, " +
	"phone_ver_flag, mail_ver_flag, google_ver_flag, google_key, gesture_ver_flag, " +
	"gesture_key, finger_ver_flag, finger_key, fish_ver_flag, fish_code, certify_num, " +
	"certify_front, certify_back, user_real_name, created_ts, " +
	"address, u_id, u_id_flag, pay_pwd, pay_pwd_flag, kyc_flag, hsf_address, hsf_num_all, hsf_num_ok, usdt_address, usdt_num_all, usdt_num_ok, birthday,user_photo " +
	"FROM user_data WHERE phone_num = $1"

const selectAccountByUserMailSQL = "" +
	"SELECT user_id, user_name, password, phone_num, email, gender, age, " +
	"phone_ver_flag, mail_ver_flag, google_ver_flag, google_key, gesture_ver_flag, " +
	"gesture_key, finger_ver_flag, finger_key, fish_ver_flag, fish_code, certify_num, " +
	"certify_front, certify_back, user_real_name, created_ts, " +
	"address, u_id, u_id_flag, pay_pwd, pay_pwd_flag, kyc_flag, hsf_address, hsf_num_all, hsf_num_ok, usdt_address, usdt_num_all, usdt_num_ok, birthday,user_photo " +
	"FROM user_data WHERE email = $1"

const updateAccountGestureSQL = "" +
	"UPDATE user_data SET gesture_key = $1 WHERE user_id = $2"

const updateAccountGestureFlagSQL = "" +
	"UPDATE user_data SET gesture_ver_flag = $1 WHERE user_id = $2"

const updateAccountGoogleCodeSQL = "" +
	"UPDATE user_data SET google_key = $2, google_ver_flag = $3 WHERE user_id = $1"

const updateAccountGoogleFlagSQL = "" +
	"UPDATE user_data SET gesture_ver_flag = $1 WHERE user_id = $2"

const updateAccountFingerSQL = "" +
	"UPDATE user_data SET finger_key = $1, finger_ver_flag = $2 WHERE user_id = $3"

const selectAccountWalletByIdSQL = "" +
						"SELECT usdt_address,usdt_num_all,usdt_num_ok,hsf_address,hsf_num_all,hsf_num_ok FROM user_data WHERE user_id = $1"
const selectAccountHtmlDataByUserIdSQL = "" + //其他用户头像账号集 -》UserId
						"SELECT user_real_name , user_photo , hsf_address , usdt_address FROM user_data WHERE user_id = $1"
const selectAccountHtmlDataByPhoneNumSQL = "" + //其他用户头像账号集 -》PhoneNum
						"SELECT user_real_name , user_photo , hsf_address , usdt_address , user_id FROM user_data WHERE phone_num = $1"
const selectAccountHtmlDataByUidSQL = "" + //其他用户头像账号集 -》Uid
	"SELECT user_real_name , user_photo , hsf_address , usdt_address FROM user_data WHERE u_id = $1"

const modifyUidSQL = "" +
	"UPDATE user_data SET u_id = $1, u_id_flag = $2 WHERE user_id = $3"
const updateAccountWalletByIdSQL = "" +
	"UPDATE user_data SET hsf_num_all = $1,hsf_num_ok = $2,usdt_num_all = $3,usdt_num_ok = $4 WHERE user_id = $5"

const updateRealNameSQL = "" +
	"UPDATE user_data SET user_real_name = $2, certify_num = $3, address = $4, birthday = $5, kyc_flag = $6 WHERE user_id = $1"

const updateEmailByIdSQL = "" +
	"UPDATE user_data SET email = $2 WHERE user_id = $1"

const updatePhoneNumByIdSQL = "" +
	"UPDATE user_data SET phone_num = $2 WHERE user_id = $1"

const setPayPasswordByIdSQL = "" +
	"UPDATE user_data SET pay_pwd = $2, pay_pwd_flag = $3 WHERE user_id = $1"

const updatePayPasswordByIdSQL = "" +
	"UPDATE user_data SET pay_pwd = $2 WHERE user_id = $1"

const updateLoginPasswordByIdSQL = "" +
	"UPDATE user_data SET password = $2 WHERE user_id = $1"

const setUserSecretSQL = "" +
	"UPDATE user_data SET user_secret = $1 WHERE phone_num = $2"

const setFishCodeByIdSQL = "" +
	"UPDATE user_data SET fish_code = $2, fish_ver_flag = $3 WHERE user_id = $1"

const setCheckSwitchByIdSQL = "" +
	"UPDATE user_data SET phone_ver_flag = $2, mail_ver_flag = $3 WHERE user_id = $1"
const selectSecretByIdSQL = "" +
	"SELECT user_secret FROM user_data WHERE user_id = $1"

const setUserAddressByPhoneNumSQL = "" +
	"UPDATE user_data SET usdt_address = $1, hsf_address = $2 WHERE phone_num = $3"

//const updateGoogleVerifyByIdSQL = "" +
//	"UPDATE user_data SET google_key = $2, google_ver_flag = $3 WHERE user_id = $1"

type accountsStatements struct {
	insertAccountStmt        *sql.Stmt
	updateAccountTokenStmt   *sql.Stmt
	selectAccountByPayIDStmt *sql.Stmt
	selectAccountByTokenStmt *sql.Stmt
	//updatePasswordByPhoneStmt           *sql.Stmt
	//updatePasswordByMailStmt            *sql.Stmt
	//selectAccountByUserIdStmt           *sql.Stmt
	//selectAccountByUserPhoneStmt        *sql.Stmt
	//selectAccountByUserMailStmt         *sql.Stmt
	//selectPasswordHashStmt              *sql.Stmt
	//updateAccountGestureStmt            *sql.Stmt //修改手势key
	//updateAccountGestureFlagStmt        *sql.Stmt //修改手势状态
	//updateGoogleCodeStmt                *sql.Stmt
	//updateGoogleFlagStmt                *sql.Stmt
	//updateFingerStmt                    *sql.Stmt
	//selectAccountWalletByIdStmt         *sql.Stmt
	//selectAccountHtmlDataByUserIdStmt   *sql.Stmt
	//selectAccountHtmlDataByPhoneNumStmt *sql.Stmt
	//selectAccountHtmlDataByUidStmt      *sql.Stmt
	//modifyUidStmt                       *sql.Stmt
	//updateAccountWalletByIdStmt         *sql.Stmt
	//updateRealNameInfoStmt              *sql.Stmt
	//updateEmailStmt                     *sql.Stmt
	//updatePhoneNumberStmt               *sql.Stmt
	//setPayPasswordStmt                  *sql.Stmt
	//updatePayPasswordStmt               *sql.Stmt
	//updateLoginPasswordStmt             *sql.Stmt
	//setFishCodeStmt                     *sql.Stmt
	//setCheckSwitchStmt                  *sql.Stmt
	//selectUserSecretByIdStmt            *sql.Stmt

	//setUserAddressByPhoneNumStmt *sql.Stmt
	//setUserSecretStmt            *sql.Stmt //插入pay钱包密钥

	//updateGoogleSwitchStmt              *sql.Stmt
}

func (s *accountsStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(userSchema)
	return err
}

func (s *accountsStatements) prepare(db *sql.DB) (err error) {
	if s.insertAccountStmt, err = db.Prepare(insertAccountSQL); err != nil {
		return
	}
	if s.updateAccountTokenStmt, err = db.Prepare(updateFilerAccountTokenSQL); err != nil {
		return
	}
	if s.selectAccountByPayIDStmt, err = db.Prepare(selectAccountByPayIDSQL); err != nil {
		return
	}
	if s.selectAccountByTokenStmt, err = db.Prepare(selectAccountByTokenSQL); err != nil {
		return
	}
	return
}

func (s *accountsStatements) insertAccount(ctx context.Context, txn *sql.Tx, payID int64, token string) (err error) {
	stmt := TxStmt(txn, s.insertAccountStmt)
	_, err = stmt.ExecContext(ctx, payID, token)
	if err != nil {
		return err
	}
	return
}

func (s *accountsStatements) selectAccountByPayID(ctx context.Context, txn *sql.Tx, payID int64) (account *types.Account, err error) {
	row := TxStmt(txn, s.selectAccountByPayIDStmt).QueryRowContext(ctx, payID)
	account = &types.Account{}
	if err = row.Scan(
		&account.FilerID,
		&account.PayID,
		&account.Token,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
	}
	return
}

func (s *accountsStatements) selectAccountByToken(ctx context.Context, txn *sql.Tx, token string) (account *types.Account, err error) {
	row := TxStmt(txn, s.selectAccountByTokenStmt).QueryRowContext(ctx, "%"+token+"%")
	account = &types.Account{}
	if err = row.Scan(
		&account.FilerID,
		&account.PayID,
		&account.Token,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
	}
	return
}

func (s *accountsStatements) updateAccountToken(ctx context.Context, txn *sql.Tx, filerID int64, token string) (err error) {
	stmt := TxStmt(txn, s.updateAccountTokenStmt)
	_, err = stmt.ExecContext(ctx, filerID, token)
	if err != nil {
		return err
	}
	return
}

func TxStmt(transaction *sql.Tx, statement *sql.Stmt) *sql.Stmt {
	if transaction != nil {
		statement = transaction.Stmt(statement)
	}
	return statement
}
