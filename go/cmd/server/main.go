package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"stribog/config"
)

func main() {
	ctx := context.Background()

	appState, err := config.InitAppState(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
		os.Exit(1)
	}
	defer appState.DB.Close()

	fmt.Println("Database connected successfully!")
	var scan string
	err = appState.DB.QueryRow(ctx, "SELECT id FROM forges").Scan(&scan)
	if err != nil {
		log.Fatalf("blabla %v", err)
	}
	fmt.Println(scan)
}
