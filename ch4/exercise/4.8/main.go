package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	cats := make(map[string]int)
	var utflen [utf8.UTFMax+1]int
	invalid := 0

	reader := bufio.NewReader(os.Stdin)
	for {
		r, n, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount:%v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		for catName, rangeTable := range unicode.Properties {
			if unicode.In(r, rangeTable) {
				cats[catName]++
			}
		}

		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("\nlen\tcount\n")
	for l, n := range utflen {
		if n > 0 {
			fmt.Printf("%d\t%d\n", l, n)
		}
	}

	fmt.Printf("\n%-30s count\n", "category")
	for cat, n := range cats {
		fmt.Printf("%-30s %d\n", cat, n)
	}

	if invalid > 0 {
		fmt.Printf("\n %d invalid UTF-8 characters.\n", invalid)
	}
}
