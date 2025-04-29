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
	entityValue := reflect.ValueOf(entity)

	if entityValue.Kind() == reflect.Invalid {
		return false
	}

	// Handle pointer case
	if entityValue.Kind() == reflect.Ptr {
		if entityValue.IsNil() {
			return false
		}
		entityValue = entityValue.Elem()
	}

	isStruct := entityValue.Kind() == reflect.Struct

	return isStruct
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
