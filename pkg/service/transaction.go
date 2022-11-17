package service

import "Avito/pkg/repository"

type TransactionService struct {
	repo repository.Transaction
}

func NewTransaction(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) AddTransaction() error {
	return s.repo.AddTransaction()
}
