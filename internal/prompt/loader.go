package prompt

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"

	"github.com/yourname/sd-auto/common"
)

// 単一ファイルを読み込む関数
func Load(filename string) ([]common.PromptItem, error) {
	var items []common.PromptItem

	// ファイルを読み込む（バイト列で取得）
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// JSONをGo構造体に変換
	err = json.Unmarshal(file, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

// srcフォルダ内の全txtファイルを読み込む関数
func LoadAll(dirPath string) (map[string][]common.PromptItem, error) {
	// map[カテゴリ名]データ のマップを作る
	allData := make(map[string][]common.PromptItem)

	// src/*.txt にマッチするファイルを全て取得
	files, err := filepath.Glob(filepath.Join(dirPath, "*.txt"))
	if err != nil {
		return nil, err
	}

	// ファイル名順にソート（順番を一定にする）
	sort.Strings(files)

	// 各ファイルを読み込む
	for _, file := range files {
		data, err := Load(file)
		if err != nil {
			return nil, err
		}

		// ファイル名からカテゴリ名を抽出
		// 例: "src/character(キャラ).txt" → "character"
		baseName := filepath.Base(file)            // "character(キャラ).txt"
		categoryName := baseName[:len(baseName)-4] // ".txt"を削除
		// さらに括弧部分を削除する場合
		if idx := len(categoryName); idx > 0 {
			// 簡易版: そのまま使う
			allData[baseName] = data
		}
	}

	return allData, nil

}
