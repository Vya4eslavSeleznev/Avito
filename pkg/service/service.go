package service

import (
	avito "Avito"
	"Avito/pkg/repository"
)

type User interface {
}

type Transaction interface {
	AddTransaction(transaction avito.Accrual) (int, error)
	ReserveTransaction(transaction avito.Reservation) (int, error)
	GetBalance(transaction avito.Balance) (float64, error)
}

type Service struct {
	User
	Transaction
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Transaction: NewTransaction(repos.Transaction),
	}
}
