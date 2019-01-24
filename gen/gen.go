package gen

import (
	"bytes"
	"fmt"
	"go/types"
	"io/ioutil"
	"os"

	"github.com/alecthomas/template"
	"github.com/go-fish/gojson/option"
	"github.com/ttacon/chalk"
	"golang.org/x/tools/imports"
)

type Builder struct {
	Package string
	Body    *bytes.Buffer
	Imports map[string]string
}

func Generate(opt *option.Option) error {
	b := new(Builder)
	b.Package = opt.Pkg.Name()
	b.Imports = map[string]string{
		"backend": "github.com/go-fish/gojson/backend",
		"errors":  "github.com/go-fish/gojson/errors",
	}
	b.Body = bytes.NewBuffer(make([]byte, 0, 4096))

	for _, name := range opt.Pkg.Scope().Names() {
		scope := opt.Pkg.Scope().Lookup(name)

		// check obj type
		if obj, ok := scope.Type().Underlying().(*types.Struct); ok {
			if !scope.Exported() {
				fmt.Fprintf(os.Stdout, chalk.Red.Color("Ignore object %s because of unexported\n"), scope.Name())
				continue
			}

			fmt.Fprintf(os.Stdout, chalk.Green.Color("Begin to generate object %s\n"), scope.Name())

			// check mode
			if opt.IsEncode() {
				if err := b.gStructEncodeWarp(name, obj, opt); err != nil {
					return err
				}
			}

			if opt.IsDecode() {
				if err := b.gStructDecodeWarp(name, obj, opt); err != nil {
					return err
				}
			}
		}
	}

	tmpl, err := template.New("gojson").Parse(gojson)
	if err != nil {
		return err
	}

	w := bytes.NewBuffer(make([]byte, 0, 4096))
	if err := tmpl.Execute(w, b); err != nil {
		return err
	}

	out, err := imports.Process(opt.Output, w.Bytes(), nil)
	if err != nil {
		return err
	}

	// flush to file
	return ioutil.WriteFile(opt.Output, out, 0600)
}
