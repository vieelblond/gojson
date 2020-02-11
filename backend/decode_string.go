package backend

import (
	"github.com/go-fish/gojson/errors"
	"github.com/go-fish/gojson/util"
)

func (d *Decoder) ReadString() ([]byte, error) {
	if c := d.NextChar(); c == 'n' {
		return nil, d.AssetNull()
	} else if c != '"' {
		return nil, errors.NewParseError(d.data[d.cursor], d.cursor)
	}

	data := d.data[d.cursor:]

	for i, c := range data {
		if c == '"' && i != 0 {
			if prev := i - 1; prev >= 0 && data[prev] == '\\' {
				continue
			}

			d.cursor = d.cursor + i + 1
			return data[:i+1], nil
		}
	}

	return nil, errors.NewParseError(d.data[d.length-1], d.length-1)
}

func (d *Decoder) DecodeString() (string, error) {
	if c := d.NextChar(); c == 'n' {
		return "", d.AssetNull()
	} else if c != '"' {
		return "", errors.NewParseError(d.data[d.cursor], d.cursor)
	}

	d.cursor++
	data := d.data[d.cursor:]

	for i, c := range data {
		if c == '"' {
			d.cursor = d.cursor + i + 1
			return util.UnsafeConvertBytesToString(data[:i]), nil
		} else if c == '\\' {
			if next := i + 1; next < len(data) && (data[next] == '\\' || data[next] == '"') {
				//copy(data[i:], data[next:])
				data = append(data[:i], data[next:]...)
			}
		}
	}

	return "", errors.NewParseError(d.data[d.length-1], d.length-1)
}

func (d *Decoder) SkipString() error {
	d.cursor++
	data := d.data[d.cursor:]

	for i, c := range data {
		if c == '"' {
			if prev := i - 1; prev >= 0 && data[prev] == '\\' {
				continue
			}

			d.cursor = d.cursor + i + 1
			return nil
		}
	}

	return errors.NewParseError(d.data[d.length-1], d.length-1)
}
