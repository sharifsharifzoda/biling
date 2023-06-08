package service

import (
	"biling/internal/repository"
	"biling/models"
)

type Account interface {
	CreateAccount(acc *models.Account) error
	GetAccounts() (models.Accounts, error)
	GetAccountById(id string) (models.Account, error)
	Transaction(tr models.Transaction) (int, error)
}

type Service struct {
	Acc Account
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Acc: NewAccountService(repo.Acc)}
}
