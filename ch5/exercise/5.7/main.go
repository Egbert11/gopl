package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int
func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		printElementNodeStart(n)
		depth++
	}
	if n.Type == html.CommentNode {
		printCommentNodeStart(n)
		depth++
	}
	if n.Type == html.TextNode {
		printTextNode(n)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		printElementNodeEnd(n)
	}
	if n.Type == html.CommentNode {
		depth--
		printCommentNodeEnd(n)
	}
	if n.Type == html.TextNode {
	}
}

func printElementNodeStart(n *html.Node) {
	fmt.Printf("%*s<%s", depth*2, "", n.Data)
	for _, a := range n.Attr {
		fmt.Printf(" %s=\"%s\"", a.Key, a.Val)
	}
	if n.FirstChild != nil {
		fmt.Print(">\n")
	} else {
		fmt.Print("/>\n")
	}
}

func printElementNodeEnd(n *html.Node) {
	if n.FirstChild != nil {
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func printCommentNodeStart(n *html.Node) {
	fmt.Printf("<!--%s", n.Data)
}

func printCommentNodeEnd(n *html.Node) {
	fmt.Print("-->\n")
}

func printTextNode(n *html.Node) {
	trimmed := strings.TrimSpace(n.Data)
	if trimmed != "" {
		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
	}
}

func getUrlContent(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return nil, err
	}
	return doc, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("you should input a url")
	}
	doc, err := getUrlContent(os.Args[1])
	if err != nil {
		log.Fatalf("getUrlContent url:%s error:%s", os.Args[1], err)
	}

	forEachNode(doc, startElement, endElement)

}