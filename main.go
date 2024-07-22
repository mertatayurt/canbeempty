package canbeempty

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// MarshalWithCanBeEmpty marshals a struct to JSON, considering the custom "can_be_empty" tag.
func MarshalWithCanBeEmpty(v interface{}) ([]byte, error) {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected a struct, got %v", typ.Kind())
	}

	result := make(map[string]interface{})

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i)
		jsonTag := field.Tag.Get("json")
		tags := strings.Split(jsonTag, ",")
		jsonFieldName := tags[0]
		canBeEmpty := len(tags) > 1 && tags[1] == "can_be_empty"

		if jsonFieldName == "-" {
			continue
		}

		if canBeEmpty || fieldValue.Interface() != reflect.Zero(fieldValue.Type()).Interface() {
			result[jsonFieldName] = fieldValue.Interface()
		}
	}

	return json.Marshal(result)
}
