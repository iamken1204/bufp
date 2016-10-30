/*
Package bufp provides a faster way to print strings to stdout.
(about 5.x ~ 9.x faster compared to fmt.Printf)
This package is inspired by http://byrd.im/competitive-go/.

    bwriter := bufp.NewStdWriter()
    // stdout must be flushed manually.
    defer bwriter.Flush()
    bufp.Bprintf(bwriter, "%d%s%f", []interface{}{123, "hello", 98.76})
*/
package bufp
