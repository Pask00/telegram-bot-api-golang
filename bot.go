package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
)

func sendToAll(array []chan *Message, message *Message) {
	for _, el := range array {
		el <- message
	}
}

type Bot struct {
	token       string
	video       []chan *Message
	photo       []chan *Message
	text        []chan *Message
	textOn      []chan *Message
	message     []chan *Message
	join        []chan *Message
	left        []chan *Message
	videoChan   bool
	photoChan   bool
	textChan    bool
	textOnChan  bool
	messageChan bool
	joinChan    bool
	leftChan    bool
}

// Create a new bot
func NewBot(token string) *Bot {
	return &Bot{
		token,
		make([]chan *Message, 0),
		make([]chan *Message, 0),
		make([]chan *Message, 0),
		make([]chan *Message, 0),
		make([]chan *Message, 0),
		make([]chan *Message, 0),
		make([]chan *Message, 0),
		false,
		false,
		false,
		false,
		false,
		false,
		false,
	}
}

func (bot *Bot) Stop() {
	bot.makeRequest("deleteWebhook", nil)
}

func (bot *Bot) Start(urlListener string) {
	values := url.Values{}

	values.Set("url", urlListener)

	fmt.Println("Bot is now running!")
	bot.makeRequest("setWebhook", values)

	r := mux.NewRouter()
	r.HandleFunc(urlListener, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":

			var result *Result
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err)
			}

			json.Unmarshal(body, &result)

			switch {
			case result.Message.NewChatMember != nil && bot.joinChan:
				sendToAll(bot.join, result.Message)
			case result.Message.LeftChatMember != nil && bot.leftChan:
				sendToAll(bot.left, result.Message)
			case result.Message.Text != "" && bot.textChan:
				sendToAll(bot.text, result.Message)
			case result.Message.Video != nil && bot.videoChan:
				sendToAll(bot.video, result.Message)
			case result.Message.Photo != nil && bot.photoChan:
				sendToAll(bot.photo, result.Message)
			}
			if result.Message.Text != "" && bot.textOnChan {
				sendToAll(bot.textOn, result.Message)
			}
			if bot.messageChan {
				sendToAll(bot.message, result.Message)
			}
		}
	})
	http.ListenAndServe(":3000", r)
}

// Start listening telegram API
// First arg -> limit
// Second arg -> timeout
func (bot *Bot) Listen(args ...interface{}) {
	values := url.Values{}

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
	values.Set("offset", strconv.Itoa(offset))
	values.Set("timeout", strconv.Itoa(timeout))
	values.Set("limit", strconv.Itoa(limit))

	fmt.Println("Bot is now running!")
	for {

		var update *GetUpdates

		val, err := json.Marshal(bot.makeRequestWithReturn("getUpdates", values))

		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(val, &update)

		if len(update.Result) > 0 {
			values.Set("offset", strconv.Itoa(update.Result[len(update.Result)-1].UpdateID+1))

			for i := 0; i < len(update.Result); i++ {
				if update.Ok {
					switch {
					case update.Result[i].Message.NewChatMember != nil && bot.joinChan:
						sendToAll(bot.join, update.Result[i].Message)
					case update.Result[i].Message.LeftChatMember != nil && bot.leftChan:
						sendToAll(bot.left, update.Result[i].Message)
					case update.Result[i].Message.Text != "" && bot.textChan:
						sendToAll(bot.text, update.Result[i].Message)
					case update.Result[i].Message.Video != nil && bot.videoChan:
						sendToAll(bot.video, update.Result[i].Message)
					case update.Result[i].Message.Photo != nil && bot.photoChan:
						sendToAll(bot.photo, update.Result[i].Message)
					}
					if update.Result[i].Message.Text != "" && bot.textOnChan {
						sendToAll(bot.textOn, update.Result[i].Message)
					}
					if bot.messageChan {
						sendToAll(bot.message, update.Result[i].Message)
					}
				} else {
					fmt.Println("Error: " + strconv.Itoa(update.ErrorCode) + "\tDescription: " + update.Description)
					fmt.Println("Bot stopped!")
					return
				}
			}
		}
	}
}
