package StickerMethods

import (
	"Telegram-Bot/Lib/TgTypes"
	"Telegram-Bot/Settings"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type AddStickerResult struct {
	Ok          bool   `json:"ok"`
	Result      bool   `json:"result"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

func AddStickerToSet(name, pngSticker, emoji string, userId int64) (bool, error) {
	query, err := json.Marshal(TgTypes.AddStickerQuery{
		UserId:     userId,
		Name:       name,
		PngSticker: pngSticker,
		Emojis:     emoji,
	})
	if err != nil {
		return false, err
	}

	resp, err := http.Post(Settings.BaseUrl+"/addStickerToSet", "application/json", bytes.NewBuffer(query))
	if err != nil {
		return false, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	data := AddStickerResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return false, err
	}

	if !data.Ok {
		return false, errors.New(data.Description)
	}

	return data.Result, nil
}
