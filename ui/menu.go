package ui

import (
	"fmt"
	"os"

	"github.com/yourname/sd-auto/common"
	"github.com/yourname/sd-auto/generator"
)

/*
解説:
	メニュー表示とルーティングを担当
	Run()がエントリーポイント
	各モードの処理はmodes.goに委譲
*/

// Run はUIのメインループを実行
func Run(allData map[string][]common.PromptItem) {
	gen := generator.New(allData)

	for {
		showMainMenu(gen, allData)
	}
}

// showMainMenu はメインメニューを表示
func showMainMenu(gen *generator.Generator, allData map[string][]common.PromptItem) {
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("🎨 プロンプト生成モード選択")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("1. 完全ランダム生成")
	fmt.Println("2. キャラ固定生成（他はランダム）")
	fmt.Println("3. 詳細設定生成（複数要素固定）")
	fmt.Println("4. データ一覧表示")
	fmt.Println("0. 終了")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Print("選択 >> ")

	choice := ReadInt()

	switch choice {
	case 1:
		CompletelyRandomMode(gen)
	case 2:
		CharacterFixedMode(gen)
	case 3:
		AdvancedFixedMode(gen, allData)
	case 4:
		ShowAllData(allData)
	case 0:
		fmt.Println("👋 終了します")
		os.Exit(0)
	default:
		fmt.Println("❌ 無効な選択です\n")
	}
}

// ShowAllData はデータ一覧を表示
func ShowAllData(allData map[string][]common.PromptItem) {
	fmt.Println("\n📋 データ一覧")

	for category, items := range allData {
		fmt.Printf("\n【%s】(%d件)\n", category, len(items))
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━")
		for i, item := range items {
			displayName := item.Ja
			if displayName == "" {
				displayName = item.En
			}
			fmt.Printf("  %2d. %s (%s)\n", i, displayName, item.Ja)
		}
		fmt.Println()
	}
}
