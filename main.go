package main

import (
	"fmt"

	"github.com/yourname/sd-auto/common"
	"github.com/yourname/sd-auto/internal/prompt"
)

// プログラム全体で使うデータ
var allData map[string][]common.PromptItem

func main() {

	// srcフォルダからデータを読み込む
	var err error
	allData, err = prompt.LoadAll("src")
	if err != nil {
		fmt.Println("❌ データ読み込みエラー:", err)
		return
	}

	// fmt.Println(allData)
	fmt.Println("📦 読み込んだカテゴリ:")
	for category := range allData {
		fmt.Printf("  - %s (%d件)\n", category, len(allData[category]))
	}
	fmt.Println()

	// // 読み込み順指定
	// fileOrder := []string{
	// 	"src/base.txt",
	// 	"src/"
	// }
	// // 呼び出し
	// data, err := prompt.Load("src/action(行為).txt")

	// if err != nil {
	// 	fmt.Println("Error!", err)
	// 	return
	// }

	// promptString := utils.ProcessingOnlyEn(data)

	// // fmt.Println(data)
}
