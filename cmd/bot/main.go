package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sh-latibov/telegram-bot-youtube/pkg/telegram"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5973915755:AAHbuNdT-mjRnrWTuyMIT2OgEcDUs8rPDIU")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient("104888-855d71bc3dfa365a36e7167")
	if err != nil {
		log.Fatal(err)
	}

	telegramBot := telegram.NewBot(bot, pocketClient, "http://localhost/")

	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
