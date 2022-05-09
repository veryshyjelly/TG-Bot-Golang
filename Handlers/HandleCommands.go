package Handlers

import (
	"Telegram-Bot/Features"
	"Telegram-Bot/Features/Autoresponder"
	"Telegram-Bot/Features/Downloader"
	"Telegram-Bot/Features/PhotoFilter"
	"Telegram-Bot/Features/StickerMaker"
	"Telegram-Bot/Lib/MediaFunctions"
	"Telegram-Bot/Lib/MessageMethods"
	"Telegram-Bot/Lib/StickerMethods"
	Functions "Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
	"Telegram-Bot/Settings"
	"log"
	"strings"
)

func HandleCommand(baseUrl string, message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	thisChatId, thisMessageId, textBody, command, joinedArgument := Functions.ParseMessage(message)
	var err error

	if strings.ToLower(message.Text) == "hi bot" {
		_, err = MessageMethods.SendTextMessage(baseUrl, "I made this bot from scratch", thisChatId, thisMessageId)
	}
	switch command {
	case "":
		err = Autoresponder.FilterMessage(baseUrl, textBody, thisChatId, thisMessageId, Settings.Delay)
	case "menu" + Settings.BotId, "menu":
		return Features.MenuCommand(baseUrl, message)
	case "add", "addfilter", "addsticker", "add" + Settings.BotId:
		err = Autoresponder.AddResponse(baseUrl, joinedArgument, thisChatId, thisMessageId, message.ReplyToMessage)
	case "revoke", "revoke" + Settings.BotId:
		_, err = Autoresponder.StopResponse(baseUrl, joinedArgument, thisChatId, thisMessageId)
	case "filters", "filters" + Settings.BotId:
		return Autoresponder.ReactionList(baseUrl, message)
	case "sticker", "sticker" + Settings.BotId:
		return StickerMaker.MakeSticker(baseUrl, thisChatId, thisMessageId, message.ReplyToMessage)
	case "photo", "photo" + Settings.BotId:
		return PhotoFilter.StickerToImage(baseUrl, thisChatId, thisMessageId, message.ReplyToMessage)
	case "play", "play" + Settings.BotId:
		return Downloader.YoutubePlay(baseUrl, joinedArgument, thisChatId, thisMessageId)
	case "remove", "remove" + Settings.BotId:
		USER, err := Functions.GetChatMember(baseUrl, thisChatId, message.From.Id)
		if err != nil {
			log.Println(err)
		}
		if USER.CanDeleteMessages == true || USER.Status == "creator" || thisChatId > 0 {
			_, err = StickerMethods.RemoveSticker(baseUrl, thisChatId, thisMessageId, message.ReplyToMessage)
		} else {
			return MessageMethods.SendTextMessage(baseUrl, "You can't remove the sticker.", thisChatId, thisMessageId)
		}
		if err != nil {
			log.Println(err)
		}

	case "data":
		if message.From.Id != Settings.OwnerId {
			return MessageMethods.SendTextMessage(baseUrl, "Sorry you can't access the data.", thisChatId, thisMessageId)
		} else {
			return MediaFunctions.SendMedia(baseUrl, "Data/reactions.json", MediaFunctions.Document, thisChatId, thisMessageId, "bot data here", false)
		}
	}
	return nil, err
}
