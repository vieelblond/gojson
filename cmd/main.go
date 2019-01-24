package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/go-fish/gojson/gen"
	"github.com/go-fish/gojson/option"
	"github.com/go-fish/gojson/version"
	"github.com/ttacon/chalk"
)

func parseOption() (*option.Option, error) {
	opt, err := option.NewOption()
	if err != nil {
		return nil, err
	}

	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <input dir|file>\n\n", os.Args[0])
		flag.PrintDefaults()
	}

	var mode string

	flag.StringVar(&opt.Output, "o", opt.Output, "Optional name of the output file to be generated.")
	flag.StringVar(&mode, "m", "all", "Mode of generate, eg: encode, decode, all")
	flag.BoolVar(&opt.Unsafe, "unsafe", false, "Use decoder without copy data")
	flag.BoolVar(&opt.Inline, "inline", true, "Use inline function in generate code")
	ver := flag.Bool("version", false, "Show version information.")

	flag.Parse()

	if *ver {
		fmt.Fprintf(os.Stdout, "%s\n\n", chalk.Magenta.Color(version.Version()))
		os.Exit(0)
	}

	switch strings.ToLower(mode) {
	case "encode":
		opt.Mode |= option.Encode

	case "decode":
		opt.Mode |= option.Decode

	case "all":
		opt.Mode |= option.All
	}

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, chalk.Red.Color("Missing <input dir|file>, need exactly one\n"))
		flag.Usage()
		os.Exit(1)
	}

	if flag.NArg() > 1 {
		fmt.Fprintf(os.Stderr, chalk.Red.Color("Too many <input dir|file>, need exactly one\n"))
		flag.Usage()
		os.Exit(1)
	}

	// append input files to option
	opt.Input = flag.Arg(0)

	// parse pkg in opt
	if err := opt.ParsePackage(); err != nil {
		return nil, err
	}

	return opt, nil
}

func main() {
	opt, err := parseOption()
	if err != nil {
		fmt.Fprintf(os.Stderr, chalk.Red.Color(fmt.Sprintf("gojson error: %s\n", err)))
		os.Exit(1)
	}

	if err := gen.Generate(opt); err != nil {
		fmt.Fprintf(os.Stderr, chalk.Red.Color(fmt.Sprintf("gojson error: %s\n", err)))
		os.Exit(1)
	}
}
