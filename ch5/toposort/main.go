// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"sort"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}
var prereqsWithMap = map[string]map[string]bool{
    "algorithms": {"data structures": true},
    "calculus":   {"linear algebra": true},

    "compilers": {
        "data structures":       true,
        "formal languages":      true,
        "computer organization": true,
    },
	
    "data structures":       {"discrete math": true},
    "databases":             {"data structures": true},
    "discrete math":         {"intro to programming": true},
    "formal languages":      {"discrete math": true},
    "networks":              {"operating systems": true},
    "operating systems":     {"data structures": true, "computer organization": true},
    "programming languages": {"data structures": true, "computer organization": true},
}

//!-table

//!+main
func main() {

	for i, course := range topSortUsingMap(prereqsWithMap){
		fmt.Printf("%d:\t%s\n", i, course)
	}

	// for i, course := range topoSort(prereqs) {
	// 	fmt.Printf("%d:\t%s\n", i+1, course)
	// }
}

func topSortUsingMap(m map[string]map[string]bool) []string {
	var result []string
	seen := make(map[string]bool)
    var visitAll func(items map[string]bool)

    visitAll = func(items map[string]bool) {
        for item, _ := range items {
            if !seen[item] {
                seen[item] = true
                visitAll(m[item])
                result = append(result, item)
            }
        }
    }

    var keys = make(map[string]bool)
    for key := range m {
        keys[key] = true
    }
    visitAll(keys)
	return result
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}





//!-main
