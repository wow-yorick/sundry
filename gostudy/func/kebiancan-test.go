package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4))
	values := []int{1, 2, 3}
	fmt.Println(sum(values...))
	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name)
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d:", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}

// func ElementsByTagName(doc *html.Node, name ...string) []*html.Node

// images := ElementsByTagName(doc, "img")
// headings := ElementsByTagName(doc, "h1", "h2", "h3")
