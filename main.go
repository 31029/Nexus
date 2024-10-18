package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	var w io.Writer
	io.WriteString()
	w = os.Stdout
	rw := w.(io.ReadWriter) // success: *os.File has both Read and Write
}