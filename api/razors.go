package api

import (
	_ "embed"
	"encoding/json"
)

//go:embed razors.json
var razorsJson []byte

type Razor struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Aka     string `json:"aka,omitempty"`
	Summary string `json:"summary"`
}

func getRazor(id int) (*Razor, error) {
	razors, err := loadRazors()
	if err != nil {
		return nil, err
	}

	var result *Razor
	for i := range razors {
		if razors[i].ID == id {
			result = &razors[i]
			break
		}
	}

	return result, nil
}

func loadRazors() ([]Razor, error) {
	var result []Razor
	if err := json.Unmarshal(razorsJson, &result); err != nil {
		return nil, err
	}

	return result, nil
}
