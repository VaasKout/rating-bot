package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"rating-bot/internal/app"
)

func main() {
	err := godotenv.Load("configs/.env")
	if err != nil {
		fmt.Println(err)
		godotenv.Load("../configs/.env")
	}
	application := app.New()
	application.Run()
}
