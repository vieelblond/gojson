package backend

import (
	"github.com/go-fish/gojson/util"
)

const chars = "0123456789abcdef"

func (e *Encoder) EncodeString(value string) {
	e.WriteByte('"')

	for i := 0; i < len(value); i++ {
		c := value[i]

		if !util.IsEscaped(c) {
			e.WriteByte(c)
			continue
		}

		switch c {
		case '\\', '"':
			e.WriteByte('\\')
			e.WriteByte(c)

		case '\n':
			e.WriteByte('\\')
			e.WriteByte('n')

		case '\f':
			e.WriteByte('\\')
			e.WriteByte('f')

		case '\b':
			e.WriteByte('\\')
			e.WriteByte('b')

		case '\r':
			e.WriteByte('\\')
			e.WriteByte('r')

		case '\t':
			e.WriteByte('\\')
			e.WriteByte('t')

		default:
			e.WriteByte(c)
		}
	}

	e.WriteByte('"')
}

func (e *Encoder) EncodeKeyString(key, value string) {
	c, ok := e.getPrevByte()
	if ok && c != '[' && c != '{' {
		e.WriteByte(',')
	}

	if key != "" {
		e.EncodeString(key)
		e.WriteByte(':')
	}
	e.EncodeString(value)
}
