package backend

import (
	"strconv"
)

func (e *Encoder) EncodeUint8(value uint8) {
	e.WriteString(strconv.FormatUint(uint64(value), 10))
}

func (e *Encoder) EncodeUint16(value uint16) {
	e.WriteString(strconv.FormatUint(uint64(value), 10))
}

func (e *Encoder) EncodeUint32(value uint32) {
	e.WriteString(strconv.FormatUint(uint64(value), 10))
}

func (e *Encoder) EncodeUint64(value uint64) {
	e.WriteString(strconv.FormatUint(value, 10))
}

func (e *Encoder) EncodeUint(value uint) {
	e.WriteString(strconv.FormatUint(uint64(value), 10))
}

func (e *Encoder) EncodeKeyUint8(key string, value uint8) {
	c, ok := e.getPrevByte()
	if ok && c != '[' && c != '{' {
		e.WriteByte(',')
	}

	if key != "" {
		e.EncodeString(key)
		e.WriteByte(':')
	}
	e.EncodeUint8(value)
}

func (e *Encoder) EncodeKeyUint16(key string, value uint16) {
	c, ok := e.getPrevByte()
	if ok && c != '[' && c != '{' {
		e.WriteByte(',')
	}

	if key != "" {
		e.EncodeString(key)
		e.WriteByte(':')
	}
	e.EncodeUint16(value)
}

func (e *Encoder) EncodeKeyUint32(key string, value uint32) {
	c, ok := e.getPrevByte()
	if ok && c != '[' && c != '{' {
		e.WriteByte(',')
	}

	if key != "" {
		e.EncodeString(key)
		e.WriteByte(':')
	}
	e.EncodeUint32(value)
}

func (e *Encoder) EncodeKeyUint64(key string, value uint64) {
	c, ok := e.getPrevByte()
	if ok && c != '[' && c != '{' {
		e.WriteByte(',')
	}

	if key != "" {
		e.EncodeString(key)
		e.WriteByte(':')
	}
	e.EncodeUint64(value)
}

func (e *Encoder) EncodeKeyUint(key string, value uint) {
	c, ok := e.getPrevByte()
	if ok && c != '[' && c != '{' {
		e.WriteByte(',')
	}

	if key != "" {
		e.EncodeString(key)
		e.WriteByte(':')
	}
	e.EncodeUint(value)
}
