package backend

import (
	"github.com/go-fish/gojson/errors"
)

func (d *Decoder) ReadBool() ([]byte, error) {
	for d.cursor < d.length {
		switch d.data[d.cursor] {
		case ' ', '\n', '\t', '\r', ',':
			d.cursor++

		case 'f':
			begin := d.cursor

			if d.cursor+4 < d.length &&
				d.data[d.cursor+1] == 'a' &&
				d.data[d.cursor+2] == 'l' &&
				d.data[d.cursor+3] == 's' &&
				d.data[d.cursor+4] == 'e' {

				d.cursor += 5
				return d.data[begin:d.cursor], nil
			}

		case 't':
			begin := d.cursor

			if d.cursor+3 < d.length &&
				d.data[d.cursor+1] == 'r' &&
				d.data[d.cursor+2] == 'u' &&
				d.data[d.cursor+3] == 'e' {

				d.cursor += 4
				return d.data[begin:d.cursor], nil
			}

		default:
			return nil, errors.NewParseError(d.data[d.cursor], d.cursor)
		}
	}

	return nil, errors.NewParseError(d.data[d.length-1], d.length-1)
}

func (d *Decoder) DecodeBool() (bool, error) {
	if c := d.NextChar(); c == 'n' {
		return false, d.AssetNull()
	} else if c == 'f' {
		return d.AssetFalse()
	} else if c == 't' {
		return d.AssetTrue()
	} else {
		return false, errors.NewParseError(d.data[d.cursor], d.cursor)
	}
}

func (d *Decoder) SkipBool() error {
	if d.Need('f') {
		d.cursor += 5
	} else if d.Char() == 't' {
		d.cursor += 4
	} else {
		return errors.NewParseError(d.data[d.cursor], d.cursor)
	}

	return nil
}
