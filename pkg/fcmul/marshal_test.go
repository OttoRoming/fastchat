package fcmul

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalString(t *testing.T) {
	s := "Hello, World"
	str, err := Marshal(s)
	if assert.NoError(t, err) {
		assert.Equal(t, `"Hello, World"`, str)
	}
}

func TestMarshalInt(t *testing.T) {
	i := 69420
	str, err := Marshal(i)
	if assert.NoError(t, err) {
		assert.Equal(t, "69420", str)
	}
}

func TestMarshalInt64Bound(t *testing.T) {
	var i int64 = math.MaxInt64
	str, err := Marshal(i)
	if assert.NoError(t, err) {
		assert.Equal(t, fmt.Sprint(math.MaxInt64), str)
	}
}

func TestMarshalStruct(t *testing.T) {
	type User struct {
		Id       int
		Username string
		Password string
	}

	data := User{
		Id:       10,
		Username: "Otto Roming",
		Password: "passw0rd",
	}

	str, err := Marshal(data)

	var data2 User
	err = Unmarshal(str, &data2)

	if assert.NoError(t, err) {
		assert.Equal(t, data, data2)
	}
}

func TestMarshalStructWithSlice(t *testing.T) {
	type User struct {
		Id       int
		Username string
		Password string
		Comments []string
	}

	data := User{
		Id:       10,
		Username: "Otto Roming",
		Password: "passw0rd",
		Comments: []string{"first", "second"},
	}

	str, err := Marshal(data)

	var data2 User
	err = Unmarshal(str, &data2)

	if assert.NoError(t, err) {
		assert.Equal(t, data, data2)
	}
}
