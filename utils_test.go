package main

import (
	"fmt"
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	exmp := "Thu, 14 Apr 2022 16:37:57 +0800"
	res := ParseTime(exmp).Format(time.RFC1123Z)
	if exmp != res {
		t.Errorf("Error while parsing: %s\t%s", exmp, res)
	}
	fmt.Println(ParseTime(exmp))
}

func TestParseEncodedString(t *testing.T) {
	exmp := `=?gb18030?B?UVHNvMasMjAyMjAyMDUyMjMxNDc=?=`
	res, _ := ParseEncodedString(exmp)
	if res != "QQ图片20220205223147" {
		t.Errorf("Error while parsing: %s\t%s", exmp, res)
	}
}

func TestRemoveOutterQuotation(t *testing.T) {
	var testPairs = map[string]string{
		"\"abd\"": "abd",
		"\"bda":   "bda",
		"bda\"":   "bda",
		"asdf":    "asdf",
	}
	for in, out := range testPairs {
		if RemoveOutterQuotation(in) != out {
			t.Errorf("Error in instance: in=%s\tout=%s", in, RemoveOutterQuotation(in))
		}
	}
}
