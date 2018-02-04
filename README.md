# teleroutine
**This is a simple telegram-bot-api written in golang.**

## Installation ##
`go get -u github.com/Pask00/teleroutine`

## Usage ##
Example:
```golang
package main

import(
  bot "github.com/Pask00/teleroutine"
)

token := "mytoken"

func main(){
  myBot := bot.NewBot(token)
  
  myBot.OnMessage(func(message *bot.Message){
    myBot.SendMessage(message.From.ID, "Hello World")
  })
  
  myBot.Listen()
}

```
Learn more through the [wiki](https://github.com/Pask00/teleroutine/wiki)!  

## License ##
**MIT**
