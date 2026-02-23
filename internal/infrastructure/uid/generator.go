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

func (g *IDGenerator) DefineValues(values ...string) error {
	for _, v := range values {
		parsed, err := uuid.Parse(v)
		if err != nil {
			return err
		}
		g.definedValues = append(g.definedValues, parsed)
	}

	return nil
}
