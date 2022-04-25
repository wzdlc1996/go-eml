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
	HTML        string
	Attachments []Attachment
}

type Attachment struct {
	Filename string
	Body     []byte
}
