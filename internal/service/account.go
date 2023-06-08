package service

import (
	"biling/internal/repository"
	"biling/models"
	"log"
)

type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}

func (a AccountService) CreateAccount(acc *models.Account) error {
	err := a.repo.CreateAccount(acc)
	if err != nil {
		log.Println("could not create a new account due to: ", err.Error())
		return err
	}
	return nil
}

func (a AccountService) GetAccounts() (models.Accounts, error) {
	accounts, err := a.repo.GetAccounts()
	if err != nil {
		log.Println("failed to get list of accounts. error is: ", err.Error())
		return nil, err
	}

	return accounts, nil
}

func (a AccountService) GetAccountById(id string) (models.Account, error) {
	account, err := a.repo.GetAccountById(id)
	if err != nil {
		log.Println("failed to get account by id. error is: ", err.Error())
		return models.Account{}, nil
	}
	return account, nil
}

func (a AccountService) Transaction(tr models.Transaction) (int, error) {
	senderAcc, err := a.repo.GetAccountById(tr.SenderAccount)
	if err != nil {
		log.Println("failed to get account by id. error is: ", err.Error())
		return 0, err
	}
	receiverAcc, err := a.repo.GetAccountById(tr.ReceiverAccount)
	if err != nil {
		log.Println("failed to get account by id. error is: ", err.Error())
		return 0, err
	}

	if err := a.repo.CreateRecords(senderAcc, receiverAcc); err != nil {
		log.Println("failed to create account records. error is: ", err.Error())
		return 0, err
	}

	if err := a.repo.Transaction(tr); err != nil {
		log.Println("unsuccessful operation. error is: ", err.Error())
		return 0, err
	}

	id, err := a.repo.CreateTransaction(tr)
	if err != nil {
		return 0, err
	}

	return id, nil
}
