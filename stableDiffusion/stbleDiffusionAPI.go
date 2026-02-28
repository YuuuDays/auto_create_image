package stablediffusion

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type SDResponse struct {
	Images []string `json:"images"`
}

func GenerateImage(prompt string, pickUpCharcterJP string) {

	// ====== ① フォルダ作成 ======
	outputDir := "./output"
	os.MkdirAll(outputDir, os.ModePerm)

	// ====== ② 次の連番を取得 ======
	nextNumber, err := getNextNumber(outputDir, pickUpCharcterJP)

	if err != nil {
		os.Exit(0)
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

	resp, err := http.Post(
		"http://127.0.0.1:7861/sdapi/v1/txt2img",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		fmt.Println("レスポンス帰ってきませんでした....")
		os.Exit(0)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result SDResponse
	json.Unmarshal(body, &result)

	fmt.Println("Status:", resp.Status)
	if len(result.Images) == 0 {
		fmt.Println("画像が帰ってきませんでした....")
		os.Exit(0)
	}

	imageBytes, _ := base64.StdEncoding.DecodeString(result.Images[0])

	// ====== ④ 保存 ======
	fileName := fmt.Sprintf("%s_%03d.png", pickUpCharcterJP, nextNumber)
	fullPath := filepath.Join(outputDir, fileName)

	err = os.WriteFile(fullPath, imageBytes, 0644)
	if err != nil {
		fmt.Println("保存できませんでした....")
		os.Exit(0)
	}

	fmt.Println("保存完了", fullPath)

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
