package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fawziridwan/go_auth/controllers"

	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func initialize() {
	// initialize environment form .env files
	if err := godotenv.Load(); err != nil {
		log.Print(".env files not found")
	}
}

func execute() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	server.Run(":9000")
}

func main() {
	initialize()
	execute()
}
