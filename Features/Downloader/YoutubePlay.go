package Downloader

import (
	"Telegram-Bot/Globals"
	Functions "Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
	"Telegram-Bot/Settings"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type YoutubePlayResult struct {
	Status  int             `json:"status"`
	Result  YoutubePlayType `json:"result"`
	Message []string        `json:"message"`
}

type YoutubePlayType struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	Thumb      string `json:"thumb"`
	Duration   string `json:"duration"`
	QualityMp4 string `json:"quality_mp4"`
	SizeMp4    string `json:"size_mp4"`
	VideoLink  string `json:"dlmp4"`
	QualityMp3 string `json:"quality_mp3"`
	SizeMp3    string `json:"size_mp3"`
	AudioLink  string `json:"dlmp3"`
}

func YoutubePlay(baseUrl, query string, chatId, messageId int64) (*TgTypes.MessageType, error) {
	resp, err := http.Get("https://violetics.pw/api/media/youtube-play?apikey=" + Settings.VioKey + "&query=" + url.QueryEscape(query))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := YoutubePlayResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if !(data.Status == 200) {
		return nil, errors.New(data.Message[0])
	}
	result := data.Result

	buttons := make([][]TgTypes.InlineKeyboardButtonType, 0)

	row1 := make([]TgTypes.InlineKeyboardButtonType, 0)
	row1 = append(row1, TgTypes.InlineKeyboardButtonType{Text: "Audio 🎵", CallbackData: "ytAudio"})
	row1 = append(row1, TgTypes.InlineKeyboardButtonType{Text: "Video 📽️", CallbackData: "ytVideo"})
	buttons = append(buttons, row1)
	row2 := make([]TgTypes.InlineKeyboardButtonType, 0)
	row2 = append(row2, TgTypes.InlineKeyboardButtonType{Text: "Exit", CallbackData: "deleteMessage"})
	buttons = append(buttons, row2)

	buttonMessage, err := Functions.SendButtons(baseUrl, result.Title, chatId, messageId, TgTypes.InlineKeyboardMarkupType{InlineKeyboard: buttons})
	if err != nil {
		return nil, err
	}

	Globals.AudioLinks[buttonMessage.MessageId] = result.AudioLink
	Globals.VideoLinks[buttonMessage.MessageId] = result.VideoLink

	return buttonMessage, nil
}
