package util

import (
	"database/sql"
	"fmt"
)

type Transaction interface {
	// Commit the transaction
	Commit() error
	// Rollback the transaction.
	Rollback() error
}

func WithTransaction(db *sql.DB, fn func(txn *sql.Tx) error) (err error) {
	txn, err := db.Begin()
	if err != nil {
		return fmt.Errorf("WithTransaction.Begin: %w", err)
	}
	succeeded := false
	defer EndTransactionWithCheck(txn, &succeeded, &err)

	err = fn(txn)
	if err != nil {
		return
	}

	succeeded = true
	return
}

func EndTransaction(txn Transaction, succeeded *bool) error {
	if *succeeded {
		return txn.Commit()
	} else {
		return txn.Rollback()
	}
}

func EndTransactionWithCheck(txn Transaction, succeeded *bool, err *error) {
	if e := EndTransaction(txn, succeeded); e != nil && *err == nil {
		*err = e
	}
}
