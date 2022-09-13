package psql

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

func GetDatabase() *pgx.Conn {
	if db != nil {
		return db
	}
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to Database: %v", err))
	}
	db = conn
	return db
}
