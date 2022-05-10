package Globals

import (
	"encoding/json"
	"io/ioutil"
)

type VioDataType struct {
	Data []VioFeatureType `json:"data"`
}

type VioFeatureType struct {
	Name         string            `json:"name"`
	Title        string            `json:"title"`
	Path         string            `json:"path"`
	HostnamePath string            `json:"hostname_path"`
	Total        int               `json:"total"`
	Stacks       []VioEndPointType `json:"stacks"`
}

type VioEndPointType struct {
	Name         string      `json:"name"`
	Title        string      `json:"title"`
	Path         string      `json:"path"`
	Methods      []string    `json:"methods"`
	Params       interface{} `json:"params"`
	HostnamePath string      `json:"hostname_path"`
}

func ParseVioApi() error {
	body, err := ioutil.ReadFile("Data/VioApi.json")
	if err != nil {
		return err
	}

	data := VioDataType{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	for _, feature := range data.Data {
		VioAPIData[feature.Name] = feature
	}

	return nil
}
