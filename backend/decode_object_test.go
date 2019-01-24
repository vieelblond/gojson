package backend

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeObject(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte(`   {"ke\"y1\"": [["TestString1","TestString2"],[123, 456],[123.123,456.456]], "key2": {"key3": "1111", "key4":123.111}}`))
	defer decoder.Release()

	testMap, err := decoder.DecodeObject()
	assert.Nil(t, err, "Err must be nil")
	assert.Len(t, testMap, 2, "testMap should be of len 2")
	assert.Len(t, testMap["ke\"y1\""], 3, "testMap should be of len 2")

	testArr := testMap["ke\"y1\""].([]interface{})
	assert.Equal(t, []interface{}{"TestString1", "TestString2"}, testArr[0], "testArr[0] must be equal to the value expected")
	assert.Equal(t, []interface{}{float64(123), float64(456)}, testArr[1], "testArr[1] must be equal to the value expected")
	assert.Equal(t, []interface{}{123.123, 456.456}, testArr[2], "testArr[2] must be equal to the value expected")

	testMap2 := testMap["key2"].(map[string]interface{})
	assert.Equal(t, "1111", testMap2["key3"], "testMap2[key3] must be equal to the value expected")
	assert.Equal(t, 123.111, testMap2["key4"], "testMap2[key4] must be equal to the value expected")
}

func TestReadObject(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte(`{"key1": [["Test{{String1}","{TestString2}"],[123, 456],[123.123,456.456]], "key2": {"key3": "1111", "key4":123.111}}`))
	defer decoder.Release()

	data, err := decoder.ReadObject()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, `{"key1": [["Test{{String1}","{TestString2}"],[123, 456],[123.123,456.456]], "key2": {"key3": "1111", "key4":123.111}}`, string(data), "data must be equal to the value expected")
}
