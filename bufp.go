/*
Package bufp provides a faster way to print strings to Stdout.
This package is inspired by http://byrd.im/competitive-go/.

    bwriter := bufp.NewStdWriter()
    // stdout must be flushed manually.
    defer bwriter.Flush()
    bufp.Bprintf(bwriter, "%d%s%b", []interface{}{123, "hello", 98.76})
*/
package bufp

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// NewStdWriter returns *bufio.Writer
// constructed by os.Stdout
func NewStdWriter() *bufio.Writer {
	return bufio.NewWriter(os.Stdout)
}

// NewWriter returns *bufio.Writer
// constructed by given writer
func NewWriter(writer io.Writer) *bufio.Writer {
	return bufio.NewWriter(writer)
}

// Bprintf prints interfaces{}
func Bprintf(writer *bufio.Writer, format string, args ...interface{}) {
	fmt.Fprintf(writer, format, args...)
}
