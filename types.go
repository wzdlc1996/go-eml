package main

import (
	"net/mail"
	"time"
)

type Message struct {
	Header mail.Header

	To   []mail.Address
	From mail.Address

	Date time.Time

	Subject     string
	Text        string
	Attachments []Attachments
}

type Attachments struct {
	Filename string
	Body     []byte
}
