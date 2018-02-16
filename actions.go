package bot

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// It makes requests to telegram API
func (bot *Bot) makeRequest(action string, values url.Values) {

	myurl := "http://api.telegram.org/bot" + bot.token + "/" + action + "?" + values.Encode()

	req, err := http.NewRequest("GET", myurl, nil)
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

	var update Response

	if err := json.NewDecoder(resp.Body).Decode(&update); err != nil {
		log.Println(err)
	}

	resp.Body.Close()

	if !update.Ok {
		fmt.Println("Error: " + strconv.Itoa(update.ErrorCode) + "\tDescription: " + update.Description)
	}
}

func (bot *Bot) makeRequestWithReturn(action string, values url.Values) Response {

	myurl := "http://api.telegram.org/bot" + bot.token + "/" + action + "?" + values.Encode()

	req, err := http.NewRequest("GET", myurl, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	var apiResp Response

	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		log.Println(err)
	}

	resp.Body.Close()

	if !apiResp.Ok {
		fmt.Println("Error: " + strconv.Itoa(apiResp.ErrorCode) + "\tDescription: " + apiResp.Description)
	}

	return apiResp
}

func (bot *Bot) SendMessageCustom(chatID interface{}, text string, args ...interface{}) {

	values := url.Values{}

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	values.Set("chat_id", chatID.(string))
	values.Set("text", text)

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
				values.Set("parse_mode", param)
			case 1:
				param, ok := arg.(bool)
				if !ok {
					panic("Type error, disableWebPagePreview can only be Bool")
				}
				values.Set("disable_web_page_preview", strconv.FormatBool(param))
			case 2:
				param, ok := arg.(bool)
				if !ok {
					panic("Type error, disableNotification can only be Bool")
				}
				values.Set("disable_notification", strconv.FormatBool(param))
			case 3:
				param, ok := arg.(int)
				if !ok {
					panic("Type error, replyToMessageID can only be Integer")
				}
				values.Set("reply_to_message_id", strconv.Itoa(param))
			}
		}
	} else {
		panic("Too many arguments")
	}
	bot.makeRequest("sendMessage", values)
}

// Send a text message
func (bot *Bot) SendMessage(chatID interface{}, text string) {
	values := url.Values{}

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	values.Set("chat_id", chatID.(string))
	values.Set("text", text)

	bot.makeRequest("sendMessage", values)
}

// Forward a message
func (bot *Bot) SendForward(chatID interface{}, from interface{}, messageID int) {
	values := url.Values{}

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	values.Set("chat_id", chatID.(string))
	values.Set("message_id", strconv.Itoa(messageID))

	switch param := from.(type) {
	case int:
		from = strconv.Itoa(param)
	case string:
		from = param
	default:
		panic("Type error, from can be only Integer or String")
	}
	values.Set("from", from.(string))

	bot.makeRequest("forwardMessage", values)
}

// Send a reply message
func (bot *Bot) SendReply(chatID interface{}, text string, messageID int) {
	values := url.Values{}

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	values.Set("chat_id", chatID.(string))
	values.Set("text", text)
	values.Set("reply_to_message_id", strconv.Itoa(messageID))

	bot.makeRequest("sendMessage", values)
}

// Send a markdown message
func (bot *Bot) SendMarkdown(chatID interface{}, text string) {
	values := url.Values{}

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	values.Set("chat_id", chatID.(string))
	values.Set("text", text)
	values.Set("parse_mode", "Markdown")

	bot.makeRequest("sendMessage", values)
}

// Send a HTML message
func (bot *Bot) SendHTML(chatID interface{}, text string) {
	values := url.Values{}

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	values.Set("chat_id", chatID.(string))
	values.Set("text", text)
	values.Set("parse_mode", "HTML")

	bot.makeRequest("sendMessage", values)
}

// Send a photo message
func (bot *Bot) SendPhoto(chatID interface{}, photo string) {
	values := url.Values{}

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	values.Set("chat_id", chatID.(string))
	values.Set("photo", photo)

	bot.makeRequest("sendPhoto", values)
}

