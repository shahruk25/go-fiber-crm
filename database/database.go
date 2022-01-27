package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

var (
	db *pgx.Conn
)

func GetConn() {

	connString := "postgres://postgres:123456@localhost:5432/crm"
	config, err := pgx.ParseConfig(connString)

	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)

	if err != nil {
		log.Fatal("err connecting to the db: ", err)
	}

	db = conn
}

func GetDB() *pgx.Conn {
	return db
}
