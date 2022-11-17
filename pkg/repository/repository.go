package repository

import "github.com/jmoiron/sqlx"

type User interface {
}

type Transaction interface {
}

type Repository struct {
	User
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
