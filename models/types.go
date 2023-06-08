package models

import "time"

type Account struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Type      int       `json:"type"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}

type AccountType struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
}

type Accounts []Account

type Transaction struct {
	Id              int     `json:"id"`
	SenderAccount   string  `json:"sender_account"`
	ReceiverAccount string  `json:"receiver_account"`
	Amount          float64 `json:"amount"`
}
