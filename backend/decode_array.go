package backend

import (
	"github.com/go-fish/gojson/errors"
)

func (d *Decoder) ReadArray() ([]byte, error) {
	if c := d.NextChar(); c == 'n' {
		return nil, d.AssetNull()
	} else if c != '[' {
		return nil, errors.NewParseError(d.data[d.cursor], d.cursor)
	}

	begin := d.cursor
	arrayOpend := 1
	d.cursor++

	for d.cursor < d.length {
		switch d.data[d.cursor] {
		case '"':
			if err := d.SkipString(); err != nil {
				return nil, err
			}

		case '[':
			d.cursor++
			arrayOpend++

		case ']':
			d.cursor++
			arrayOpend--

			if arrayOpend == 0 {
				return d.data[begin:d.cursor], nil
			}

		default:
			d.cursor++
		}
	}

	return nil, errors.NewParseError(d.data[d.length-1], d.length-1)
}

func (d *Decoder) DecodeArray() ([]interface{}, error) {
	v := make([]interface{}, 0, 8)

	if c := d.NextChar(); c == 'n' {
		return nil, d.AssetNull()
	} else if c != '[' {
		return nil, errors.NewParseError(d.data[d.cursor], d.cursor)
	}

	d.cursor++

	if d.Need(']') {
		d.cursor++
		return nil, nil
	}

	for d.cursor < d.length {
		value, err := d.DecodeValue()
		if err != nil {
			return nil, err
		}

		v = append(v, value)

		if d.Need(']') {
			d.cursor++
			return v, nil
		}
	}

	return nil, errors.NewParseError(d.data[d.length-1], d.length-1)
}

func (d *Decoder) SkipArray() error {
	arrayOpend := 1
	d.cursor++

	for d.cursor < d.length {
		switch d.data[d.cursor] {
		case '"':
			if err := d.SkipString(); err != nil {
				return err
			}

		case '[':
			d.cursor++
			arrayOpend++

		case ']':
			d.cursor++
			arrayOpend--

			if arrayOpend == 0 {
				return nil
			}

		default:
			d.cursor++
		}
	}

	return errors.NewParseError(d.data[d.length-1], d.length-1)
}

func (d *Decoder) IsArrayOpen() bool {
	if d.Need('[') {
		d.cursor++
		return true
	}

	return false
}

func (d *Decoder) IsArrayClose() bool {
	if d.Need(']') {
		d.cursor++
		return true
	}

	return false
}
