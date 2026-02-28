package ui

import (
	"fmt"
	"strings"

	"github.com/yourname/sd-auto/common"
	"github.com/yourname/sd-auto/generator"
)

// CompletelyRandomMode は完全ランダム生成モード
func CompletelyRandomMode(gen *generator.Generator) {
	fmt.Println("\n🎲 完全ランダム生成モード")
	fmt.Print("何個生成しますか？ >> ")
	count := ReadInt()

	fmt.Println("\n生成結果:")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━")
	for i := 0; i < count; i++ {
		prompt := gen.GenerateRandom()
		fmt.Printf("%d. %s\n", i+1, prompt)
	}
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
}

// CharacterFixedMode はキャラクター固定生成モード
func CharacterFixedMode(gen *generator.Generator) {
	fmt.Println("\n👤 キャラクター固定生成モード")

	// キャラクターカテゴリを探す
	characterCategory, characterItems := gen.FindCharacterCategory()

	if len(characterItems) == 0 {
		fmt.Println("❌ キャラクターデータが見つかりません")
		return
	}

	// キャラクター一覧を日本語で表示
	fmt.Println("\n📋 キャラクター一覧:")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━")
	for i, item := range characterItems {
		displayName := item.Ja
		if displayName == "" {
			displayName = item.Ja
		}
		fmt.Printf("  %2d. %s\n", i, displayName)
	}
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━")

	// キャラクター選択
	fmt.Print("\nキャラクター番号を選択 (-1でランダム) >> ")
	charIdx := ReadInt()

	var fixedCharacter string
	if charIdx >= 0 && charIdx < len(characterItems) {
		fixedCharacter = characterItems[charIdx].Ja
		displayName := characterItems[charIdx].Ja
		if displayName == "" {
			displayName = fixedCharacter
		}
		fmt.Printf("✅ 「%s」に固定しました\n", displayName)
	} else {
		fmt.Println("✅ キャラクターもランダムにします")
	}

	// 生成数を聞く
	fmt.Print("\n何個生成しますか？ >> ")
	count := ReadInt()

	// 固定要素マップを作成
	fixedElements := make(map[string]string)
	if fixedCharacter != "" {
		fixedElements[characterCategory] = fixedCharacter
	}

	// 生成
	fmt.Println("\n生成結果:")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━")
	for i := 0; i < count; i++ {
		prompt := gen.GenerateWithFixed(fixedElements)
		fmt.Printf("%d. %s\n", i+1, prompt)
	}
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
}

// AdvancedFixedMode は詳細設定生成モード
func AdvancedFixedMode(gen *generator.Generator, allData map[string][]common.PromptItem) {
	fmt.Println("\n🔧 詳細設定生成モード")

	// 固定する要素を選択
	fixedElements := make(map[string]string)

	fmt.Println("\n固定したい要素を選んでください")
	fmt.Println("（何も固定しない場合は -1 を入力）")

	// カテゴリごとに固定するか聞く
	for category, items := range allData {
		fmt.Printf("\n📌 %s を固定しますか？ (y/n/skip) >> ", category)
		answer := ReadString()

		if strings.ToLower(answer) == "y" || strings.ToLower(answer) == "yes" {
			// このカテゴリの項目を日本語で表示
			fmt.Printf("\n【%s】選択肢:\n", category)
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━")
			for i, item := range items {
				displayName := item.Ja
				if displayName == "" {
					displayName = item.En
				}
				fmt.Printf("  %2d. %s\n", i, displayName)
			}
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━")

			fmt.Print("番号を選択 (-1でスキップ) >> ")
			idx := ReadInt()

			if idx >= 0 && idx < len(items) {
				fixedElements[category] = items[idx].En
				displayName := items[idx].Ja
				if displayName == "" {
					displayName = items[idx].En
				}
				fmt.Printf("✅ 「%s」に固定しました\n", displayName)
			}
		}
	}

	// 生成数を聞く
	fmt.Print("\n何個生成しますか？ >> ")
	count := ReadInt()

	// 生成
	fmt.Println("\n生成結果:")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━")
	for i := 0; i < count; i++ {
		prompt := gen.GenerateWithFixed(fixedElements)
		fmt.Printf("%d. %s\n", i+1, prompt)
	}
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━\n")
}
