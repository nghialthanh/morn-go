package clause

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// --------------------------------- READ OPERATION METHODS ---------------------------------//

// FindOne finds a single document in the collection
// With condition is a map[string]interface{} or bson.M take from Where method
// Warning:
// - entity must be a pointer to a struct
func (c *Clause) MFindOne(entity interface{}) error {
	var opts *options.FindOneOptionsBuilder = options.FindOne()
	if c.opts != nil {
		opts = c.opts.ToFindOne()
	}

	res := c.collection.FindOne(c.ctx, c.condition, opts)
	if res == nil || res.Err() != nil {
		return res.Err()
	}

	err := c.convResultToObj(entity, res)
	if err != nil {
		return err
	}
	return nil
}

func (c *Clause) FindOne() (*mongo.SingleResult, error) {
	var opts *options.FindOneOptionsBuilder = options.FindOne()
	if c.opts != nil {
		opts = c.opts.ToFindOne()
	}

	res := c.collection.FindOne(c.ctx, c.condition, opts)
	if res == nil || res.Err() != nil {
		return nil, res.Err()
	}

	return res, nil
}

// FindMany finds multiple documents in the collection
// With condition is a map[string]interface{} or bson.M take from Where method
// Warning:
// - entity must be a pointer to a slice of struct
func (c *Clause) MFindMany(entity interface{}) error {
	var opts *options.FindOptionsBuilder = options.Find()
	if c.opts != nil {
		opts = c.opts.ToFind()
	}

	if c.offset > 0 {
		opts = opts.SetSkip(int64(c.offset))
	}
	if c.limit > 0 {
		opts = opts.SetLimit(int64(c.limit))
	}
	if c.sort != nil {
		opts = opts.SetSort(c.sort)
	}

	res, err := c.collection.Find(c.ctx, c.condition, opts)
	if err != nil {
		return err
	}

	err = c.convResultToObj(entity, res)
	if err != nil {
		return err
	}
	return nil
}

func (c *Clause) FindMany() (*mongo.Cursor, error) {
	var opts *options.FindOptionsBuilder = options.Find()
	if c.opts != nil {
		opts = c.opts.ToFind()
	}

	if c.offset > 0 {
		opts = opts.SetSkip(int64(c.offset))
	}
	if c.limit > 0 {
		opts = opts.SetLimit(int64(c.limit))
	}
	if c.sort != nil {
		opts = opts.SetSort(c.sort)
	}

	res, err := c.collection.Find(c.ctx, c.condition, opts)
	if err != nil {
		return nil, err
	}

	return res, nil
}
