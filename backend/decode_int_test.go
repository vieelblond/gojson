package backend

import (
	"testing"

	"github.com/go-fish/gojson/errors"
	"github.com/stretchr/testify/assert"
)

func TestDecoderInt8Basic(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("127"))
	defer decoder.Release()

	v, err := decoder.DecodeInt8()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int8(127), v, "v must be equal to the value expected")
}

func TestDecoderNegativeInt8Basic(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("-127"))
	defer decoder.Release()

	v, err := decoder.DecodeInt8()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int8(-127), v, "v must be equal to the value expected")
}

func TestDecoderInt8Overflow(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("128"))
	defer decoder.Release()

	_, err := decoder.DecodeInt8()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewInt8OverflowError(0), err, "err must be equal to the value expected")
}

func TestDecoderNegativeInt8Overflow(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("-128"))
	defer decoder.Release()

	_, err := decoder.DecodeInt8()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewInt8OverflowError(1), err, "err must be equal to the value expected")
}

func TestDecoderInt8Null(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("null"))
	defer decoder.Release()

	v, err := decoder.DecodeInt8()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int8(0), v, "v must be equal to the value expected")
}

func TestDecoderInt8(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte(" \n\n 127"))
	defer decoder.Release()

	v, err := decoder.DecodeInt8()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int8(127), v, "v must be equal to the value expected")
}
func TestDecoderInt16Basic(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("32767"))
	defer decoder.Release()

	v, err := decoder.DecodeInt16()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int16(32767), v, "v must be equal to the value expected")
}

func TestDecoderNegativeInt16Basic(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("-32767"))
	defer decoder.Release()

	v, err := decoder.DecodeInt16()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int16(-32767), v, "v must be equal to the value expected")
}

func TestDecoderInt16Overflow(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("32768"))
	defer decoder.Release()

	_, err := decoder.DecodeInt16()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewInt16OverflowError(0), err, "err must be equal to the value expected")
}

func TestDecoderNegativeInt16Overflow(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("-32768"))
	defer decoder.Release()

	_, err := decoder.DecodeInt16()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewInt16OverflowError(1), err, "err must be equal to the value expected")
}

func TestDecoderInt16Null(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("null"))
	defer decoder.Release()

	v, err := decoder.DecodeInt16()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int16(0), v, "v must be equal to the value expected")
}

func TestDecoderInt16(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte(" \n\n -32767"))
	defer decoder.Release()

	v, err := decoder.DecodeInt16()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int16(-32767), v, "v must be equal to the value expected")
}

func TestDecoderInt32Basic(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("2147483647"))
	defer decoder.Release()

	v, err := decoder.DecodeInt32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int32(2147483647), v, "v must be equal to the value expected")
}

func TestDecoderNegativeInt32Basic(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("-2147483647"))
	defer decoder.Release()

	v, err := decoder.DecodeInt32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int32(-2147483647), v, "v must be equal to the value expected")
}

func TestDecoderInt32Overflow(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("2147483648"))
	defer decoder.Release()

	_, err := decoder.DecodeInt32()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewInt32OverflowError(0), err, "err must be equal to the value expected")
}

func TestDecoderNegativeInt32Overflow(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("-2147483648"))
	defer decoder.Release()

	_, err := decoder.DecodeInt32()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewInt32OverflowError(1), err, "err must be equal to the value expected")
}

func TestDecoderInt32Null(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("null"))
	defer decoder.Release()

	v, err := decoder.DecodeInt32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int32(0), v, "v must be equal to the value expected")
}

func TestDecoderInt32(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte(" \n\n -2147483647"))
	defer decoder.Release()

	v, err := decoder.DecodeInt32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int32(-2147483647), v, "v must be equal to the value expected")
}

func TestDecoderInt64Basic(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("9223372036854775807"))
	defer decoder.Release()

	v, err := decoder.DecodeInt64()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int64(9223372036854775807), v, "v must be equal to the value expected")
}

func TestDecoderNegativeInt64Basic(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("-9223372036854775807"))
	defer decoder.Release()

	v, err := decoder.DecodeInt64()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int64(-9223372036854775807), v, "v must be equal to the value expected")
}

func TestDecoderInt64Overflow(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("9223372036854775808"))
	defer decoder.Release()

	_, err := decoder.DecodeInt64()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewInt64OverflowError(0), err, "err must be equal to the value expected")
}

func TestDecoderNegativeInt64Overflow(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("-9223372036854775808"))
	defer decoder.Release()

	_, err := decoder.DecodeInt64()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewInt64OverflowError(1), err, "err must be equal to the value expected")
}

func TestDecoderInt64Null(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("null"))
	defer decoder.Release()

	v, err := decoder.DecodeInt64()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int64(0), v, "v must be equal to the value expected")
}

func TestDecoderInt64(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte(" \n\n -9223372036854775807"))
	defer decoder.Release()

	v, err := decoder.DecodeInt64()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int64(-9223372036854775807), v, "v must be equal to the value expected")
}
