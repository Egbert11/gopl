package main

import (
	"fmt"
	"os"
	"strings"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	} else {
		return comma(s[:n-3]) + "," + s[n-3:]
	}
}

func main()  {
	for _, arg := range os.Args[1:] {
		var prefix string
		var result string = arg
		var suffix string

		if arg[0] == '-' {
			prefix = "-"
			result = arg[1:]
		}

		dot := strings.LastIndex(result, ".")
		if dot >= 0 {
			suffix = result[dot:]
			result = result[:dot]
		}
		fmt.Println(prefix + comma(result) + suffix)
	}
}