package fcmul

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/OttoRoming/fastchat/pkg/fcmul/token"
	// "github.com/stretchr/testify/assert"
)

func TestUnmarshalWithoutPointer(t *testing.T) {
	s := ""
	err := Unmarshal(`"Hello, World"`, s)
	assert.Error(t, err)
}

func TestUnmarshalString(t *testing.T) {
	s := ""
	err := Unmarshal(`"Hello, World"`, &s)
	if assert.NoError(t, err) {
		assert.Equal(t, "Hello, World", s)
	}
}

func TestUnmarshalInt(t *testing.T) {
	i := -1
	err := Unmarshal("69420", &i)
	if assert.NoError(t, err) {
		assert.Equal(t, 69420, i)
	}
}

func TestUnmarshalStruct(t *testing.T){
	var data struct {
		Id int
		Username string
		Password string
	}

	err := Unmarshal(`
		{
			"Id" -> 10
			"Username" -> "Otto Roming"
			"Password" -> "passw0rd"
		}
	`, &data)

	if assert.NoError(t, err) {
		assert.Equal(t, 10, data.Id)
		assert.Equal(t, "Otto Roming", data.Username)
		assert.Equal(t, "passw0rd", data.Password)
	}
}

func TestUnmarshalStructUnusedField(t *testing.T){
	var data struct {
		Id int
		Username string
		Password string
	}

	err := Unmarshal(`
		{
			"Id" -> 10
			"Username" -> "Otto Roming"
			"Password" -> "passw0rd"
			"Password2" -> "passw0rd"
		}
	`, &data)

	if assert.NoError(t, err) {
		assert.Equal(t, 10, data.Id)
		assert.Equal(t, "Otto Roming", data.Username)
		assert.Equal(t, "passw0rd", data.Password)
	}
}

func TestUnmarshalStructMissingField(t *testing.T){
	var data struct {
		Id int
		Username string
		Password string
	}

	err := Unmarshal(`
		{
			"Id" -> 10
			"Username" -> "Otto Roming"
		}
	`, &data)

	assert.Error(t, err)
}
