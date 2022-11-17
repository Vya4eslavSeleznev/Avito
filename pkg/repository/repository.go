package repository

import (
	avito "Avito"
	"github.com/jmoiron/sqlx"
)

type User interface {
}

type Transaction interface {
	AddTransaction(transaction avito.Accrual) (int, error)
}

type Repository struct {
	User
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Transaction: NewTransactionRepository(db),
	}
}
