package Handlers

import (
	"Telegram-Bot/Features"
	"Telegram-Bot/Features/Autoresponder"
	"Telegram-Bot/Features/Downloader"
	"Telegram-Bot/Features/PhotoFilter"
	"Telegram-Bot/Features/Photomaker"
	"Telegram-Bot/Features/StickerMaker"
	"Telegram-Bot/Lib/MessageMethods"
	"Telegram-Bot/Lib/StickerMethods"
	Functions "Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
	"Telegram-Bot/Settings"
)

func HandleCommand(message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	if message == nil {
		return nil, nil
	}

	thisChatId, thisMessageId, textBody, command, joinedArgument := Functions.ParseMessage(message)
	var err error

	switch command {
	case "":
		err = Autoresponder.FilterMessage(textBody, thisChatId, thisMessageId, Settings.Delay)

	case "menu" + Settings.BotId, "menu":
		return Features.MenuCommand(message)

	case "add", "add" + Settings.BotId:
		err = Autoresponder.AddResponse(joinedArgument, thisChatId, thisMessageId, message.ReplyToMessage)

	case "revoke", "revoke" + Settings.BotId:
		_, err = Autoresponder.StopResponse(joinedArgument, thisChatId, thisMessageId)

	case "filters", "filters" + Settings.BotId:
		return Autoresponder.ReactionList(message)

	case "sticker", "sticker" + Settings.BotId:
		return StickerMaker.MakeSticker(thisChatId, thisMessageId, message.ReplyToMessage, joinedArgument)

	case "photo", "photo" + Settings.BotId:
		return PhotoFilter.StickerToImage(thisChatId, thisMessageId, message.ReplyToMessage)

	case "play", "play" + Settings.BotId:
		return Downloader.YoutubePlay(joinedArgument, thisChatId, thisMessageId)

	case "pfilter", "pfilter" + Settings.BotId:
		return PhotoFilter.HandleFilterPrompt(message)

	case "pmaker", "pmaker" + Settings.BotId:
		return Photomaker.HandlePMakerPrompt(message)

	case "remove", "remove" + Settings.BotId:
		USER, err := Functions.GetChatMember(thisChatId, message.From.Id)
		if err != nil {
			return nil, err
		}
		if USER.CanDeleteMessages == true || USER.Status == "creator" || thisChatId > 0 {
			_, err = StickerMethods.RemoveSticker(thisChatId, thisMessageId, message.ReplyToMessage)
		} else {
			return MessageMethods.SendTextMessage("You can't remove the sticker.", thisChatId, thisMessageId)
		}
		if err != nil {
			return nil, err
		}

	case "data":
		if message.From.Id != Settings.OwnerId {
			return MessageMethods.SendTextMessage("Sorry you can't access the data.", thisChatId, thisMessageId)
		} else {
			return Functions.SendMedia("Data/reactions.json", Functions.Document, thisChatId, thisMessageId, "bot data here", false)
		}
	}
	return nil, err
}
