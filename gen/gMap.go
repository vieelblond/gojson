package gen

import (
	"fmt"
	"go/types"
	"strings"

	"github.com/go-fish/gojson/option"
	"github.com/go-fish/gojson/util"
)

func (b *Builder) gMapEncode(fn string, obj *types.Map, opt *option.Option) {
	b.line("if len(%s) == 0 {", fn)
	b.line("enc.WriteNull()")
	b.line("} else {")

	key := util.GenerateID("key")
	value := util.GenerateID("value")

	b.line("enc.WriteByte('{')")
	b.line("for %s, %s := range %s {", key, value, fn)

	switch x := obj.Elem().Underlying().(type) {
	case *types.Struct:
		b.line("enc.WriteKey(%s)", key)
		b.gStructEncode(value, new(FieldTag), x, opt)

	case *types.Map:
		b.line("enc.WriteKey(%s)", key)
		b.gMapEncode(value, x, opt)

	case *types.Array:
		b.line("enc.WriteKey(%s)", key)
		b.gArrayEncode(value, x, opt)

	case *types.Slice:
		b.line("enc.WriteKey(%s)", key)
		b.gSliceEncode(value, x, opt)

	case *types.Pointer:
		b.line("enc.WriteKey(%s)", key)
		b.line("if %s != nil {", value)
		b.gPointerEncode(value, new(FieldTag), x, opt)
		b.line("}")

	case *types.Interface:
		b.line("if err := enc.EncodeKeyValue(%s, %s); err != nil {", key, value)
		b.line("return nil, err")
		b.line("}")

	case *types.Basic:
		alias := value
		if typ := b.typeString(obj.Elem(), opt); typ != x.Name() {
			alias = fmt.Sprintf("%s(%s)", typ, value)
		}

		switch x.Kind() {
		case types.String:
			b.line("enc.EncodeKeyString(%s, %s)", key, alias)

		case types.Int:
			b.line("enc.EncodeKeyInt(%s, %s)", key, alias)

		case types.Int8:
			b.line("enc.EncodeKeyInt8(%s, %s)", key, alias)

		case types.Int16:
			b.line("enc.EncodeKeyInt16(%s, %s)", key, alias)

		case types.Int32:
			b.line("enc.EncodeKeyInt32(%s, %s)", key, alias)

		case types.Int64:
			b.line("enc.EncodeKeyInt64(%s, %s)", key, alias)

		case types.Uint:
			b.line("enc.EncodeKeyUint(%s, %s)", key, alias)

		case types.Uint8:
			b.line("enc.EncodeKeyUint8(%s, %s)", key, alias)

		case types.Uint16:
			b.line("enc.EncodeKeyUint16(%s, %s)", key, alias)

		case types.Uint32:
			b.line("enc.EncodeKeyUint32(%s, %s)", key, alias)

		case types.Uint64:
			b.line("enc.EncodeKeyUint64(%s, %s)", key, alias)

		case types.Float32:
			b.line("enc.EncodeKeyFloat32(%s, %s)", key, alias)

		case types.Float64:
			b.line("enc.EncodeKeyFloat64(%s, %s)", key, alias)

		case types.Bool:
			b.line("enc.EncodeKeyBool(%s, %s)", key, alias)
		}

	default:
		b.line("if err := enc.EncodeKeyValue(%s, %s); err != nil {", key, value)
		b.line("return nil, err")
		b.line("}")
	}

	b.line("}")
	b.line("")
	b.line("enc.WriteByte('}')")
	b.line("}")
}

