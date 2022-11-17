package repository

import (
	avito "Avito"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (r *TransactionPostgres) AddTransaction(transaction avito.Transaction) error {

	query := fmt.Sprintf("INSERT INTO %s (user_id, amount, type, service_id, order_id, date)"+
		"values ($1, $2, $3, $4, $5, $6)", transactionTable)

	row := r.db.QueryRow(query, transaction.UserId, transaction.Amount, transaction.Type, transaction.ServiceId,
		transaction.OrderId, transaction.Date)

	if err := row.Scan(); err != nil {
		return err
	}

	return nil
}
