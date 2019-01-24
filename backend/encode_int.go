package backend

import (
	"strconv"
)

func (e *Encoder) EncodeInt8(value int8) {
	e.WriteString(strconv.FormatInt(int64(value), 10))
}

func (e *Encoder) EncodeInt16(value int16) {
	e.WriteString(strconv.FormatInt(int64(value), 10))
}

func (e *Encoder) EncodeInt32(value int32) {
	e.WriteString(strconv.FormatInt(int64(value), 10))
}

func (e *Encoder) EncodeInt64(value int64) {
	e.WriteString(strconv.FormatInt(value, 10))
}

func (e *Encoder) EncodeInt(value int) {
	e.WriteString(strconv.FormatInt(int64(value), 10))
}

func (e *Encoder) EncodeKeyInt8(key string, value int8) {
	c, ok := e.getPrevByte()
	if ok && c != '[' && c != '{' {
		e.WriteByte(',')
	}

	if key != "" {
		e.EncodeString(key)
		e.WriteByte(':')
	}
	e.EncodeInt8(value)
}

func (e *Encoder) EncodeKeyInt16(key string, value int16) {
	c, ok := e.getPrevByte()
	if ok && c != '[' && c != '{' {
		e.WriteByte(',')
	}

	if key != "" {
		e.EncodeString(key)
		e.WriteByte(':')
	}
	e.EncodeInt16(value)
}

func (e *Encoder) EncodeKeyInt32(key string, value int32) {
	c, ok := e.getPrevByte()
	if ok && c != '[' && c != '{' {
		e.WriteByte(',')
	}

	if key != "" {
		e.EncodeString(key)
		e.WriteByte(':')
	}
	e.EncodeInt32(value)
}

func (e *Encoder) EncodeKeyInt64(key string, value int64) {
	c, ok := e.getPrevByte()
	if ok && c != '[' && c != '{' {
		e.WriteByte(',')
	}

	if key != "" {
		e.EncodeString(key)
		e.WriteByte(':')
	}
	e.EncodeInt64(value)
}

func (e *Encoder) EncodeKeyInt(key string, value int) {
	c, ok := e.getPrevByte()
	if ok && c != '[' && c != '{' {
		e.WriteByte(',')
	}

	if key != "" {
		e.EncodeString(key)
		e.WriteByte(':')
	}
	e.EncodeInt(value)
}
