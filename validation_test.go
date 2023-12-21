// Copyright 2015 xeipuuv ( https://github.com/xeipuuv )
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// author           xeipuuv
// author-github    https://github.com/xeipuuv
// author-mail      xeipuuv@gmail.com
//
// repository-name  gojsonschema
// repository-desc  An implementation of JSON Schema, based on IETF's draft v4 - Go language.
//
// description      Extends Schema and subSchema, implements the validation phase.
//
// created          28-02-2013

package gojsonschema

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const SchemaTest string = `
{
	"type": "object",
	"properties": {
		"string_readonly": {
			"type": ["string", "null"],
			"readOnly": true
		},
		"integer_readonly": {
			"type": ["integer", "null"],
			"readOnly": true
		},
		"number_readonly": {
			"type": ["number", "null"],
			"readOnly": true
		},
		"object_readonly": {
			"type": ["object", "null"],
			"properties": {},
			"readOnly": true
		},
		"array_readonly": {
			"type": ["array", "null"],
			"items": {
				"type": "number"
			},	
			"readOnly": true
		}
	}
}
	
`

func TestValidation(t *testing.T) {
	var testCases = []struct {
		schema      string
		json        string
		errorNumber int
	}{
		{
			schema: SchemaTest,
			json: `{
				"string_readonly": "update"
			}`,
			errorNumber: 1,
		},
		{
			schema: SchemaTest,
			json: `{
				"string_readonly": null
			}`,
			errorNumber: 1,
		},
		{
			schema:      SchemaTest,
			json:        `{}`,
			errorNumber: 0,
		},
		{
			schema: SchemaTest,
			json: `{
				"integer_readonly": 1
			}`,
			errorNumber: 1,
		},
		{
			schema: SchemaTest,
			json: `{
				"integer_readonly": null
			}`,
			errorNumber: 1,
		},
		{
			schema:      SchemaTest,
			json:        `{}`,
			errorNumber: 0,
		},
		{
			schema: SchemaTest,
			json: `{
				"number_readonly": 1.5
			}`,
			errorNumber: 1,
		},
		{
			schema: SchemaTest,
			json: `{
				"number_readonly": null
			}`,
			errorNumber: 1,
		},
		{
			schema:      SchemaTest,
			json:        `{}`,
			errorNumber: 0,
		},
		{
			schema: SchemaTest,
			json: `{
				"object_readonly": {}
			}`,
			errorNumber: 1,
		},
		{
			schema: SchemaTest,
			json: `{
				"object_readonly": null
			}`,
			errorNumber: 1,
		},
		{
			schema:      SchemaTest,
			json:        `{}`,
			errorNumber: 0,
		},
	}

	for _, testCase := range testCases {
		schema := NewStringLoader(testCase.schema)
		data := NewStringLoader(testCase.json)

		res, err := Validate(schema, data)
		require.NoError(t, err)
		assert.Len(t, res.errors, testCase.errorNumber)
		if testCase.errorNumber > 0 {
			assert.False(t, res.Valid())
		} else {
			assert.True(t, res.Valid())
		}
	}
}
