package botErrors

import (
	"log"
	"os"
)

type BotErrors struct{}

func NewBotErrors() *BotErrors {
	return &BotErrors{}
}

func (be *BotErrors) CheckError(err error) {
	if err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}
}

func (be *BotErrors) GetError(err string) {
	log.Println("Error:", err)
	os.Exit(1)
}
