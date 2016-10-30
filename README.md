# bufp

Package `bufp` provides a faster way to print values to [`io.Writer`](https://golang.org/pkg/io/#Writer).   
(about 5.x ~ 9.x faster compared to `fmt.Printf`)   
This package is inspired by [Using Go for competitive programming](http://byrd.im/competitive-go/).

```go
bwriter := bufp.NewStdWriter()
// stdout must be flushed manually.
defer bwriter.Flush()
bufp.Printf(bwriter, "%d%s%f", []interface{}{123, "hello", 98.76})
```
