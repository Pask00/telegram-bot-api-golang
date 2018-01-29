# telegram-bot-api-golang
**This is a simple telegram-bot-api written in golang.**

## Installation ##
`go get -u github.com/Pask00/telegram-bot-api-golang`

## Usage ##
Example:
```golang
package main

import(
  bot "github.com/Pask00/telegram-bot-api-golang"
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
## License ##
**MIT**
