// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)

	}
	

}

func outline(url string) error {
	var depth int
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	//s:=`<p>Links:</p><div>Hi</div>`
	fmt.Println(resp.Body)
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	//!+startend
var startElement func(n *html.Node)
var endElement func(n *html.Node)
var findTag func(n *html.Node)
var findText func(n *html.Node)
selectedTag := ""
text:=""
startElement = func (n *html.Node) {
	localN := n
	if localN.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
		findTag(n)
	}
}
findTag = func (n *html.Node)  {
	tag:=n.Data
	switch tag {
	case "body","html", "head", "style", "script","title","input", "form","br","img","video":{
	}
	default: {
	selectedTag = n.Data
}}
	
}
findText = func(n *html.Node) {
	tag:=n.Data
	if selectedTag != "" {
		text += "\n"+tag
		selectedTag = ""
	}
}

endElement = func (n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, " ", n.Data)
	}
	findText(n)
}

//!-startend

	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call
fmt.Printf("%s \n", text)
	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		//fmt.Print(n.Data)
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
	//	fmt.Println("current ",c.Data)
		forEachNode(c, pre, post)
	}

	if post != nil {
		//fmt.Println("post ",n.Data)
		post(n)
	}
}

//!-forEachNode

