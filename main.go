package main

import (
	"flag"
	"log"
	"github.com/cherap93/keeper-tg-bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(tgBotHost, mustToken()) // клиент telegram готов
	// fetcher и процессор общаются с API Telegram
	// fetcher := fetcher.New() Получатель информации
	// processor := processor.New() Обрабатывает и отправляет ссылку обратно, либо ругается, 
	// что не понял команду 
 	// Consumer.Start(fetcher, processor)
}

// получения токена из флага
func mustToken() string {
	// bot -tg-bot-token 'my token'
	
	token := flag.String(
		"tg-bot-token", 
		"", 
		"token for access to tg bot",
	)
	flag.Parse()
	if *token == "" {  // нежелательно так делать
		log.Fatal("token is not specified")
	}
	return *token
} 
