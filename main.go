package main

import (
	"./bot"
)

var token = "527204875:AAEQWGuX6H0NJfDzxGwTNcrbtA3QzmV8zjw"

func main() {
	mybot := bot.NewBot(token)

	mybot.OnText(func(message *bot.Message) {
		mybot.SendMessage(23, "ciao")
	})

	mybot.Listen()
}
