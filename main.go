package main

import (
	"./bot"
)

var token = "token"

func main() {
	mybot := bot.NewBot(token)

	mybot.OnText(func(message *bot.Message) {
		mybot.SendMessage(23, "ciao")
	})

	mybot.Listen()
}
