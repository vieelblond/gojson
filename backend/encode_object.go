package backend

func (e *Encoder) EncodeObject(obj map[string]interface{}) {
	if len(obj) == 0 {
		e.WriteNull()
	}

	e.WriteByte('{')
	for key, value := range obj {
		e.EncodeKeyValue(key, value)
	}
	e.WriteByte('}')
}
