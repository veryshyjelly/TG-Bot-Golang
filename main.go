package main

import (
	"Telegram-Bot/Handlers"
	"Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
	"Telegram-Bot/Settings"
	"fmt"
	"log"
	"time"
)

func main() {
	baseUrl := "https://api.telegram.org/bot" + Settings.ApiToken

	var offset, limit int64 = 0, 100
	for {
		response, err := Functions.GetMessage(baseUrl, offset, limit)
		if err != nil {
			log.Fatalln(err)
		}

		for _, messages := range response {
			go func(messages TgTypes.UpdateType) {
				fmt.Println(messages)
				_, err = Handlers.HandleCommand(baseUrl, &messages.Message)
				if err != nil {
					log.Println(err)
				}

				if messages.CallbackQuery.Id != "" {
					_, err = Handlers.HandleCallback(baseUrl, messages.CallbackQuery)
				}
				if err != nil {
					log.Println(err)
				}
			}(messages)

			offset = messages.UpdateId + 1
		}
		time.Sleep(80 * time.Millisecond)
	}
}
