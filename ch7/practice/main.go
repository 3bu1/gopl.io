package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"text/tabwriter"
	"time"

	"golang.org/x/net/html"
)

// a simple version of strings.Reader implement io.Reader interface
type StringReader struct {
    s string
    i int64
}

type Profile struct {
	Name string
	Age int
	Mobile string
}

var pro = []*Profile{
	{"Tribhvan",31, "7760808111",},
	{"Tribhvan1",33, "7760808112",},
	{"Tribhvan2",34, "7760808113",},
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
	fmt.Println("By Title, Artist")
	printTracks(useSortByColumns())

	fmt.Println("\nUse sort.Stable. By Title, Artist")
	printTracks(useSortStable())

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



// Sorting sorts a music playlist into a variety of orders.
var stdout io.Writer = os.Stdout

//!+main
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func tracks() []*Track {
	return []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

//!-main

//!+printTracks
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

//!-printTracks

//!+titlecode
type byTitle []*Track

func (x byTitle) Len() int           { return len(x) }
func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x byTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-titlecode

//!+artistcode
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-artistcode

//!+yearcode
type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

//!-yearcode

//!+multicolumns

type less func(x, y *Track) bool

func colTitle(x, y *Track) bool  { return x.Title < y.Title }
func colArtist(x, y *Track) bool { return x.Artist < y.Artist }
func colAlbum(x, y *Track) bool  { return x.Album < y.Album }
func colYear(x, y *Track) bool   { return x.Year < y.Year }
func colLength(x, y *Track) bool { return x.Length < y.Length }

type byColumns struct {
	tracks  []*Track
	columns []less
}

func sortByColumns(t []*Track, f ...less) *byColumns {
	return &byColumns{
		tracks:  t,
		columns: f,
	}
}

func (x byColumns) Len() int      { return len(x.tracks) }
func (x byColumns) Swap(i, j int) { x.tracks[i], x.tracks[j] = x.tracks[j], x.tracks[i] }
func (x byColumns) Less(i, j int) bool {
	a, b := x.tracks[i], x.tracks[j]
	var k int
	// compare columns one by one except the last
	for k = 0; k < len(x.columns)-1; k++ {
		f := x.columns[k]
		switch {
		case f(a, b):
			return true
		case f(b, a):
			return false
		}
	}
	// all equal, use last column as final judgement
	return x.columns[k](a, b)
}

//!-multicolumns

func useSortByColumns() []*Track {
	t := tracks()
	sort.Sort(sortByColumns(t, colTitle, colArtist))
	return t
}

func useSortStable() []*Track {
	t := tracks()
	sort.Stable(byArtist(t))
	sort.Stable(byTitle(t))
	return t
}

