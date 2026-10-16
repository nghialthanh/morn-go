package clause

import (
	"errors"
	"strconv"

	"github.com/nghialthanh/morn-go/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// UpdateOne updates a single document in the collection
// With updater is a map[string]interface{} or bson.M or struct of collection
// Filter is a map[string]interface{} or bson.M take from Where method
// Warning:
// - Passing a struct may reduce performance due to the use of the reflect library.
// - If pass struct please check type of field and omitempty tag
func (c *Clause) MUpdateOne(updater interface{}) error {
	updaterObj, err := c.setDocument(updater)
	if err != nil {
		return err
	}

	res, err := c.collection.UpdateOne(c.ctx, c.condition, updaterObj, c.updateOneOptions())
	if err != nil {
		return err
	}

	if res.ModifiedCount == 0 {
		c.logger.Warn("No document updated")
	}

	return nil
}

// UpdateMany updates multiple documents in the collection
// With updater is a map[string]interface{} or bson.M or struct of collection
// Filter is a map[string]interface{} or bson.M take from Where method
// Warning:
// - Passing a struct may reduce performance due to the use of the reflect library.
// - If pass struct please check type of field and omitempty tag
func (c *Clause) MUpdateMany(updater interface{}) error {
	res, err := c.updateMany(updater)
	if err != nil {
		return err
	}

	if res.ModifiedCount == 0 {
		c.logger.Warn("No document updated")
	}

	return nil
}

func (c *Clause) UpdateMany(updater interface{}) (*mongo.UpdateResult, error) {
	res, err := c.updateMany(updater)
	if err != nil {
		return nil, err
	}

	if res.ModifiedCount == 0 {
		c.logger.Warn("No document updated")
	}

	return res, nil
}

func (c *Clause) updateMany(updater interface{}) (*mongo.UpdateResult, error) {
	updaterObj, err := c.setDocument(updater)
	if err != nil {
		return nil, err
	}

	return c.collection.UpdateMany(c.ctx, c.condition, updaterObj, c.updateManyOptions())
}

// IncreaseValue increases the value of a field in the collection
// With field is a string in the format of field:value
// Example: "age:5" or "age:-1"
// Warning:
// - If condition not mapping with any document, the operation will create a new document with the field and value
func (c *Clause) MIncreaseValue(entity interface{}, field string, upsert bool) error {
	key, value, err := utils.ConvKeyValue(field)
	if err != nil {
		return err
	}
	valInt, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return errors.New("value must be a number")
	}

	opts := c.findOneAndUpdateOptions().SetUpsert(upsert)
	res := c.collection.FindOneAndUpdate(c.ctx, c.condition, bson.M{
		"$inc": bson.M{key: valInt},
	}, opts)
	if res == nil || res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return errors.New("no document found")
		}
		return res.Err()
	}

	if entity != nil {
		return c.convResultToObj(entity, res)
	}

	return nil
}

// FindOneAndUpdate finds a single document and updates it
// With updater is a map[string]interface{} or bson.M or struct of collection
// Filter is a map[string]interface{} or bson.M take from Where method
// Record after update will be returned in entity field
func (c *Clause) MFindOneAndUpdate(updater interface{}, entity interface{}) error {
	res, err := c.findOneAndUpdate(updater)
	if err != nil {
		return err
	}

	if entity != nil {
		return c.convResultToObj(entity, res)
	}

	return nil
}

func (c *Clause) FindOneAndUpdate(updater interface{}) (*mongo.SingleResult, error) {
	return c.findOneAndUpdate(updater)
}

func (c *Clause) findOneAndUpdate(updater interface{}) (*mongo.SingleResult, error) {
	updaterObj, err := c.setDocument(updater)
	if err != nil {
		return nil, err
	}

	res := c.collection.FindOneAndUpdate(c.ctx, c.condition, updaterObj, c.findOneAndUpdateOptions())
	if res == nil || res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, errors.New("no document found")
		}
		return nil, res.Err()
	}

	return res, nil
}
