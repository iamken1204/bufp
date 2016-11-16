package main

import (
	"bufio"
	"os"

	"github.com/iamken1204/bufp"
)

func main() {

	bw := bufp.NewStdBWriter()
	defer bw.W.Flush()
	bw.Brintf("%v %v %v\n", 123, "gop", 99.999)

	bw2 := bufp.NewWriter(bufio.NewWriter(os.Stdout))
	defer bw2.Flush()
	bufp.Printf(bw2, "%v %v %v\n", 123, "gop", 99.999)

}
