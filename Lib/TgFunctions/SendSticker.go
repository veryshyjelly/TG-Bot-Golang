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

type SendStickerResult struct {
	Ok     bool                `json:"ok"`
	Result TgTypes.MessageType `json:"result"`
}

func SendSticker(baseUrl, documentPath string, chatId, replyId int64, isProtected bool) *TgTypes.MessageType {
	client := &http.Client{Timeout: time.Minute * 20}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	file, err := os.Open(documentPath)
	if err != nil {
		log.Fatalln(err)
	}

	sendQuery := make(map[string]interface{})
	sendQuery["chat_id"], sendQuery["reply_to_message_id"], sendQuery["title"], sendQuery["protect_content"] = chatId, replyId, file.Name(), isProtected

	for k, v := range sendQuery {
		fw, err := writer.CreateFormField(k)
		_, err = io.Copy(fw, strings.NewReader(fmt.Sprint(v)))
		if err != nil {
			log.Fatalln(err)
		}
	}

	fw, err := writer.CreateFormFile("sticker", file.Name())
	_, err = io.Copy(fw, file)
	if err != nil {
		log.Fatalln(err)
	}
	writer.Close()
	file.Close()
	req, err := http.NewRequest("POST", baseUrl+"/sendSticker", bytes.NewReader(body.Bytes()))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType()) // Very very important step
	rsp, _ := client.Do(req)
	sendResult, err := ioutil.ReadAll(rsp.Body)
	//fmt.Println(string(sendResult))

	data := SendStickerResult{}
	err = json.Unmarshal(sendResult, &data)
	if err != nil {
		log.Fatalln(err)
	}

	return &data.Result
}

func SendStickerByUrl(baseUrl, stickerUrl string, chatId, replyId int64, isProtected bool) *TgTypes.MessageType {
	sendQuery := new(TgTypes.SendStickerQuery)
	sendQuery.ChatId, sendQuery.Sticker, sendQuery.ReplyToMessageId, sendQuery.ProtectContent = chatId, stickerUrl, replyId, isProtected
	query, err := json.Marshal(sendQuery)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post(baseUrl+"/sendSticker", "application/json", bytes.NewBuffer(query))
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	data := SendStickerResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}
	return &data.Result
}
