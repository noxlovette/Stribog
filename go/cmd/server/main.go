package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func getEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

	v, ok := os.LookupEnv(key)
	if !ok {
		log.Default().Fatalf("env variable %v not set", key)
	}
	return v
}
func main() {
	dbpool, err := pgxpool.New(context.Background(), getEnvVar("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var forge string
	err = dbpool.QueryRow(context.Background(), "SELECT id FROM forges").Scan(&forge)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(forge)
}
