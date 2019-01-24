package backend

import (
	"encoding/json"
	"fmt"
)

func (e *Encoder) EncodeValue(value interface{}) error {
	switch x := value.(type) {
	case string:
		e.EncodeString(x)

	case bool:
		e.EncodeBool(x)

	case int:
		e.EncodeInt(x)

	case int8:
		e.EncodeInt8(x)

	case int16:
		e.EncodeInt16(x)

	case int32:
		e.EncodeInt32(x)

	case int64:
		e.EncodeInt64(x)

	case uint:
		e.EncodeUint(x)

	case uint8:
		e.EncodeUint8(x)

	case uint16:
		e.EncodeUint16(x)

	case uint32:
		e.EncodeUint32(x)

	case uint64:
		e.EncodeUint64(x)

	case float32:
		e.EncodeFloat32(x)

	case float64:
		e.EncodeFloat64(x)

	case json.Marshaler:
		data, err := x.MarshalJSON()
		if err != nil {
			return err
		}

		e.WriteBytes(data)

	default:
		return fmt.Errorf("Unsupported value type %T", x)
	}

	return nil
}

func (e *Encoder) EncodeKeyValue(key string, value interface{}) error {
	c, ok := e.getPrevByte()
	if ok && c != '[' && c != '{' {
		e.WriteByte(',')
	}

	if key != "" {
		e.EncodeString(key)
		e.WriteByte(':')
	}
	return e.EncodeValue(value)
}
