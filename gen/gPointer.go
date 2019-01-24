package gen

import (
	"fmt"
	"go/types"

	"github.com/go-fish/gojson/option"
	"github.com/go-fish/gojson/util"
)

func (b *Builder) gPointerEncode(fn string, self *FieldTag, obj *types.Pointer, opt *option.Option) {
	b.line("if %s == nil {", fn)
	b.line("enc.WriteNull()")
	b.line("} else {")

	switch x := obj.Elem().Underlying().(type) {
	case *types.Struct:
		b.gStructEncode(fn, self, x, opt)

	case *types.Map:
		b.gMapEncode(fn, x, opt)

	case *types.Array:
		b.gArrayEncode(fn, x, opt)

	case *types.Slice:
		b.gSliceEncode(fn, x, opt)

	case *types.Pointer:
		b.gPointerEncode(fn, self, x, opt)

	case *types.Interface:
		b.line("if err := enc.EncodeValue(%s); err != nil {", fn)
		b.line("return nil, err")
		b.line("}")

	case *types.Basic:
		alias := fn
		if _, ok := obj.Elem().(*types.Named); ok {
			alias = fmt.Sprintf("%s(%s)", x.Name(), fn)
		}

		switch x.Kind() {
		case types.String:
			b.line("enc.EncodeKeyString(%q, %s)", self.name, alias)

		case types.Int:
			b.line("enc.EncodeKeyInt(%q, %s)", self.name, alias)

		case types.Int8:
			b.line("enc.EncodeKeyInt8(%q, %s)", self.name, alias)

		case types.Int16:
			b.line("enc.EncodeKeyInt16(%q, %s)", self.name, alias)

		case types.Int32:
			b.line("enc.EncodeKeyInt32(%q, %s)", self.name, alias)

		case types.Int64:
			b.line("enc.EncodeKeyInt64(%q, %s)", self.name, alias)

		case types.Uint:
			b.line("enc.EncodeKeyUint(%q, %s)", self.name, alias)

		case types.Uint8:
			b.line("enc.EncodeKeyUint8(%q, %s)", self.name, alias)

		case types.Uint16:
			b.line("enc.EncodeKeyUint16(%q, %s)", self.name, alias)

		case types.Uint32:
			b.line("enc.EncodeKeyUint32(%q, %s)", self.name, alias)

		case types.Uint64:
			b.line("enc.EncodeKeyUint64(%q, %s)", self.name, alias)

		case types.Float32:
			b.line("enc.EncodeKeyFloat32(%q, %s)", self.name, alias)

		case types.Float64:
			b.line("enc.EncodeKeyFloat64(%q, %s)", self.name, alias)

		case types.Bool:
			b.line("enc.EncodeKeyBool(%q, %s)", self.name, alias)
		}

	default:
		b.line("if err := enc.EncodeKeyValue(%q, %s); err != nil {", self.name, fn)
		b.line("return nil, err")
		b.line("}")
	}

	b.line("}")
}

func (b *Builder) gPointerDecode(fn string, self *FieldTag, obj *types.Pointer, opt *option.Option) {
	b.line("if dec.IsNull() {")
	b.line("%s = nil", fn)
	b.line("} else {")
	b.line("%s = new(%s)", fn, b.typeString(obj.Elem(), opt))
	b.line("")
	switch x := obj.Elem().Underlying().(type) {
	case *types.Struct:
		b.gStructDecode(fn, self, x, opt)

	case *types.Map:
		b.gMapDecode(fn, x, opt)

	case *types.Array:
		b.gArrayDecode(fn, x, opt)

	case *types.Slice:
		b.gSliceDecode(fn, x, opt)

	case *types.Pointer:
		b.gPointerDecode(fn, self, x, opt)

	case *types.Interface:
		value := util.GenerateID("value")
		b.line("%s, err := dec.DecodeValue()", value)
		b.line("if err != nil {")
		b.line("return err")
		b.line("}")
		b.line("")
		b.line("%s = &(%s)", fn, value)

	case *types.Basic:
		value := util.GenerateID("value")
		alias := value
		if typ := b.typeString(obj.Elem(), opt); typ != x.Name() {
			alias = fmt.Sprintf("%s(%s)", typ, value)
		}

		switch x.Kind() {
		case types.String:
			b.line("%s, err := dec.DecodeString()", value)

		case types.Int:
			b.line("%s, err := dec.DecodeInt()", value)

		case types.Int8:
			b.line("%s, err := dec.DecodeInt8()", value)

		case types.Int16:
			b.line("%s, err := dec.DecodeInt16()", value)

		case types.Int32:
			b.line("%s, err := dec.DecodeInt32()", value)

		case types.Int64:
			b.line("%s, err := dec.DecodeInt64()", value)

		case types.Uint:
			b.line("%s, err := dec.DecodeUint()", value)

		case types.Uint8:
			b.line("%s, err := dec.DecodeUint8()", value)

		case types.Uint16:
			b.line("%s, err := dec.DecodeUint16()", value)

		case types.Uint32:
			b.line("%s, err := dec.DecodeUint32()", value)

		case types.Uint64:
			b.line("%s, err := dec.DecodeUint64()", value)

		case types.Float32:
			b.line("%s, err := dec.DecodeFloat32()", value)

		case types.Float64:
			b.line("%s, err := dec.DecodeFloat64()", value)

		case types.Bool:
			b.line("%s, err := dec.DecodeBool()", value)
		}

		b.line("if err != nil {")
		b.line("return err")
		b.line("}")
		b.line("")

		if alias != value {
			v := util.GenerateID("value")
			b.line("%s := %s", v, alias)
			b.line("%s = &(%s)", fn, v)
		} else {
			b.line("%s = &(%s)", fn, value)
		}

	default:
		value := util.GenerateID("value")
		b.line("%s, err := dec.DecodeValue()", value)
		b.line("if err != nil {")
		b.line("return err")
		b.line("}")
		b.line("")
		b.line("%s = &(%s)", fn, value)
	}

	b.line("}")
}
