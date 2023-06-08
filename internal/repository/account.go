package repository

import (
	"biling/database"
	"biling/models"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (a AccountRepository) CreateAccount(acc *models.Account) error {
	row := a.db.QueryRow(database.CreateAccount, acc.Id, acc.Name, acc.Type, acc.Balance)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

func (a AccountRepository) GetAccounts() (models.Accounts, error) {
	var accounts models.Accounts
	rows, err := a.db.Query(database.GetAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var acc models.Account
		err := rows.Scan(&acc)
		if err != nil {
			log.Println("failed to scan due to: ", err.Error())
			return nil, err
		}
		accounts = append(accounts, acc)
	}

	return accounts, nil
}

func (a AccountRepository) GetAccountById(id string) (models.Account, error) {
	var acc models.Account
	row := a.db.QueryRow(database.GetAccountById, id)
	err := row.Scan(&acc.Id, &acc.Name, &acc.Type, &acc.Balance)
	if err != nil {
		return models.Account{}, err
	}

	return acc, nil
}

func (a AccountRepository) CreateRecords(sender, receiver models.Account) error {
	//_, err := a.db.Exec(database.CreateRecordsTable, sender.Id)
	//_, err = a.db.Exec(database.CreateRecordsTable, receiver.Id)
	//
	//if err != nil {
	//	return err
	//}

	return nil
}

func (a AccountRepository) Transaction(tr models.Transaction) error {
	var result string
	row := a.db.QueryRow(database.CallProcedure, tr.SenderAccount, tr.ReceiverAccount, tr.Amount)
	err := row.Scan(&result)
	if err != nil {
		return err
	}
	if result == "ok" {
		return nil
	}

	err = errors.New(result)
	fmt.Println(err)
	return err
}

func (a AccountRepository) CreateTransaction(tr models.Transaction) (int, error) {
	var id int
	row := a.db.QueryRow(database.CreateTransaction, tr.ReceiverAccount, tr.SenderAccount, tr.Amount)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
