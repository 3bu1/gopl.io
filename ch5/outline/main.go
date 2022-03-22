// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 123.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

//!+
func main() {
	ElementMaps := make(map[string]int)

	f,ferr := os.Open("/home/tribhuvan/workspace/lab/gopl.io/ch5/practice/index.html")
	if ferr != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", ferr)
		//os.Exit(1)
	}
	doc, err := html.Parse(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
//	x := &ElementMaps
	outline(nil, doc, ElementMaps)
	fmt.Printf("ElementMaps size : %d, element: %v \n",len(ElementMaps), ElementMaps)
}

func outline(stack []string, n *html.Node, ElementMaps map[string]int) {
	if n.Type == html.ElementNode {
		elementCountOutline(stack, n.Data, ElementMaps)
		
		
		//ElementMaps[stack]
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c, ElementMaps)
	}
}

func elementCountOutline(stack []string, tag string, ElementMaps map[string]int)  {
	stack = append(stack, tag) // push tag
	ElementMaps[tag]++
	fmt.Println(stack)
}

//!-
