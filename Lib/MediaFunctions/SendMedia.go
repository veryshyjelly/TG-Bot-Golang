package MediaFunctions

import (
	"Telegram-Bot/Lib/TgTypes"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
)

type MediaType string

const (
	Video     MediaType = "/sendVideo"
	Audio     MediaType = "/sendAudio"
	Photo     MediaType = "/sendPhoto"
	Document  MediaType = "/sendDocument"
	Animation MediaType = "/sendAnimation"
	Sticker   MediaType = "/sendSticker"
)

type SendMediaResult struct {
	Ok          bool                `json:"ok"`
	Result      TgTypes.MessageType `json:"result"`
	ErrorCode   int                 `json:"error_code"`
	Description string              `json:"description"`
}

func SendMediaByUrl(baseUrl, Url string, mediaType MediaType, chatId, replyId int64, caption string, isProtected bool) (*TgTypes.MessageType, error) {
	sendQuery := map[string]interface{}{
		"chat_id":                              chatId,
		"caption":                              caption,
		"parse_mode":                           "HTML",
		"protect_content":                      isProtected,
		strings.ToLower(string(mediaType[5:])): Url,
		"reply_to_message_id":                  replyId,
	}

	query, err := json.Marshal(sendQuery)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(baseUrl+string(mediaType), "application/json", bytes.NewBuffer(query))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := SendMediaResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if !data.Ok {
		return nil, errors.New(data.Description)
	}

	return &data.Result, nil
}

func SendMedia(baseUrl, mediaPath string, mediaType MediaType, chatId, replyId int64, caption string, isProtected bool) (*TgTypes.MessageType, error) {
	client := &http.Client{Timeout: time.Minute * 20}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	file, err := os.Open(mediaPath)
	if err != nil {
		return nil, err
	}

	sendQuery := map[string]interface{}{
		"chat_id":             chatId,
		"caption":             caption,
		"parse_mode":          "HTML",
		"protect_content":     isProtected,
		"reply_to_message_id": replyId,
	}

	for k, v := range sendQuery {
		fw, err := writer.CreateFormField(k)
		_, err = io.Copy(fw, strings.NewReader(fmt.Sprint(v)))
		if err != nil {
			return nil, err
		}
	}

	fw, err := writer.CreateFormFile(strings.ToLower(string(mediaType[5:])), file.Name())
	_, err = io.Copy(fw, file)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", baseUrl+string(mediaType), bytes.NewReader(body.Bytes()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType()) // Very very important step
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	sendResult, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	data := SendMediaResult{}
	err = json.Unmarshal(sendResult, &data)
	if err != nil {
		return nil, err
	}

	if !data.Ok {
		return nil, errors.New(data.Description)
	}

	return &data.Result, err
}
