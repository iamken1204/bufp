/*
Package bufp provides a faster way to print values to io.Writer.
(about 5.x ~ 9.x faster compared to fmt.Printf)
This package is inspired by http://byrd.im/competitive-go/.

    bw := bufp.NewStdBWriter()
    // stdout must be flushed manually.
    defer bw.W.Flush()
    bw.Brintf("%v %v %v\n", 123, "gop", 99.999)
*/
package bufp
