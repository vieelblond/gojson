package backend

import (
	"github.com/go-fish/gojson/errors"
)

func (d *Decoder) ReadValue() ([]byte, error) {
	if d.IsNull() {
		return nil, nil
	}

	for d.cursor < d.length {
		switch d.data[d.cursor] {
		case '"':
			return d.ReadString()

		case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return d.ReadFloat()

		case 'f', 't':
			return d.ReadBool()

		case '{':
			return d.ReadObject()

		case '[':
			return d.ReadArray()

		default:
			return nil, errors.NewParseError(d.data[d.cursor], d.cursor)
		}
	}

	return nil, errors.NewParseError(d.data[d.cursor], d.cursor)
}

func (d *Decoder) DecodeValue() (interface{}, error) {
	if d.IsNull() {
		return nil, nil
	}

	for d.cursor < d.length {
		switch d.data[d.cursor] {
		case '"':
			return d.DecodeString()

		case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return d.DecodeFloat64()

		case 'f', 't':
			return d.DecodeBool()

		case '{':
			return d.DecodeObject()

		case '[':
			return d.DecodeArray()

		default:
			return nil, errors.NewParseError(d.data[d.cursor], d.cursor)
		}
	}

	return nil, errors.NewParseError(d.data[d.cursor], d.cursor)
}

func (d *Decoder) SkipValue() error {
	for d.cursor < d.length {
		switch d.data[d.cursor] {
		case ' ', '\n', '\t', '\r', ',':
			d.cursor++

		case 'n':
			d.cursor += 4
			return nil

		case '"':
			return d.SkipString()

		case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return d.SkipFloat64()

		case 'f', 't':
			return d.SkipBool()

		case '{':
			return d.SkipObject()

		case '[':
			return d.SkipArray()

		default:
			return errors.NewParseError(d.data[d.cursor], d.cursor)
		}
	}

	return errors.NewParseError(d.data[d.cursor], d.cursor)
}
