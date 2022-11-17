package service

import (
	avito "Avito"
	"Avito/pkg/repository"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransaction(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) AddTransaction(transaction avito.Accrual) (int, error) {
	return s.repo.AddTransaction(transaction)
}
