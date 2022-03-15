// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	// for _, url := range os.Args[1:] {
	// }
	
	outline("https://google.com")

}

func outline(url string) error {
	var depth int
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	s:=`<p>Links:</p>`
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		return err
	}
	//!+startend
var startElement func(n *html.Node)
var endElement func(n *html.Node)
startElement = func (n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

endElement = func (n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, " ", n.Data)
	}
}

//!-startend

	//fmt.Println(doc.Data)
	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		//fmt.Print("pre ",n.Data)
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		//fmt.Println("current ",c.Data)
		forEachNode(c, pre, post)
	}

	if post != nil {
		//fmt.Println("post ",n.Data)
		post(n)
	}
}

//!-forEachNode

