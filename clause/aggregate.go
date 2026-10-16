package clause

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (c *Clause) MAggregate(entity interface{}, pipeline []bson.M) error {
	res, err := c.collection.Aggregate(c.ctx, pipeline, c.aggregateOptions())
	if err != nil {
		return err
	}

	return c.convResultToObj(entity, res)
}

func (c *Clause) Aggregate(pipeline []bson.M) (*mongo.Cursor, error) {
	return c.collection.Aggregate(c.ctx, pipeline, c.aggregateOptions())
}