func (b *Builder) gMapDecode(fn string, x types.Type, opt *option.Option) {
	if obj, _ := x.Underlying().(*types.Map); obj != nil {
		b.line("if dec.IsNull() {")
		b.line("%s = nil", fn)
		b.line("} else if !dec.IsObjectOpen() {")
		b.line("return errors.NewParseError(dec.Char(), dec.Cursor())")
		b.line("} else {")

		// empty map
		b.line("if dec.IsObjectClose() {")
		b.line("%s = nil", fn)
		b.line("} else {")

		// initialize map
		b.line("if %s == nil {", fn)
		b.line("%s = make(%s)", fn, b.typeString(x, opt))
		b.line("}")

		object := util.GenerateID("obj")
		b.line("for %s := 1; %s > 0; {", object, object)

		key := util.GenerateID("key")
		value := util.GenerateID("value")

		b.line("%s, err := dec.NextKey()", key)
		b.line("if err != nil {")
		b.line("return err")
		b.line("}")
		b.line("")

		alias := key
		if typ := b.typeString(obj.Key(), opt); typ != "string" {
			alias = fmt.Sprintf("%s(%s)", typ, key)
		}

		switch x := obj.Elem().Underlying().(type) {
		case *types.Struct:
			b.line("var %s %s", value, b.typeString(obj.Elem(), opt))
			b.gStructDecode(value, new(FieldTag), x, opt)
			b.line("%s[%s] = %s", fn, alias, value)

		case *types.Array:
			b.gArrayDecode(value, x, opt)
			b.line("%s[%s] = %s", fn, alias, value)

		case *types.Slice:
			b.gSliceDecode(value, x, opt)
			b.line("%s[%s] = %s", fn, alias, value)

		case *types.Pointer:
			b.line("var %s %s", value, b.typeString(x, opt))
			b.gPointerDecode(value, new(FieldTag), x, opt)
			b.line("if %s != nil {", value)
			b.line("%s[%s] = %s", fn, alias, value)
			b.line("}")

		case *types.Interface:
			b.line("%s, err := dec.DecodeValue()", value)
			b.line("if err != nil {")
			b.line("return err")
			b.line("}")
			b.line("")
			b.line("%s[%s] = %s", fn, alias, value)

		case *types.Basic:
			elem := value
			if typ := b.typeString(obj.Elem(), opt); typ != x.Name() {
				elem = fmt.Sprintf("%s(%s)", typ, value)
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
			b.line("%s[%s] = %s", fn, alias, elem)

		default:
			value := util.GenerateID("value")
			b.line("%s, err := dec.DecodeValue()", value)
			b.line("if err != nil {")
			b.line("return err")
			b.line("}")
			b.line("")
			b.line("%s[%s] = %s", fn, alias, value)
		}

		b.line("if dec.IsObjectClose() {")
		b.line("%s--", object)
		b.line("}")

		b.line("}")
		b.line("}")
		b.line("}")
	}
}

func (b *Builder) gFieldEncode(fn string, parent, self *FieldTag, obj *types.Struct, field *types.Var, opt *option.Option) error {
	if b.needPrint(parent, self) {

		// add import
		if !opt.IsLocal(field.Pkg()) {
			pkg := field.Pkg().Path()
			if index := strings.Index(pkg, "/vendor/"); index > 0 {
				pkg = pkg[index+8:]
			}

			b.Imports[field.Pkg().Name()] = pkg
		}

		switch x := field.Type().Underlying().(type) {
		case *types.Struct:
			fn = fmt.Sprintf("%s.%s", fn, field.Name())

			if !self.inline && ((!opt.Inline && opt.IsLocal(field.Pkg())) || opt.IsMarshaler(field.Type())) {
				b.line("enc.WriteKey(%q)", self.name)
				tmpData := util.GenerateID("data")
				b.line("%s, err := %s.MarshalJSON()", tmpData, fn)
				b.line("if err != nil {")
				b.line("return nil, err")
				b.line("}")
				b.line("")
				b.line("enc.WriteBytes(%s)", tmpData)
			} else {
				if !self.inline {
					b.line("enc.WriteKey(%q)", self.name)
					b.line("enc.WriteByte('{')")
					self.keys = nil
				} else {
					self.keys = b.getKeys(obj)
				}

				b.gStructEncode(fn, self, x, opt)

				if !self.inline {
					b.line("enc.WriteByte('}')")
				}
			}

		case *types.Map:
			fn = fmt.Sprintf("%s.%s", fn, field.Name())

			if self.omitempty {
				b.line("if len(%s) > 0 {", fn)
				b.line("enc.WriteKey(%q)", self.name)
				b.gMapEncode(fn, x, opt)
				b.line("}")
			} else {
				b.line("enc.WriteKey(%q)", self.name)
				b.gMapEncode(fn, x, opt)
			}

		case *types.Array:
			fn = fmt.Sprintf("%s.%s", fn, field.Name())

			if self.omitempty {
				b.line("if len(%s) > 0 {", fn)
				b.line("enc.WriteKey(%q)", self.name)
				b.gArrayEncode(fn, x, opt)
				b.line("}")
			} else {
				b.line("enc.WriteKey(%q)", self.name)
				b.gArrayEncode(fn, x, opt)
			}

		case *types.Slice:
			fn = fmt.Sprintf("%s.%s", fn, field.Name())

			if self.omitempty {
				b.line("if len(%s) > 0 {", fn)
				b.line("enc.WriteKey(%q)", self.name)
				b.gSliceEncode(fn, x, opt)
				b.line("}")
			} else {
				b.line("enc.WriteKey(%q)", self.name)
				b.gSliceEncode(fn, x, opt)
			}

		case *types.Pointer:
			if !self.inline {
				self.keys = nil
			} else {
				self.keys = b.getKeys(obj)
			}

			f := types.NewVar(field.Pos(), field.Pkg(), field.Name(), x.Elem())
			b.line("if %s.%s != nil {", fn, field.Name())
			b.gFieldEncode(fn, parent, self, obj, f, opt)
			b.line("}")

		case *types.Interface:
			fn = fmt.Sprintf("%s.%s", fn, field.Name())

			if self.omitempty {
				b.line("if %s != nil {", fn)
				b.line("if err := enc.EncodeKeyValue(%q, %s); err != nil {", self.name, fn)
				b.line("return nil, err")
				b.line("}")
				b.line("}")
			} else {
				b.line("if err := enc.EncodeKeyValue(%q, %s); err != nil {", self.name, fn)
				b.line("return nil, err")
				b.line("}")
			}

		case *types.Basic:
			fn = fmt.Sprintf("%s.%s", fn, field.Name())
			if typ := b.typeString(field.Type(), opt); typ != x.Name() {
				fn = fmt.Sprintf("%s(%s)", x.Name(), fn)
			}

			switch x.Kind() {
			case types.String:
				if self.omitempty {
					b.line("if %s != \"\" {", fn)
					b.line("enc.EncodeKeyString(%q, %s)", self.name, fn)
					b.line("}")
				} else {
					b.line("enc.EncodeKeyString(%q, %s)", self.name, fn)
				}

			case types.Int:
				if self.omitempty {
					b.line("if %s != 0 {", fn)
					b.line("enc.EncodeKeyInt(%q, %s)", self.name, fn)
					b.line("}")
				} else {
					b.line("enc.EncodeKeyInt(%q, %s)", self.name, fn)
				}

			case types.Int8:
				if self.omitempty {
					b.line("if %s != 0 {", fn)
					b.line("enc.EncodeKeyInt8(%q, %s)", self.name, fn)
					b.line("}")
				} else {
					b.line("enc.EncodeKeyInt8(%q, %s)", self.name, fn)
				}

			case types.Int16:
				if self.omitempty {
					b.line("if %s != 0 {", fn)
					b.line("enc.EncodeKeyInt16(%q, %s)", self.name, fn)
					b.line("}")
				} else {
					b.line("enc.EncodeKeyInt16(%q, %s)", self.name, fn)
				}

			case types.Int32:
				if self.omitempty {
					b.line("if %s != 0 {", fn)
					b.line("enc.EncodeKeyInt32(%q, %s)", self.name, fn)
					b.line("}")
				} else {
					b.line("enc.EncodeKeyInt32(%q, %s)", self.name, fn)
				}

			case types.Int64:
				if self.omitempty {
					b.line("if %s != 0 {", fn)
					b.line("enc.EncodeKeyInt64(%q, %s)", self.name, fn)
					b.line("}")
				} else {
					b.line("enc.EncodeKeyInt64(%q, %s)", self.name, fn)
				}

			case types.Uint:
				if self.omitempty {
					b.line("if %s != 0 {", fn)
					b.line("enc.EncodeKeyUint(%q, %s)", self.name, fn)
					b.line("}")
				} else {
					b.line("enc.EncodeKeyUint(%q, %s)", self.name, fn)
				}

			case types.Uint8:
				if self.omitempty {
					b.line("if %s != 0 {", fn)
					b.line("enc.EncodeKeyUint8(%q, %s)", self.name, fn)
					b.line("}")
				} else {
					b.line("enc.EncodeKeyUint8(%q, %s)", self.name, fn)
				}

			case types.Uint16:
				if self.omitempty {
					b.line("if %s != 0 {", fn)
					b.line("enc.EncodeKeyUint16(%q, %s)", self.name, fn)
					b.line("}")
				} else {
					b.line("enc.EncodeKeyUint16(%q, %s)", self.name, fn)
				}

			case types.Uint32:
				if self.omitempty {
					b.line("if %s != 0 {", fn)
					b.line("enc.EncodeKeyUint32(%q, %s)", self.name, fn)
					b.line("}")
				} else {
					b.line("enc.EncodeKeyUint32(%q, %s)", self.name, fn)
				}

			case types.Uint64:
				if self.omitempty {
					b.line("if %s != 0 {", fn)
					b.line("enc.EncodeKeyUint64(%q, %s)", self.name, fn)
					b.line("}")
				} else {
					b.line("enc.EncodeKeyUint64(%q, %s)", self.name, fn)
				}

			case types.Float32:
				if self.omitempty {
					b.line("if %s != 0 {", fn)
					b.line("enc.EncodeKeyFloat32(%q, %s)", self.name, fn)
					b.line("}")
				} else {
					b.line("enc.EncodeKeyFloat32(%q, %s)", self.name, fn)
				}

			case types.Float64:
				if self.omitempty {
					b.line("if %s != 0 {", fn)
					b.line("enc.EncodeKeyFloat64(%q, %s)", self.name, fn)
					b.line("}")
				} else {
					b.line("enc.EncodeKeyFloat64(%q, %s)", self.name, fn)
				}

			case types.Bool:
				if self.omitempty {
					b.line("if %s {", fn)
					b.line("enc.EncodeKeyBool(%q, %s)", self.name, fn)
					b.line("}")
				} else {
					b.line("enc.EncodeKeyBool(%q, %s)", self.name, fn)
				}
			}

			b.line("")

		default:
			b.line("if err := enc.EncodeKeyValue(%q, %s); err != nil {", self.name, fn)
			b.line("return nil, err")
			b.line("}")
		}
	}

	return nil
}

func (b *Builder) gStructEncode(fn string, parent *FieldTag, obj *types.Struct, opt *option.Option) {
	if b.isRoot(fn) {
		b.line("enc.WriteByte('{')")
	}

	// generate fields
	for i := 0; i < obj.NumFields(); i++ {
		tag := b.parseFieldTag(obj.Tag(i), obj.Field(i))

		if tag.ignore {
			continue
		}

		b.gFieldEncode(fn, parent, tag, obj, obj.Field(i), opt)
	}

	if b.isRoot(fn) {
		b.line("enc.WriteByte('}')")
	}
}

func (b *Builder) gFieldDecode(fn string, parent, self *FieldTag, obj *types.Struct, field *types.Var, opt *option.Option) {
	if b.needPrint(parent, self) {
		if !opt.IsLocal(field.Pkg()) {
			pkg := field.Pkg().Path()
			if index := strings.Index(pkg, "/vendor/"); index > 0 {
				pkg = pkg[index+8:]
			}

			b.Imports[field.Pkg().Name()] = pkg
		}

		switch x := field.Type().Underlying().(type) {
		case *types.Struct:
			fn = fmt.Sprintf("%s.%s", fn, field.Name())

			if !self.inline && ((!opt.Inline && opt.IsLocal(field.Pkg())) || opt.IsUnmarshaler(field.Type())) {
				b.line("case %q:", self.name)
				data := util.GenerateID("data")
				b.line("%s, err := dec.ReadValue()", data)
				b.line("if err != nil {")
				b.line("return nil")
				b.line("}")
				b.line("")

				// set pointer to nil when read empty data
				if self.pointer {
					b.line("if len(%s) == 0 {", data)
					b.line("%s = nil", fn)
					b.line("}")
				}

				b.line("if err := %s.UnmarshalJSON(%s); err != nil {", fn, data)
				b.line("return err")
				b.line("}")
				b.line("")
			} else {
				object := util.GenerateID("obj")

				if !self.inline {
					b.line("case %q:", self.name)

					b.line("if dec.IsNull() {")
					if self.pointer {
						b.line("%s = nil", fn)
					} else {
						b.line("%s = %s{}", fn, b.typeString(field.Type(), opt))
					}
					b.line("} else if !dec.IsObjectOpen() {")
					b.line("return errors.NewParseError(dec.Char(), dec.Cursor())")
					b.line("} else {")

					// empty object
					b.line("if dec.IsObjectClose() {")
					if self.pointer {
						b.line("%s = nil", fn)
					} else {
						b.line("%s = %s{}", fn, b.typeString(field.Type(), opt))
					}
					b.line("} else {")

					if self.pointer {
						b.line("if %s == nil {", fn)
						b.line("%s = new(%s)", fn, b.typeString(field.Type(), opt))
						b.line("}")
						b.line("")
					}

					b.line("for %s := 1; %s > 0; {", object, object)

					// read key
					key := util.GenerateID("key")
					b.line("%s, err := dec.NextKey()", key)
					b.line("if err != nil {")
					b.line("return err")
					b.line("}")
					b.line("")

					b.line("switch %s {", key)

					self.keys = nil
				} else {
					self.keys = b.getKeys(obj)
				}

				b.gStructDecode(fn, self, x, opt)

				if !self.inline {
					b.line("default:")
					b.line("if err := dec.SkipValue(); err != nil {")
					b.line("return err")
					b.line("}")
					b.line("}")

					// check whether object closed
					b.line("if dec.IsObjectClose() {")
					b.line("%s--", object)
					b.line("}")

					b.line("}")
					b.line("}")
					b.line("}")
				}
			}

		case *types.Map:
			fn = fmt.Sprintf("%s.%s", fn, field.Name())

			b.line("case %q:", self.name)
			b.gMapDecode(fn, field.Type(), opt)

		case *types.Array:
			fn = fmt.Sprintf("%s.%s", fn, field.Name())
			b.gArrayDecode(fn, x, opt)

		case *types.Slice:
			fn = fmt.Sprintf("%s.%s", fn, field.Name())

			b.line("case %q:", self.name)
			b.gSliceDecode(fn, field.Type(), opt)

		case *types.Pointer:
			if !self.inline {
				self.keys = nil
			} else {
				self.keys = b.getKeys(obj)
			}

			self.pointer = true

			f := types.NewVar(field.Pos(), field.Pkg(), field.Name(), x.Elem())
			b.gFieldDecode(fn, parent, self, obj, f, opt)

		case *types.Interface:
			fn = fmt.Sprintf("%s.%s", fn, field.Name())

			value := util.GenerateID("value")
			alias := value
			if typ := b.typeString(field.Type(), opt); typ != "interface{}" {
				alias = fmt.Sprintf("%s(%s)", typ, value)
			}

			b.line("case %q:", self.name)
			b.line("%s, err := dec.DecodeValue()", value)
			b.line("if err != nil {")
			b.line("return err")
			b.line("}")
			b.line("")
			b.line("%s = %s", fn, alias)

		case *types.Basic:
			b.line("case %q:", self.name)
			value := util.GenerateID("value")
			alias := value
			if typ := b.typeString(field.Type(), opt); typ != x.Name() {
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
			b.line("%s.%s = %s", fn, field.Name(), alias)

		default:
			value := util.GenerateID("value")
			b.line("%s, err := dec.DecodeValue()", value)
			b.line("if err != nil {")
			b.line("return err")
			b.line("}")
			b.line("")
			b.line("%s.%s = %s", fn, field.Name(), value)
		}
	}
}

func (b *Builder) gStructDecode(fn string, parent *FieldTag, obj *types.Struct, opt *option.Option) {
	object := util.GenerateID("obj")

	if b.isRoot(fn) {
		//initialize pointer field
		// b.initPointerField(fn, obj, opt)

		// b.line("if dec.IsNull() {")
		// b.line("return nil")
		// b.line("} else if !dec.IsObjectOpen() {")
		// b.line("return errors.NewParseError(dec.Char(), dec.Cursor())")
		// b.line("} else {")
		b.line("if char := dec.NextChar(); char == 'n' {")
		b.line("if err := dec.AssetNull(); err != nil {")
		b.line("return err")
		b.line("}")
		b.line("")
		b.line("} else if char != '{' {")
		b.line("return errors.NewParseError(dec.Char(), dec.Cursor())")
		b.line("} else {")
		b.line("dec.Next()")

		// empty object
		b.line("if dec.IsObjectClose() {")
		b.line("return nil")
		b.line("} else {")

		b.line("for %s := 1; %s > 0;  {", object, object)

		// read key
		key := util.GenerateID("key")
		b.line("%s, err := dec.NextKey()", key)
		b.line("if err != nil {")
		b.line("return err")
		b.line("}")
		b.line("")

		b.line("switch %s {", key)
	}

	for i := 0; i < obj.NumFields(); i++ {
		tag := b.parseFieldTag(obj.Tag(i), obj.Field(i))

		if tag.ignore {
			continue
		}

		b.gFieldDecode(fn, parent, tag, obj, obj.Field(i), opt)
		b.line("")
	}

	if b.isRoot(fn) {
		b.line("default:")
		b.line("if err := dec.SkipValue(); err != nil {")
		b.line("return err")
		b.line("}")
		b.line("}")

		// check whether object closed
		b.line("if dec.IsObjectClose() {")
		b.line("%s--", object)
		b.line("}")

		b.line("}")
		b.line("}")
		b.line("}")
	}
}
