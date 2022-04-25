package Functions

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

type SendDocumentQuery struct {
	ChatId                      int64                       `json:"chat_id"`
	Document                    string                      `json:"document"`        // Or multipart file
	Thumb                       string                      `json:"thumb,omitempty"` // Or multipart file
	Caption                     string                      `json:"caption,omitempty"`
	ParseMode                   string                      `json:"parse_mode,omitempty"`
	CaptionEntities             []TgTypes.MessageEntityType `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool                        `json:"disable_content_type_detection,omitempty"`
	DisableNotification         bool                        `json:"disable_notification,omitempty"`
	ProtectContent              bool                        `json:"protect_content,omitempty"`
	ReplyToMessageId            int64                       `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply    bool                        `json:"allow_sending_without_reply,omitempty"`
	//ReplyMarkup                 InlineKeyboardMarkupType `json:"reply_markup,omitempty"`
}

type SendDocumentResult struct {
	Ok          bool                `json:"ok"`
	Result      TgTypes.MessageType `json:"result"`
	ErrorCode   int                 `json:"error_code"`
	Description string              `json:"description"`
}

func SendDocumentByUrl(baseUrl, documentUrl string, chatId, replyId int64, caption string, isProtected bool) (*TgTypes.MessageType, error) {
	query, err := json.Marshal(SendDocumentQuery{
		ChatId:           chatId,
		Document:         documentUrl,
		Thumb:            "",
		Caption:          caption,
		ParseMode:        "HTML",
		ProtectContent:   isProtected,
		ReplyToMessageId: replyId,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(baseUrl+"/sendDocument", "application/json", bytes.NewBuffer(query))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := SendDocumentResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if !data.Ok {
		return nil, errors.New(data.Description)
	}

	return &data.Result, nil
}

func SendDocument(baseUrl, documentPath string, chatId, replyId int64, caption string, isProtected bool) (*TgTypes.MessageType, error) {
	client := &http.Client{Timeout: time.Minute * 20}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	file, err := os.Open(documentPath)
	if err != nil {
		return nil, err
	}

	sendQuery := make(map[string]interface{})
	sendQuery["chat_id"], sendQuery["reply_to_message_id"], sendQuery["title"], sendQuery["caption"], sendQuery["protect_content"] = chatId, replyId, file.Name(), caption, isProtected

	for k, v := range sendQuery {
		fw, err := writer.CreateFormField(k)
		_, err = io.Copy(fw, strings.NewReader(fmt.Sprint(v)))
		if err != nil {
			return nil, err
		}
	}

	fw, err := writer.CreateFormFile("document", file.Name())
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

	req, err := http.NewRequest("POST", baseUrl+"/sendDocument", bytes.NewReader(body.Bytes()))
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

	data := SendDocumentResult{}
	err = json.Unmarshal(sendResult, &data)
	if err != nil {
		return nil, err
	}

	if !data.Ok {
		return nil, errors.New(data.Description)
	}

	return &data.Result, err
}
