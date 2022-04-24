package Functions

import (
	"Telegram-Bot/Lib/TgTypes"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
)

type SendAudioResult struct {
	Ok     bool                `json:"ok"`
	Result TgTypes.MessageType `json:"result"`
}

func SendAudio(baseUrl, photoPath string, chatId, replyId int64, caption string, isProtected bool) *TgTypes.MessageType {
	client := &http.Client{Timeout: time.Minute * 15}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	file, err := os.Open(photoPath)
	if err != nil {
		log.Fatalln(err)
	}

	sendQuery := make(map[string]interface{})
	sendQuery["chat_id"], sendQuery["reply_to_message_id"], sendQuery["title"], sendQuery["caption"], sendQuery["protect_content"] = chatId, replyId, file.Name(), caption, isProtected

	for k, v := range sendQuery {
		fw, err := writer.CreateFormField(k)
		_, err = io.Copy(fw, strings.NewReader(fmt.Sprint(v)))
		if err != nil {
			log.Fatalln(err)
		}
	}

	fw, err := writer.CreateFormFile("audio", file.Name())
	_, err = io.Copy(fw, file)
	if err != nil {
		log.Fatalln(err)
	}
	writer.Close()
	file.Close()
	req, err := http.NewRequest("POST", baseUrl+"/sendAudio", bytes.NewReader(body.Bytes()))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType()) // Very very important step
	rsp, _ := client.Do(req)
	sendResult, err := ioutil.ReadAll(rsp.Body)
	//fmt.Println(string(sendResult))

	data := SendAudioResult{}
	err = json.Unmarshal(sendResult, &data)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return &data.Result
}

func SendAudioByUrl(baseUrl, AudioUrl string, chatId, replyId int64, caption string, isProtected bool) *TgTypes.MessageType {
	sendQuery := new(TgTypes.SendAudioQuery)
	sendQuery.ChatId, sendQuery.Audio, sendQuery.ReplyToMessageId, sendQuery.Caption, sendQuery.ProtectContent = chatId, AudioUrl, replyId, caption, isProtected
	query, err := json.Marshal(sendQuery)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post(baseUrl+"/sendAudio", "application/json", bytes.NewBuffer(query))
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	data := SendVideoResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}
	return &data.Result
}
