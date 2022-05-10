package Functions

import (
	"Telegram-Bot/Lib/TgTypes"
	"Telegram-Bot/Settings"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UpdateResult struct {
	Ok          bool                 `json:"ok"`
	Result      []TgTypes.UpdateType `json:"result"`
	ErrorCode   int                  `json:"error_code"`
	Description string               `json:"description"`
}

func GetMessage(offset, limit int64) ([]TgTypes.UpdateType, error) {
	resp, err := http.Get(Settings.BaseUrl + "/getUpdates" + "?offset=" + fmt.Sprint(offset) + "&limit=" + fmt.Sprint(limit))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := UpdateResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if !data.Ok {
		return nil, errors.New(data.Description)
	}

	return data.Result, nil
}
