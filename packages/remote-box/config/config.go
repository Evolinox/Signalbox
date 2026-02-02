package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig(path string) {
	env := os.Getenv("ENV")

	if "" == env {
		err := godotenv.Load(path + ".env")
		if err != nil {
			fmt.Println("Error loading .env file")
			os.Exit(3)
		}
	}
}
