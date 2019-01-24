package benchmark

import (
	json "encoding/json"
	"testing"

	"github.com/francoispqt/gojay"

	"github.com/buger/jsonparser"
	"github.com/go-fish/gojson"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	"github.com/stretchr/testify/assert"
	"github.com/ugorji/go/codec"
)

func nothing(_ ...interface{}) {}

var smallObject SmallPayload
var mediumObject MediumPayload
var largeObject LargePayload

func TestGOJSONUnmarshalLarge(t *testing.T) {
	var obj LargePayload

	err := gojson.Unmarshal(largeFixture, &obj)
	assert.Nil(t, err, "Err must be nil")

	var obj1 LargePayload
	err = easyjson.Unmarshal(largeFixture, &obj1)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, obj1, obj, "obj must be equal to the value expected")
}

func TestGOJSONUnmarshalMedium(t *testing.T) {
	var obj MediumPayload

	err := gojson.Unmarshal(mediumFixture, &obj)
	assert.Nil(t, err, "Err must be nil")

	var obj1 MediumPayload
	err = easyjson.Unmarshal(mediumFixture, &obj1)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, obj1, obj, "obj must be equal to the value expected")
}

func TestGOJSONMarshalLarge(t *testing.T) {
	var obj LargePayload

	err := gojson.Unmarshal(largeFixture, &obj)
	assert.Nil(t, err, "Err must be nil")

	data, err := gojson.Marshal(&obj)
	assert.Nil(t, err, "Err must be nil")

	var obj1 LargePayload
	err = easyjson.Unmarshal(data, &obj1)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, obj1, obj, "obj must be equal to the value expected")
}

func BenchmarkGOJSONUnmarshalLarge(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(largeFixture)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var obj LargePayload
		gojson.Unmarshal(largeFixture, &obj)

		for _, u := range obj.Users {
			nothing(u.Username)
		}

		for _, t := range obj.Topics.Topics {
			nothing(t.Id, t.Slug)
		}
	}
}

func BenchmarkJsonParserUnmarshalLarge(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(largeFixture)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jsonparser.ArrayEach(largeFixture, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			jsonparser.Get(value, "username")
			nothing()
		}, "users")

		jsonparser.ArrayEach(largeFixture, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			jsonparser.GetInt(value, "id")
			jsonparser.Get(value, "slug")
			nothing()
		}, "topics", "topics")
	}
}

func BenchmarkGoJayUnmarshalLarge(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(largeFixture)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var obj LargePayload
		// gojay.UnmarshalJSONObject(largeFixture, &obj)
		gojay.Unsafe.UnmarshalJSONObject(largeFixture, &obj)

		for _, u := range obj.Users {
			nothing(u.Username)
		}

		for _, t := range obj.Topics.Topics {
			nothing(t.Id, t.Slug)
		}
	}
}

func BenchmarkEasyJsonUnmarshalLarge(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(largeFixture)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lexer := &jlexer.Lexer{Data: largeFixture}
		data := new(LargePayload)
		data.UnmarshalEasyJSON(lexer)

		for _, u := range data.Users {
			nothing(u.Username)
		}

		for _, t := range data.Topics.Topics {
			nothing(t.Id, t.Slug)
		}
	}
}

func BenchmarkCodecUnmarshalLarge(b *testing.B) {

	b.ReportAllocs()
	b.SetBytes(int64(len(largeFixture)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var obj LargePayload
		decoder := codec.NewDecoderBytes(largeFixture, new(codec.JsonHandle))
		obj.CodecDecodeSelf(decoder)

		for _, u := range obj.Users {
			nothing(u.Username)
		}

		for _, t := range obj.Topics.Topics {
			nothing(t.Id, t.Slug)
		}
	}
}

func BenchmarkJSONUnmarshalLarge(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(largeFixture)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var obj LargePayload

		json.Unmarshal(largeFixture, &obj)

		for _, u := range obj.Users {
			nothing(u.Username)
		}

		for _, t := range obj.Topics.Topics {
			nothing(t.Id, t.Slug)
		}
	}
}

func BenchmarkJsonIterUnmarshalLarge(b *testing.B) {
	iter := jsoniter.ParseBytes(jsoniter.ConfigFastest, largeFixture)

	b.ReportAllocs()
	b.SetBytes(int64(len(largeFixture)))
	b.ResetTimer()

	// var json = jsoniter.ConfigFastest
	for i := 0; i < b.N; i++ {
		iter.ResetBytes(largeFixture)
		count := 0
		for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
			if "topics" != field {
				iter.Skip()
				continue
			}
			for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
				if "topics" != field {
					iter.Skip()
					continue
				}
				for iter.ReadArray() {
					iter.Skip()
					count++
				}
				break
			}
			break
		}

	}
}

func BenchmarkGOJSONUnmarshalMedium(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(mediumFixture)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var obj MediumPayload
		gojson.Unmarshal(mediumFixture, &obj)

		nothing(obj.Person.Name.FullName, obj.Person.Github.Followers, obj.Company)

		for _, el := range obj.Person.Gravatar.Avatars {
			nothing(el.Url)
		}
	}
}

func BenchmarkJsonParserUnmarshalMedium(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(mediumFixture)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jsonparser.Get(mediumFixture, "person", "name", "fullName")
		jsonparser.GetInt(mediumFixture, "person", "github", "followers")
		jsonparser.Get(mediumFixture, "company")

		jsonparser.ArrayEach(mediumFixture, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			jsonparser.Get(value, "url")
			nothing()
		}, "person", "gravatar", "avatars")
	}
}

