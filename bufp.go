package bufp

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// BWriter contains bufio.Writer
type BWriter struct {
	W *bufio.Writer
}

// Brintf prints out arguments to assigned writer.
func (bw *BWriter) Brintf(format string, args ...interface{}) {
	Printf(bw.W, format, args...)
}

// NewStdBWriter returns a BWriter with StdWriter
func NewStdBWriter() *BWriter {
	w := bufio.NewWriter(os.Stdout)
	return &BWriter{w}
}

// NewWriter returns *bufio.Writer
// constructed by given writer.
func NewWriter(writer io.Writer) *bufio.Writer {
	return bufio.NewWriter(writer)
}

// Printf prints out arguments to given writer.
func Printf(writer *bufio.Writer, format string, args ...interface{}) {
	fmt.Fprintf(writer, format, args...)
}
