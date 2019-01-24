package backend

import "strconv"

func (e *Encoder) EncodeFloat32(value float32) {
	e.WriteString(strconv.FormatFloat(float64(value), 'f', -1, 64))
}

func (e *Encoder) EncodeKeyFloat32(key string, value float32) {
	c, ok := e.getPrevByte()
	if ok && c != '[' && c != '{' {
		e.WriteByte(',')
	}

	e.EncodeString(key)
	e.WriteByte(':')
	e.EncodeFloat32(value)
}

func (e *Encoder) EncodeFloat64(value float64) {
	e.WriteString(strconv.FormatFloat(value, 'f', -1, 64))
}

func (e *Encoder) EncodeKeyFloat64(key string, value float64) {
	c, ok := e.getPrevByte()
	if ok && c != '[' && c != '{' {
		e.WriteByte(',')
	}

	if key != "" {
		e.EncodeString(key)
		e.WriteByte(':')
	}
	e.EncodeFloat64(value)
}
