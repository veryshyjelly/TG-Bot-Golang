package Functions

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type AnswerCallbackQueryType struct {
	CallbackQueryId string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	Url             string `json:"url,omitempty"`
}

type AnswerCallbackResult struct {
	Ok          bool   `json:"ok"`
	Result      bool   `json:"result"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

func AnswerCallbackQuery(baseUrl, queryId, text string, showAlert bool) (bool, error) {
	query, err := json.Marshal(AnswerCallbackQueryType{
		CallbackQueryId: queryId,
		Text:            text,
		ShowAlert:       showAlert,
	})
	if err != nil {
		return false, err
	}

	resp, err := http.Post(baseUrl+"/answerCallbackQuery", "application/json", bytes.NewBuffer(query))
	if err != nil {
		return false, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	data := AnswerCallbackResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return false, err
	}

	if !data.Ok {
		return false, errors.New(data.Description)
	}

	return data.Result, err
}
