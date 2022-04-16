package main

import (
	"encoding/base64"
	"regexp"
	"strings"
	"time"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func ParseTime(ts string) time.Time {
	if ts == "" {
		return time.Time{}
	}
	formats := []string{
		time.RFC1123Z,
	}
	for _, fm := range formats {
		date, err := time.Parse(fm, ts)
		if err == nil {
			return date
		}
	}
	return time.Time{}
}

// ParseEncodedString parses the string with form like "=?gb18030?B?...", where ... is usually the regular Base64 code
// while the charset is in the head.
func ParseEncodedString(es string) (string, error) {
	var err error
	re := regexp.MustCompile(`^=\?(.*?)\?(.*?)\?(.*)\?=$`)
	res := re.FindAllStringSubmatch(es, -1)
	if len(res) == 0 || len(res[0]) != 4 {
		return "", err
	}
	chset := res[0][1]   // is the charset type of the string
	encType := res[0][2] // must be either B or Q, means base64 or quoted-printable
	cont := res[0][3]    // is the content of string

	chDecoder := chsetMap[chset].NewDecoder()

	switch strings.ToUpper(encType) {
	case "B":
		var resRaw []byte
		if resRaw, err = base64.StdEncoding.DecodeString(cont); err != nil {
			return "", err
		}
		if resRaw, err = chDecoder.Bytes(resRaw); err != nil {
			return "", err
		}
		return string(resRaw), nil
	case "Q":
		// TODO: need to add the handler for quoted-printable strings
	default:
		//
	}
	return "", err
}

var chsetMap = map[string]encoding.Encoding{
	"gb18030": simplifiedchinese.GB18030,
	"gbk":     simplifiedchinese.GBK,
}
