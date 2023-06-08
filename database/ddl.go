package database

const (
	CreateAccountTypesTable = `CREATE TABLE IF NOT EXISTS account_types (
    	id SERIAL PRIMARY KEY NOT NULL,
    	type VARCHAR NOT NULL
	);`
	CreateAccountTable = `CREATE TABLE IF NOT EXISTS accounts (
    	id VARCHAR PRIMARY KEY UNIQUE NOT NULL,
    	name VARCHAR NOT NULL,
    	type INTEGER REFERENCES account_types (id),
    	balance DECIMAL DEFAULT 0,
    	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    	updated_at TIMESTAMP,
    	deleted_at TIMESTAMP
	);`
	CreateTransactionsTable = `CREATE TABLE IF NOT EXISTS transactions (
    	id SERIAL PRIMARY KEY NOT NULL,
    	receiver_account VARCHAR REFERENCES accounts (id),
    	sender_account VARCHAR REFERENCES accounts (id),
    	amount DECIMAL NOT NULL,
    	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`
)

const (
	CreateRecordsTable = `CREATE TABLE IF NOT EXISTS $1_records (
    	id SERIAL PRIMARY KEY NOT NULL,
    	account_id VARCHAR REFERENCES accounts (id),
    	operation_type VARCHAR NOT NULL,
    	amount DECIMAL NOT NULL,
    	initial_balance DECIMAL,
    	operation_data TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`
)
