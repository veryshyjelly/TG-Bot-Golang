package Functions

import (
	"Telegram-Bot/Lib/TgTypes"
	"Telegram-Bot/Settings"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var (
	Default           = "default"
	AllPrivate        = "all_private_chats"
	AllGroup          = "all_group_chats"
	AllAdministrators = "all_chat_administrators"
)

type CommandQuery struct {
	Commands     []TgTypes.BotCommandType `json:"commands"`
	scope        string                   `json:"scope,omitempty"`
	LanguageCode string                   `json:"language_code,omitempty"`
}

type SetCommandsResult struct {
	Ok          bool   `json:"ok"`
	Result      bool   `json:"result"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

func SetMyCommands(commands []TgTypes.BotCommandType, scope string) (bool, error) {
	query, err := json.Marshal(CommandQuery{
		Commands: commands,
		scope:    scope,
	})
	if err != nil {
		return false, err
	}

	resp, err := http.Post(Settings.BaseUrl+"/setMyCommands", "application/json", bytes.NewBuffer(query))
	if err != nil {
		return false, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	data := SetCommandsResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return false, err
	}

	if !data.Ok {
		return false, errors.New(data.Description)
	}

	return data.Result, nil
}

func DeleteMyCommands(scope string) (bool, error) {
	query, err := json.Marshal(CommandQuery{
		scope: scope,
	})
	if err != nil {
		return false, err
	}

	resp, err := http.Post(Settings.BaseUrl+"/deleteMyCommands", "application/json", bytes.NewBuffer(query))
	if err != nil {
		return false, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	data := SetCommandsResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return false, err
	}

	if !data.Ok {
		return false, errors.New(data.Description)
	}

	return data.Result, nil

}
