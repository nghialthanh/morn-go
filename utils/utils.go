package utils

import (
	"errors"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func ConvToBson(entity interface{}) (bson.M, error) {
	if entity == nil {
		return bson.M{}, nil
	}

	sel, err := bson.Marshal(entity)
	if err != nil {
		return nil, err
	}

	obj := bson.M{}
	bson.Unmarshal(sel, &obj)

	return obj, nil
}

// ConvKeyValue converts a field string to a key and value
// Example: "age:5" -> "age", "5"
// Example: "age:-1" -> "age", "-1"
// Example: "age" -> "age", ""
func ConvKeyValue(field string) (string, string, error) {
	fieldArr := strings.Split(field, ":")
	if len(fieldArr) != 2 {
		return "", "", errors.New("field must be in the format of field:value")
	}
	return fieldArr[0], fieldArr[1], nil
}

func IsStructType(template interface{}, entity interface{}) bool {
	entityType, ok := structType(entity)
	if !ok {
		return false
	}
	templateType, templateOK := structType(template)
	if !templateOK {
		return false
	}
	return entityType == templateType
}

func structType(v interface{}) (reflect.Type, bool) {
	if v == nil {
		return nil, false
	}
	value := reflect.ValueOf(v)
	if !value.IsValid() {
		return nil, false
	}
	for value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return nil, false
		}
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return nil, false
	}
	return value.Type(), true
}

func ConvSlice(entityList interface{}) ([]interface{}, error) {
	entityValue := reflect.ValueOf(entityList)

	if entityValue.Kind() == reflect.Ptr {
		if entityValue.IsNil() {
			return nil, errors.New("input is nil")
		}
		entityValue = entityValue.Elem()
	}

	if entityValue.Kind() != reflect.Slice {
		return nil, errors.New("input must be a slice")
	}
	slice := make([]interface{}, entityValue.Len())
	for i := 0; i < entityValue.Len(); i++ {
		slice[i] = entityValue.Index(i).Interface()
	}
	return slice, nil
}
