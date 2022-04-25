package main

import (
	"fmt"
	"testing"
	"time"
)

func TestparseTime(t *testing.T) {
	exmp := "Thu, 14 Apr 2022 16:37:57 +0800"
	res := parseTime(exmp).Format(time.RFC1123Z)
	if exmp != res {
		t.Errorf("Error while parsing: %s\t%s", exmp, res)
	}
	fmt.Println(parseTime(exmp))
}

func TestparseEncodedString(t *testing.T) {
	exmp := `=?gb18030?B?UVHNvMasMjAyMjAyMDUyMjMxNDc=?=`
	res, _ := parseEncodedString(exmp)
	if res != "QQ图片20220205223147" {
		t.Errorf("Error while parsing: %s\t%s", exmp, res)
	}
}

func TestremoveOutterQuotation(t *testing.T) {
	var testPairs = map[string]string{
		"\"abd\"": "abd",
		"\"bda":   "bda",
		"bda\"":   "bda",
		"asdf":    "asdf",
	}
	for in, out := range testPairs {
		if removeOutterQuotation(in) != out {
			t.Errorf("Error in instance: in=%s\tout=%s", in, removeOutterQuotation(in))
		}
	}
}
