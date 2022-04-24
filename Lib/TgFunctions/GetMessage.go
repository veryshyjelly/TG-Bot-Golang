package Functions

import (
	"Telegram-Bot/Lib/TgTypes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type UpdateResult struct {
	Ok     bool                 `json:"ok"`
	Result []TgTypes.UpdateType `json:"result"`
}

func GetMessage(baseUrl string, offset, limit int64) []TgTypes.UpdateType {
	resp, err := http.Get(baseUrl + "/getUpdates" + "?offset=" + fmt.Sprint(offset) + "&limit=" + fmt.Sprint(limit))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	data := UpdateResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}

	if !data.Ok {
		return nil
	}

	return data.Result
}
