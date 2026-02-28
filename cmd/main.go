package main

import (
	"fmt"

	"github.com/yourname/sd-auto/internal/prompt"
)

func main() {
	data, err := prompt.Load("src/action(行為).txt")

	if err != nil {
		fmt.Println("Error!", err)
		return
	}

	fmt.Println(data)
}
