package gen

import (
	"fmt"
	"go/types"

	"github.com/go-fish/gojson/option"
	"github.com/go-fish/gojson/util"
)

func (b *Builder) gArrayEncode(fn string, obj *types.Array, opt *option.Option) {
	b.line("if len(%s) == 0 {", fn)
	b.line("enc.WriteNull()")
	b.line("} else {")

	value := util.GenerateID("value")

	b.line("enc.WriteByte('[')")
	b.line("for _, %s := range %s {", value, fn)
	b.line("enc.WriteComma()")

	switch x := obj.Elem().Underlying().(type) {
	case *types.Struct:
		b.gStructEncode(value, new(FieldTag), x, opt)

	case *types.Map:
		b.gMapEncode(value, x, opt)

	case *types.Array:
		b.gArrayEncode(value, x, opt)

	case *types.Slice:
		b.gSliceEncode(value, x, opt)

	case *types.Pointer:
		b.gPointerEncode(value, new(FieldTag), x, opt)

	case *types.Interface:
		b.line("if err := enc.EncodeValue(%s); err != nil {", value)
		b.line("return nil, err")
		b.line("}")

	case *types.Basic:
		alias := value
		if _, ok := obj.Elem().(*types.Named); ok {
			alias = fmt.Sprintf("%s(%s)", x.Name(), value)
		}

		switch x.Kind() {
		case types.String:
			b.line("enc.EncodeString(%s)", alias)

		case types.Int:
			b.line("enc.EncodeInt(%s)", alias)

		case types.Int8:
			b.line("enc.EncodeInt8(%s)", alias)

		case types.Int16:
			b.line("enc.EncodeInt16(%s)", alias)

		case types.Int32:
			b.line("enc.EncodeInt32(%s)", alias)

		case types.Int64:
			b.line("enc.EncodeInt64(%s)", alias)

		case types.Uint:
			b.line("enc.EncodeUint(%s)", alias)

		case types.Uint8:
			b.line("enc.EncodeUint8(%s)", alias)

		case types.Uint16:
			b.line("enc.EncodeUint16(%s)", alias)

		case types.Uint32:
			b.line("enc.EncodeUint32(%s)", alias)

		case types.Uint64:
			b.line("enc.EncodeUint64(%s)", alias)

		case types.Float32:
			b.line("enc.EncodeFloat32(%s)", alias)

		case types.Float64:
			b.line("enc.EncodeFloat64(%s)", alias)

		case types.Bool:
			b.line("enc.EncodeBool(%s)", alias)
		}

	default:
		b.line("if err := enc.EncodeValue(%s); err != nil {", value)
		b.line("return nil, err")
		b.line("}")
	}

	b.line("}")
	b.line("")
	b.line("enc.WriteByte(']')")
	b.line("}")
}

