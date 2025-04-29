package clause

import (
	"github.com/nghialthanh/morn-go/utils"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// CreateOne creates a single document in the collection
// With entity is a map[string]interface{} or bson.M or struct of collection
// Warning:
// - Passing a struct may reduce performance due to the use of the reflect library.
func (c *Clause) MCreateOne(entity interface{}) (interface{}, error) {
	var opts *options.InsertOneOptionsBuilder = options.InsertOne()
	if c.opts != nil {
		opts = c.opts.ToInsertOne()
	}

	createField := ""
	if c.option.CreateAtField != "" {
		createField = c.option.CreateAtField
	}

	obj, err := c.convTypeInput(entity, createField)
	if err != nil {
		return nil, err
	}

	res, err := c.collection.InsertOne(c.ctx, obj, opts)

	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}

// CreateMany insert many object into db
// With entityList is a slice of map[string]interface{} or bson.M or struct of collection
// Warning:
// - Passing a struct may reduce performance due to the use of the reflect library.
func (c *Clause) MCreateMany(entityList interface{}) ([]interface{}, error) {
	var opts *options.InsertManyOptionsBuilder = options.InsertMany()
	if c.opts != nil {
		opts = c.opts.ToInsertMany()
	}

	list, err := utils.ConvSlice(entityList)
	if err != nil {
		return nil, err
	}

	createField := ""
	if c.option.CreateAtField != "" {
		createField = c.option.CreateAtField
	}

	var objList []interface{}
	for _, item := range list {
		obj, err := c.convTypeInput(item, createField)
		if err != nil {
			return nil, err
		}

		objList = append(objList, obj)
	}

	res, err := c.collection.InsertMany(c.ctx, objList, opts)
	if err != nil {
		return nil, err
	}

	return res.InsertedIDs, nil
}
