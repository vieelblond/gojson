package backend

import (
	"encoding/base64"
)

func (d *Decoder) DecodeBytes() ([]byte, error) {
	v, err := d.DecodeString()
	if err != nil {
		return nil, err
	}

	return base64.StdEncoding.DecodeString(v)
}
