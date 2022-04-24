package features

import (
	Functions "Telegram-Bot/Lib/TgFunctions"
	"Telegram-Bot/Lib/TgTypes"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sunshineplan/imgconv"
	"image"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func ResizeImage(data io.Reader) (*bytes.Buffer, error) {

	myImage, err := imgconv.Decode(data)
	//myImage, _, err := image.Decode(data)
	if err != nil {
		return nil, err
	}
	im, _, err := imgconv.DecodeConfig(data)

	var marks image.Image
	if im.Width > im.Height {
		marks = imgconv.Resize(myImage, imgconv.ResizeOption{Width: 512})
	} else {
		marks = imgconv.Resize(myImage, imgconv.ResizeOption{Height: 512})
	}

	res := new(bytes.Buffer)

	err = imgconv.Write(res, marks, imgconv.FormatOption{Format: imgconv.PNG})

	return res, nil
}

func SendResizeImage(baseUrl, apiToken string, message *TgTypes.MessageType) *TgTypes.MessageType {

	imagePath := Functions.GetFile(baseUrl, message.Document.FileId).FilePath
	imageLink := "https://api.telegram.org/file/bot" + apiToken + "/" + url.QueryEscape(imagePath)

	//fmt.Println(vioBase + "?" + vioApi + "&" + author + "&" + pack + "&" + imageLink)
	res, err := http.Get(imageLink)
	if err != nil {
		log.Fatalln(err)
	}
	data := new(bytes.Buffer)
	//data, err = ioutil.ReadAll(res.Body)
	xyz := Functions.SendTextMessage(baseUrl, "Download in process...", message.Chat.Id, message.MessageId)
	convChannel := make(chan bool)
	go func() {
		_, err = io.Copy(data, res.Body)
		convChannel <- true
	}()
	if <-convChannel {
		go Functions.DeleteMessage(baseUrl, xyz.Chat.Id, xyz.MessageId)
		xyz = Functions.SendTextMessage(baseUrl, "Conversion in process...", message.Chat.Id, message.MessageId)
	}
	if err != nil {
		log.Fatalln(err)
	}

	resImage, err := ResizeImage(data)
	go Functions.DeleteMessage(baseUrl, xyz.Chat.Id, xyz.MessageId)

	if err != nil {
		go Functions.SendTextMessage(baseUrl, fmt.Sprint(err), message.Chat.Id, message.MessageId)
		return nil
	}

	xyz = Functions.SendTextMessage(baseUrl, "Conversion done...", message.Chat.Id, message.MessageId)
	go Functions.DelayDelete(baseUrl, 10, xyz.MessageId, xyz.Chat.Id)

	client := &http.Client{Timeout: time.Minute * 10}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	sendQuery := make(map[string]interface{})
	sendQuery["chat_id"], sendQuery["reply_to_message_id"], sendQuery["caption"], sendQuery["protect_content"] = message.Chat.Id, message.MessageId, "here", false

	for k, v := range sendQuery {
		fw, err := writer.CreateFormField(k)
		_, err = io.Copy(fw, strings.NewReader(fmt.Sprint(v)))
		if err != nil {
			log.Fatalln(err)
		}
	}

	if err != nil {
		log.Fatalln(err)
	}
	fw, err := writer.CreateFormFile("document", "resizing.png")
	_, err = io.Copy(fw, resImage)
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

	returnData := Functions.SendPhotoResult{}
	err = json.Unmarshal(sendResult, &returnData)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return &returnData.Result
}
