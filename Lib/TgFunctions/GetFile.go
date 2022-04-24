package Functions

import (
	"Telegram-Bot/Lib/TgTypes"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type GetFileResult struct {
	Ok     bool             `json:"ok"`
	Result TgTypes.FileType `json:"result"`
}

func GetFile(baseUrl, fileId string) TgTypes.FileType {
	query, err := json.Marshal(TgTypes.GetFileQuery{FileId: fileId})
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(baseUrl+"/getFile", "application/json", bytes.NewBuffer(query))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	data := GetFileResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}

	if !data.Ok {
		return TgTypes.FileType{}
	}

	return data.Result
}
