package bot

import (
	"log"
	"os"

	"github.com/Toolnado/bot-parser/internal/botErrors"
)

var (
	TAG_NAME = ""
	ATR_NAME = ""

	GET_TORS_PROXY = false
	GET_USER_AGENT = false
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

func (b *Bot) RarseHtml(html string) {
	if TAG_NAME != "" {
		for _, result := range b.GetObject(html) {
			log.Println(result)
		}
	} else {
		log.Println(html)
	}
}

func (b *Bot) ParseArgs(args []string) {
	var (
		flagTag bool = false
		flagAtr bool = false
	)

	if len(args) == 1 {
		b.errors.GetError("Np url specified")
	} else if len(args) == 2 {
		if args[1] == "-i" || args[1] == "--info" {
			b.GetInfo()
		} else {
			return
		}
	} else if len(args) > 2 {
		for _, value := range args[2:] {
			if value == "-tp" || value == "--tor-proxy" {
				GET_TORS_PROXY = true
			} else if value == "ua" || value == "--user-agent" {
				GET_USER_AGENT = true
			} else if value == "-t" || value == "--tag" {
				flagTag = true
			} else if value == "-a" || value == "--attr" {
				flagAtr = true
			} else if flagTag {
				TAG_NAME = value
				flagTag = false
			} else if flagAtr {
				ATR_NAME = value
				flagAtr = false
			}
		}
	}

}

func (b *Bot) GetInfo() {
	log.Println(`
	Modules:
		-tp || --tor-proxy  -> Turn on tor proxy
		-ua || --user-agent -> Turn on user-agent
		-t  || --tag        -> Tag name
		-a  || --attr       -> Attribute name
	Example:
		parse google.com --tag a --attr href -tp -ua
	`)

	os.Exit(0)
}

func (b *Bot) GetObject(html string) []string {
	return []string{}
}
