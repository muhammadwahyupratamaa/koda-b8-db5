package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func Conn() *pgx.Conn {
	_ = godotenv.Load()

	connStr := os.Getenv("DATABASE_URL")

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return conn
}