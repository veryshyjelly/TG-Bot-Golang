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

type StickerSetQuery struct {
	Name string `json:"name"`
}

type StickerSetResult struct {
	Ok          bool `json:"ok"`
	Result      TgTypes.StickerSetType
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

func GetStickerSet(Name string) (*TgTypes.StickerSetType, error) {
	query, err := json.Marshal(StickerSetQuery{Name: Name})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(Settings.BaseUrl+"/getStickerSet", "application/json", bytes.NewBuffer(query))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := StickerSetResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if !data.Ok {
		return nil, errors.New(data.Description)
	}

	return &data.Result, nil
}
