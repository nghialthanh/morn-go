package clause

import (
	"errors"
	"strconv"

	"github.com/nghialthanh/morn-go/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// UpdateOne updates a single document in the collection
// With updater is a map[string]interface{} or bson.M or struct of collection
// Filter is a map[string]interface{} or bson.M take from Where method
// Warning:
// - Passing a struct may reduce performance due to the use of the reflect library.
// - If pass struct please check type of field and omitempty tag
func (c *Clause) MUpdateOne(updater interface{}) error {
	var opts *options.UpdateOneOptionsBuilder = options.UpdateOne()
	if c.opts != nil {
		opts = c.opts.ToUpdateOne()
	}

	updateField := ""
	if c.option.UpdateAtField != "" {
		updateField = c.option.UpdateAtField
	}

	updaterObj, err := c.convTypeInput(updater, updateField)
	if err != nil {
		return err
	}

	updaterObj = bson.M{
		"$set": updaterObj,
	}

	res, err := c.collection.UpdateOne(c.ctx, c.condition, updaterObj, opts)

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
	var opts *options.UpdateManyOptionsBuilder = options.UpdateMany()
	if c.opts != nil {
		opts = c.opts.ToUpdateMany()
	}

	updateField := ""
	if c.option.UpdateAtField != "" {
		updateField = c.option.UpdateAtField
	}

	updaterObj, err := c.convTypeInput(updater, updateField)
	if err != nil {
		return err
	}

	updaterObj = bson.M{
		"$set": updaterObj,
	}

	res, err := c.collection.UpdateMany(c.ctx, c.condition, updaterObj, opts)
	if err != nil {
		return err
	}

	if res.ModifiedCount == 0 {
		c.logger.Warn("No document updated")
	}

	return nil
}

func (c *Clause) UpdateMany(updater interface{}) (*mongo.UpdateResult, error) {
	var opts *options.UpdateManyOptionsBuilder = options.UpdateMany()
	if c.opts != nil {
		opts = c.opts.ToUpdateMany()
	}

	updateField := ""
	if c.option.UpdateAtField != "" {
		updateField = c.option.UpdateAtField
	}

	updaterObj, err := c.convTypeInput(updater, updateField)
	if err != nil {
		return nil, err
	}

	updaterObj = bson.M{
		"$set": updaterObj,
	}

	res, err := c.collection.UpdateMany(c.ctx, c.condition, updaterObj, opts)
	if err != nil {
		return nil, err
	}

	if res.ModifiedCount == 0 {
		c.logger.Warn("No document updated")
	}

	return res, nil
}

// IncreaseValue increases the value of a field in the collection
// With field is a string in the format of field:value
// Example: "age:5" or "age:-1"
// Warning:
// - If condition not mapping with any document, the operation will create a new document with the field and value
func (c *Clause) MIncreaseValue(entity interface{}, field string, upsert bool) error {
	var opts *options.FindOneAndUpdateOptionsBuilder = options.FindOneAndUpdate()
	if c.opts != nil {
		opts = c.opts.ToFindOneAndUpdate()
	}

	key, value, err := utils.ConvKeyValue(field)
	if err != nil {
		return err
	}
	valInt, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return errors.New("value must be a number")
	}

	opts = opts.SetUpsert(upsert)
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
		err = c.convResultToObj(entity, res)
		if err != nil {
			return err
		}
	}

	return nil
}

// FindOneAndUpdate finds a single document and updates it
// With updater is a map[string]interface{} or bson.M or struct of collection
// Filter is a map[string]interface{} or bson.M take from Where method
// Record after update will be returned in entity field
func (c *Clause) MFindOneAndUpdate(updater interface{}, entity interface{}) error {

	var opts *options.FindOneAndUpdateOptionsBuilder = options.FindOneAndUpdate()
	if c.opts != nil {
		opts = c.opts.ToFindOneAndUpdate()
	}

	updateField := ""
	if c.option.UpdateAtField != "" {
		updateField = c.option.UpdateAtField
	}

	updaterObj, err := c.convTypeInput(updater, updateField)
	if err != nil {
		return err
	}

	updaterObj = bson.M{
		"$set": updaterObj,
	}

	res := c.collection.FindOneAndUpdate(c.ctx, c.condition, updaterObj, opts)
	if res == nil || res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return errors.New("no document found")
		}
		return res.Err()
	}

	if entity != nil {
		err = c.convResultToObj(entity, res)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Clause) FindOneAndUpdate(updater interface{}) (*mongo.SingleResult, error) {

	var opts *options.FindOneAndUpdateOptionsBuilder = options.FindOneAndUpdate()
	if c.opts != nil {
		opts = c.opts.ToFindOneAndUpdate()
	}

	updateField := ""
	if c.option.UpdateAtField != "" {
		updateField = c.option.UpdateAtField
	}

	updaterObj, err := c.convTypeInput(updater, updateField)
	if err != nil {
		return nil, err
	}

	updaterObj = bson.M{
		"$set": updaterObj,
	}

	res := c.collection.FindOneAndUpdate(c.ctx, c.condition, updaterObj, opts)
	if res == nil || res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, errors.New("no document found")
		}
		return nil, res.Err()
	}

	return res, nil
}
