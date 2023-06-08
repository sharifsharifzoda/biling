package database

const (
	CreateAccount  = `INSERT INTO accounts(id, name, type, balance) VALUES (($1), ($2), ($3), ($4));`
	GetAccounts    = `SELECT a.id, a.name, a.type, a.balance FROM accounts AS a;`
	GetAccountById = `SELECT a.id, a.name, a.type, a.balance FROM accounts AS a WHERE a.id = $1;`

	CreateTransaction = `INSERT INTO transactions(receiver_account, sender_account, amount) 
							VALUES (($1), ($2), ($3)) RETURNING id;`
)

const (
	CallProcedure = `call transfer($1, $2, $3, null);`
)
