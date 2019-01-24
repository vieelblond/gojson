package backend

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeBytes(t *testing.T) {
	var data = []byte("Test123")

	info, err := json.Marshal(data)
	assert.Nil(t, err, "Err must be nil")

	decoder := NewDecoder()
	decoder.SetData(info)
	defer decoder.Release()

	v, err := decoder.DecodeBytes()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, data, v, "v must be equal to the value expected")
}
