/*
Package bufp provides a faster way to print values to io.Writer.
(about 5.x ~ 9.x faster compared to fmt.Printf)
This package is inspired by http://byrd.im/competitive-go/.

    bwriter := bufp.NewStdWriter()
    // stdout must be flushed manually.
    defer bwriter.Flush()
    bufp.Printf(bwriter, "%d%s%f", 123, "hello", 98.76)
*/
package bufp
