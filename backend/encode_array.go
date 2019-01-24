package backend

func (e *Encoder) EncodeArray(obj []interface{}) {
	if len(obj) == 0 {
		e.WriteNull()
		return
	}

	e.WriteByte('[')
	for _, value := range obj {
		e.EncodeKeyValue("", value)
	}
	e.WriteByte(']')
}
