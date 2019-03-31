package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)


func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil {
		ret := pre(n)
		if ret == false {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode(c, pre, post)
		if node != nil  {
			return node
		}
	}

	if post != nil {
		ret := post(n)
		if ret == false {
			return n
		}
	}
	return nil
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


func ElementById(n *html.Node, id string) *html.Node {
	pre := func (n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}
		for _, a := range n.Attr {
			//fmt.Printf("id:%s val:%s\n", a.Key, a.Val)
			if a.Key == "id" && a.Val == id {
				//fmt.Printf("searchElementById id:%s val:%s", a.Key, a.Val)
				return false
			}
		}
		return true
	}
	return forEachNode(n, pre, nil)
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("you should input a url and id")
	}
	doc, err := getUrlContent(os.Args[1])
	if err != nil {
		log.Fatalf("getUrlContent url:%s error:%s", os.Args[1], err)
	}
	id := os.Args[2]
	node := ElementById(doc, id)
	if node != nil {
		fmt.Println(*node)
	}
}