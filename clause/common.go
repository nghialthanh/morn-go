package clause

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/nghialthanh/morn-go/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// --------------------------------- PRIVATE METHODS ---------------------------------//
func (c *Clause) convTypeInput(entity interface{}, optsField string) (bson.M, error) {
	timeNow := time.Now()
	var obj bson.M
	switch entity.(type) {
	case bson.M:
		obj = entity.(bson.M)
	case map[string]interface{}:
		if optsField != "" {
			entity.(map[string]interface{})[optsField] = timeNow
		}
		result := make(bson.M)
		for key, value := range entity.(map[string]interface{}) {
			result[key] = value
		}
		obj = result
	default:
		if !utils.IsStructType(c.template, entity) {
			return nil, errors.New("struct of entity is not match with template")
		}
		bsonObj, err := utils.ConvToBson(entity)
		if err != nil {
			return nil, err
		}
		if optsField != "" {
			bsonObj[optsField] = timeNow
		}
		obj = bsonObj
	}
	return obj, nil
}

func (c *Clause) convSorted(sort string) (bson.M, error) {
	key, value, err := utils.ConvKeyValue(sort)
	if err != nil {
		return nil, err
	}

	if value != "asc" && value != "desc" {
		return nil, errors.New("direction must be either asc or desc")
	}
	valueSorted := 1
	if value == "desc" {
		valueSorted = -1
	}

	return bson.M{key: valueSorted}, nil
}

func (c *Clause) convResultToObj(obj interface{}, result interface{}) error {
	ctx := c.ctx
	if ctx == nil {
		ctx = context.Background()
	}

	// Ensure obj is a pointer
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() != reflect.Ptr {
		return fmt.Errorf("obj must be a pointer")
	}
	if objValue.IsNil() {
		return fmt.Errorf("obj cannot be nil")
	}
	objElem := objValue.Elem()

	// Handle different result types
	switch res := result.(type) {
	case *mongo.Cursor:
		defer res.Close(ctx)
		switch objElem.Kind() {
		case reflect.Slice:
			sliceType := objElem.Type()
			slice := reflect.MakeSlice(sliceType, 0, 0)

			for res.Next(ctx) {
				elem := reflect.New(sliceType.Elem()).Interface()
				if err := res.Decode(elem); err != nil {
					return fmt.Errorf("failed to decode cursor into slice element: %v", err)
				}
				slice = reflect.Append(slice, reflect.ValueOf(elem).Elem())
			}
			objElem.Set(slice)

		case reflect.Struct:
			if res.Next(ctx) {
				if err := res.Decode(obj); err != nil {
					return fmt.Errorf("failed to decode cursor into struct: %v", err)
				}
			} else {
				return fmt.Errorf("no documents found in cursor")
			}

		default:
			return fmt.Errorf("unsupported obj type for cursor: %v", objElem.Kind())
		}
		if err := res.Err(); err != nil {
			return fmt.Errorf("cursor error: %v", err)
		}

	case *mongo.SingleResult:
		if objElem.Kind() != reflect.Struct {
			return fmt.Errorf("SingleResult requires obj to be a struct, got: %v", objElem.Kind())
		}
		if err := res.Decode(obj); err != nil {
			return fmt.Errorf("failed to decode SingleResult: %v", err)
		}

	case *mongo.InsertOneResult:
		switch objElem.Kind() {
		case reflect.String:
			if id, ok := res.InsertedID.(string); ok {
				objElem.SetString(id)
			} else if oid, ok := res.InsertedID.(bson.ObjectID); ok {
				objElem.SetString(oid.Hex())
			} else {
				return fmt.Errorf("InsertedID is not a string or ObjectID: %T", res.InsertedID)
			}
		case reflect.Struct:
			// If obj has an ID field, set it
			if idField := objElem.FieldByName("ID"); idField.IsValid() && idField.CanSet() {
				if id, ok := res.InsertedID.(string); ok {
					idField.SetString(id)
				} else if oid, ok := res.InsertedID.(bson.ObjectID); ok {
					idField.SetString(oid.Hex())
				} else {
					return fmt.Errorf("InsertedID is not compatible with ID field: %T", res.InsertedID)
				}
			} else {
				return fmt.Errorf("obj struct has no settable ID field")
			}
		default:
			return fmt.Errorf("unsupported obj type for InsertOneResult: %v", objElem.Kind())
		}

	default:
		return fmt.Errorf("unsupported result type: %T", result)
	}

	return nil
}
