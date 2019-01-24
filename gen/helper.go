package gen

import (
	"fmt"
	"go/types"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/go-fish/gojson/option"
)

type FieldTag struct {
	inline    bool
	omitempty bool
	ignore    bool
	pointer   bool
	name      string
	keys      []string
}

func (b *Builder) parseFieldTag(tag string, field *types.Var) *FieldTag {
	var ft FieldTag

	ft.inline = field.Anonymous()
	ft.name = field.Name()
	ft.ignore = !field.Exported()

	v, ok := reflect.StructTag(strings.Trim(tag, "`")).Lookup("json")
	if !ok {
		return &ft
	}

	if v == "-" {
		ft.ignore = true
		return &ft
	}

	for i, t := range strings.Split(v, ",") {
		if i == 0 && t != "" {
			ft.name = t
			continue
		}

		switch t {
		case "inline":
			ft.inline = true

		case "omitempty":
			ft.omitempty = true
		}
	}

	return &ft
}

func (b *Builder) line(data string, args ...interface{}) {
	b.Body.WriteString(fmt.Sprintf(data, args...))
	b.Body.WriteByte('\n')
}

func (b *Builder) isRoot(fn string) bool {
	return !strings.Contains(fn, ".")
}

func (b *Builder) needPrint(parent, self *FieldTag) bool {
	if !parent.inline {
		return true
	}

	for _, key := range parent.keys {
		if self.name == key {
			return false
		}
	}

	return true
}

func (b *Builder) getKeys(obj *types.Struct) []string {
	keys := make([]string, 0, 8)

	if v, ok := obj.Underlying().(*types.Struct); ok {
		for i := 0; i < v.NumFields(); i++ {
			tag := b.parseFieldTag(v.Tag(i), v.Field(i))

			if !tag.ignore && !tag.inline {
				keys = append(keys, tag.name)
			}
		}
	}

	return keys
}

func (b *Builder) typeString(typ types.Type, opt *option.Option) string {
	switch x := typ.(type) {
	case *types.Named:
		if opt.IsLocal(x.Obj().Pkg()) {
			return x.Obj().Name()
		}

		// add import
		if !opt.IsLocal(x.Obj().Pkg()) {
			pkg := x.Obj().Pkg().Path()
			if index := strings.Index(pkg, "/vendor/"); index > 0 {
				pkg = pkg[index+8:]
			}

			b.Imports[x.Obj().Pkg().Name()] = pkg
		}

		return fmt.Sprintf("%s.%s", x.Obj().Pkg().Name(), x.Obj().Name())

	case *types.Basic:
		return x.Name()

	case *types.Slice:
		return fmt.Sprintf("[]%s", b.typeString(x.Elem(), opt))

	case *types.Array:
		return fmt.Sprintf("[%d]%s", x.Len(), b.typeString(x.Elem(), opt))

	case *types.Map:
		return fmt.Sprintf("map[%s]%s", b.typeString(x.Key(), opt), b.typeString(x.Elem(), opt))

	case *types.Pointer:
		return "*" + b.typeString(x.Elem(), opt)

	default:
		return filepath.Base(strings.Replace(typ.String(), "..", "", -1))
	}
}

// init all pointer field at once
func (b *Builder) initPointerField(fn string, obj *types.Struct, opt *option.Option) {
	for i := 0; i < obj.NumFields(); i++ {
		field := obj.Field(i)
		tag := b.parseFieldTag(obj.Tag(i), field)

		if tag.ignore {
			continue
		}

		switch x := field.Type().Underlying().(type) {
		case *types.Pointer:
			child := fmt.Sprintf("%s.%s", fn, field.Name())
			b.line("if %s == nil {", child)
			b.line("%s = new(%s)", child, b.typeString(x.Elem(), opt))
			b.line("}")
			b.line("")

			if o, ok := x.Elem().Underlying().(*types.Struct); ok {
				b.initPointerField(child, o, opt)
			}

		case *types.Struct:
			child := fmt.Sprintf("%s.%s", fn, field.Name())
			b.initPointerField(child, x, opt)
		}
	}
}
