package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseBytes(b []byte) []byte {
	if utf8.RuneCount(b) == 1 {
		return b
	}
	_, s := utf8.DecodeRune(b)
	return append(reverseBytes(b[s:]), b[:s]...)
}

func main()  {
	b := []byte("你好啊，世界")
	fmt.Println(string(b))
	fmt.Println(string(reverseBytes(b)))
}
