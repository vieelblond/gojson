package gen

import (
	"testing"

	"github.com/go-fish/gojson/option"
)

func TestGenerate(t *testing.T) {
	opt, err := option.NewOption()
	if err != nil {
		panic(err)
	}

	opt.Input = "../benchmark/test/test.go"
	opt.Output = "../benchmark/test/test.generate.go"
	opt.Mode = 4
	opt.Unsafe = true

	if err := opt.ParsePackage(); err != nil {
		panic(err)
	}

	if err := Generate(opt); err != nil {
		panic(err)
	}
}