func (bot *Bot) SendPhotoCustom(chatID interface{}, photo string, args ...interface{}) {
	values := url.Values{}

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	values.Set("chat_id", chatID.(string))
	values.Set("photo", photo)

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
				values.Set("caption", param)
			case 1:
				param, ok := arg.(bool)
				if !ok {
					panic("Type error, disableNotification can only be Bool")
				}
				values.Set("disable_notification", strconv.FormatBool(param))
			case 2:
				param, ok := arg.(int)
				if !ok {
					panic("Type error, replyToMessageID can only be Integer")
				}
				values.Set("reply_to_message_id", strconv.Itoa(param))
			}
		}
	} else {
		panic("Too many arguments")
	}
	bot.makeRequest("sendPhoto", values)
}

// Send an audio message
func (bot *Bot) SendAudio(chatID interface{}, audio string) {
	values := url.Values{}

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	values.Set("chat_id", chatID.(string))
	values.Set("audio", audio)

	bot.makeRequest("sendAudio", values)
}

func (bot *Bot) SendAudioCustom(chatID interface{}, audio string, args ...interface{}) {
	values := url.Values{}

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	values.Set("chat_id", chatID.(string))
	values.Set("audio", audio)

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
				values.Set("caption", param)
			case 1:
				param, ok := arg.(int)
				if !ok {
					panic("Type error, duration can only be Integer")
				}
				values.Set("duration", strconv.Itoa(param))
			case 2:
				param, ok := arg.(string)
				if !ok {
					panic("Type performer, performer can only be String")
				}
				values.Set("performer", param)
			case 3:
				param, ok := arg.(string)
				if !ok {
					panic("Type title, title can only be String")
				}
				values.Set("title", param)
			case 4:
				param, ok := arg.(bool)
				if !ok {
					panic("Type error, disableNotification can only be Bool")
				}
				values.Set("disable_notification", strconv.FormatBool(param))
			case 5:
				param, ok := arg.(int)
				if !ok {
					panic("Type error, replyToMessageID can only be Integer")
				}
				values.Set("reply_to_message_id", strconv.Itoa(param))
			}
		}
	} else {
		panic("Too many arguments")
	}
	bot.makeRequest("sendAudio", values)
}

//Send a document message
func (bot *Bot) SendDocument(chatID interface{}, document string) {
	values := url.Values{}

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	values.Set("chat_id", chatID.(string))
	values.Set("document", document)

	bot.makeRequest("sendDocument", values)
}

func (bot *Bot) SendDocumentCustom(chatID interface{}, document string, args ...interface{}) {
	values := url.Values{}

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	values.Set("chat_id", chatID.(string))
	values.Set("document", document)

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
				values.Set("caption", param)
			case 1:
				param, ok := arg.(bool)
				if !ok {
					panic("Type error, disableNotification can only be Bool")
				}
				values.Set("disable_notification", strconv.FormatBool(param))
			case 2:
				param, ok := arg.(int)
				if !ok {
					panic("Type error, replyToMessageID can only be Integer")
				}
				values.Set("reply_to_message_id", strconv.Itoa(param))
			}
		}
	} else {
		panic("Too many arguments")
	}
	bot.makeRequest("sendDocument", values)
}

// Send a video message
func (bot *Bot) SendVideo(chatID interface{}, video string) {
	values := url.Values{}

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	values.Set("chat_id", chatID.(string))
	values.Set("video", video)

	bot.makeRequest("sendVideo", values)
}

func (bot *Bot) SendVideoCustom(chatID interface{}, video string, args ...interface{}) {
	values := url.Values{}

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	values.Set("chat_id", chatID.(string))
	values.Set("video", video)

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
				values.Set("duration", strconv.Itoa(param))
			case 1:
				param, ok := arg.(int)
				if !ok {
					panic("Type error, width can only be Integer")
				}
				values.Set("width", strconv.Itoa(param))
			case 2:
				param, ok := arg.(int)
				if !ok {
					panic("Type error, height can only be Integer")
				}
				values.Set("height", strconv.Itoa(param))
			case 3:
				param, ok := arg.(string)
				if !ok {
					panic("Type error, caption can only be String")
				}
				values.Set("caption", param)
			case 4:
				param, ok := arg.(bool)
				if !ok {
					panic("Type error, disableNotification can only be Bool")
				}
				values.Set("disable_notification", strconv.FormatBool(param))
			case 5:
				param, ok := arg.(int)
				if !ok {
					panic("Type error, replyToMessageID can only be Integer")
				}
				values.Set("reply_to_message_id", strconv.Itoa(param))
			}
		}
	} else {
		panic("Too many arguments")
	}
	bot.makeRequest("sendVideo", values)
}

