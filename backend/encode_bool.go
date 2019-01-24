package backend

func (e *Encoder) EncodeBool(value bool) {
	if value {
		e.WriteString("true")
		return
	}

	e.WriteString("false")
}

func (e *Encoder) EncodeKeyBool(key string, value bool) {
	c, ok := e.getPrevByte()
	if ok && c != '[' && c != '{' {
		e.WriteByte(',')
	}

	if key != "" {
		e.EncodeString(key)
		e.WriteByte(':')
	}
	e.EncodeBool(value)
}
