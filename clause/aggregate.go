package clause

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (c *Clause) MAggregate(entity interface{}, pipeline []bson.M) error {
	var opts *options.AggregateOptionsBuilder = options.Aggregate()
	if c.opts != nil {
		opts = c.opts.ToAggregate()
	}

	res, err := c.collection.Aggregate(c.ctx, pipeline, opts)
	if err != nil {
		return err
	}

	err = c.convResultToObj(entity, res)
	if err != nil {
		return err
	}

	return nil
}

func (c *Clause) Aggregate(pipeline []bson.M) (*mongo.Cursor, error) {
	var opts *options.AggregateOptionsBuilder = options.Aggregate()
	if c.opts != nil {
		opts = c.opts.ToAggregate()
	}

	res, err := c.collection.Aggregate(c.ctx, pipeline, opts)
	if err != nil {
		return nil, err
	}

	return res, nil
}
