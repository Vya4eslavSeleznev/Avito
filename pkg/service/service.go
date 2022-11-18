package service

import (
	avito "Avito"
	"Avito/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Transaction interface {
	AddTransaction(transaction avito.Accrual) (int, error)
	ReserveTransaction(transaction avito.Reservation) (int, error)
	GetBalance(transaction avito.Balance) (float64, error)
	RevenueRecognition(transaction avito.Accounting) (int, error)
}

type Service struct {
	Transaction
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Transaction: NewTransaction(repos.Transaction),
	}
}
