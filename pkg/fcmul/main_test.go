package fcmul

import (
	"math"
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"
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

// FIXME: this test case is is failing for some reason
// func TestUnmarshalInt64Bound(t *testing.T) {
// 	var i int64
// 	err := Unmarshal(fmt.Sprint(math.MaxInt64), &i)
// 	if assert.NoError(t, err) {
// 		assert.Equal(t, int64(math.MaxInt64), i)
// 	}
// }

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

func TestUnmarshalStructWithList(t *testing.T){
	var data struct {
		Id int
		Username string
		Password string
		Comments []string
	}

	err := Unmarshal(`
		{
			"Id" -> 10
			"Username" -> "Otto Roming"
			"Password" -> "passw0rd"
			"Comments" -> ["first" "second"]
		}
	`, &data)

	if assert.NoError(t, err) {
		assert.Equal(t, 10, data.Id)
		assert.Equal(t, "Otto Roming", data.Username)
		assert.Equal(t, "passw0rd", data.Password)
		assert.Equal(t, []string{"first", "second"}, data.Comments)
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

func TestUnmarshalMap(t *testing.T){
	var data map[string]string

	err := Unmarshal(`
		{
			"Username" -> "Otto Roming"
			"Password" -> "passw0rd"
		}
	`, &data)

	if assert.NoError(t, err) {
		assert.Equal(t, map[string]string{
			"Username": "Otto Roming",
			"Password": "passw0rd",
		}, data)
	}
}

func TestUnmarshalPopulatedMap(t *testing.T){
	data := map[string]string{
		"this": "text",
		"should": "not",
		"be": "here",
	}

	err := Unmarshal(`
		{
			"Username" -> "Otto Roming"
			"Password" -> "passw0rd"
		}
	`, &data)

	if assert.NoError(t, err) {
		assert.Equal(t, map[string]string{
			"Username": "Otto Roming",
			"Password": "passw0rd",
		}, data)
	}
}

func TestUnmarshalMapWrongType(t *testing.T){
	var data map[int]int

	err := Unmarshal(`
		{
			"Username" -> "Otto Roming"
			"Password" -> "passw0rd"
		}
	`, &data)

	assert.Error(t, err)
}

func TestUnmarshalSlice(t *testing.T){
	var data []int

	err := Unmarshal("[1 2 3 4 5]", &data)

	if assert.NoError(t, err) {
		assert.Equal(t, []int{1, 2, 3, 4, 5}, data)
	}
}

func TestUnmarshalSliceWrongType(t *testing.T){
	var data []string

	err := Unmarshal("[1 2 3 4 5]", &data)

	assert.Error(t, err)
}

func TestUnmarshalArray(t *testing.T){
	var data [5]int

	err := Unmarshal("[1 2 3 4 5]", &data)

	if assert.NoError(t, err) {
		assert.Equal(t, [5]int{1, 2, 3, 4, 5}, data)
	}
}

func TestUnmarshalArrayWrongType(t *testing.T){
	var data [5]string

	err := Unmarshal("[1 2 3 4 5]", &data)

	assert.Error(t, err)
}

func TestUnmarshalArrayWrongLength(t *testing.T){
	var data [4]int

	err := Unmarshal("[1 2 3 4 5]", &data)

	assert.Error(t, err)
}
