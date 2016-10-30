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
