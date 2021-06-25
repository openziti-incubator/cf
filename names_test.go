package cf

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSnakeCaseNameConverter(t *testing.T) {
	in := reflect.StructField{Name: "OhWow"}
	assert.Equal(t, "oh_wow", SnakeCaseNameConverter(in))
}