// Send a voice message
func (bot *Bot) SendVoice(chatID interface{}, voice string) {
	values := url.Values{}

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	values.Set("chat_id", chatID.(string))
	values.Set("voice", voice)

	bot.makeRequest("sendVoice", values)
}

func (bot *Bot) SendVoiceCustom(chatID interface{}, voice string, args ...interface{}) {
	values := url.Values{}

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	values.Set("chat_id", chatID.(string))
	values.Set("voice", voice)

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
				values.Set("caption", param)
			case 1:
				param, ok := arg.(int)
				if !ok {
					panic("Type error, duration can only be Integer")
				}
				values.Set("duration", strconv.Itoa(param))
			case 2:
				param, ok := arg.(bool)
				if !ok {
					panic("Type error, disableNotification can only be Bool")
				}
				values.Set("disable_notification", strconv.FormatBool(param))
			case 3:
				param, ok := arg.(int)
				if !ok {
					panic("Type error, replyToMessageID can only be Integer")
				}
				values.Set("reply_to_message_id", strconv.Itoa(param))
			}
		}
	} else {
		panic("Too many arguments")
	}
	bot.makeRequest("sendVoice", values)
}

// Edit a message
func (bot *Bot) EditMessage(chatID interface{}, messageID int, text string) {
	values := url.Values{}

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	values.Set("chat_id", chatID.(string))
	values.Set("message_id", strconv.Itoa(messageID))
	values.Set("text", text)

	bot.makeRequest("editMessage", values)
}

// Delete a message
func (bot *Bot) DeleteMessage(chatID interface{}, messageID int) {
	values := url.Values{}

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	values.Set("chat_id", chatID.(string))
	values.Set("message_id", strconv.Itoa(messageID))

	bot.makeRequest("deleteMessage", values)
}

func (bot *Bot) Ban(chatID interface{}, message *Message) {
	values := url.Values{}

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	values.Set("chat_id", chatID.(string))
	values.Set("user_id", strconv.Itoa(message.ReplyToMessage.From.ID))

	bot.makeRequest("kickChatMember", values)
}

func (bot *Bot) GetChat(chatID interface{}) *Chat {
	values := url.Values{}

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	values.Set("chat_id", chatID.(string))

	var chat *Chat

	val, err := json.Marshal(bot.makeRequestWithReturn("getChat", values).Result)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(val, &chat)

	return chat
}

func (bot *Bot) GetMe() *User {
	var user *User

	val, err := json.Marshal(bot.makeRequestWithReturn("getMe", nil).Result)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(val, &user)

	return user
}

func (bot *Bot) GetChatMember(chatID interface{}, userID int) *ChatMember {
	values := url.Values{}

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	values.Set("chat_id", chatID.(string))
	values.Set("user_id", strconv.Itoa(userID))

	var chatMember *ChatMember

	val, err := json.Marshal(bot.makeRequestWithReturn("getChatMember", values).Result)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(val, &chatMember)

	return chatMember
}

func (bot *Bot) ReplyKeyboardMarkup(chatID interface{}, messageID int, text string, texts ...string) {
	values := url.Values{}

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	var options string

	for index, words := range texts {
		if len(texts) == index+1 {
			options += "[\"" + words + "\"]"
		} else {
			options += "[\"" + words + "\"],"
		}
	}
	values.Set("reply_markup", `{"keyboard":[`+options+`], "selective": true}`)
	values.Set("chat_id", chatID.(string))
	values.Set("reply_to_message_id", strconv.Itoa(messageID))
	values.Set("text", text)

	fmt.Println(values.Encode())

	bot.makeRequest("sendMessage", values)

}

func (bot *Bot) RemoveKeyboard(chatID interface{}, messageID int, text string) {
	values := url.Values{}

	switch param := chatID.(type) {
	case int:
		chatID = strconv.Itoa(param)
	case string:
		chatID = param
	default:
		panic("Type error, chatID can be only Integer or String")
	}

	values.Set("reply_markup", `{"remove_keyboard":true, "selective": true}`)
	values.Set("chat_id", chatID.(string))
	values.Set("reply_to_message_id", strconv.Itoa(messageID))
	values.Set("text", text)

	fmt.Println(values.Encode())

	bot.makeRequest("sendMessage", values)

}
