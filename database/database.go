package database

import (
	"biling/configs"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func GetDBConnection(cfg configs.DatabaseConnConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Dushanbe",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	log.Printf("Connection success host:%s port:%s", cfg.Host, cfg.Port)

	Init(conn)

	return conn, nil
}

func Init(db *sql.DB) {
	DDLs := []string{
		CreateAccountTypesTable,
		CreateAccountTable,
		CreateTransactionsTable,
	}

	for i, ddl := range DDLs {
		_, err := db.Exec(ddl)
		if err != nil {
			log.Fatalf("failed to create table #%d due to: %s", i, err.Error())
		}
	}
}
