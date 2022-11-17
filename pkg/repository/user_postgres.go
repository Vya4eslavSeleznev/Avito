package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) AddUser() (int, error) {

	// TODO query
	var id int
	query := fmt.Sprintf("INSERT INTO %s (id) values ($1) RETURNING id", userTable)
	row := r.db.QueryRow(query, 44)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
