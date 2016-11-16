# bufp
[![Go Report Card](https://goreportcard.com/badge/github.com/iamken1204/bufp)](https://goreportcard.com/report/github.com/iamken1204/bufp)
[![Build Status](https://travis-ci.org/iamken1204/bufp.svg?branch=master)](https://travis-ci.org/iamken1204/bufp)   

Package `bufp` provides a faster way to print values to [`io.Writer`](https://golang.org/pkg/io/#Writer).   
(about 5.x ~ 9.x faster compared to `fmt.Printf`)   
This package is inspired by [Using Go for competitive programming](http://byrd.im/competitive-go/).

```go
bw := bufp.NewStdBWriter()
// stdout must be flushed manually.
defer bw.W.Flush()
bw.Brintf("%v %v %v\n", 123, "gop", 99.999)
```
_or_
```go
bw := bufp.NewWriter(bufio.NewWriter(os.Stdout))
// stdout must be flushed manually.
defer bw.Flush()
bufp.Printf(bw, "%v %v %v\n", 123, "gop", 99.999)
```
