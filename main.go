package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/mail"
	"os"
)

func main() {
	emlFile, _ := os.ReadFile("./testEmls/图片.eml")
	rawmess, _ := ParseRaw(emlFile)
	fmt.Println(rawmess.Subject)
	emlReader := bytes.NewReader(emlFile)
	msg, _ := mail.ReadMessage(emlReader)
	body, _ := ioutil.ReadAll(msg.Body)
	fmt.Println(string(body))
}
