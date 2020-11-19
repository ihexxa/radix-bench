# radix-bench

_A simple benchmark of Go/Golang radix tree implementations._

### How to test

```go
go test -bench=. -benchmem
or
go test -bench=. -benchmem -args -n=<sample count>
```

### Implementations

[armon/go-radix](https://github.com/armon/go-radix)
[ihexxa/q-radix](https://github.com/ihexxa/q-radix)
[kellydunn/go-art](https://github.com/kellydunn/go-art)

### Test Results

**Note**: This benchmark provides a quick reference of implementations' performance.
Special features or useful APIs provided by each implementations are not listed here.

**Case1: Insert one element and get it, repeat this for 10000 times:**

```
BenchmarkRadix/go-radix_insert_and_get-12                    300           4332316 ns/op          166434 B/op      10164 allocs/op
BenchmarkRadix/go-art_insert_and_get-12                      200           7110222 ns/op         1416413 B/op     115582 allocs/op
BenchmarkRadix/q-radix_insert_and_get-12                     500           3834489 ns/op          161996 B/op      10048 allocs/op
```

### License

MIT
