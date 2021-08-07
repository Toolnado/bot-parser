package bot

import (
	"github.com/Toolnado/bot-parser/internal/botErrors"
)

type Errors interface {
	CheckError(err error)
	GetError(err string)
}
type Bot struct {
	errors Errors
}

func NewBot() *Bot {
	return &Bot{
		errors: botErrors.NewBotErrors(),
	}
}
