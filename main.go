package main

import (
	"Telegram-Bot/Commands"
	"Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
	"fmt"
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
		response := Functions.GetMessage(baseUrl, offset, limit)

		for _, messages := range response {
			go func(messages TgTypes.UpdateType) {
				thisChatId, thisMessageId, textBody, command, joinedArgument := Functions.ParseMessage(&messages.Message)
				fmt.Println(messages)
				if strings.ToLower(messages.Message.Text) == "hi bot" {
					Functions.SendTextMessage(baseUrl, "I made this bot from scratch", thisChatId, thisMessageId)
				}

				switch command {
				case "":
					Commands.FilterMessage(baseUrl, textBody, thisChatId, thisMessageId, delay)
				case "menu" + botName, "menu":
					Commands.MenuCommand(baseUrl, &messages.Message)
				case "add", "addfilter", "addsticker":
					if messages.Message.ReplyToMessage == nil {
						Functions.SendTextMessage(baseUrl, "Please reply to a message", thisChatId, thisMessageId)
					} else {
						Commands.AddResponse(baseUrl, joinedArgument, thisChatId, thisMessageId, messages.Message.ReplyToMessage)
					}
				case "revoke":
					Commands.StopResponse(baseUrl, joinedArgument, thisChatId, thisMessageId)
				case "filters":
					Commands.ReactionList(baseUrl, &messages.Message)
				case "resize":
					if messages.Message.Document.FileId == "" && (messages.Message.ReplyToMessage == nil || messages.Message.ReplyToMessage.Document.FileId == "") {
						Functions.SendTextMessage(baseUrl, "Where is the image? reply to an image document (uncompressed).", thisChatId, thisMessageId)
					} else {
						if messages.Message.Document.FileId != "" {
							go Commands.SendResizeImage(baseUrl, apiToken, &messages.Message)
						} else {
							go Commands.SendResizeImage(baseUrl, apiToken, messages.Message.ReplyToMessage)
						}
					}
				case "sticker":
					if messages.Message.Document.FileId == "" && (messages.Message.ReplyToMessage == nil || messages.Message.ReplyToMessage.Document.FileId == "") {
						Functions.SendTextMessage(baseUrl, "Where is the image? reply to an image document (uncompressed).", thisChatId, thisMessageId)
					} else {
						if messages.Message.Document.FileId != "" {
							go Commands.MakeSticker(baseUrl, apiToken, &messages.Message)
						} else {
							go Commands.MakeSticker(baseUrl, apiToken, messages.Message.ReplyToMessage)
						}
					}
				case "remove":
					USER := Functions.GetChatMember(baseUrl, thisChatId, messages.Message.From.Id)
					if USER.CanDeleteMessages == true || USER.Status == "creator" {
						Functions.RemoveSticker(baseUrl, thisChatId, thisMessageId, messages.Message.ReplyToMessage)
					} else {
						Functions.SendTextMessage(baseUrl, "You can't remove the sticker.", thisChatId, thisMessageId)
					}
				}

				if messages.CallbackQuery.Id != "" {
					if messages.CallbackQuery.Data == "stickerMenu" {
						Commands.StickerMenu(baseUrl, messages.CallbackQuery.Id)
					} else if messages.CallbackQuery.Data == "filterMenu" {
						Commands.FilterMenu(baseUrl, messages.CallbackQuery.Id)
					} else {
						Functions.AnswerCallbackQuery(baseUrl, messages.CallbackQuery.Id, "Answering Query", true)
					}
				}
			}(messages)
			offset = messages.UpdateId + 1
		}
		time.Sleep(80 * time.Millisecond)
	}
}
