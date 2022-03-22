// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 125.

// Findlinks2 does an HTTP GET on each URL, parses the
// result as HTML, and prints the links within it.
//
// Usage:
//	findlinks url ...
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// visit appends to links each link found in n, and returns the result.
func visitAllNodes(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode &&  (n.Data == "a" || n.Data == "link" ) {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Printf("n.Data: %s, a.val: %s \n", n.Data,a.Val)
				links = append(links, a.Val)
			}
		}
	}
	if n.Type == html.ElementNode && (n.Data == "img" || n.Data == "script") {
		for _, a := range n.Attr {
			if a.Key == "src" {
				fmt.Printf("n.Data: %s, a.val: %s \n", n.Data,a.Val)
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visitAllNodes(links, c)
	}
	return links
}

var mapping = map[string]string{"a": "href", "img": "src", "script": "src", "link": "href"}

// visit appends to links each link found in n and returns the result.
func visit(target string, links []string, n *html.Node) []string {
    if n.Type == html.ElementNode && n.Data == target {
        for _, a := range n.Attr {
            if a.Key == mapping[target] {
                links = append(links, a.Val)
            }
        }
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        links = visit(target, links, c)
    }
    return links
}

//!+
func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

// findLinks performs an HTTP GET request for url, parses the
// response as HTML, and extracts and returns the links.
func findLinks(url string) ([]string, error) {
	result := []string{}
	visitAllNodesFlag:=true 
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	if visitAllNodesFlag {
		result = visitAllNodes(nil,doc)		
	}else{
		result = visit("img",nil, doc)
	}
	
	return result,err
}

//!-
