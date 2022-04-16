package main

import (
	"bytes"
	"net/mail"
)

func ParseRaw(emlRaw []byte) (m mail.Message, err error) {
	emlReader := bytes.NewReader(emlRaw)
	msg, err := mail.ReadMessage(emlReader)
	m = *msg
	if err != nil {
		return
	}
	return
}
