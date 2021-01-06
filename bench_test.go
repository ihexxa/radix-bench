package test

import (
	"flag"
	"fmt"
	"testing"

	"github.com/ihexxa/randstr"

	"github.com/armon/go-radix"
	qradix "github.com/ihexxa/q-radix"
	"github.com/kellydunn/go-art"
)

var testSampleCount = flag.Int("n", 100, "how many samples will be generated for each benchmark test, default: 100")

func BenchmarkRadix(b *testing.B) {
	var testKeyVals map[string]string

	fmt.Println("use 'go test -bench=. -args -n=<sample count>' to set custom sample count")

	benchGoRadixGetOnly := func(b *testing.B) {
		r := radix.New()

		for key, val := range testKeyVals {
			r.Insert(key, val)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for key  := range testKeyVals {
				r.Get(key)
			}
		}
	}

	benchGoArtGetOnly := func(b *testing.B) {
		tree := art.NewArtTree()

		for key, val := range testKeyVals {
			tree.Insert([]byte(key), []byte(val))
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for key := range testKeyVals {
				tree.Search([]byte(key))
			}
		}
	}

	benchQRadixGetOnly := func(b *testing.B) {
		tree := qradix.NewRTree()

		for key, val := range testKeyVals {
			tree.Insert(key, val)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for key  := range testKeyVals {
				tree.Get(key)
			}
		}
	}

	benchGoRadixInsertAndGet := func(b *testing.B) {
		r := radix.New()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for key, val := range testKeyVals {
				r.Insert(key, val)
				r.Get(key)
			}
		}
	}

	benchGoArtInsertAndGet := func(b *testing.B) {
		tree := art.NewArtTree()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for key, val := range testKeyVals {
				tree.Insert([]byte(key), []byte(val))
				tree.Search([]byte(key))
			}
		}
	}

	benchQRadixInsertAndGet := func(b *testing.B) {
		tree := qradix.NewRTree()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for key, val := range testKeyVals {
				tree.Insert(key, val)
				tree.Get(key)
			}
		}
	}

	zhRandStr := randstr.NewRandStr([]string{"你","好","世","界","壹","贰","叁","肆","伍","陆","柒","捌","玖","零"}, false, 16)
	ruRandStr := randstr.NewRandStr([]string{"А","Б","В","Г","Д","Е","Ж","З","И","К","Л","М","Н","О","П","Р","С","Т","У","Ф","Х","Ц","Ч","Ш"}, false, 16)
	charTypes := []string{
		"alphabets",
		"nums",
		"alnums",
		"ch chars",
		"ru chars",
	}

	for _, charType := range charTypes {
		// reset testKeyVals
		testKeyVals = map[string]string{}

		var str string
		for i := 0; i < *testSampleCount; i++ {
			switch charType {
			case "alphabets":
				str = zhRandStr.Alphabets()
			case "nums":
				str = zhRandStr.Numbers()
			case "alnums":
				str = zhRandStr.Alnums()
			case "ch chars":
				str = zhRandStr.Gen()
			case "ru chars":
				str = ruRandStr.Gen()
			}
			fmt.Println(str)
			testKeyVals[str] = str
		}

		b.Run(fmt.Sprintf("go-radix get only (characters=%s)", charType), benchGoRadixGetOnly)
		b.Run(fmt.Sprintf("go-art get only (characters=%s)", charType), benchGoArtGetOnly)
		b.Run(fmt.Sprintf("q-radix get only (characters=%s)", charType), benchQRadixGetOnly)

		b.Run(fmt.Sprintf("go-radix insert and get (characters=%s)", charType), benchGoRadixInsertAndGet)
		b.Run(fmt.Sprintf("go-art insert and get (characters=%s)", charType), benchGoArtInsertAndGet)
		b.Run(fmt.Sprintf("q-radix insert and get (characters=%s)", charType), benchQRadixInsertAndGet)
	}
}
