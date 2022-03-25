package main

import (
	"bytes"
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// a simple version of strings.Reader implement io.Reader interface
type StringReader struct {
    s string
    i int64
}

func (r *StringReader) Read(p []byte) (n int, err error) {
    // if p is nil or empty, return (0, nil)
    if len(p) == 0 {
        return 0, nil
    }

    // copy() guarantee copy min(len(p),len(r.s[r.i:])) bytes
    n = copy(p, r.s[r.i:])
    if r.i += int64(n); r.i >= int64(len(r.s)) {
        err = io.EOF
    }
    return
}

// NewReader return a StringReader with s
func NewReader(s string) *StringReader {
    return &StringReader{s, 0}
}


func main()  {
	// NewReader("Hi")

	//var sr StringReader
	//x:= NewReader("<p>Links:</p><div>Hi</div>")
	// fmt.Println("x ",x.s)
	// sr.Read([]byte(x.s))
	// x:=NewReader("")
	doc, err := html.Parse(NewReader("<p>Links:</p><div><a href='xxx'>Hi</a></div>"))
	if err != nil {
		panic(err)
	}
	fmt.Println("visit(nil, doc) ", visit(nil, doc))
	s := "hello 世界"
	b := &bytes.Buffer{}
	r := LimitReader(NewReader(s), 5)
	n, nb := b.ReadFrom(r)
	fmt.Println("n: ", n, nb)

	if n != 5 {
	//fmt.Println("r.r: ", r.r)
		
	}

	
}

func visit(links []string, n *html.Node) []string {
	
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	if n.FirstChild != nil {

		links = visit(links, n.FirstChild)
	}
	if 	n.NextSibling != nil{
		links = visit(links, n.NextSibling)
	}
	return links
}


type LimitReaderType struct{
	r io.Reader
	c int64
}


func (l *LimitReaderType)Read(p []byte) (n int, err error){
	if l.c <= 0 {
		return 0, io.EOF
	}
	//defer l.Close()
	if int64(len(p)) > int64(l.c) {
		p = p[:l.c]
	}
	n, err = l.r.Read(p)
	l.c -= int64(n)
	return 
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitReaderType{r,n}
}

// func (l *LimitReaderType)Close()(err error)  {
// 	return l.cl.Close()
// }