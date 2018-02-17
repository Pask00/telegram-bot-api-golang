package bot

import "strings"

// Listener for a specific text message
func (bot *Bot) On(text string, fun func(*Message)) chan *Message {
	bot.textOnChan = true
	newChan := make(chan *Message)
	bot.textOn = append(bot.textOn, newChan)
	go func() {
		for {
			v, ok := <-newChan
			if !ok {
				bot.textOn = remove(bot.textOn, newChan)
				return
			} else {
				if strings.Index(v.Text, text) == 0 {
					fun(v)
				}
			}

		}
	}()
	return newChan
}

func (bot *Bot) OnJoin(fun func(*Message)) chan *Message {
	bot.joinChan = true
	newChan := make(chan *Message)
	bot.join = append(bot.join, newChan)
	go func() {
		for {
			v, ok := <-newChan
			if !ok {
				bot.join = remove(bot.join, newChan)
				return
			} else {
				fun(v)
			}

		}
	}()
	return newChan
}

func (bot *Bot) OnLeft(fun func(*Message)) chan *Message {
	bot.leftChan = true
	newChan := make(chan *Message)
	bot.left = append(bot.left, newChan)
	go func() {
		for {
			v, ok := <-newChan
			if !ok {
				bot.left = remove(bot.left, newChan)
				return
			} else {
				fun(v)
			}

		}
	}()
	return newChan
}

// Listener for a generic message
func (bot *Bot) OnMessage(fun func(*Message)) chan *Message {
	bot.messageChan = true
	newChan := make(chan *Message)
	bot.message = append(bot.message, newChan)
	go func() {
		for {
			v, ok := <-newChan
			if !ok {
				bot.message = remove(bot.message, newChan)
				return
			} else {
				fun(v)
			}

		}
	}()
	return newChan
}

func remove(s []chan *Message, r chan *Message) []chan *Message {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// Listener for a text message
func (bot *Bot) OnText(fun func(*Message)) chan *Message {
	bot.textChan = true
	newChan := make(chan *Message)
	bot.text = append(bot.text, newChan)
	go func() {
		for {
			v, ok := <-newChan
			if !ok {
				bot.text = remove(bot.text, newChan)
				return
			} else {
				fun(v)
			}

		}
	}()
	return newChan
}

// Listener for a video message
func (bot *Bot) OnVideo(fun func(*Message)) chan *Message {
	bot.videoChan = true
	newChan := make(chan *Message)
	bot.video = append(bot.video, newChan)
	go func() {
		for {
			v, ok := <-newChan
			if !ok {
				bot.video = remove(bot.video, newChan)
				return
			} else {
				fun(v)
			}
		}
	}()
	return newChan
}

// Listener for a photo message
func (bot *Bot) OnPhoto(fun func(*Message)) chan *Message {
	bot.photoChan = true
	newChan := make(chan *Message)
	bot.photo = append(bot.photo, newChan)
	go func() {
		for {
			v, ok := <-newChan
			if !ok {
				bot.photo = remove(bot.photo, newChan)
				return
			} else {
				fun(v)
			}
		}
	}()
	return newChan
}
