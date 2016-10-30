/*
Package bufp provides a faster way to print strings to Stdout.
(about 5.x ~ 9.x compared to `fmt.Printf`)
This package is inspired by http://byrd.im/competitive-go/ .

    bwriter := bufp.NewStdWriter()
    // stdout must be flushed manually.
    defer bwriter.Flush()
    bufp.Bprintf(bwriter, "%d%s%b", []interface{}{123, "hello", 98.76})
*/
package bufp
