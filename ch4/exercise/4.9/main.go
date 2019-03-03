package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		word := input.Text()
		counts[word]++
	}

	if err := input.Err(); err != nil {
		fmt.Printf("word freq:%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("word\tfrequency\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

}
