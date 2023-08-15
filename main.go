package main

import (
	tgClient "Go-Tg-Bot/clients/telegram"
	eventConsumer "Go-Tg-Bot/consumer/event-consumer"
	"Go-Tg-Bot/events/telegram"
	"Go-Tg-Bot/storage/files"
	"flag"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

// 5528334737:AAHtgqeBOVZp59AREfRX17IR2OiHndA6rd0
func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := eventConsumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"5528334737:AAHtgqeBOVZp59AREfRX17IR2OiHndA6rd0",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
