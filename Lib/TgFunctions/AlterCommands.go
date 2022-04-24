package Functions

import (
	"Telegram-Bot/Lib/TgTypes"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
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

func SetMyCommands(baseUrl string, commands []TgTypes.BotCommandType, scope string) bool {
	sendQuery := new(CommandQuery)
	sendQuery.Commands, sendQuery.scope = commands, scope
	query, err := json.Marshal(sendQuery)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post(baseUrl+"/setMyCommands", "application/json", bytes.NewBuffer(query))
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	var result bool
	err = json.Unmarshal(body, &result)
	return result
}

func DeleteMyCommands(baseUrl, scope string) bool {
	sendQuery := new(CommandQuery)
	sendQuery.scope = scope
	query, err := json.Marshal(sendQuery)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post(baseUrl+"/deleteMyCommands", "application/json", bytes.NewBuffer(query))
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	var result bool
	err = json.Unmarshal(body, &result)
	return result
}
