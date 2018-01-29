package bot

import (
	"strings"
)

// Listener for a specific text message
func (bot *Bot) On(text string, fun function) {
	bot.textOnChan = true
	go func() {
		for {
			select {
			case s := <-bot.textOn:
				if strings.Index(s.Text, text) == 0 {
					fun(s)
				}
			}
		}
	}()
}

// Listener for a generic message
func (bot *Bot) OnMessage(fun function) {
	bot.messageChan = true
	go func() {
		for {
			select {
			case s := <-bot.message:
				fun(s)
			}
		}
	}()
}

// Listener for a text message
func (bot *Bot) OnText(fun function) {
	bot.textChan = true
	go func() {
		for {
			select {
			case s := <-bot.text:
				fun(s)
			}
		}
	}()
}

// Listener for a video message
func (bot *Bot) OnVideo(fun function) {
	bot.videoChan = true
	go func() {
		for {
			select {
			case s := <-bot.video:
				fun(s)
			}
		}
	}()
}

// Listener for a photo message
func (bot *Bot) OnPhoto(fun function) {
	bot.photoChan = true
	go func() {
		for {
			select {
			case s := <-bot.photo:
				fun(s)
			}
		}
	}()
}
