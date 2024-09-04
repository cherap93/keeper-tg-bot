package telegram

import "github.com/cherap93/keeper-tg-bot/clients/telegram"

type Processor struct {
	tg *telegram.Client
	offset int
	// Storage
}

func New(client *telegram.Client, storage)