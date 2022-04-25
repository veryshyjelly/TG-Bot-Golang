package Functions

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type DeleteMessageQuery struct {
	ChatId    int64 `json:"chat_id"`
	MessageId int64 `json:"message_id"`
}

type DeleteResult struct {
	Ok          bool   `json:"ok"`
	Result      bool   `json:"result"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

func DeleteMessage(baseUrl string, ChatId int64, messageId int64) (bool, error) {
	query, err := json.Marshal(DeleteMessageQuery{
		ChatId:    ChatId,
		MessageId: messageId,
	})
	if err != nil {
		return false, err
	}

	resp, err := http.Post(baseUrl+"/deleteMessage", "application/json", bytes.NewBuffer(query))
	if err != nil {
		return false, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	data := DeleteResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return false, err
	}

	if !data.Ok {
		return false, errors.New(data.Description)
	}

	return data.Result, nil
}

func DelayDelete(baseUrl string, delay int, messageId, chatId int64) (bool, error) {
	time.Sleep(time.Second * time.Duration(delay))
	//fmt.Println("Deleted")
	return DeleteMessage(baseUrl, chatId, messageId)
}
