// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//!+bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

//!-bytecounter


type WordCounter int
type LineCounter int
type ScanFunc func (p []byte, EOF bool) (advance int, token []byte, err error)

func ScanBytes(p []byte, fn ScanFunc) (cnt int)  {
	for true {
		advance, token, _ := fn(p, true)
		if len(token) == 0{
			break
		}
		p= p[advance:]
		cnt++
	}
	return cnt
}

func (w *WordCounter)Write(p []byte) (int,error)  {
	cnt := ScanBytes(p, bufio.ScanWords)
	*w += WordCounter(cnt)
	return cnt, nil
}

func (l *LineCounter)Write(p []byte)(int, error)  {
	cnt := ScanBytes(p, bufio.ScanLines)
	*l += LineCounter(cnt)
	return cnt, nil
}

func (wordC *WordCounter) StringCount(p string) (int, error)  {
	pword := strings.Split(p, " ")
	pcount := 0
	errorVar := fmt.Errorf("input cant be empty") 
	if p == "" {
		return 0, errorVar
	}
	for _, pv := range pword {
		if pv != "" {
			pcount++
		}

	}
	*wordC = WordCounter(pcount)
	return pcount, errorVar
}

func (c WordCounter) String() string {
    return fmt.Sprintf("contains %d words", c)
}

func main() {
	//!+main
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Println("len(name)",len(name))
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
	var w WordCounter
	//w.StringCount("hello worldHi")
	fmt.Fprintf(&w, "this is a sentence")
	fmt.Println("word length",w)
	w = 0
    fmt.Fprintf(&w, "This")
    fmt.Println(w)
//	fmt.Fprintf(&w, "This is an sentence.")
  //  fmt.Println(c)
	
  var l LineCounter
  fmt.Println(l)

  fmt.Fprintf(&l, `This is another\n
line`)
  fmt.Println(l)

  l = 0
  fmt.Fprintf(&l, "This is another\nline")
  fmt.Println(l)

  fmt.Fprintf(&l, "This is one line")
  fmt.Println(l)

 
  //!-main

var cwx CounterWriter
	fmt.Println(cwx)
	//fmt.Fprintf(cwx.writer, "this is counterwriter writing !!")
	cwx.Write([]byte("this is counterwriter writing !!"))
	fmt.Println(cwx)
}

// type countingWriter int64

// func (cw *countingWriter) CountingWriter(w io.Writer)(io.Writer, *int64)  {

// 	nw := bufio.NewWriter(w)
// 	var c ByteCounter
// 	c.Write([]byte("hello"))
// 	cw = countingWriter(len(w))
// 	// cnt := ScanBytes(w, bufio.NewWriter)
// 	// *cw += countingWriter(cnt)
	
// 	return nw,(*int64)(cw)
// }

type CounterWriter struct {
    counter int64
    writer  io.Writer
}

// must be pointer type in order to count
func (cw *CounterWriter) Write(p []byte) (int, error) {
    cw.counter += int64(len(p))
    return cw.writer.Write(p)
}

// newWriter is a Writer Wrapper, return original Writer
// and a Counter which record bytes have written
func CountingWriter(w io.Writer) (io.Writer, *int64) {
    cw := CounterWriter{0, w}
    return &cw, &cw.counter
}
