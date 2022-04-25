package Commands

import (
	Functions "Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type ChatStickerSet struct {
	ChatId int64  `json:"chat_id"`
	Name   string `json:"name"`
	Count  int    `json:"count"`
}

type CreatedSticker struct {
	Data []ChatStickerSet `json:"data"`
}

func MakeSticker(baseUrl, apiToken string, message *TgTypes.MessageType) *TgTypes.MessageType {
	var upStickerFile string
	var imageLink string
	var wasSticker bool

	if message.Document.FileId != "" {
		imagePath := Functions.GetFile(baseUrl, message.Document.FileId).FilePath
		imageLink = "https://api.telegram.org/file/bot" + apiToken + "/" + url.QueryEscape(imagePath)
	} else if message.Sticker.FileId != "" && !message.Sticker.IsAnimated {
		wasSticker = true
		imagePath := Functions.GetFile(baseUrl, message.Sticker.FileId).FilePath
		imageLink = "https://api.violetics.pw/api/converter/webp-to-image?apikey=73dc-6d6c-9e3e&image=" + "https://api.telegram.org/file/bot" + apiToken + "/" + url.QueryEscape(imagePath)
	} else {
		Functions.SendTextMessage(baseUrl, "Where is the image? reply to an image document (uncompressed) or a sticker (static).", message.Chat.Id, message.MessageId)
	}

	//fmt.Println(imageLink)
	res, err := http.Get(imageLink)
	if err != nil {
		log.Fatalln(err)
	}
	data := new(bytes.Buffer)
	_, err = io.Copy(data, res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var resImage *bytes.Buffer

	if wasSticker {
		resImage = data
	} else {
		resImage, err = ResizeImage(data)
	}
	if err != nil {
		go Functions.SendTextMessage(baseUrl, fmt.Sprint(err), message.Chat.Id, message.MessageId)
		return nil
	}

	upStickerFile = Functions.UploadStickerFile(baseUrl, message.From.Id, resImage).FileId

	storage, _ := ioutil.ReadFile("Data/createdStickers.json")
	stickerData := CreatedSticker{}
	err = json.Unmarshal(storage, &stickerData)

	var sentMessage *TgTypes.MessageType
	var done bool
	for k, chats := range stickerData.Data {
		if chats.ChatId == message.Chat.Id {
			if Functions.AddStickerToSet(baseUrl, chats.Name, upStickerFile, "😂", 1653921867) {
				chats.Count++
				stickerData.Data[k] = chats
				sentMessage = Functions.SendTextMessage(baseUrl, "Sticker added to <a href=\"https://t.me/addstickers/"+chats.Name+"\">Pack</a>", message.Chat.Id, message.MessageId)
			} else {
				sentMessage = Functions.SendTextMessage(baseUrl, "Adding of sticker failed.", message.Chat.Id, message.MessageId)
			}
			done = true
			break
		}
	}

	if !done {
		packName := "x" + fmt.Sprint(uint64(message.Chat.Id)) + "_by_AB22TGBot"
		fmt.Println("Packname", packName)
		title := message.Chat.Title + " Daemon-Bot"
		if Functions.CreateStickerSet(baseUrl, packName, title, "🌝", upStickerFile, 1653921867) {
			if stickerData.Data == nil {
				stickerData.Data = []ChatStickerSet{ChatStickerSet{
					ChatId: message.Chat.Id,
					Name:   packName,
					Count:  1,
				}}
			} else {
				stickerData.Data = append(stickerData.Data, ChatStickerSet{
					ChatId: message.Chat.Id,
					Name:   packName,
					Count:  1,
				})
			}

			sentMessage = Functions.SendTextMessage(baseUrl, "Sticker added to <a href=\"https://t.me/addstickers/"+packName+"\">Pack</a>", message.Chat.Id, message.MessageId)
		} else {
			sentMessage = Functions.SendTextMessage(baseUrl, "Adding of sticker failed.", message.Chat.Id, message.MessageId)
		}
	}
	//fmt.Println(stickerData)
	byteData, err := json.MarshalIndent(stickerData, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}

	err = ioutil.WriteFile("Data/createdStickers.json", byteData, 0)
	if err != nil {
		log.Fatalln(err)
	}

	return sentMessage
}
