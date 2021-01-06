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

**Note**: This benchmark just provides a quick reference of common features' performance.

**Case 1: Insert 100 strings and get them:**

```
BenchmarkRadix/go-radix_insert_and_get_(characters=alphabets)-12           74684             16060 ns/op            1600 B/op        100 allocs/op
BenchmarkRadix/go-art_insert_and_get_(characters=alphabets)-12             51928             23727 ns/op            7792 B/op        574 allocs/op
BenchmarkRadix/q-radix_insert_and_get_(characters=alphabets)-12            55930             22213 ns/op            1600 B/op        100 allocs/op
BenchmarkRadix/go-radix_insert_and_get_(characters=nums)-12                63859             19206 ns/op            1600 B/op        100 allocs/op
BenchmarkRadix/go-art_insert_and_get_(characters=nums)-12                  30204             37734 ns/op            8960 B/op        720 allocs/op
BenchmarkRadix/q-radix_insert_and_get_(characters=nums)-12                 75013             16754 ns/op            1600 B/op        100 allocs/op
BenchmarkRadix/go-radix_insert_and_get_(characters=alnums)-12              74883             16010 ns/op            1600 B/op        100 allocs/op
BenchmarkRadix/go-art_insert_and_get_(characters=alnums)-12                52698             22619 ns/op            7680 B/op        560 allocs/op
BenchmarkRadix/q-radix_insert_and_get_(characters=alnums)-12               52933             23385 ns/op            1600 B/op        100 allocs/op
BenchmarkRadix/go-radix_insert_and_get_(characters=ch_chars)-12            46417             25614 ns/op            1600 B/op        100 allocs/op
BenchmarkRadix/go-art_insert_and_get_(characters=ch_chars)-12              19676             58943 ns/op           43217 B/op       1402 allocs/op
BenchmarkRadix/q-radix_insert_and_get_(characters=ch_chars)-12             56268             21150 ns/op            1600 B/op        100 allocs/op
BenchmarkRadix/go-radix_insert_and_get_(characters=ru_chars)-12            54969             22031 ns/op            1600 B/op        100 allocs/op
BenchmarkRadix/go-art_insert_and_get_(characters=ru_chars)-12              33145             36015 ns/op           22624 B/op        828 allocs/op
BenchmarkRadix/q-radix_insert_and_get_(characters=ru_chars)-12             55494             21235 ns/op            1600 B/op        100 allocs/op
```

**Case 2: Get 100 strings which are existing in the tree:**

```
BenchmarkRadix/go-radix_get_only_(characters=alphabets)-12                181988              5849 ns/op               0 B/op          0 allocs/op
BenchmarkRadix/go-art_get_only_(characters=alphabets)-12                  129264              8944 ns/op            1496 B/op        187 allocs/op
BenchmarkRadix/q-radix_get_only_(characters=alphabets)-12                 117786             10474 ns/op               0 B/op          0 allocs/op
BenchmarkRadix/go-radix_get_only_(characters=nums)-12                     163506              6957 ns/op               0 B/op          0 allocs/op
BenchmarkRadix/go-art_get_only_(characters=nums)-12                        77432             15666 ns/op            2080 B/op        260 allocs/op
BenchmarkRadix/q-radix_get_only_(characters=nums)-12                      175273              6687 ns/op               0 B/op          0 allocs/op
BenchmarkRadix/go-radix_get_only_(characters=alnums)-12                   201602              5691 ns/op               0 B/op          0 allocs/op
BenchmarkRadix/go-art_get_only_(characters=alnums)-12                     127227              8870 ns/op            1440 B/op        180 allocs/op
BenchmarkRadix/q-radix_get_only_(characters=alnums)-12                    111726             10870 ns/op               0 B/op          0 allocs/op
BenchmarkRadix/go-radix_get_only_(characters=ch_chars)-12                 127326              9555 ns/op               0 B/op          0 allocs/op
BenchmarkRadix/go-art_get_only_(characters=ch_chars)-12                    44758             26536 ns/op           17608 B/op        601 allocs/op
BenchmarkRadix/q-radix_get_only_(characters=ch_chars)-12                  119182              9418 ns/op               0 B/op          0 allocs/op
BenchmarkRadix/go-radix_get_only_(characters=ru_chars)-12                 157010              7356 ns/op               0 B/op          0 allocs/op
BenchmarkRadix/go-art_get_only_(characters=ru_chars)-12                    79587             14909 ns/op            8112 B/op        314 allocs/op
BenchmarkRadix/q-radix_get_only_(characters=ru_chars)-12                  116301              9787 ns/op               0 B/op          0 allocs/op
```
### License

MIT
