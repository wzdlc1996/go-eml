package main

import (
	"fmt"
	"os"
)

func main() {
	emlFile, _ := os.ReadFile("./testEmls/图片.eml")
	rawmess, _ := ParseRaw(emlFile)
	fmt.Println(string(rawmess.Subject))
	for _, att := range rawmess.Attachments {
		path := fmt.Sprintf("testEmls/%s", att.Filename)
		_ = os.WriteFile(path, att.Body, 0666)
	}
}
