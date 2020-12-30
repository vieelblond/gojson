package gen

import (
	"go/types"
	"strings"

	"github.com/go-fish/gojson/option"
)

func (b *Builder) gStructEncodeWarp(fn string, obj *types.Struct, opt *option.Option) error {
	sn := strings.ToLower(fn[:1])
	b.line("func (%s *%s) MarshalJSON() ([]byte, error) {", sn, fn)
	b.line("enc := backend.NewEncoder()")
	b.line("")
	b.gStructEncode(sn, new(FieldTag), obj, opt)
	b.line("data := enc.Bytes()")
	b.line("enc.Release()")
	b.line("")
	b.line("return data, nil")
	b.line("}")
	b.line("")
	return nil
}

func (b *Builder) gStructDecodeWarp(fn string, obj *types.Struct, opt *option.Option) error {
	sn := strings.ToLower(fn[:1])
	b.line("func (%s *%s) UnmarshalJSON(data []byte) error {", sn, fn)
	b.line("if len(data) == 0 {")
	b.line("return nil")
	b.line("}")
	b.line("")
	b.line("dec := backend.NewDecoder()")

	if opt.Unsafe {
		b.line("dec.SetUnsafeData(data)")
	} else {
		b.line("dec.SetData(data)")
	}

	b.line("")
	b.gStructDecode(sn, new(FieldTag), obj, opt)
	b.line("dec.Release()")
	b.line("")
	b.line("return nil")
	b.line("}")
	b.line("")
	return nil
}
