// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 123.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

//!+
func main() {
    for _, url := range os.Args[1:] {
		//outline(url)

	fmt.Println(url)
    resp, _ := http.Get(url)
	
	defer resp.Body.Close()
	//s:=`<p>Links:</p><div>Hi</div>`
	//fmt.Println(resp.Body)
	doc, err := html.Parse(resp.Body)
   // doc, err := html.Parse(os.Stdin)
    if err != nil {
        fmt.Fprintf(os.Stderr, "outline: %v\n", err)
        os.Exit(1)
    }
    text(doc)
}
}

func text(n *html.Node) {
    // content in <style> or <script> are ignored
    if n.Type == html.ElementNode &&
        (n.Data == "style" || n.Data == "script") {
        return
    }

    if n.Type == html.TextNode {
        fmt.Println(n.Data)
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        text(c)
    }
}

//!-
