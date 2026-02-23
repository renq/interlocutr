package uid_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/renq/interlocutr/internal/infrastructure/uid"
	"github.com/stretchr/testify/assert"
)

func TestReturnsUUIDv7(t *testing.T) {
	generator := uid.NewIDGenerator()

	id, err := generator.Generate()

	assert.NoError(t, err)

	assert.Equal(t, uuid.Version(7), id.Version())
}

func TestUsesCustomValueIfDefined(t *testing.T) {
	definedValues := []string{"019c8c7b-c223-7408-bb50-263bf57372d3", "019c8c7d-0ec3-72d2-b113-f95eb38dbdb8"}

	generator := uid.NewIDGenerator()
	err := generator.DefineValues(definedValues...)
	assert.NoError(t, err)

	uuid1 := uuid.MustParse(definedValues[0])
	uuid2 := uuid.MustParse(definedValues[1])

	// generator returns predefined values in order
	generated1, _ := generator.Generate()
	generated2, _ := generator.Generate()

	assert.Equal(t, uuid1, generated1)
	assert.Equal(t, uuid2, generated2)

	// 3rd one is randomly generated
	generated3, _ := generator.Generate()
	assert.Equal(t, uuid.Version(7), generated3.Version())
	assert.NotEqual(t, generated1, generated3)
	assert.NotEqual(t, generated2, generated3)
}
