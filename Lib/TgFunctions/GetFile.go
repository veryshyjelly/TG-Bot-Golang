package Functions

import (
	"Telegram-Bot/Lib/TgTypes"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type GetFileQuery struct {
	FileId string `json:"file_id"`
}

type GetFileResult struct {
	Ok          bool             `json:"ok"`
	Result      TgTypes.FileType `json:"result"`
	ErrorCode   int              `json:"error_code"`
	Description string           `json:"description"`
}

func GetFile(baseUrl, fileId string) (*TgTypes.FileType, error) {
	query, err := json.Marshal(GetFileQuery{
		FileId: fileId,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(baseUrl+"/getFile", "application/json", bytes.NewBuffer(query))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := GetFileResult{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if !data.Ok {
		return nil, errors.New(data.Description)
	}

	return &data.Result, nil
}
