package main

import (
	"aiInWhitelists/gemini"
	"aiInWhitelists/telegram"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println(".env loaded")
	gemini.InitAi()
	telegram.InitTelegramBot()
}
