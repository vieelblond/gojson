package util

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"math"
	"math/big"
	mathrand "math/rand"
	"strconv"
	"time"
	"unicode/utf8"
	"unsafe"

	"github.com/go-fish/gojson/errors"
)

func IsNumber(c byte) bool {
	return c >= 0x30 && c <= 0x39
}

func IsFloat(c byte) bool {
	return IsNumber(c) || c == 0x2E
}

func IsSkip(c byte) bool {
	return c == ' ' || c == '\n' || c == '\t' || c == '\r'
}

func IsEscaped(c byte) bool {
	return c == '\\' || c == '"' || c < 0x20 || c >= utf8.RuneSelf
}

func UnsafeConvertBytesToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

const maxInt64Length = 19
const maxInt32Length = 10
const maxInt16Length = 5
const maxInt8Length = 3
const maxUint64Length = 20
const maxUint32Length = 10
const maxUint16Length = 5
const maxUint8Length = 3

func ConvertBytesToInt8(data []byte, pos int) (int8, error) {
	var val int8

	if count := len(data); count < maxInt8Length {
		for i := 0; i < count; i++ {
			val = (val << 3) + (val << 1) + int8(data[i]-'0')
		}

		return val, nil
	} else if count == maxInt8Length {
		for i := 0; i < count; i++ {
			val = (val << 3) + (val << 1)
			num := int8(data[i] - '0')

			if int8(math.MaxInt8)-val < num {
				return 0, errors.NewInt8OverflowError(pos)
			}

			val += num
		}

		return val, nil
	}

	return 0, errors.NewInt8OverflowError(pos)
}

func ConvertBytesToInt16(data []byte, pos int) (int16, error) {
	var val int16

	if count := len(data); count < maxInt16Length {
		for i := 0; i < count; i++ {
			val = (val << 3) + (val << 1) + int16(data[i]-'0')
		}

		return val, nil
	} else if count == maxInt16Length {
		for i := 0; i < count; i++ {
			val = (val << 3) + (val << 1)
			num := int16(data[i] - '0')

			if int16(math.MaxInt16)-val < num {
				return 0, errors.NewInt16OverflowError(pos)
			}

			val += num
		}

		return val, nil
	}

	return 0, errors.NewInt16OverflowError(pos)
}

func ConvertBytesToInt32(data []byte, pos int) (int32, error) {
	var val int32

	if count := len(data); count < maxInt32Length {
		for i := 0; i < count; i++ {
			val = (val << 3) + (val << 1) + int32(data[i]-'0')
		}

		return val, nil
	} else if count == maxInt32Length {
		for i := 0; i < count; i++ {
			val = (val << 3) + (val << 1)
			num := int32(data[i] - '0')

			if int32(math.MaxInt32)-val < num {
				return 0, errors.NewInt32OverflowError(pos)
			}

			val += num
		}

		return val, nil
	}

	return 0, errors.NewInt32OverflowError(pos)
}

func ConvertBytesToInt64(data []byte, pos int) (int64, error) {
	var val int64

	if count := len(data); count < maxInt64Length {
		for i := 0; i < count; i++ {
			val = (val << 3) + (val << 1) + int64(data[i]-'0')
		}

		return val, nil
	} else if count == maxInt64Length {
		for i := 0; i < count; i++ {
			val = (val << 3) + (val << 1)
			num := int64(data[i] - '0')

			if int64(math.MaxInt64)-val < num {
				return 0, errors.NewInt64OverflowError(pos)
			}

			val += num
		}

		return val, nil
	}

	return 0, errors.NewInt64OverflowError(pos)
}

func ConvertBytesToUint8(data []byte, pos int) (uint8, error) {
	var val uint8

	if count := len(data); count < maxUint8Length {
		for i := 0; i < count; i++ {
			val = (val << 3) + (val << 1) + uint8(data[i]-'0')
		}

		return val, nil
	} else if count == maxUint8Length {
		for i := 0; i < count; i++ {
			val = (val << 3) + (val << 1)
			num := uint8(data[i] - '0')

			if uint8(math.MaxUint8)-val < num {
				return 0, errors.NewUint8OverflowError(pos)
			}

			val += num
		}

		return val, nil
	}

	return 0, errors.NewUint8OverflowError(pos)
}

