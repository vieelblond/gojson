package backend

import (
	"github.com/go-fish/gojson/errors"
	"github.com/go-fish/gojson/util"
)

func (d *Decoder) NextKey() (string, error) {
	if !d.Need('"') {
		return "", errors.NewParseError(d.data[d.cursor], d.cursor)
	}

	d.cursor++
	data := d.data[d.cursor:]

	for i, c := range data {
		if c == '"' {
			d.cursor = d.cursor + i + 1

			var found byte
			for _, c := range data[i+1:] {
				if c == ':' {
					found |= 1
					break
				}

				d.cursor++
			}

			if found&1 > 0 {
				d.cursor++
				return util.UnsafeConvertBytesToString(data[:i]), nil
			}

			return "", errors.NewParseError(d.data[d.cursor], d.cursor)
		} else if c == '\\' {
			if next := i + 1; next < len(data) && (data[next] == '\\' || data[next] == '"') {
				// _ = append(data[:i], data[next:]...)
				copy(data[i:], data[next:])
			}
		}
	}

	return "", errors.NewParseError(d.data[d.length-1], d.length-1)
}

func (d *Decoder) ReadObject() ([]byte, error) {
	if c := d.NextChar(); c == 'n' {
		return nil, d.AssetNull()
	} else if c != '{' {
		return nil, errors.NewParseError(d.data[d.cursor], d.cursor)
	}

	objectOpened := 1
	begin := d.cursor
	d.cursor++

	for d.cursor < d.length {
		switch d.data[d.cursor] {
		case '"':
			if err := d.SkipString(); err != nil {
				return nil, err
			}

		case '{':
			d.cursor++
			objectOpened++

		case '}':
			d.cursor++
			objectOpened--

			if objectOpened == 0 {
				return d.data[begin:d.cursor], nil
			}

		default:
			d.cursor++
		}
	}

	return nil, errors.NewParseError(d.data[d.length-1], d.length-1)
}

func (d *Decoder) DecodeObject() (map[string]interface{}, error) {
	v := map[string]interface{}{}

	if c := d.NextChar(); c == 'n' {
		return nil, d.AssetNull()
	} else if c != '{' {
		return nil, errors.NewParseError(d.data[d.cursor], d.cursor)
	}

	d.cursor++

	if d.Need('}') {
		d.cursor++
		return nil, nil
	}

	for d.cursor < d.length {
		key, err := d.NextKey()
		if err != nil {
			return nil, err
		}

		value, err := d.DecodeValue()
		if err != nil {
			return nil, err
		}

		v[key] = value

		if d.Need('}') {
			d.cursor++
			return v, nil
		}
	}

	return nil, errors.NewParseError(d.data[d.length-1], d.length-1)
}

func (d *Decoder) SkipObject() error {
	objectOpened := 1
	d.cursor++

	for d.cursor < d.length {
		switch d.data[d.cursor] {
		case '"':
			if err := d.SkipString(); err != nil {
				return err
			}

		case '{':
			d.cursor++
			objectOpened++

		case '}':
			d.cursor++
			objectOpened--

			if objectOpened == 0 {
				return nil
			}

		default:
			d.cursor++
		}
	}

	return errors.NewParseError(d.data[d.length-1], d.length-1)
}

func (d *Decoder) IsObjectOpen() bool {
	if d.Need('{') {
		d.cursor++
		return true
	}

	return false
}

func (d *Decoder) IsObjectClose() bool {
	if d.Need('}') {
		d.cursor++
		return true
	}

	return false
}
