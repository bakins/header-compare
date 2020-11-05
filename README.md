# Comparison of HTTP headers access in Go

## To run benchmarks

Clone this repo, cd into it, then run:

```shell
go test -v -bench=. -run=xxx -benchmem -benchtime=10s .
```

## Sample results

```
goos: darwin
goarch: amd64
pkg: github.com/bakins/header-compare
BenchmarkGetCommon
BenchmarkGetCommon-12               	563411564	        21.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetCommonLowercase
BenchmarkGetCommonLowercase-12      	222627403	        53.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetCommonDirect
BenchmarkGetCommonDirect-12         	1000000000	         4.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetUncommon
BenchmarkGetUncommon-12             	439755751	        27.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetUncommonLowercase
BenchmarkGetUncommonLowercase-12    	100000000	       101 ns/op	      32 B/op	       1 allocs/op
BenchmarkGetUncommonDirect
BenchmarkGetUncommonDirect-12       	1000000000	         4.03 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/bakins/header-compare	65.484s
```

## Discussion

go [net/http headers](https://golang.org/pkg/net/http/#Header) provide a way to access HTTP headers in a case-insensitive
manner.  This is helpful, particularly when working with both HTTP/1.1 and HTTP/2 - as HTTP/2 uses lowercase keys while HTTP/1.1
generally uses a "canonical form" with mixed cases.

In go, HTTP headers are [MIMEHeaders](https://golang.org/pkg/net/textproto/#MIMEHeader).  When using [Header.Get](https://golang.org/pkg/net/http/#Header.Get)
the key is converted to [canonical case](https://golang.org/pkg/net/textproto/#CanonicalMIMEHeaderKey). Common header keys
are "cached" internally, so there is no need to copy keys and convert in place - however, the key is still scanned to ensure
it is "common."  Uncommon or custom header keys are never cached.

Not: you should always use canonical case to access headers to ensure case insensitivity is used.

These benchmarks show different ways to access header values.  The fastest way for both uncommon and common ways is to pre-canonicalize
the key and access the header map directly.  

Example:

```go
func getHeader(h http.Header, key string) string {
	if h == nil {
		return ""
	}

	v := h[key]
	
	if len(v) == 0 {
		return ""
	}

	return v[0]
}

// works for common and custom headers
var myHeaderKey = http.CanonicalHeaderKey("x-my-header")

func doSomething(r *http.Header) {
    val := getHeader(r.Header, myHeaderKey)
}
```

For most code, accessing headers is rarely a major cpu/memory consumer.  However, if you need to eek out that last little
bit of performance, I hope this is helpful.
