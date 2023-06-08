package repository

import (
	"biling/models"
	"database/sql"
)

type Account interface {
	CreateAccount(acc *models.Account) error
	GetAccounts() (models.Accounts, error)
	GetAccountById(id string) (models.Account, error)
	Transaction(tr models.Transaction) error
	CreateRecords(sender, receiver models.Account) error
	CreateTransaction(tr models.Transaction) (int, error)
}

type Repository struct {
	Acc Account
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{Acc: NewAccountRepository(db)}
}
