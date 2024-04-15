package gojsonschema

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func GetOpenAPI(ls JSONLoader) (*openapi3.Schema, error) {
	schema, err := NewSchema(ls)
	if err != nil {
		return nil, err
	}
	oapi := schema.GetOpenAPI()
	return oapi, nil
}

func (d *Schema) GetOpenAPI() *openapi3.Schema {
	return recursiveOpenApi(d.rootSchema)
}

func recursiveOpenApi(currentSubSchema *subSchema) *openapi3.Schema {
	var schema *openapi3.Schema
	if currentSubSchema.types.Contains(TYPE_OBJECT) {
		// Loop over properties and call this function
		schema = openapi3.NewObjectSchema()
		for _, property := range currentSubSchema.propertiesChildren {
			subSchema := recursiveOpenApi(property)
			schema.WithProperty(property.property, subSchema)
		}
	} else if currentSubSchema.types.Contains(TYPE_ARRAY) {
		schema = openapi3.NewArraySchema()
		item := recursiveOpenApi(currentSubSchema.itemsChildren[0])
		schema.WithItems(item)
	} else if currentSubSchema.types.Contains(TYPE_BOOLEAN) {
		schema = openapi3.NewBoolSchema()
	} else if currentSubSchema.types.Contains(TYPE_NUMBER) {
		schema = openapi3.NewFloat64Schema()
	} else if currentSubSchema.types.Contains(TYPE_INTEGER) {
		schema = openapi3.NewIntegerSchema()
	} else if currentSubSchema.types.Contains(TYPE_STRING) {
		schema = openapi3.NewStringSchema()

		if currentSubSchema.format != "" {
			schema.WithFormat(currentSubSchema.format)
		}
	}

	if schema != nil {
		if len(currentSubSchema.rawEnum) > 0 {
			schema.WithEnum(currentSubSchema.rawEnum...)
		}
		if currentSubSchema.readOnly {
			schema.ReadOnly = true
		}
		if currentSubSchema.types.Contains(TYPE_NULL) {
			schema.Nullable = true
		}
	}

	return schema
}
