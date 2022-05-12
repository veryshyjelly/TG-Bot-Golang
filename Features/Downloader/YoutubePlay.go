package Downloader

import (
	"Telegram-Bot/Globals"
	Functions "Telegram-Bot/Lib/MessageMethods"
	"Telegram-Bot/Lib/TgTypes"
	"Telegram-Bot/Settings"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type YoutubePlayResult struct {
	Status  int             `json:"status"`
	Result  YoutubePlayType `json:"result,omitempty"`
	Message interface{}     `json:"message"`
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

func YoutubePlay(query string, chatId, messageId int64) (*TgTypes.MessageType, error) {
	if query == "" {
		return Functions.SendTextMessage("query is empty. Add keywords or youtube link after command.", chatId, messageId)
	}
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
		return nil, errors.New(fmt.Sprint(data.Message))
	}
	result := data.Result

	buttons := [][]TgTypes.InlineKeyboardButtonType{
		{Globals.AudioButton, Globals.VideoButton},
		{Globals.YtLinkButton, Globals.ExitButton},
	}

	buttonMessage, err := Functions.SendButtonImage(result.Thumb, result.Title, chatId, messageId, TgTypes.InlineKeyboardMarkupType{InlineKeyboard: buttons}, true)
	if err != nil {
		return nil, err
	}

	Globals.AudioLinks[buttonMessage.MessageId] = result.AudioLink
	Globals.VideoLinks[buttonMessage.MessageId] = result.VideoLink

	return buttonMessage, nil
}
