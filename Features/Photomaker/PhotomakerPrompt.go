package Photomaker

import (
	"Telegram-Bot/Globals"
	"Telegram-Bot/Lib/MessageMethods"
	"Telegram-Bot/Lib/TgTypes"
	"strconv"
)

func PMakerButtons(offset int) TgTypes.InlineKeyboardMarkupType {
	initialOffset := offset
	buttons := make([][]TgTypes.InlineKeyboardButtonType, 0)
	dataList, dataSize := Globals.VioAPIData["photomaker"].Stacks, Globals.VioAPIData["photomaker"].Total

	for i := 0; i < 5 && offset < dataSize; i++ {
		row := make([]TgTypes.InlineKeyboardButtonType, 0)
	LOOP:
		for j := 0; j < 3 && offset < dataSize; offset++ {
			for k, _ := range dataList[offset].Params {
				if k != "image" && k != "image2" && k != "animation" && k != "colour" {
					continue LOOP
				}
			}
			row = append(row, TgTypes.InlineKeyboardButtonType{Text: dataList[offset].Title, CallbackData: "pMaker " + dataList[offset].Name})
			j++
		}
		buttons = append(buttons, row)
	}

	row := make([]TgTypes.InlineKeyboardButtonType, 0)
	if offset > 15 {
		row = append(row, TgTypes.InlineKeyboardButtonType{Text: "Back", CallbackData: "pMakerNext " + strconv.Itoa(initialOffset-15)})
	}
	row = append(row, Globals.ExitButton)
	if offset < dataSize {
		row = append(row, TgTypes.InlineKeyboardButtonType{Text: "Next", CallbackData: "pMakerNext " + strconv.Itoa(offset)})
	}
	buttons = append(buttons, row)

	return TgTypes.InlineKeyboardMarkupType{InlineKeyboard: buttons}
}

func HandlePMakerPrompt(message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	if message.ReplyToMessage == nil || message.ReplyToMessage.Document.FileId == "" {
		return MessageMethods.SendTextMessage("reply to an image (uncompressed).", message.Chat.Id, message.MessageId)
	}

	replyMarkup := PMakerButtons(0)

	buttonMessage, err := MessageMethods.SendButtons("Choose Option:", message.Chat.Id, message.MessageId, replyMarkup, true)
	if err != nil {
		return nil, err
	}

	Globals.PhotoFilterQueue[buttonMessage.MessageId] = message.ReplyToMessage.Document.FileId

	return buttonMessage, nil
}

func HandlePMakerNext(offset string, queryId string, message *TgTypes.MessageType) (*TgTypes.MessageType, error) {
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return nil, err
	}

	replyMarkup := PMakerButtons(offsetInt)

	return MessageMethods.EditMessageMarkup(message.Chat.Id, message.MessageId, replyMarkup, "")
}
