package generator

import (
	"math/rand"
	"strings"

	"github.com/yourname/sd-auto/common"
)

// Generator はプロンプト生成を担当
type Generator struct {
	AllData map[string][]common.PromptItem
}

// New は新しいGeneratorを作成
func New(data map[string][]common.PromptItem) *Generator {
	return &Generator{AllData: data}
}

// GenerateRandom は完全ランダムなプロンプトを生成
func (g *Generator) GenerateRandom() string {
	var parts []string

	for _, items := range g.AllData {
		if len(items) == 0 {
			continue
		}
		randomItem := items[rand.Intn(len(items))]
		parts = append(parts, randomItem.En)
	}

	return strings.Join(parts, ", ")
}

// GenerateWithFixed は固定要素ありでプロンプトを生成
func (g *Generator) GenerateWithFixed(fixed map[string]string) string {
	var parts []string

	for category, items := range g.AllData {
		if len(items) == 0 {
			continue
		}

		// 固定されている場合はそれを使う
		if fixedValue, ok := fixed[category]; ok {
			parts = append(parts, fixedValue)
		} else {
			// 固定されていない場合はランダム
			randomItem := items[rand.Intn(len(items))]
			parts = append(parts, randomItem.En)
		}
	}

	return strings.Join(parts, ", ")
}

// FindCharacterCategory はキャラクターカテゴリを探す
func (g *Generator) FindCharacterCategory() (string, []common.PromptItem) {
	for category, items := range g.AllData {
		if strings.Contains(strings.ToLower(category), "character") ||
			strings.Contains(category, "キャラ") {
			return category, items
		}
	}
	return "", nil
}