func ConvertBytesToUint16(data []byte, pos int) (uint16, error) {
	var val uint16

	if count := len(data); count < maxUint16Length {
		for i := 0; i < count; i++ {
			val = (val << 3) + (val << 1) + uint16(data[i]-'0')
		}

		return val, nil
	} else if count == maxUint16Length {
		for i := 0; i < count; i++ {
			val = (val << 3) + (val << 1)
			num := uint16(data[i] - '0')

			if uint16(math.MaxUint16)-val < num {
				return 0, errors.NewUint16OverflowError(pos)
			}

			val += num
		}

		return val, nil
	}

	return 0, errors.NewUint16OverflowError(pos)
}

func ConvertBytesToUint32(data []byte, pos int) (uint32, error) {
	var val uint32

	if count := len(data); count < maxUint32Length {
		for i := 0; i < count; i++ {
			val = (val << 3) + (val << 1) + uint32(data[i]-'0')
		}

		return val, nil
	} else if count == maxUint32Length {
		for i := 0; i < count; i++ {
			val = (val << 3) + (val << 1)
			num := uint32(data[i] - '0')

			if uint32(math.MaxUint32)-val < num {
				return 0, errors.NewUint32OverflowError(pos)
			}

			val += num
		}

		return val, nil
	}

	return 0, errors.NewUint32OverflowError(pos)
}

func ConvertBytesToUint64(data []byte, pos int) (uint64, error) {
	var val uint64

	if count := len(data); count < maxInt64Length {
		for i := 0; i < count; i++ {
			val = (val << 3) + (val << 1) + uint64(data[i]-'0')
		}

		return val, nil
	} else if count == maxUint64Length {
		for i := 0; i < count; i++ {
			val = (val << 3) + (val << 1)
			num := uint64(data[i] - '0')

			if uint64(math.MaxUint64)-val < num {
				return 0, errors.NewUint64OverflowError(pos)
			}

			val += num
		}

		return val, nil
	}

	return 0, errors.NewUint64OverflowError(pos)
}

func ConvertBytesToFloat32(data []byte, pos int) (float32, error) {
	var integer, decimal uint64

	count := len(data)
	p := 0

	for i := 0; i < count; i++ {
		if data[i] == '.' {
			p = len(data) - i - 1
			continue
		}

		if p == 0 {
			integer = (integer << 3) + (integer << 1) + uint64(data[i]-'0')
			continue
		}

		integer = (integer << 3) + (integer << 1)
		decimal = (decimal << 3) + (decimal << 1) + uint64(data[i]-'0')
	}

	return float32(integer+decimal) / float32(math.Pow10(p)), nil
}

func ConvertBytesToFloat64(data []byte, pos int) (float64, error) {
	var integer, decimal uint64

	count := len(data)
	p := 0

	for i := 0; i < count; i++ {
		if data[i] == '.' {
			p = len(data) - i - 1
			continue
		}

		if p == 0 {
			integer = (integer << 3) + (integer << 1) + uint64(data[i]-'0')
			continue
		}

		integer = (integer << 3) + (integer << 1)
		decimal = (decimal << 3) + (decimal << 1) + uint64(data[i]-'0')
	}

	return float64(integer+decimal) / float64(math.Pow10(p)), nil
}

func GenerateID(prefix string) string {
	b := make([]byte, 8)
	for {
		if _, err := io.ReadFull(rand.Reader, b); err != nil {
			panic(err) // This shouldn't happen
		}
		id := hex.EncodeToString(b)
		// if we try to parse the truncated for as an int and we don't have
		// an error then the value is all numeric and causes issues when
		// used as a hostname. ref #3869
		if _, err := strconv.ParseInt(id, 10, 64); err == nil {
			continue
		}

		return prefix + id
	}
}

func init() {
	// safely set the seed globally so we generate random ids. Tries to use a
	// crypto seed before falling back to time.
	var seed int64
	if cryptoseed, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64)); err != nil {
		// This should not happen, but worst-case fallback to time-based seed.
		seed = time.Now().UnixNano()
	} else {
		seed = cryptoseed.Int64()
	}

	mathrand.Seed(seed)
}
