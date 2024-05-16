package main

import (
	"fmt"
	"weather-tracker/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("File not found", err)
		return
	}
	router.Routing()
}
