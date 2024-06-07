package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/mattn/go-sqlite3"
)

var ConnPool *pgxpool.Pool

func InitDB() {
	var err error
	ConnPool, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		panic("Could not connect to database.")
	}

}
