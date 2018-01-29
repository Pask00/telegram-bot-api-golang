package bot

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Bot struct {
	token       string
	video       chan *Message
	photo       chan *Message
	text        chan *Message
	textOn      chan *Message
	message     chan *Message
	videoChan   bool
	photoChan   bool
	textChan    bool
	textOnChan  bool
	messageChan bool
}

type function func(*Message)

// Create a new bot
func NewBot(token string) *Bot {
	return &Bot{
		token,
		make(chan *Message),
		make(chan *Message),
		make(chan *Message),
		make(chan *Message),
		make(chan *Message),
		false,
		false,
		false,
		false,
		false,
	}
}

// Start listening telegram API
// First arg -> limit
// Second arg -> timeout
func (bot *Bot) Listen(args ...interface{}) {
	var offset = 0
	var limit = 100
	var timeout = 0
	if len(args) < 2 {
		for index, arg := range args {
			switch index {
			case 0:
				param, ok := arg.(int)
				if !ok {
					panic("Limit parameter cannot be a string")
				}
				limit = param
			case 1:
				param, ok := arg.(int)
				if !ok {
					panic("Timeout parameter cannot be a string")
				}
				timeout = param
			}
		}
	} else {
		panic("Too many arguments")
	}
	fmt.Println("Bot is now running!")
	for {
		url := "https://api.telegram.org/bot" + bot.token + "/getUpdates?offset=" + strconv.Itoa(offset) + "&limit=" + strconv.Itoa(limit) + "&timeout" + strconv.Itoa(timeout)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal("Request: ", err)
			return
		}

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal("Response: ", err)
			return
		}

		var update GetUpdates

		if err := json.NewDecoder(resp.Body).Decode(&update); err != nil {
			log.Println(err)
		}

		resp.Body.Close()

		if len(update.Result) > 0 {
			offset = update.Result[len(update.Result)-1].UpdateID + 1

			for i := 0; i < len(update.Result); i++ {
				if update.Ok {
					switch {
					case update.Result[i].Message.Text != "" && bot.textChan:
						bot.text <- update.Result[i].Message
					case update.Result[i].Message.Video != nil && bot.videoChan:
						bot.video <- update.Result[i].Message
					case update.Result[i].Message.Photo != nil && bot.photoChan:
						bot.photo <- update.Result[i].Message
					}
					if update.Result[i].Message.Text != "" && bot.textOnChan {
						bot.textOn <- update.Result[i].Message
					}
					if bot.messageChan {
						bot.message <- update.Result[i].Message
					}
				} else {
					fmt.Println("Error: " + strconv.Itoa(update.ErrorCode) + "\tDescription: " + update.Description)
				}
			}
		}
	}
}
