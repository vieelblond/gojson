package benchmark

import (
	"testing"

	backend "github.com/go-fish/gojson/backend"
)

func BenchmarkNeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := []byte("  \n\r\n\t\n    ,\t\r\n 1")
		decoder := backend.NewDecoder()
		decoder.SetData(data)
		decoder.Need('1')
		decoder.Release()
	}
}
