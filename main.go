package main

import (
	"fmt"
	"os"
)

func main() {
	emlFile, _ := os.ReadFile("./testEmls/图片.eml")
	rawmess, _ := ParseRaw(emlFile)
	fmt.Println(rawmess.Header)
}
