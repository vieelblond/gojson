package backend

import "sync"

type Encoder struct {
	data []byte
	err  error
}

var encoderPool = &sync.Pool{New: func() interface{} { return &Encoder{data: make([]byte, 0, 4096)} }}

func NewEncoder() *Encoder {
	return encoderPool.Get().(*Encoder)
}

func (e *Encoder) Release() {
	e.reset()
	encoderPool.Put(e)
}

func (e *Encoder) reset() {
	e.data = e.data[:0]
	e.err = nil
}

func (e *Encoder) WriteByte(value byte) {
	e.data = append(e.data, value)
}

func (e *Encoder) WriteBytes(values []byte) {
	e.data = append(e.data, values...)
}

func (e *Encoder) WriteString(value string) {
	e.data = append(e.data, value...)
}

func (e *Encoder) WriteKey(key string) {
	c, ok := e.getPrevByte()
	if ok && c != '[' && c != '{' {
		e.WriteByte(',')
	}

	e.EncodeString(key)
	e.WriteByte(':')
}

func (e *Encoder) WriteNull() {
	e.WriteString("null")
}

func (e *Encoder) WriteComma() {
	if c, ok := e.getPrevByte(); ok && c != '[' && c != '{' {
		e.WriteByte(',')
	}
}

func (e *Encoder) getPrevByte() (byte, bool) {
	if len(e.data) == 0 {
		return 0, false
	}

	return e.data[len(e.data)-1], true
}

func (e *Encoder) Bytes() []byte {
	return e.data
}
