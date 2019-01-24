package backend

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecoderFloat32Basic(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("127.11"))
	defer decoder.Release()

	v, err := decoder.DecodeFloat32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, float32(127.11), v, "v must be equal to the value expected")
}

func TestDecoderNegativeFloat32Basic(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("-127.11"))
	defer decoder.Release()

	v, err := decoder.DecodeFloat32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, float32(-127.11), v, "v must be equal to the value expected")
}

func TestDecoderFloat32Null(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("null"))
	defer decoder.Release()

	v, err := decoder.DecodeFloat32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, float32(0), v, "v must be equal to the value expected")
}

func TestDecoderFloat32(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte(" \n\n 127"))
	defer decoder.Release()

	v, err := decoder.DecodeFloat32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, float32(127), v, "v must be equal to the value expected")
}

func TestReadFloat(t *testing.T) {
	decoder := NewDecoder()
	decoder.SetData([]byte("127.11"))
	defer decoder.Release()

	data, err := decoder.ReadFloat()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, "127.11", string(data), "v must be equal to the value expected")
}
