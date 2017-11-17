package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/andersfylling/imt2681bot/hooks/currencyconversation"
	"github.com/s1kx/unison"
)

func main() {
	token := os.Getenv("IMT_BOT_TOKEN")
	if token == "" {
		logrus.Error("Missing discord bot token in env: IMT_BOT_TOKEN")
		return
	}

	// Create bot structure
	settings := &unison.BotSettings{
		Token: token,

		Commands: []*unison.Command{},
		EventHooks: []*unison.EventHook{
			currencyconversation.Hook,
		},
		Services: []*unison.Service{},
	}

	// check for a command prefix, otherwise invoke by mention
	prefix := os.Getenv("IMT_BOT_COMMAND_PREFIX")
	if prefix != "" {
		settings.CommandPrefix = prefix
		logrus.Info("Bot command prefix is set to " + prefix)
	} else {
		settings.CommandInvokedByMention = true
		logrus.Info("Bot commands are invoked by mention")
		logrus.Info("For a specific command set env IMT_BOT_COMMAND_PREFIX")
	}

	// Start the bot
	err := unison.RunBot(settings)
	if err != nil {
		logrus.Error(err)
	}
}
