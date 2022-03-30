package handlers

import (
	"fmt"

	tele "gopkg.in/telebot.v3"
)

func StakeHandler(m *tele.Message) error {
	fmt.Println(m.Payload)
	return nil
}
