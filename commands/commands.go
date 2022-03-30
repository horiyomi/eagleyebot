package commands

import (
	"fmt"

	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

// StartCommand is the command for get info about the bot
func StartCommand(c tele.Context) error {

	template := `
	Welcome to Eagle Eye Bot @eagl_eye_bot

	Commands:
		/balances	shows all linked wallets balances and info
		/networks	shows all available network nodes
		/stake		allow bot to stake on your behalf. Args <NETWORK> <AMOUNT>
		/help		shows all available commands
	`

	return c.Send(template)
}

// BalanceThresholdCommand shows all linked wallet balances and there alert threshold if any
func BalanceThresholdCommand(c tele.Context) error {
	return c.Send(`
		BTC: $40,000
		ETH: $3,400
		ADA: $6,000
	`)
}

// NetworkListCommand - Fetch available connected nodes
func NetworkListCommand(c tele.Context) error {
	return c.Send(`All Networks Available
		Cardano
		Ethereum
		Cosmos
	
	`)
}

// HelpCommand shows all available command
func HelpCommand(c tele.Context) error {
	temp := `
	Commands:
		/balances	shows all linked wallets balances and info
		/networks	shows all available network nodes
		/stake		allow bot to stake on your behalf. Args <NETWORK> <AMOUNT>
		/help		shows all available commands
	`

	return c.Send(temp)
}

// StakeCommand allow staking on behalf user
func StakeCommand(c tele.Context) error {

	args := c.Args()
	msg := fmt.Sprintf("Expected 2 arguements. %d was provided", len(args))
	if len(args) < 2 {
		logrus.Fatal(msg)
		return c.Send(msg)
	}

	if len(args) > 2 {
		msg = fmt.Sprintf("Expected 2 arguements. %d was provided", len(args))
		logrus.Fatal(msg)
		return c.Send(msg)
	}

	// network, stake := args[0], args[1]

	// nNode := GetNetwork(network)
	// resp, err := nNode.Stake(stake)

	// if err != nil {
	// 	logrus.Fatal(err)
	// 	return err
	// }

	respFormat := `
		Network Name: Sampe
		Node: 100-AU
		Validator: Ios

	`
	return c.Send(respFormat)
}
