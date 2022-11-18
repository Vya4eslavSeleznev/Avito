package repository

import (
	avito "Avito"
	"github.com/jmoiron/sqlx"
)

type Transaction interface {
	AddTransaction(transaction avito.Accrual) (int, error)
	ReserveTransaction(transaction avito.Reservation) (int, error)
	GetBalance(transaction avito.Balance) (float64, error)
	RevenueRecognition(transaction avito.Accounting) (int, error)
}

type Repository struct {
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Transaction: NewTransactionRepository(db),
	}
}
