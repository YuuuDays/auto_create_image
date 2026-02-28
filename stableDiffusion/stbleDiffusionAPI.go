package stablediffusion

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type SDResponse struct {
	Images []string `json:"images"`
}

func GenerateImage(ctx context.Context, prompt string, pickUpCharcterJP string) {

	// ======  フォルダ作成 ======
	outputDir := "./output"
	os.MkdirAll(outputDir, os.ModePerm)

	// ====== 次の連番を取得 ======
	nextNumber, err := getNextNumber(outputDir, pickUpCharcterJP)

	if err != nil {
		return
	}

	// nagative_prompt ベタ打ちの為要変更検討
	// sizeは64の倍数
	payload := map[string]interface{}{
		"prompt":          prompt,
		"negative_prompt": "bad quality,worst quality,worst detail,sketch,censored, artist name, signature, watermark,patreon username, patreon logo",
		"width":           816,
		"height":          1024,
		"cfg_scale":       7,
		"seed":            -1, // -1でランダム
	}

	jsonData, _ := json.Marshal(payload)

	// ====== HTTPリクエスト作成 ======
	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		"http://127.0.0.1:7860/sdapi/v1/txt2img",
		bytes.NewBuffer(jsonData),
	)
	req.Header.Set("Content-Type", "application/json")

	// API呼び出しの開始時刻を記録
	startTime := time.Now()

	// ====== リクエスト送信 ======
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() != nil {
			fmt.Println("リクエストキャンセルされました")
			return
		}
		fmt.Println("レスポンス帰ってきませんでした....")
		return
	}
	defer resp.Body.Close()

	// ====== レスポンス読み込み ======
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("レスポンス読み込み失敗...")
		return
	}

	var result SDResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("JSONパース失敗...")
		return
	}

	if len(result.Images) == 0 {
		fmt.Println("画像が帰ってきませんでした....")
		return
	}

	// ====== Base64デコード ======
	imageBytes, err := base64.StdEncoding.DecodeString(result.Images[0])
	if err != nil {
		fmt.Println("Base64デコード失敗...")
		return
	}

	// ====== 保存 ======
	fileName := fmt.Sprintf("%s_%03d.png", pickUpCharcterJP, nextNumber)
	fullPath := filepath.Join(outputDir, fileName)

	err = os.WriteFile(fullPath, imageBytes, 0644)
	if err != nil {
		fmt.Println("保存できませんでした....")
		return
	}
	// 経過時間を計算
	elapsed := time.Since(startTime)

	// 方法1: 分と秒に分解して表示
	minutes := int(elapsed.Minutes())
	seconds := int(elapsed.Seconds()) % 60

	fmt.Printf("保存完了 (生成時間: %d分%d秒)\n", minutes, seconds)

}

// =====================================
// 次の連番を取得する関数
// =====================================
func getNextNumber(dir, characterName string) (int, error) {

	files, err := os.ReadDir(dir)
	if err != nil {
		return 1, nil
	}

	maxNumber := 0

	for _, file := range files {
		name := file.Name()

		// 例: Alice_003.png
		if strings.HasPrefix(name, characterName+"_") && strings.HasSuffix(name, ".png") {

			base := strings.TrimSuffix(name, ".png")
			parts := strings.Split(base, "_")

			if len(parts) < 2 {
				continue
			}

			num, err := strconv.Atoi(parts[len(parts)-1])
			if err != nil {
				continue
			}

			if num > maxNumber {
				maxNumber = num
			}
		}
	}

	return maxNumber + 1, nil
}
