package main

import (
	"fmt"
	"os"

	"github.com/yourname/sd-auto/common"
	"github.com/yourname/sd-auto/config"
	"github.com/yourname/sd-auto/prompt"
	"github.com/yourname/sd-auto/ui"
)

// プログラム全体で使うデータ
var allData map[string][]common.PromptItem

/*
	```

**解説:**
- 各モードの実装を独立した関数に
- `generator`パッケージを使って生成
- UIロジックとビジネスロジックを分離

---

## ベストプラクティスのポイント

### 1. **単一責任の原則**
```
main.go          → エントリーポイント
generator/       → プロンプト生成ロジック
ui/              → ユーザーインターフェース
prompt/          → データ読み込み
*/
func main() {

	// 設定ファイル読み込み
	cfg, err := config.Load("config/order.json")
	if err != nil {
		fmt.Println("❌ 設定ファイル読み込みエラー:", err)
		os.Exit(1)
	}

	// srcフォルダからデータを読み込む
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

	// UI開始
	ui.Run(allData, cfg.PromptOrder)

}

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
