package telegram

import (
	"fmt"
	"log"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jon4hz/coffee-tracker/internal/database"
)

func startHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	return nil
}

func coffeeHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	p, err := database.NewProject(strings.TrimPrefix(ctx.Message.Text, "/coffee "), ctx.EffectiveUser.Id)
	if err != nil {
		log.Println(err)
		return err
	}
	err = p.AddCoffee()
	if err != nil {
		log.Println(err)
		return err
	}
	coffees := p.GetCoffees()
	if coffees == 1 {
		_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf("%d coffee", coffees), nil)
	} else {
		_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf("%d coffees", coffees), nil)
	}
	if err != nil {
		log.Printf("failed to send message: %s", err)
	}
	return err
}
