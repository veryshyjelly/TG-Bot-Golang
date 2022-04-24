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

type SendVideoResult struct {
	Ok     bool                `json:"ok"`
	Result TgTypes.MessageType `json:"result"`
}

func SendVideo(baseUrl, videoPath string, chatId, replyId int64, caption string, isProtected bool) *TgTypes.MessageType {
	client := &http.Client{Timeout: time.Minute * 15}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	file, err := os.Open(videoPath)
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

	fw, err := writer.CreateFormFile("video", file.Name())
	_, err = io.Copy(fw, file)
	if err != nil {
		log.Fatalln(err)
	}
	writer.Close()
	file.Close()
	req, err := http.NewRequest("POST", baseUrl+"/sendVideo", bytes.NewReader(body.Bytes()))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType()) // Very very important step
	rsp, _ := client.Do(req)
	if rsp == nil {
		log.Fatalln("Rsp is nil")
	}
	sendResult, err := ioutil.ReadAll(rsp.Body)
	//fmt.Println(string(sendResult))

	data := SendVideoResult{}
	err = json.Unmarshal(sendResult, &data)
	if err != nil {
		log.Fatalln(err)
	}

	return &data.Result
	//return TgTypes.MessageType{}
}

func SendVideoByUrl(baseUrl, VideoUrl string, chatId, replyId int64, caption string, isProtected bool) *TgTypes.MessageType {
	sendQuery := new(TgTypes.SendVideoQuery)
	sendQuery.ChatId, sendQuery.Video, sendQuery.ReplyToMessageId, sendQuery.Caption, sendQuery.ProtectContent = chatId, VideoUrl, replyId, caption, isProtected
	query, err := json.Marshal(sendQuery)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post(baseUrl+"/sendVideo", "application/json", bytes.NewBuffer(query))
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
