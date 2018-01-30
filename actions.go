package bot

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// It makes requests to telegram API
func makeRequest(url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	var update SendMessage

	if err := json.NewDecoder(resp.Body).Decode(&update); err != nil {
		log.Println(err)
	}

	resp.Body.Close()

	if !update.Ok {
		fmt.Println("Error: " + strconv.Itoa(update.ErrorCode) + "\tDescription: " + update.Description)
	}
}

func (bot *Bot) SendMessageCustom(chatID interface{}, text string, args ...interface{}) {

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}
	var url = "https://api.telegram.org/bot" + bot.token + "/sendMessage?chat_id=" + chatID.(string) + "&text=" + text
	if len(args) < 5 {
		for index, arg := range args {
			if arg == nil {
				continue
			}
			switch index {
			case 0:
				param, ok := arg.(string)
				if !ok {
					panic("Type error, parseMode can only be String")
				}
				url += "&parse_mode=" + param
			case 1:
				param, ok := arg.(bool)
				if !ok {
					panic("Type error, disableWebPagePreview can only be Bool")
				}
				url += "&disable_web_page_preview=" + strconv.FormatBool(param)
			case 2:
				param, ok := arg.(bool)
				if !ok {
					panic("Type error, disableNotification can only be Bool")
				}
				url += "&disable_notification=" + strconv.FormatBool(param)
			case 3:
				param, ok := arg.(int)
				if !ok {
					panic("Type error, replyToMessageID can only be Integer")
				}
				url += "&reply_to_message_id=" + strconv.Itoa(param)
			}
		}
	} else {
		panic("Too many arguments")
	}
	makeRequest(url)
}

// Send a text message
func (bot *Bot) SendMessage(chatID interface{}, text string) {

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	url := "https://api.telegram.org/bot" + bot.token + "/sendMessage?chat_id=" + chatID.(string) + "&text=" + text
	makeRequest(url)
}

// Forward a message
func (bot *Bot) SendForward(chatID interface{}, from interface{}, messageID int) {

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	switch param := from.(type) {
	case int:
		from = strconv.Itoa(param)
	case string:
		from = param
	default:
		panic("Type error, from can be only Integer or String")
	}

	url := "https://api.telegram.org/bot" + bot.token + "/forwardMessage?chat_id=" + chatID.(string) + "&from_chat_id=" + from.(string) + "&message_id=" + strconv.Itoa(messageID)
	makeRequest(url)
}

// Send a reply message
func (bot *Bot) SendReply(chatID interface{}, text string, messageID int) {

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	url := "https://api.telegram.org/bot" + bot.token + "/sendMessage?chat_id=" + chatID.(string) + "&text=" + text + "&reply_to_message_id=" + strconv.Itoa(messageID)
	makeRequest(url)
}

// Send a markdown message
func (bot *Bot) SendMarkdown(chatID interface{}, text string) {

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	url := "https://api.telegram.org/bot" + bot.token + "/sendMessage?chat_id=" + chatID.(string) + "&text=" + text + "&parse_mode=Markdown"
	makeRequest(url)
}

// Send a HTML message
func (bot *Bot) SendHTML(chatID interface{}, text string) {

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	url := "https://api.telegram.org/bot" + bot.token + "/sendMessage?chat_id=" + chatID.(string) + "&text=" + text + "&parse_mode=HTML"
	makeRequest(url)
}

// Send a photo message
func (bot *Bot) SendPhoto(chatID interface{}, photo string) {

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	url := "https://api.telegram.org/bot" + bot.token + "/sendPhoto?chat_id=" + chatID.(string) + "&photo=" + photo
	makeRequest(url)
}

func (bot *Bot) SendPhotoCustom(chatID interface{}, photo string, args ...interface{}) {

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	var url = "https://api.telegram.org/bot" + bot.token + "/sendPhoto?chat_id=" + chatID.(string) + "&photo=" + photo
	if len(args) < 4 {
		for index, arg := range args {
			if arg == nil {
				continue
			}
			switch index {
			case 0:
				param, ok := arg.(string)
				if !ok {
					panic("Type error, caption can only be String")
				}
				url += "&caption=" + param
			case 1:
				param, ok := arg.(bool)
				if !ok {
					panic("Type error, disableNotification can only be Bool")
				}
				url += "&disable_notification=" + strconv.FormatBool(param)
			case 2:
				param, ok := arg.(int)
				if !ok {
					panic("Type error, replyToMessageID can only be Integer")
				}
				url += "&reply_to_message_id=" + strconv.Itoa(param)
			}
		}
	} else {
		panic("Too many arguments")
	}
	makeRequest(url)
}

// Send an audio message
func (bot *Bot) SendAudio(chatID interface{}, audio string) {

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	url := "https://api.telegram.org/bot" + bot.token + "/sendAudio?chat_id=" + chatID.(string) + "&audio=" + audio
	makeRequest(url)
}

func (bot *Bot) SendAudioCustom(chatID interface{}, audio string, args ...interface{}) {

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	var url = "https://api.telegram.org/bot" + bot.token + "/sendAudio?chat_id=" + chatID.(string) + "&audio=" + audio
	if len(args) < 6 {
		for index, arg := range args {
			if arg == nil {
				continue
			}
			switch index {
			case 0:
				param, ok := arg.(string)
				if !ok {
					panic("Type error, caption can only be String")
				}
				url += "&caption=" + param
			case 1:
				param, ok := arg.(int)
				if !ok {
					panic("Type error, duration can only be Integer")
				}
				url += "&duration=" + strconv.Itoa(param)
			case 2:
				param, ok := arg.(string)
				if !ok {
					panic("Type performer, performer can only be String")
				}
				url += "&performer=" + param
			case 3:
				param, ok := arg.(string)
				if !ok {
					panic("Type title, title can only be String")
				}
				url += "&title=" + param
			case 4:
				param, ok := arg.(bool)
				if !ok {
					panic("Type error, disableNotification can only be Bool")
				}
				url += "&disable_notification=" + strconv.FormatBool(param)
			case 5:
				param, ok := arg.(int)
				if !ok {
					panic("Type error, replyToMessageID can only be Integer")
				}
				url += "&reply_to_message_id=" + strconv.Itoa(param)
			}
		}
	} else {
		panic("Too many arguments")
	}
	makeRequest(url)
}

