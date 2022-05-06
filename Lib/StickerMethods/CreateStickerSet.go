package StickerMethods

import (
	"Telegram-Bot/Lib/TgTypes"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type CreateStickerSetResult struct {
	Ok          bool   `json:"ok"`
	Result      bool   `json:"result"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

func CreateStickerSet(baseUrl, name, title, emoji, sticker string, userId int64) (bool, error) {
	query, err := json.Marshal(TgTypes.CreateStickerSetQuery{
		UserId:     userId,
		Name:       name,
		Title:      title,
		PngSticker: sticker,
		Emojis:     emoji,
	})
	if err != nil {
		return false, err
	}

	resp, err := http.Post(baseUrl+"/createNewStickerSet", "application/json", bytes.NewBuffer(query))
	if err != nil {
		return false, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	data := CreateStickerSetResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return false, err
	}

	//fmt.Println(string(body))
	if !data.Ok {
		return false, errors.New(data.Description)
	}

	return data.Result, nil
}
