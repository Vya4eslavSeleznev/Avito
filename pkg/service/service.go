package service

import "Avito/pkg/repository"

type User interface {
}

type Transaction interface {
}

type Service struct {
	User
	Transaction
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
