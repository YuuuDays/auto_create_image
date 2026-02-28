package main

import (
	"fmt"

	"github.com/yourname/sd-auto/internal/prompt"
	"github.com/yourname/sd-auto/internal/utils"
)

func main() {

	// 呼び出し
	data, err := prompt.Load("src/action(行為).txt")

	if err != nil {
		fmt.Println("Error!", err)
		return
	}

	utils.Processing(data)

	// fmt.Println(data)
}
