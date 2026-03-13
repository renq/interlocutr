package tests

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func bufferToJson(t *testing.T, body *bytes.Buffer) map[string]any {
	var responseBody map[string]any
	if e := json.Unmarshal(body.Bytes(), &responseBody); e != nil {
		assert.NoError(t, e, "response is not valid json: %s", body.String())
	}

	return responseBody
}

func bufferToStruct(t *testing.T, body *bytes.Buffer, out any) {
	if e := json.Unmarshal(body.Bytes(), out); e != nil {
		assert.NoError(t, e)
	}
}
