package main

import (
	"Telegram-Bot/Globals"
	"Telegram-Bot/Handlers"
	"Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
	"fmt"
	"log"
	"time"
)

func main() {
	err := Globals.ParseVioApi()
	if err != nil {
		log.Fatalln(err)
	}

	var offset, limit int64 = 0, 100
	for {
		response, err := Functions.GetMessage(offset, limit)
		if err != nil {
			log.Fatalln(err)
		}

		for _, messages := range response {
			go func(messages TgTypes.UpdateType) {
				fmt.Println(messages)
				_, err = Handlers.HandleCommand(&messages.Message)
				if err != nil {
					log.Println(err)
				}

				if messages.CallbackQuery.Id != "" {
					_, err = Handlers.HandleCallback(&messages.CallbackQuery)
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
