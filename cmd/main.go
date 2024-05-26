package main

import (
	"flag"
	"log"
	"telegram-adviser/clients/telegram"
)

func main() {

	tgClient := telegram.NewClient(mustHost(), mustToken())
	// fetcher = fetch.NewFetcher()
	// processor = processor.NewProcessor()
	// server = server.NewServer()
	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String(
		"token",
		"",
		"token for telegram bot",
	)
	flag.Parse()
	if *token == "" {
		log.Fatal("token is empty")
	}
	return *token
}

func mustHost() string {
	host := flag.String(
		"api.telegram.bot",
		"",
		"host for telegram bot",
	)
	flag.Parse()
	if *host == "" {
		log.Fatal("host is empty")
	}
	return *host
}
