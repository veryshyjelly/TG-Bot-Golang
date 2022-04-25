package Functions

import (
	"Telegram-Bot/Lib/TgTypes"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type CreateStickerSetResult struct {
	Ok     bool `json:"ok"`
	Result bool `json:"result"`
}

func CreateStickerSet(baseUrl, name, title, emoji, sticker string, userId int64) bool {

	query, err := json.Marshal(TgTypes.CreateStickerSetQuery{UserId: userId, Name: name, Title: title, PngSticker: sticker, Emojis: emoji})
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post(baseUrl+"/createNewStickerSet", "application/json", bytes.NewBuffer(query))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	data := CreateStickerSetResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))
	if !data.Ok {
		return false
	}

	return data.Result
}
