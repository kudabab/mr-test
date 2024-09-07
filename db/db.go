package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn //указатель на объект соединения

func InitDB() {
	var err error
	dbURL := "postgres://owner:123@localhost:5432/go_medium_api"

	db, err = pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	err = db.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to ping to database: %v", err)
	}
	fmt.Println("succes connected to DB")
}

func GetDB() *pgx.Conn {
	return db
}
