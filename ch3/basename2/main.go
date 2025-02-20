// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 72.

// Basename2 reads file names from stdin and prints the base name of each one.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
	// NOTE: ignoring potential errors from input.Err()
}

// basename removes directory components and a trailing .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
//!+
func basename(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	fmt.Println(slash)
	s = s[slash+1:]
	fmt.Println(s)
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

//!-
