package repository

import (
	"github.com/jmoiron/sqlx"
)

type User interface {
	AddUser() (int, error)
}

type Transaction interface {
	AddTransaction() error
}

type Repository struct {
	User
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:        NewUserPostgres(db),
		Transaction: NewTransactionPostgres(db),
	}
}
