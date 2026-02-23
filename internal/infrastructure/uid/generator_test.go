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

func TestCanPregenerateNextIDs(t *testing.T) {
	generator := uid.NewIDGenerator()
	uuids := generator.GetNextValues(2)

	uuid1 := uuids[0]
	uuid2 := uuids[1]

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
