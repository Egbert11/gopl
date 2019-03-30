package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func outline(count map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		count[n.Data] += 1
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(count, c)
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	var count = make(map[string]int)
	outline(count, doc)
	total := 0
	for k, v := range(count) {
		fmt.Printf("%s: %d\n", k, v)
		total += v
	}
	fmt.Printf("total: %d\n", total)
}

