package repository

import (
	avito "Avito"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) AddTransaction(transaction avito.Transaction) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (user_id, amount, type, service_id, order_id, date) values ($1, $2, $3, $4, $5, $6) RETURNING id", transactionTable)

	row := r.db.QueryRow(query, transaction.UserId, transaction.Amount, transaction.Type, transaction.ServiceId,
		transaction.OrderId, transaction.Date)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
