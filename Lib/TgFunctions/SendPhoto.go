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

type SendPhotoResult struct {
	Ok     bool                `json:"ok"`
	Result TgTypes.MessageType `json:"result"`
}

func SendPhoto(baseUrl, photoPath string, chatId, replyId int64, caption string, isProtected bool) *TgTypes.MessageType {
	client := &http.Client{Timeout: time.Minute * 10}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	sendQuery := make(map[string]interface{})
	sendQuery["chat_id"], sendQuery["reply_to_message_id"], sendQuery["caption"], sendQuery["protect_content"] = chatId, replyId, caption, isProtected

	for k, v := range sendQuery {
		fw, err := writer.CreateFormField(k)
		_, err = io.Copy(fw, strings.NewReader(fmt.Sprint(v)))
		if err != nil {
			log.Fatalln(err)
		}
	}

	file, err := os.Open(photoPath)
	if err != nil {
		log.Fatalln(err)
	}
	fw, err := writer.CreateFormFile("photo", file.Name())
	_, err = io.Copy(fw, file)
	if err != nil {
		log.Fatalln(err)
	}
	writer.Close()
	file.Close()
	req, err := http.NewRequest("POST", baseUrl+"/sendPhoto", bytes.NewReader(body.Bytes()))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType()) // Very very important step
	rsp, _ := client.Do(req)
	sendResult, err := ioutil.ReadAll(rsp.Body)
	//fmt.Println(string(sendResult))

	data := SendPhotoResult{}
	err = json.Unmarshal(sendResult, &data)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return &data.Result
}

func SendPhotoByUrl(baseUrl, photoUrl string, chatId, replyId int64, caption string, isProtected bool) *TgTypes.MessageType {
	sendQuery := new(TgTypes.SendPhotoQuery)
	sendQuery.ChatId, sendQuery.Photo, sendQuery.ReplyToMessageId, sendQuery.Caption, sendQuery.ProtectContent = chatId, photoUrl, replyId, caption, isProtected
	query, err := json.Marshal(sendQuery)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post(baseUrl+"/sendPhoto", "application/json", bytes.NewBuffer(query))
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(string(body))
	data := SendPhotoResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}
	return &data.Result
}

func SendPhotoByReader(baseUrl string, photoPath *bytes.Buffer, message *TgTypes.MessageType, caption string, isProtected bool) *TgTypes.MessageType {
	client := &http.Client{Timeout: time.Minute * 10}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	sendQuery := make(map[string]interface{})
	sendQuery["chat_id"], sendQuery["reply_to_message_id"], sendQuery["caption"], sendQuery["protect_content"] = message.Chat.Id, message.MessageId, caption, false

	for k, v := range sendQuery {
		fw, err := writer.CreateFormField(k)
		_, err = io.Copy(fw, strings.NewReader(fmt.Sprint(v)))
		if err != nil {
			log.Fatalln(err)
		}
	}

	fw, err := writer.CreateFormFile("document", "resizing.png")
	_, err = io.Copy(fw, photoPath)
	if err != nil {
		log.Fatalln(err)
	}
	writer.Close()
	req, err := http.NewRequest("POST", baseUrl+"/sendDocument", bytes.NewReader(body.Bytes()))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType()) // Very very important step
	rsp, _ := client.Do(req)
	sendResult, err := ioutil.ReadAll(rsp.Body)
	//fmt.Println(string(sendResult))

	returnData := SendPhotoResult{}
	err = json.Unmarshal(sendResult, &returnData)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return &returnData.Result
}
