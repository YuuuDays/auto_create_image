package prompt

import (
	"encoding/json"
	"os"
)

type PromptItem struct {
	En string `json:"en"`
	Ja string `json:"ja"`
}

func Load(filename string) ([]PromptItem, error) {
	var items []PromptItem

	file, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &items)
	if err != nil {
		return nil, err
	}

	return items, nil

}
