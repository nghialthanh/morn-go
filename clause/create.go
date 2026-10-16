package clause

import "github.com/nghialthanh/morn-go/utils"

// CreateOne creates a single document in the collection
// With entity is a map[string]interface{} or bson.M or struct of collection
// Warning:
// - Passing a struct may reduce performance due to the use of the reflect library.
func (c *Clause) MCreateOne(entity interface{}) (interface{}, error) {
	obj, err := c.insertDocument(entity)
	if err != nil {
		return nil, err
	}

	res, err := c.collection.InsertOne(c.ctx, obj, c.insertOneOptions())
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
	list, err := utils.ConvSlice(entityList)
	if err != nil {
		return nil, err
	}

	var objList []interface{}
	for _, item := range list {
		obj, err := c.insertDocument(item)
		if err != nil {
			return nil, err
		}
		objList = append(objList, obj)
	}

	res, err := c.collection.InsertMany(c.ctx, objList, c.insertManyOptions())
	if err != nil {
		return nil, err
	}

	return res.InsertedIDs, nil
}
