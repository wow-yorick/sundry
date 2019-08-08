package main

import (
	"bytes"
	"flag"
	"io"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

const debug = false

func main() {
	// var w io.Writer
	// w = os.Stdout
	// w = new(bytes.Buffer)

	// var rwc io.ReadWriteCloser
	// rwc = os.Stdout
	// os.Stdout.Write([]byte("hello"))
	// os.Stdout.Close()

	// var w io.Writer
	// w = os.Stdout
	// w.Write([]byte("hello"))
	// w.Close()
	// flag.Parse()
	// fmt.Printf("sleeping for %v ...", *period)
	// time.Sleep(*period)
	// fmt.Println()
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
	if debug {

	}
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
