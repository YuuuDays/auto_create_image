package prompt

import (
	"encoding/json"
	"os"

	"github.com/yourname/sd-auto/common"
)

func Load(filename string) ([]common.PromptItem, error) {
	var items []common.PromptItem

	file, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	// この時点ではfileはバイナリなのでGoの構造体へ置換
	err = json.Unmarshal(file, &items)
	if err != nil {
		return nil, err
	}

	// fmt.Println(items)
	// fmt.Println(reflect.TypeOf(items))
	return items, nil

}
