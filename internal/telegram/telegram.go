package telegram

import (
	"log"
	"net/http"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

type Bot struct {
	bot   *gotgbot.Bot
	owner int64
}

func NewBot(token string, owner int64) (*Bot, error) {
	bot, err := gotgbot.NewBot(token, &gotgbot.BotOpts{
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})
	if err != nil {
		return nil, err
	}
	return &Bot{
		bot:   bot,
		owner: owner,
	}, nil
}

// Start starts the bot. This method is blocking
func (b *Bot) Start() error {
	updater := ext.NewUpdater(nil)
	setHandlers(updater.Dispatcher)

	err := updater.StartPolling(b.bot, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: gotgbot.GetUpdatesOpts{
			AllowedUpdates: []string{
				"message",
				"callback_query",
			},
		},
	})
	if err != nil {
		return err
	}
	log.Printf("%s has been started...", b.bot.User.Username)
	updater.Idle()
	return nil
}

func setHandlers(d *ext.Dispatcher) {
	d.AddHandler(handlers.NewCommand("start", startHandler))
	d.AddHandler(handlers.NewCommand("coffee", coffeeHandler))
}
