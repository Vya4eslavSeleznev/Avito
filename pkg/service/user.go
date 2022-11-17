package service

import "Avito/pkg/repository"

type UserService struct {
	repo repository.User
}

func NewUser(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) AddUser() (int, error) {
	return s.repo.AddUser()
}
