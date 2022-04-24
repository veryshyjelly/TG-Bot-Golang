package Functions

import (
	"Telegram-Bot/Lib/TgTypes"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type AnswerCallbackResult struct {
	Ok     bool `json:"ok"`
	Result bool `json:"result"`
}

func AnswerCallbackQuery(baseUrl, queryId, text string, showAlert bool) bool {
	query, err := json.Marshal(TgTypes.AnswerCallbackQuery{
		CallbackQueryId: queryId,
		Text:            text,
		ShowAlert:       showAlert,
	})
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post(baseUrl+"/answerCallbackQuery", "application/json", bytes.NewBuffer(query))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	data := AnswerCallbackResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}

	if !data.Ok {
		return false
	}

	return data.Result
}
