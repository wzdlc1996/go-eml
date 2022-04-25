package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"mime/quotedprintable"
	"net/mail"
	"strings"
)

func ParseRaw(emlRaw []byte) (m Message, err error) {
	emlReader := bytes.NewReader(emlRaw)
	msg, _ := mail.ReadMessage(emlReader)

	m.Header = msg.Header

	m.Date = parseTime(msg.Header.Get("Date"))
	m.From, _ = parseAddress(msg.Header.Get("From"))
	m.To, _ = parseAddressList(msg.Header["To"])
	//contentType := msg.Header.Get("Content-Type")

	m.Subject, _ = parseEncodedString(msg.Header.Get("Subject"))

	ctype, params, err := mime.ParseMediaType(msg.Header.Get("Content-Type"))
	var text string
	var att []Attachment
	if strings.Contains(ctype, "multipart") {
		bd, ok := params["boundary"]
		if ok {
			text, att, err = parseBody(msg.Body, bd)
		}
	}

	m.Text = text
	m.Attachments = att

	return
}

func parseAddress(addrPair string) (addr mail.Address, e error) {
	addrPairSlc := strings.Split(addrPair, " ")
	nameStr := addrPairSlc[0]
	nameStr = removeOutterQuotation(nameStr)
	plainNameStr, err := parseEncodedString(nameStr)
	if err != nil {
		return
	}
	addrStr := addrPairSlc[1]

	plainAddrPair := strings.Join([]string{plainNameStr, addrStr}, " ")
	addrpt, e := mail.ParseAddress(plainAddrPair)
	return *addrpt, e
}

func parseAddressList(addrPairs []string) (addrs []mail.Address, e error) {
	for _, addrPair := range addrPairs {
		addr, err := parseAddress(addrPair)
		if err != nil {
			return
		}
		addrs = append(addrs, addr)

	}
	return addrs, nil
}

func parseBody(msgBody io.Reader, boundary string) (text string, attachments []Attachment, err error) {
	mpReader := multipart.NewReader(msgBody, boundary)
	for {
		part, err := mpReader.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			return text, attachments, err
		}

		ctype, params, err := mime.ParseMediaType(part.Header.Get("Content-Type"))
		_ = fmt.Sprint(ctype, params)
		if err != nil {
			return text, attachments, err
		}

		if isAttatchment(part) {
			att, err := decodeAttachment(part)
			if err != nil {
				return text, attachments, err
			}
			attachments = append(attachments, att)
		}
		// TODO add html and plain text parsing, they are in multipart/...
	}
	return
}

func isAttatchment(part *multipart.Part) bool {
	return part.FileName() != ""
}

func decodeAttachment(part *multipart.Part) (att Attachment, err error) {
	att.Filename, err = parseEncodedString(part.FileName())
	if err != nil {
		return
	}

	encoding := part.Header.Get("Content-Transfer-Encoding")
	switch encoding {
	case "base64":
		dec := base64.NewDecoder(base64.StdEncoding, part)
		b, err := ioutil.ReadAll(dec)
		if err != nil {
			return Attachment{}, err
		}
		att.Body = b
	case "quoted-printable":
		// TODO: error controlling for quotedprintable
		att.Body, _ = ioutil.ReadAll(quotedprintable.NewReader(part))
	default:
		att.Body, _ = ioutil.ReadAll(part)
	}
	return att, nil
}
