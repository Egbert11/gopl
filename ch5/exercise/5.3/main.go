package main

import (
	"fmt"
	"os"
	"unicode"

	"golang.org/x/net/html"
)

func output(n *html.Node) {
	if n.Type == html.TextNode && n.Parent.Data != "script" && n.Parent.Data != "style" {
		if containNonSpace([]rune(n.Data)){
			fmt.Println(n.Data)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		output(c)
	}
}

func containNonSpace(s []rune) bool {
	retVal := false
	for _, c := range s {
		if !unicode.IsSpace(c) {
			retVal = true
			break
		}
	}
	return retVal
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	output(doc)
}