func BenchmarkGoJayUnmarshalMedium(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(mediumFixture)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var obj MediumPayload
		gojay.UnmarshalJSONObject(mediumFixture, &obj)

		nothing(obj.Person.Name.FullName, obj.Person.Github.Followers, obj.Company)

		for _, el := range obj.Person.Gravatar.Avatars {
			nothing(el.Url)
		}
	}
}

func BenchmarkEasyJsonUnmarshalMedium(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(mediumFixture)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lexer := &jlexer.Lexer{Data: mediumFixture}
		data := new(MediumPayload)
		data.UnmarshalEasyJSON(lexer)

		nothing(data.Person.Name.FullName, data.Person.Github.Followers, data.Company)

		for _, el := range data.Person.Gravatar.Avatars {
			nothing(el.Url)
		}
	}
}

func BenchmarkCodecUnmarshalMedium(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(mediumFixture)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		decoder := codec.NewDecoderBytes(mediumFixture, new(codec.JsonHandle))
		data := new(MediumPayload)
		json.Unmarshal(mediumFixture, &data)
		data.CodecDecodeSelf(decoder)

		nothing(data.Person.Name.FullName, data.Person.Github.Followers, data.Company)

		for _, el := range data.Person.Gravatar.Avatars {
			nothing(el.Url)
		}
	}
}

func BenchmarkJSONUnmarshalMedium(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(mediumFixture)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var obj MediumPayload
		json.Unmarshal(mediumFixture, &obj)

		nothing(obj.Person.Name.FullName, obj.Person.Github.Followers, obj.Company)

		for _, el := range obj.Person.Gravatar.Avatars {
			nothing(el.Url)
		}
	}
}

func BenchmarkJsonIterUnmarshalMedium(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(mediumFixture)))
	b.ResetTimer()

	var json = jsoniter.ConfigFastest
	for i := 0; i < b.N; i++ {
		var obj MediumPayload
		json.Unmarshal(mediumFixture, &obj)

		nothing(obj.Person.Name.FullName, obj.Person.Github.Followers, obj.Company)

		for _, el := range obj.Person.Gravatar.Avatars {
			nothing(el.Url)
		}
	}
}

func BenchmarkGOJSONUnmarshalSmall(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(smallFixture)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var obj SmallPayload
		gojson.Unmarshal(smallFixture, &obj)

		nothing(obj.Uuid, obj.Tz, obj.Ua, obj.St)
	}
}

func BenchmarkJsonParserUnmarshalSmall(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(smallFixture)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jsonparser.Get(smallFixture, "uuid")
		jsonparser.GetInt(smallFixture, "tz")
		jsonparser.Get(smallFixture, "ua")
		jsonparser.GetInt(smallFixture, "st")

		nothing()
	}
}

func BenchmarkGoJayUnmarshalSmall(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(smallFixture)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var obj SmallPayload
		gojay.UnmarshalJSONObject(smallFixture, &obj)

		nothing(obj.Uuid, obj.Tz, obj.Ua, obj.St)
	}
}

func BenchmarkEasyJsonUnmarshalSmall(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(smallFixture)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lexer := &jlexer.Lexer{Data: smallFixture}
		data := new(SmallPayload)
		data.UnmarshalEasyJSON(lexer)

		nothing(data.Uuid, data.Tz, data.Ua, data.St)
	}
}

func BenchmarkCodecUnmarshalSmall(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(smallFixture)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		decoder := codec.NewDecoderBytes(smallFixture, new(codec.JsonHandle))
		data := new(SmallPayload)
		json.Unmarshal(smallFixture, &data)
		data.CodecDecodeSelf(decoder)

		nothing(data.Uuid, data.Tz, data.Ua, data.St)
	}
}

func BenchmarkJSONUnmarshalSmall(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(smallFixture)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var obj SmallPayload
		json.Unmarshal(smallFixture, &obj)

		nothing(obj.Uuid, obj.Tz, obj.Ua, obj.St)
	}
}

func BenchmarkJsonIterUnmarshalSmall(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(smallFixture)))
	b.ResetTimer()

	var json = jsoniter.ConfigFastest
	for i := 0; i < b.N; i++ {
		var obj SmallPayload
		json.Unmarshal(smallFixture, &obj)

		nothing(obj.Uuid, obj.Tz, obj.Ua, obj.St)
	}
}

func BenchmarkGOJSONMarshal(b *testing.B) {

	b.Run("large", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(largeFixture)))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			data, err := gojson.Marshal(&largeObject)
			nothing(data, err)
		}
	})

	b.Run("medium", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(mediumFixture)))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			data, err := gojson.Marshal(&mediumObject)
			nothing(data, err)
		}
	})

	b.Run("small", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(smallFixture)))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			data, err := gojson.Marshal(&smallObject)
			nothing(data, err)
		}
	})
}

func BenchmarkEasyJSONMarshal(b *testing.B) {
	b.Run("large", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(largeFixture)))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			data, err := easyjson.Marshal(&largeObject)
			nothing(data, err)
		}
	})

	b.Run("medium", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(mediumFixture)))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			data, err := easyjson.Marshal(&mediumObject)
			nothing(data, err)
		}
	})

	b.Run("small", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(smallFixture)))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			data, err := easyjson.Marshal(&smallObject)
			nothing(data, err)
		}
	})
}

func BenchmarkJSONMarshal(b *testing.B) {
	b.Run("large", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(largeFixture)))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			data, err := json.Marshal(&largeObject)
			nothing(data, err)
		}
	})

	b.Run("medium", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(mediumFixture)))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			data, err := json.Marshal(&mediumObject)
			nothing(data, err)
		}
	})

	b.Run("small", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(smallFixture)))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			data, err := json.Marshal(&smallObject)
			nothing(data, err)
		}
	})
}
