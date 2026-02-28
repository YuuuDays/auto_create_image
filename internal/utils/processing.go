package utils

import (
	"fmt"

	"github.com/yourname/sd-auto/common"
)

// 取り出したjson(構造体)を取り出し加工

func Processing(jsonStruct []common.PromptItem) {

	for i, number := range jsonStruct {
		fmt.Println(i, number)
	}
}
