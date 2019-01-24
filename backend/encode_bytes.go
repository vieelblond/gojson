package backend

import (
	"encoding/base64"
)

func (e *Encoder) EncodeBytes(value []byte) {
	e.WriteString(base64.StdEncoding.EncodeToString(value))
}

func (e *Encoder) EncoderKeyBytes(key string, value []byte) {
	c, ok := e.getPrevByte()
	if ok && c != '[' && c != '{' {
		e.WriteByte(',')
	}

	if key != "" {
		e.EncodeString(key)
		e.WriteByte(':')
	}
	e.EncodeBytes(value)
}
