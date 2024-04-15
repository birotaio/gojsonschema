package gojsonschema

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const Openapi string = `
{
	"type": "object",
	"properties": {
		"string": {
			"type": ["string", "null"],
			"enum": [
				"hello",
				"world",
				null
			]
		}
	}
}
	
`

func TestOpenAPI(t *testing.T) {
	var testCases = []struct {
		schema string
	}{
		{
			schema: Openapi,
		},
	}

	// Check enum is well formated

	for _, testCase := range testCases {
		schema := NewStringLoader(testCase.schema)

		res, err := GetOpenAPI(schema)
		require.NoError(t, err)

		a := &openapi3.T{}

		a.Components = &openapi3.Components{}
		a.Components.Schemas = make(openapi3.Schemas)
		a.Components.Schemas["hihi"] = openapi3.NewSchemaRef("", res)
		data, err := a.MarshalJSON()
		require.NoError(t, err)
		assert.JSONEq(t, `{"components":{"schemas":{"hihi":{"properties":{"string":{"enum":["hello","world",null],"nullable":true,"type":"string"}},"type":"object"}}},"info":null,"openapi":"","paths":null}`, string(data))
	}
}
