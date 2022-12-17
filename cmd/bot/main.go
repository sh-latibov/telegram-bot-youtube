package main

import (
	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sh-latibov/telegram-bot-youtube/pkg/config"
	"github.com/sh-latibov/telegram-bot-youtube/pkg/repository"
	"github.com/sh-latibov/telegram-bot-youtube/pkg/repository/boltdb"
	"github.com/sh-latibov/telegram-bot-youtube/pkg/server"
	"github.com/sh-latibov/telegram-bot-youtube/pkg/telegram"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(cfg)

	bot, err := tgbotapi.NewBotAPI("5973915755:AAHbuNdT-mjRnrWTuyMIT2OgEcDUs8rPDIU")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	pocketClient, err := pocket.NewClient("104888-855d71bc3dfa365a36e7167")
	if err != nil {
		log.Fatal(err)
	}

	db, err := initDb()
	if err != nil {
		log.Fatal(err)
	}

	tokenRepository := boltdb.NewTokenRepository(db)

	telegramBot := telegram.NewBot(bot, pocketClient, tokenRepository, "http://localhost/")

	authorizationServer := server.NewAuthorizationServer(pocketClient, tokenRepository, "https://t.me/tishtidish_bot")

	go func() {
		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := authorizationServer.Start(); err != nil {
		log.Fatal(err)
	}
}

func initDb() (*bolt.DB, error) {
	db, err := bolt.Open("bot.db", 8600, nil)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.AccessTokens))
		if err != nil {
			return nil
		}

		_, err = tx.CreateBucketIfNotExists([]byte(repository.RequestTokens))
		if err != nil {
			return nil
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return db, nil
}
