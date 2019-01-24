package backend

import (
	"testing"

	"github.com/go-fish/gojson/errors"
	"github.com/stretchr/testify/assert"
)

func TestDecoderUint8Basic(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("255"))
	defer decoder.Release()

	v, err := decoder.DecodeUint8()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint8(255), v, "v must be equal to the value expected")
}

func TestDecoderUint8Overflow(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("256"))
	defer decoder.Release()

	_, err := decoder.DecodeUint8()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewUint8OverflowError(0), err, "err must be equal to the value expected")
}

func TestDecoderUint8Null(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("null"))
	defer decoder.Release()

	v, err := decoder.DecodeUint8()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint8(0), v, "v must be equal to the value expected")
}

func TestDecoderUint8(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte(" \n\n 255"))
	defer decoder.Release()

	v, err := decoder.DecodeUint8()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint8(255), v, "v must be equal to the value expected")
}
func TestDecoderUint16Basic(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("65535"))
	defer decoder.Release()

	v, err := decoder.DecodeUint16()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint16(65535), v, "v must be equal to the value expected")
}

func TestDecoderUint16Overflow(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("65536"))
	defer decoder.Release()

	_, err := decoder.DecodeUint16()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewUint16OverflowError(0), err, "err must be equal to the value expected")
}

func TestDecoderUint16Null(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("null"))
	defer decoder.Release()

	v, err := decoder.DecodeUint16()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint16(0), v, "v must be equal to the value expected")
}

func TestDecoderUint16(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte(" \n\n 65535"))
	defer decoder.Release()

	v, err := decoder.DecodeUint16()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint16(65535), v, "v must be equal to the value expected")
}

func TestDecoderUint32Basic(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("4294967295"))
	defer decoder.Release()

	v, err := decoder.DecodeUint32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint32(4294967295), v, "v must be equal to the value expected")
}

func TestDecoderUint32Overflow(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("4294967296"))
	defer decoder.Release()

	_, err := decoder.DecodeUint32()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewUint32OverflowError(0), err, "err must be equal to the value expected")
}

func TestDecoderUint32Null(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("null"))
	defer decoder.Release()

	v, err := decoder.DecodeUint32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint32(0), v, "v must be equal to the value expected")
}

func TestDecoderUint32(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte(" \n\n 4294967295"))
	defer decoder.Release()

	v, err := decoder.DecodeUint32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint32(4294967295), v, "v must be equal to the value expected")
}

func TestDecoderUint64Basic(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("18446744073709551615"))
	defer decoder.Release()

	v, err := decoder.DecodeUint64()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint64(18446744073709551615), v, "v must be equal to the value expected")
}

func TestDecoderUint64Overflow(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("18446744073709551616"))
	defer decoder.Release()

	_, err := decoder.DecodeUint64()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewUint64OverflowError(0), err, "err must be equal to the value expected")
}

func TestDecoderUint64Null(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("null"))
	defer decoder.Release()

	v, err := decoder.DecodeUint64()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint64(0), v, "v must be equal to the value expected")
}

func TestDecoderUint64(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte(" \n\n 18446744073709551615"))
	defer decoder.Release()

	v, err := decoder.DecodeUint64()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint64(18446744073709551615), v, "v must be equal to the value expected")
}
