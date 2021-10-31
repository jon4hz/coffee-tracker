package cmd

import (
	"log"

	"github.com/jon4hz/coffee-tracker/internal/config"
	"github.com/jon4hz/coffee-tracker/internal/database"
	"github.com/jon4hz/coffee-tracker/internal/telegram"
	"github.com/prometheus/common/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Version: version.Version,
	Use:     "coffee-tracker",
	Short:   "Tracker your coffees",
	Long:    "Track how many coffees you drink per project",
	RunE: func(cmd *cobra.Command, args []string) error {
		return Root()
	},
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

var Root = func() error {

	cfg := config.Get()
	database.Connect()
	database.Migrate()

	// start api

	b, err := telegram.NewBot(cfg.Telegram.BotToken, cfg.Telegram.OwnerID)
	if err != nil {
		log.Fatalf("failed to create the bot: %s", err)
	}
	err = b.Start()
	if err != nil {
		log.Fatalf("failed to start the bot: %s", err)
	}

	return nil
}
