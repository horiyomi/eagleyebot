package main

import (
	"os"
	"time"

	"github.com/horiyomi/eagleyebot/commands"
	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

func main() {
	pref := tele.Settings{
		Token:  os.Getenv("TELEGRAM_APITOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		logrus.Fatal(err)
		return
	}

	b.Handle("/hello", commands.StartCommand)
	b.Handle("/balances", commands.BalanceThresholdCommand)
	// b.Handle("/stake", commands.StakeCommand)
	b.Handle("/networks", commands.NetworkListCommand)
	b.Handle("/help", commands.HelpCommand)

	b.Start()
}
