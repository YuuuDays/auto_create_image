package ui

import (
	"fmt"

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
func Run(allData map[string][]common.PromptItem, order []string) ([]string, string) {
	gen := generator.New(allData, order)

	return showMainMenu(gen, allData)

}

// showMainMenu はメインメニューを表示
func showMainMenu(gen *generator.Generator, allData map[string][]common.PromptItem) ([]string, string) {
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("🎨 プロンプト生成モード選択")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("1. 完全ランダム生成")
	fmt.Println("2. キャラ固定生成（他はランダム）")
	// fmt.Println("3. 詳細設定生成（複数要素固定）")
	// fmt.Println("4. データ一覧表示")
	fmt.Println("0. 終了")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Print("選択 >> ")

	// choice := ReadInt()

	var prompts []string      // 生成されたプロンプトを保持
	var pickupCharcter string //選択したキャラ名(JP)

	// キャラ固定生成をデフォに
	prompts, pickupCharcter = CharacterFixedMode(gen)
	return prompts, pickupCharcter
}
