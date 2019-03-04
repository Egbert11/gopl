package main

import (
	"bytes"
	"fmt"
)

func reverse(s string) string {
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}
	offset := n % 3
	if offset > 0 {
		buf.WriteString(s[:offset])
		buf.WriteByte(',')
	}
	s = s[offset:]
	steps := n / 3
	for i := 0; i < steps; i++ {
		buf.WriteString(s[3*i:3*(i+1)])
		if i < steps - 1 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}

func main()  {
	fmt.Println(comma("1234567890"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("11234"))
	fmt.Println(comma("112"))
	fmt.Println(comma("11"))
	fmt.Println(comma("1"))
}