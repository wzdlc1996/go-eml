package main

import (
	"bytes"
	"net/mail"
	"strings"
)

func ParseRaw(emlRaw []byte) (m Message, err error) {
	emlReader := bytes.NewReader(emlRaw)
	msg, _ := mail.ReadMessage(emlReader)

	m.Header = msg.Header

	m.Date = ParseTime(msg.Header.Get("Date"))
	m.From, _ = parseAddress(msg.Header.Get("From"))
	m.To, _ = parseAddressList(msg.Header["To"])

	m.Subject, err = ParseEncodedString(msg.Header.Get("Subject"))
	return
}

func parseAddress(addrPair string) (addr mail.Address, e error) {
	addrPairSlc := strings.Split(addrPair, " ")
	nameStr := addrPairSlc[0]
	nameStr = RemoveOutterQuotation(nameStr)
	plainNameStr, err := ParseEncodedString(nameStr)
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
