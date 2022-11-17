package service

import "Avito/pkg/repository"

type User interface {
	AddUser() (int, error)
}

type Transaction interface {
	AddTransaction() error
}

type Service struct {
	User
	Transaction
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:        NewUser(repos.User),
		Transaction: NewTransaction(repos.Transaction),
	}
}
