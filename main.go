package main

import (
	"Telegram-Bot/Commands"
	"Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	apiToken := "5238894566:AAEvB2BgSISLA_tl_Cs6bpDU2lR66JbkMmI"
	baseUrl := "https://api.telegram.org/bot" + apiToken
	botName := "@AB22TGBot"
	//var chat2 int64 = -703095609
	//var chat1 int64 = -1001684508017
	var offset, limit int64 = 0, 100
	delay := 10
	for {
		response, err := Functions.GetMessage(baseUrl, offset, limit)
		if err != nil {
			log.Fatalln(err)
		}

		for _, messages := range response {
			go func(messages TgTypes.UpdateType) {
				thisChatId, thisMessageId, textBody, command, joinedArgument := Functions.ParseMessage(&messages.Message)
				fmt.Println(messages)
				if strings.ToLower(messages.Message.Text) == "hi bot" {
					_, err = Functions.SendTextMessage(baseUrl, "I made this bot from scratch", thisChatId, thisMessageId)
				}

				switch command {
				case "":
					err = Commands.FilterMessage(baseUrl, textBody, thisChatId, thisMessageId, delay)
				case "menu" + botName, "menu":
					_, err = Commands.MenuCommand(baseUrl, &messages.Message)
				case "add", "addfilter", "addsticker", "add" + botName:
					if messages.Message.ReplyToMessage == nil {
						_, err = Functions.SendTextMessage(baseUrl, "Please reply to a message", thisChatId, thisMessageId)
					} else {
						err = Commands.AddResponse(baseUrl, joinedArgument, thisChatId, thisMessageId, messages.Message.ReplyToMessage)
					}
				case "revoke", "revoke" + botName:
					_, err = Commands.StopResponse(baseUrl, joinedArgument, thisChatId, thisMessageId)
				case "filters", "filters" + botName:
					_, err = Commands.ReactionList(baseUrl, &messages.Message)
				case "resize", "resize" + botName:
					if messages.Message.Document.FileId == "" && (messages.Message.ReplyToMessage == nil || messages.Message.ReplyToMessage.Document.FileId == "") {
						_, err = Functions.SendTextMessage(baseUrl, "Where is the image? reply to an image document (uncompressed).", thisChatId, thisMessageId)
					} else {
						if messages.Message.Document.FileId != "" {
							_, err = Commands.SendResizeImage(baseUrl, apiToken, &messages.Message)
						} else {
							_, err = Commands.SendResizeImage(baseUrl, apiToken, messages.Message.ReplyToMessage)
						}
					}
				case "sticker", "sticker" + botName:
					if messages.Message.Document.FileId != "" {
						_, err = Commands.MakeSticker(baseUrl, apiToken, &messages.Message)
					} else if messages.Message.ReplyToMessage == nil {
						_, err = Functions.SendTextMessage(baseUrl, "Where is the image? reply to an image document (uncompressed) or a sticker (static).", thisChatId, thisMessageId)
					} else {
						_, err = Commands.MakeSticker(baseUrl, apiToken, messages.Message.ReplyToMessage)
					}
				case "remove", "remove" + botName:
					USER, err := Functions.GetChatMember(baseUrl, thisChatId, messages.Message.From.Id)
					if err != nil {
						log.Println(err)
					}
					if USER.CanDeleteMessages == true || USER.Status == "creator" {
						_, err = Functions.RemoveSticker(baseUrl, thisChatId, thisMessageId, messages.Message.ReplyToMessage)
					} else {
						_, err = Functions.SendTextMessage(baseUrl, "You can't remove the sticker.", thisChatId, thisMessageId)
					}
					if err != nil {
						log.Println(err)
					}
				}

				if messages.CallbackQuery.Id != "" {
					if messages.CallbackQuery.Data == "stickerMenu" {
						err = Commands.StickerMenu(baseUrl, messages.CallbackQuery.Id)
					} else if messages.CallbackQuery.Data == "filterMenu" {
						err = Commands.FilterMenu(baseUrl, messages.CallbackQuery.Id)
					} else {
						_, err = Functions.AnswerCallbackQuery(baseUrl, messages.CallbackQuery.Id, "Answering Query", true)
					}
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
