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
	"strings"
	"time"
)

type UploadStickerResult struct {
	Ok     bool             `json:"ok"`
	Result TgTypes.FileType `json:"result"`
}

func UploadStickerFile(baseUrl string, userId int64, file *bytes.Buffer) *TgTypes.FileType {

	client := &http.Client{Timeout: time.Minute * 10}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	sendQuery := make(map[string]interface{})
	sendQuery["user_id"] = userId

	for k, v := range sendQuery {
		fw, err := writer.CreateFormField(k)
		_, err = io.Copy(fw, strings.NewReader(fmt.Sprint(v)))
		if err != nil {
			log.Fatalln(err)
		}
	}

	fw, err := writer.CreateFormFile("png_sticker", "upsticker.png")
	_, err = io.Copy(fw, file)
	if err != nil {
		log.Fatalln(err)
	}
	writer.Close()
	req, err := http.NewRequest("POST", baseUrl+"/uploadStickerFile", bytes.NewReader(body.Bytes()))

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType()) // Very very important step
	rsp, _ := client.Do(req)
	sendResult, err := ioutil.ReadAll(rsp.Body)
	//fmt.Println(string(sendResult))

	returnData := UploadStickerResult{}
	err = json.Unmarshal(sendResult, &returnData)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return &returnData.Result
}
