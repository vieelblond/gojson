package backend

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecoderStringSliceBasic(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte(`["TestString1","TestString2"]`))
	defer decoder.Release()

	testArr, err := decoder.DecodeArray()
	assert.Nil(t, err, "Err must be nil")
	assert.Len(t, testArr, 2, "testArr should be of len 2")
	assert.Equal(t, "TestString1", testArr[0], "testArr[0] must be equal to the value expected")
	assert.Equal(t, "TestString2", testArr[1], "testArr[1] must be equal to the value expected")
}

func TestDecoderTwoDimensionalStringSliceBasic(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte(`[["TestString1","TestString2"],["TestString3", "TestString4"]]`))
	defer decoder.Release()

	testArr, err := decoder.DecodeArray()
	assert.Nil(t, err, "Err must be nil")
	assert.Len(t, testArr, 2, "testArr should be of len 2")
	assert.Equal(t, []interface{}{"TestString1", "TestString2"}, testArr[0], "testArr[0] must be equal to the value expected")
	assert.Equal(t, []interface{}{"TestString3", "TestString4"}, testArr[1], "testArr[1] must be equal to the value expected")
}

func TestDecoderTwoDimensionalInterfaceSliceBasic(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte(`[["TestString1","TestString2"],[123, 456],[123.123,456.456]]`))
	defer decoder.Release()

	testArr, err := decoder.DecodeArray()
	assert.Nil(t, err, "Err must be nil")
	assert.Len(t, testArr, 3, "testArr should be of len 2")
	assert.Equal(t, []interface{}{"TestString1", "TestString2"}, testArr[0], "testArr[0] must be equal to the value expected")
	assert.Equal(t, []interface{}{float64(123), float64(456)}, testArr[1], "testArr[1] must be equal to the value expected")
	assert.Equal(t, []interface{}{123.123, 456.456}, testArr[2], "testArr[2] must be equal to the value expected")
}

func TestReadArray(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte(`[["Test[String1]","TestString2"],[123, 456],[123.123,456.456]]`))
	defer decoder.Release()

	data, err := decoder.ReadArray()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, `[["Test[String1]","TestString2"],[123, 456],[123.123,456.456]]`, string(data), "data must be equal to the value expected")
}

func TestSkipArray(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte(`[["Test[String1]","TestString2"],[123, 456],[123.123,456.456]]]`))
	defer decoder.Release()

	err := decoder.SkipArray()
	assert.Nil(t, err, "Err must be nil")
}
