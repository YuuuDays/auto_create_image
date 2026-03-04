package common

// レスポンスのシード値,画像名,時間を記録する用

type SDResponse struct {
	Images     []string               `json:"images"`
	Parameters map[string]interface{} `json:"parameters"`
	Info       string                 `json:"info"`
}

type InfoData struct {
	Seed int64 `json:"seed"`
}

type ImageRecord struct {
	Seed     int64  `json:"seed"`
	Prompt   string `json:"prompt"`
	File     string `json:"file"`
	Time     string `json:"time"`
	Duration int    `json:"duration"` // 秒
}
