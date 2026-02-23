package uid

import "github.com/google/uuid"

type IDGenerator struct {
	definedValues []uuid.UUID
}

func NewIDGenerator() IDGenerator {
	return IDGenerator{
		definedValues: make([]uuid.UUID, 0),
	}
}

func (g *IDGenerator) Generate() (uuid.UUID, error) {
	if len(g.definedValues) > 0 {
		defer func() {
			g.definedValues = g.definedValues[1:]
		}()

		return g.definedValues[0], nil
	}

	return uuid.NewV7()
}

func (g *IDGenerator) GetNextValues(n int) []uuid.UUID {
	result := make([]uuid.UUID, n)
	for k := range n {
		id, _ := uuid.NewV7()
		g.definedValues = append(g.definedValues, id)
		result[k] = id
	}

	return result
}
