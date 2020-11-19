package test

import (
	"flag"
	"fmt"
	"math/rand"
	"testing"

	"github.com/armon/go-radix"
	qradix "github.com/ihexxa/q-radix"
	"github.com/kellydunn/go-art"
)

var testSampleCount = flag.Int("n", 1000, "how many samples will be generated for each benchmark test, default: 1000")

func BenchmarkRadix(b *testing.B) {
	var testKeyVals = map[string]string{}
	for i := 0; i < *testSampleCount; i++ {
		str := fmt.Sprintf("%d", rand.Int())
		testKeyVals[str] = str
	}

	fmt.Printf("each benchmark case contains %d time insert/get\n", len(testKeyVals))
	fmt.Println("use 'go test -bench=. -args -n=<sample count>' to set custom sample count")

	benchGoRadixInsert := func(b *testing.B) {
		r := radix.New()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for key, val := range testKeyVals {
				r.Insert(key, val)
				r.Get(key)
			}
		}
	}

	benchGoArtInsert := func(b *testing.B) {
		tree := art.NewArtTree()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for key, val := range testKeyVals {
				tree.Insert([]byte(key), []byte(val))
				tree.Search([]byte(key))
			}
		}
	}

	benchQRadixInsert := func(b *testing.B) {
		tree := qradix.NewRTree()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for key, val := range testKeyVals {
				tree.Insert(key, val)
				tree.Get(key)
			}
		}
	}

	b.Run("go-radix insert and get", benchGoRadixInsert)
	b.Run("go-art insert and get", benchGoArtInsert)
	b.Run("q-radix insert and get", benchQRadixInsert)
}
