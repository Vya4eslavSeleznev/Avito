package repository

import (
	avito "Avito"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) AddTransaction(transaction avito.Accrual) (int, error) {
	var id int
	var transactionType int

	if transaction.Amount > 0 {
		transactionType = 1
	} else if transaction.Amount < 0 {
		transactionType = -1
	} else {
		return -1, nil
	}

	query := fmt.Sprintf("INSERT INTO %s (user_id, amount, type, service_id, order_id, date) values ($1, $2, $3, $4, $5, $6) RETURNING id", transactionTable)

	row := r.db.QueryRow(query, transaction.UserId, transaction.Amount, transactionType, nil,
		nil, time.Now())

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *TransactionRepository) ReserveTransaction(transaction avito.Reservation) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (user_id, amount, type, service_id, order_id, date) values ($1, $2, $3, $4, $5, $6) RETURNING id", transactionTable)

	row := r.db.QueryRow(query, transaction.UserId, transaction.Amount, 0, transaction.ServiceId,
		transaction.OrderId, time.Now())

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *TransactionRepository) GetBalance(id avito.Balance) (float64, error) {
	var balance float64

	query := fmt.Sprintf(
		"SELECT SUM("+
			"(SELECT SUM(amount) FROM %[1]s WHERE type = 1) -"+
			"(SELECT SUM(amount) FROM %[1]s WHERE type = 0) +"+
			"(SELECT SUM(amount) FROM %[1]s WHERE type = -1)) AS Balance "+
			"FROM %[1]s WHERE user_id = $1", transactionTable)

	row := r.db.QueryRow(query, id.UserId)

	if err := row.Scan(&balance); err != nil {
		return 0, err
	}

	return balance, nil
}
