package cmd

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Not exist .env")
	}
}

func main() {
	db, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		panic("No database connection")
	}

	defer func() {
		if err := db.Close(context.Background()); err != nil {
			log.Fatalln("Failed to close database", err)
		}

	}()

}
