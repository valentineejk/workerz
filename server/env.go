package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() (err error) {

	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("failed to load .env file")
	}

	if os.Getenv("ENV") == "DEV" {

		if err := godotenv.Load(); err != nil {
			fmt.Println("failed to load .env file")
		}

		fmt.Println("Loading Development env...")

		return nil

	} else {
		fmt.Println("Loading Production env...")
	}
	return nil

}