func (b *Builder) gArrayDecode(fn string, obj *types.Array, opt *option.Option) {
	b.line("if dec.IsNull() {")
	b.line("%s = nil", fn)
	b.line("} else if !dec.IsArrayOpen() {")
	b.line("return errors.NewParseError(dec.Char(), dec.Cursor())")
	b.line("} else {")

	// empty array
	b.line("if dec.IsArrayClose() {")
	b.line("%s = nil", fn)
	b.line("} else {")

	index := util.GenerateID("index")
	array := util.GenerateID("array")
	b.line("%s := 1", array)
	b.line("%s := 0", index)
	b.line("for %s > 0 {", array)
	b.line("if %s < %d {", index, obj.Len())

	value := util.GenerateID("value")

	switch x := obj.Elem().Underlying().(type) {
	case *types.Struct:
		b.line("var %s %s", value, b.typeString(obj.Elem(), opt))
		b.gStructDecode(value, new(FieldTag), x, opt)
		b.line("%s[%s] = %s", fn, index, value)

	case *types.Map:
		b.gMapDecode(value, x, opt)
		b.line("%s[%s] = %s", fn, index, value)

	case *types.Array:
		b.gArrayDecode(value, x, opt)
		b.line("%s[%s] = %s", fn, index, value)

	case *types.Slice:
		b.gSliceDecode(value, x, opt)
		b.line("%s[%s] = %s", fn, index, value)

	case *types.Pointer:
		b.line("var %s %s", value, b.typeString(x, opt))
		b.gPointerDecode(value, new(FieldTag), x, opt)
		b.line("if %s != nil {", value)
		b.line("%s[%s] = %s", fn, index, value)
		b.line("}")

	case *types.Interface:
		b.line("%s, err := dec.DecodeValue()", value)
		b.line("if err != nil {")
		b.line("return err")
		b.line("}")
		b.line("")
		b.line("%s[%s] = %s", fn, index, value)

	case *types.Basic:
		elem := value
		if typ := b.typeString(obj.Elem(), opt); typ != x.Name() {
			elem = fmt.Sprintf("%s(%s)", elem, value)
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
		b.line("%s[%s] = %s", fn, index, elem)

	default:
		value := util.GenerateID("value")
		b.line("%s, err := dec.DecodeValue()", value)
		b.line("if err != nil {")
		b.line("return err")
		b.line("}")
		b.line("")
		b.line("%s[%s] = %s", fn, index, value)
	}

	b.line("} else {")
	b.line("return fmt.Errorf(\"index out of range at pos %%d\", dec.Cursor())")
	b.line("}")
	b.line("if dec.IsArrayClose() {")
	b.line("%s--", array)
	b.line("}")
	b.line("}")
	b.line("}")
	b.line("}")
}

func (b *Builder) gSliceEncode(fn string, obj *types.Slice, opt *option.Option) {
	b.line("if len(%s) == 0 {", fn)
	b.line("enc.WriteNull()")
	b.line("} else {")

	// hack []byte
	if basic, ok := obj.Elem().Underlying().(*types.Basic); ok && basic.Kind() == types.Byte {
		b.line("enc.EncodeBytes(%s)", fn)
		return
	}

	value := util.GenerateID("value")

	b.line("enc.WriteByte('[')")
	b.line("for _, %s := range %s {", value, fn)
	b.line("enc.WriteComma()")

	switch x := obj.Elem().Underlying().(type) {
	case *types.Struct:
		b.gStructEncode(value, new(FieldTag), x, opt)

	case *types.Map:
		b.gMapEncode(value, x, opt)

	case *types.Array:
		b.gArrayEncode(value, x, opt)

	case *types.Slice:
		b.gSliceEncode(value, x, opt)

	case *types.Pointer:
		b.gPointerEncode(value, new(FieldTag), x, opt)

	case *types.Interface:
		b.line("if err := enc.EncodeValue(%s); err != nil {", value)
		b.line("return nil, err")
		b.line("}")

	case *types.Basic:
		alias := value
		if _, ok := obj.Elem().(*types.Named); ok {
			alias = fmt.Sprintf("%s(%s)", x.Name(), value)
		}

		switch x.Kind() {
		case types.String:
			b.line("enc.EncodeString(%s)", alias)

		case types.Int:
			b.line("enc.EncodeInt(%s)", alias)

		case types.Int8:
			b.line("enc.EncodeInt8(%s)", alias)

		case types.Int16:
			b.line("enc.EncodeInt16(%s)", alias)

		case types.Int32:
			b.line("enc.EncodeInt32(%s)", alias)

		case types.Int64:
			b.line("enc.EncodeInt64(%s)", alias)

		case types.Uint:
			b.line("enc.EncodeUint(%s)", alias)

		case types.Uint8:
			b.line("enc.EncodeUint8(%s)", alias)

		case types.Uint16:
			b.line("enc.EncodeUint16(%s)", alias)

		case types.Uint32:
			b.line("enc.EncodeUint32(%s)", alias)

		case types.Uint64:
			b.line("enc.EncodeUint64(%s)", alias)

		case types.Float32:
			b.line("enc.EncodeFloat32(%s)", alias)

		case types.Float64:
			b.line("enc.EncodeFloat64(%s)", alias)

		case types.Bool:
			b.line("enc.EncodeBool(%s)", alias)
		}

	default:
		b.line("if err := enc.EncodeValue(%s); err != nil {", value)
		b.line("return nil, err")
		b.line("}")
	}

	b.line("}")
	b.line("")
	b.line("enc.WriteByte(']')")
	b.line("}")
}

func (b *Builder) gSliceDecode(fn string, x types.Type, opt *option.Option) {
	if obj, _ := x.Underlying().(*types.Slice); obj != nil {
		// hack []byte
		if basic, ok := obj.Elem().Underlying().(*types.Basic); ok && basic.Kind() == types.Byte {
			b.line("if dec.IsNull() {")
			b.line("%s = nil", fn)
			b.line(" } else {")
			value := util.GenerateID("value")
			b.line("%s, err := dec.DecodeBytes()", value)
			b.line("if err != nil {")
			b.line("return err")
			b.line("}")
			b.line("")
			b.line("%s = %s", fn, value)
			return
		}

		// b.line("if dec.IsNull() {")
		// b.line("%s = nil", fn)
		// b.line("} else if !dec.IsArrayOpen() {")
		// b.line("return errors.NewParseError(dec.Char(), dec.Cursor())")
		b.line("if char := dec.NextChar(); char == 'n' {")
		b.line("if err := dec.AssetNull(); err != nil {")
		b.line("return err")
		b.line("}")
		b.line("")
		b.line("%s = nil", fn)
		b.line("} else if char != '[' {")
		b.line("return errors.NewParseError(dec.Char(), dec.Cursor())")
		b.line("} else {")
		b.line("dec.Next()")

		// empty slice
		b.line("if dec.IsArrayClose() {")
		b.line("%s = nil", fn)
		b.line("} else {")

		// initialize slice
		b.line("if %s == nil {", fn)
		b.line("%s = make(%s, 0, 8)", fn, b.typeString(x, opt))
		b.line("}")
		b.line("")

		array := util.GenerateID("array")
		b.line("for %s := 1; %s > 0; {", array, array)

		value := util.GenerateID("value")

		switch x := obj.Elem().Underlying().(type) {
		case *types.Struct:
			b.line("var %s %s", value, b.typeString(obj.Elem(), opt))
			b.gStructDecode(value, new(FieldTag), x, opt)
			b.line("%s = append(%s, %s)", fn, fn, value)

		case *types.Map:
			b.gMapDecode(value, x, opt)
			b.line("%s = append(%s, %s)", fn, fn, value)

		case *types.Array:
			b.gArrayDecode(value, x, opt)
			b.line("%s = append(%s, %s)", fn, fn, value)

		case *types.Slice:
			b.gSliceDecode(value, x, opt)
			b.line("%s = append(%s, %s)", fn, fn, value)

		case *types.Pointer:
			b.line("var %s %s", value, b.typeString(x, opt))
			b.gPointerDecode(value, new(FieldTag), x, opt)
			b.line("if %s != nil {", value)
			b.line("%s = append(%s, %s)", fn, fn, value)
			b.line("}")

		case *types.Interface:
			b.line("%s, err := dec.DecodeValue()", value)
			b.line("if err != nil {")
			b.line("return err")
			b.line("}")
			b.line("")
			b.line("%s = append(%s, %s)", fn, fn, value)

		case *types.Basic:
			elem := value
			if typ := b.typeString(obj.Elem(), opt); typ != x.Name() {
				elem = fmt.Sprintf("%s(%s)", elem, value)
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
			b.line("%s = append(%s, %s)", fn, fn, elem)

		default:
			value := util.GenerateID("value")
			b.line("%s, err := dec.DecodeValue()", value)
			b.line("if err != nil {")
			b.line("return err")
			b.line("}")
			b.line("")
			b.line("%s = append(%s, %s)", fn, fn, value)
		}

		b.line("if dec.IsArrayClose() {")
		b.line("%s--", array)
		b.line("}")
		b.line("}")
		b.line("}")
		b.line("}")
	}
}
