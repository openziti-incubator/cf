package cf

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestBasic(t *testing.T) {
	basic := &struct {
		StringValue string
	}{}

	var data = map[string]interface{}{
		"StringValue": "oh, wow!",
	}

	err := Load(data, basic)
	assert.Nil(t, err)
	assert.Equal(t, "oh, wow!", basic.StringValue)
}

func TestRenaming(t *testing.T) {
	renamed := &struct {
		SomeInt int `cf:"some_int,+required"`
	}{}

	var data = map[string]interface{}{
		"some_int": 46,
	}

	err := Load(data, renamed)
	assert.Nil(t, err)
	assert.Equal(t, 46, renamed.SomeInt)
}

func TestStringArray(t *testing.T) {
	withArray := &struct {
		StringArray []string
	}{}

	var data = map[string]interface{}{
		"StringArray": []string{"one", "two", "three"},
	}

	err := Load(data, withArray)
	assert.Nil(t, err)
	assert.EqualValues(t, []string{"one", "two", "three"}, withArray.StringArray)
}

func TestRequired(t *testing.T) {
	required := &struct {
		Required int `cf:"+required"`
	}{}

	data := make(map[string]interface{})

	err := Load(data, required)
	assert.NotNil(t, err)
}

type nestedType struct {
	Name  string
	Count int
}

func newNestedType() *nestedType {
	return &nestedType{Name: "oh, wow!", Count: 33}
}

func TestNestedPtr(t *testing.T) {
	root := &struct {
		Id     string
		Nested *nestedType
	}{}

	var data = map[string]interface{}{
		"Id": "TestNested",
		"Nested": map[string]interface{}{
			"Name": "Different",
		},
	}

	SetGlobalInstantiator(reflect.TypeOf(nestedType{}), func() interface{} { return newNestedType() })

	err := Load(data, root)
	assert.Nil(t, err)
	assert.Equal(t, "TestNested", root.Id)
	assert.NotNil(t, root.Nested)
	assert.Equal(t, "Different", root.Nested.Name)
	assert.Equal(t, 33, root.Nested.Count)
}

func TestNestedValue(t *testing.T) {
	root := &struct {
		Id     string
		Nested nestedType
	}{}

	var data = map[string]interface{}{
		"Id": "TestNested",
		"Nested": map[string]interface{}{
			"Name": "Different",
		},
	}

	SetGlobalInstantiator(reflect.TypeOf(nestedType{}), func() interface{} { return newNestedType() })

	err := Load(data, root)
	assert.Nil(t, err)
	assert.Equal(t, "TestNested", root.Id)
	assert.NotNil(t, root.Nested)
	assert.Equal(t, "Different", root.Nested.Name)
	assert.Equal(t, 33, root.Nested.Count)
}
