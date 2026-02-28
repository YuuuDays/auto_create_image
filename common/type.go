package common

// txtからjson配列に変える際のstruct構造
type PromptItem struct {
	En string `json:"en"`
	Ja string `json:"ja"`
}
