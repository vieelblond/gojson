package backend

import (
	"sync"

	"github.com/go-fish/gojson/errors"
)

type Decoder struct {
	data   []byte
	cursor int
	length int
	err    error
}

var decoderPool = &sync.Pool{New: func() interface{} { return new(Decoder) }}

func NewDecoder() *Decoder {
	return decoderPool.Get().(*Decoder)
}

func (d *Decoder) reset() {
	d.data = nil
	d.cursor = 0
	d.length = 0
	d.err = nil
}

func (d *Decoder) SetData(data []byte) {
	// for safry to use unsafe pointer
	d.data = make([]byte, len(data))
	copy(d.data, data)

	d.length = len(data)
	d.cursor = 0
}

func (d *Decoder) SetUnsafeData(data []byte) {
	d.data = data
	d.length = len(data)
	d.cursor = 0
}

func (d *Decoder) Release() {
	d.reset()
	decoderPool.Put(d)
}

func (d *Decoder) Need(b byte) bool {
	for d.cursor < d.length {
		switch d.data[d.cursor] {
		case ' ', '\n', '\t', '\r', ',':
			d.cursor++

		default:
			return d.data[d.cursor] == b
		}
	}

	return false
}

func (d *Decoder) NextChar() byte {
	for d.cursor < d.length {
		switch d.data[d.cursor] {
		case ' ', '\n', '\t', '\r', ',':
			d.cursor++

		default:
			return d.data[d.cursor]
		}
	}

	return 0
}

func (d *Decoder) Next() {
	d.cursor++
}

func (d *Decoder) IsNull() bool {
	if d.Need('n') && d.cursor+3 < d.length &&
		d.data[d.cursor+1] == 'u' &&
		d.data[d.cursor+2] == 'l' &&
		d.data[d.cursor+3] == 'l' {

		d.cursor += 4
		return true
	}

	return false
}

func (d *Decoder) AssetNull() error {
	if d.cursor+3 < d.length &&
		d.data[d.cursor+1] == 'u' &&
		d.data[d.cursor+2] == 'l' &&
		d.data[d.cursor+3] == 'l' {

		d.cursor += 4
		return nil
	}

	return errors.NewParseError(d.data[d.cursor], d.cursor)
}

func (d *Decoder) AssetFalse() (bool, error) {
	if d.cursor+4 < d.length &&
		d.data[d.cursor+1] == 'a' &&
		d.data[d.cursor+2] == 'l' &&
		d.data[d.cursor+3] == 's' &&
		d.data[d.cursor+4] == 'e' {

		d.cursor += 5
		return false, nil
	}

	return false, errors.NewParseError(d.data[d.cursor], d.cursor)
}

func (d *Decoder) AssetTrue() (bool, error) {
	if d.cursor+3 < d.length &&
		d.data[d.cursor+1] == 'r' &&
		d.data[d.cursor+2] == 'u' &&
		d.data[d.cursor+3] == 'e' {

		d.cursor += 4
		return true, nil
	}

	return false, errors.NewParseError(d.data[d.cursor], d.cursor)
}

func (d *Decoder) Cursor() int {
	return d.cursor
}

func (d *Decoder) Char() byte {
	return d.data[d.cursor]
}
