package option

import (
	"fmt"
	"go/importer"
	"go/types"
	"os"
	"path/filepath"
)

type Mode uint8

const (
	Decode Mode = 1 << iota
	Encode
	All
	None
)

// Option defines options on files to be convert.
type Option struct {
	Input  string
	Output string
	Mode   Mode

	// Unsafe used to decied whether we use copy in decoder, too make sure the result of Unmarshal will not change even if source data is changed.
	Unsafe bool

	// Inline used to decied whether we use inline functions in generated code to increase the performance.
	Inline      bool
	Marshaler   *types.Interface
	Unmarshaler *types.Interface
	Pkg         *types.Package
}

func NewOption() (*Option, error) {
	opt := &Option{
		Output: "gojson.generate.go",
		Mode:   All,
		Unsafe: false,
		Inline: true,
	}

	// initialize Marshaler && Unmarshaler
	pkg, err := importer.For("source", nil).Import("encoding/json")
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize Marshaler && Unmarshaler, error: %s", err)
	}

	for _, name := range pkg.Scope().Names() {
		obj := pkg.Scope().Lookup(name)

		switch x := obj.Type().Underlying().(type) {
		case *types.Interface:
			if name == "Marshaler" {
				opt.Marshaler = x
			} else if name == "Unmarshaler" {
				opt.Unmarshaler = x
			}
		}
	}

	if opt.Marshaler == nil || opt.Unmarshaler == nil {
		return nil, fmt.Errorf("Failed to initialize Marshaler && Unmarshaler, error: Not found Marshaler or Unmarshaler")
	}

	return opt, nil
}

func (o *Option) IsEncode() bool {
	return o.Mode&Encode > 0 || o.Mode&All > 0
}

func (o *Option) IsDecode() bool {
	return o.Mode&Decode > 0 || o.Mode&All > 0
}

func (o *Option) IsMarshaler(v types.Type) bool {
	if fn, _ := types.MissingMethod(v, o.Marshaler, true); fn == nil {
		return true
	}

	if fn, _ := types.MissingMethod(types.NewPointer(v).Underlying(), o.Marshaler, true); fn == nil {
		return true
	}

	return false
}

func (o *Option) IsUnmarshaler(v types.Type) bool {
	if fn, _ := types.MissingMethod(v, o.Unmarshaler, true); fn == nil {
		return true
	}

	if fn, _ := types.MissingMethod(types.NewPointer(v).Underlying(), o.Unmarshaler, true); fn == nil {
		return true
	}

	return false
}

func (o *Option) ParsePackage() error {
	err := os.Remove(o.Output)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	input := o.Input

	fi, err := os.Stat(o.Input)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		input = filepath.Dir(input)
	}

	o.Pkg, err = importer.For("source", nil).Import(input)
	if err != nil {
		return err
	}

	if o.Pkg.Scope() == nil || len(o.Pkg.Scope().Names()) == 0 {
		return fmt.Errorf("No types to generate")
	}

	return nil
}

func (o *Option) IsLocal(pkg *types.Package) bool {
	return o.Pkg.Name() == pkg.Name() && o.Pkg.Path() == pkg.Path()
}
