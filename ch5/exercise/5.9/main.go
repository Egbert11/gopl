package main

import (
	"fmt"
	"strings"
)

func expand(s string, f func(s string) string) string {
	return strings.Replace(s, "foo", f("foo"), -1)
}

func main() {
	s := "foo string containing foo here foo"
	fmt.Println(expand(s, strings.ToUpper))
}