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

type AddStickerResult struct {
	Ok     bool `json:"ok"`
	Result bool `json:"result"`
}

func AddStickerToSet(baseUrl, name, pngSticker, emoji string, userId int64) bool {
	query, err := json.Marshal(TgTypes.AddStickerQuery{UserId: userId, Name: name, PngSticker: pngSticker, Emojis: emoji})
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post(baseUrl+"/addStickerToSet", "application/json", bytes.NewBuffer(query))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	data := AddStickerResult{}
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
