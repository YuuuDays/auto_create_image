package utils

import (
	"fmt"
	"math/rand"

	"github.com/yourname/sd-auto/common"
)

// 取り出したjson(構造体)を取り出し加工

// En番のランダム一つを返す
func ProcessingOnlyEn(jsonStruct []common.PromptItem) string {

	// for _, jsonString := range jsonStruct {
	// 	// fmt.Println(i, number)
	// 	// fmt.Println(reflect.TypeOf(number))
	// 	fmt.Println(jsonString.En)
	// 	fmt.Println(reflect.TypeOf(number.En))
	// }
	if len(jsonStruct) == 0 {
		fmt.Println("要素0でした")
		return ""
	}
	item := jsonStruct[rand.Intn(len(jsonStruct))]
	fmt.Println("お試し", item)
	fmt.Println("お試し", item.En)
	// ここでjsonStringの中身をランダムで一つ選んで返す

	return item.En
}