//Send a document message
func (bot *Bot) SendDocument(chatID interface{}, document string) {

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	url := "https://api.telegram.org/bot" + bot.token + "/sendDocument?chat_id=" + chatID.(string) + "&document=" + document
	makeRequest(url)
}

func (bot *Bot) SendDocumentCustom(chatID interface{}, document string, args ...interface{}) {

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	var url = "https://api.telegram.org/bot" + bot.token + "/sendDocument?chat_id=" + chatID.(string) + "&document=" + document
	if len(args) < 3 {
		for index, arg := range args {
			if arg == nil {
				continue
			}
			switch index {
			case 0:
				param, ok := arg.(string)
				if !ok {
					panic("Type error, caption can only be String")
				}
				url += "&caption=" + param
			case 1:
				param, ok := arg.(bool)
				if !ok {
					panic("Type error, disableNotification can only be Bool")
				}
				url += "&disable_notification=" + strconv.FormatBool(param)
			case 2:
				param, ok := arg.(int)
				if !ok {
					panic("Type error, replyToMessageID can only be Integer")
				}
				url += "&reply_to_message_id=" + strconv.Itoa(param)
			}
		}
	} else {
		panic("Too many arguments")
	}
	makeRequest(url)
}

// Send a video message
func (bot *Bot) SendVideo(chatID interface{}, video string) {

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	url := "https://api.telegram.org/bot" + bot.token + "/sendVideo?chat_id=" + chatID.(string) + "&video=" + video
	makeRequest(url)
}

func (bot *Bot) SendVideoCustom(chatID interface{}, video string, args ...interface{}) {

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	var url = "https://api.telegram.org/bot" + bot.token + "/sendVideo?chat_id=" + chatID.(string) + "&video=" + video
	if len(args) < 6 {
		for index, arg := range args {
			if arg == nil {
				continue
			}
			switch index {

			case 0:
				param, ok := arg.(int)
				if !ok {
					panic("Type error, duration can only be Integer")
				}
				url += "&duration=" + strconv.Itoa(param)
			case 1:
				param, ok := arg.(int)
				if !ok {
					panic("Type error, width can only be Integer")
				}
				url += "&width=" + strconv.Itoa(param)
			case 2:
				param, ok := arg.(int)
				if !ok {
					panic("Type error, height can only be Integer")
				}
				url += "&height=" + strconv.Itoa(param)
			case 3:
				param, ok := arg.(string)
				if !ok {
					panic("Type error, caption can only be String")
				}
				url += "&caption=" + param
			case 4:
				param, ok := arg.(bool)
				if !ok {
					panic("Type error, disableNotification can only be Bool")
				}
				url += "&disable_notification=" + strconv.FormatBool(param)
			case 5:
				param, ok := arg.(int)
				if !ok {
					panic("Type error, replyToMessageID can only be Integer")
				}
				url += "&reply_to_message_id=" + strconv.Itoa(param)
			}
		}
	} else {
		panic("Too many arguments")
	}
	makeRequest(url)
}

// Send a voice message
func (bot *Bot) SendVoice(chatID interface{}, voice string) {

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	url := "https://api.telegram.org/bot" + bot.token + "/sendVoice?chat_id=" + chatID.(string) + "&voice=" + voice
	makeRequest(url)
}

func (bot *Bot) SendVoiceCustom(chatID interface{}, voice string, args ...interface{}) {

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	var url = "https://api.telegram.org/bot" + bot.token + "/sendVoice?chat_id=" + chatID.(string) + "&voice=" + voice
	if len(args) < 4 {
		for index, arg := range args {
			if arg == nil {
				continue
			}
			switch index {
			case 0:
				param, ok := arg.(string)
				if !ok {
					panic("Type error, caption can only be String")
				}
				url += "&caption=" + param
			case 1:
				param, ok := arg.(int)
				if !ok {
					panic("Type error, duration can only be Integer")
				}
				url += "&duration=" + strconv.Itoa(param)
			case 2:
				param, ok := arg.(bool)
				if !ok {
					panic("Type error, disableNotification can only be Bool")
				}
				url += "&disable_notification=" + strconv.FormatBool(param)
			case 3:
				param, ok := arg.(int)
				if !ok {
					panic("Type error, replyToMessageID can only be Integer")
				}
				url += "&reply_to_message_id=" + strconv.Itoa(param)
			}
		}
	} else {
		panic("Too many arguments")
	}
	makeRequest(url)
}

// Edit a message
func (bot *Bot) EditMessage(chatID interface{}, messageID int, text string) {
	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	var url = "https://api.telegram.org/bot" + bot.token + "/editMessageText?chat_id=" + chatID.(string) + "&message_id=" + strconv.Itoa(messageID) + "&text=" + text

	makeRequest(url)
}

// Delete a message
func (bot *Bot) DeleteMessage(chatID interface{}, messageID int) {
	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	var url = "https://api.telegram.org/bot" + bot.token + "/deleteMessage?chat_id=" + chatID.(string) + "&message_id=" + strconv.Itoa(messageID)

	makeRequest(url)
}
