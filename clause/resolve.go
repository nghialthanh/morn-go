package clause

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (c *Clause) findOneOptions() *options.FindOneOptionsBuilder {
	if c.opts != nil {
		return c.opts.ToFindOne()
	}
	return options.FindOne()
}

func (c *Clause) findOptions() *options.FindOptionsBuilder {
	opts := options.Find()
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
	return opts
}

func (c *Clause) countOptions() *options.CountOptionsBuilder {
	opts := options.Count()
	if c.opts != nil {
		opts = c.opts.ToCount()
	}
	if c.offset > 0 {
		opts = opts.SetSkip(int64(c.offset))
	}
	if c.limit > 0 {
		opts = opts.SetLimit(int64(c.limit))
	}
	return opts
}

func (c *Clause) updateOneOptions() *options.UpdateOneOptionsBuilder {
	if c.opts != nil {
		return c.opts.ToUpdateOne()
	}
	return options.UpdateOne()
}

func (c *Clause) updateManyOptions() *options.UpdateManyOptionsBuilder {
	if c.opts != nil {
		return c.opts.ToUpdateMany()
	}
	return options.UpdateMany()
}

func (c *Clause) findOneAndUpdateOptions() *options.FindOneAndUpdateOptionsBuilder {
	if c.opts != nil {
		return c.opts.ToFindOneAndUpdate()
	}
	return options.FindOneAndUpdate()
}

func (c *Clause) deleteOneOptions() *options.DeleteOneOptionsBuilder {
	if c.opts != nil {
		return c.opts.ToDeleteOne()
	}
	return options.DeleteOne()
}

func (c *Clause) deleteManyOptions() *options.DeleteManyOptionsBuilder {
	if c.opts != nil {
		return c.opts.ToDeleteMany()
	}
	return options.DeleteMany()
}

func (c *Clause) insertOneOptions() *options.InsertOneOptionsBuilder {
	if c.opts != nil {
		return c.opts.ToInsertOne()
	}
	return options.InsertOne()
}

func (c *Clause) insertManyOptions() *options.InsertManyOptionsBuilder {
	if c.opts != nil {
		return c.opts.ToInsertMany()
	}
	return options.InsertMany()
}

func (c *Clause) aggregateOptions() *options.AggregateOptionsBuilder {
	if c.opts != nil {
		return c.opts.ToAggregate()
	}
	return options.Aggregate()
}

func (c *Clause) indexOptions() *options.IndexOptionsBuilder {
	if c.opts != nil {
		return c.opts.ToCreateIndex()
	}
	return options.Index()
}

func (c *Clause) setDocument(updater interface{}) (bson.M, error) {
	obj, err := c.convTypeInput(updater, c.option.UpdateAtField)
	if err != nil {
		return nil, err
	}
	return bson.M{"$set": obj}, nil
}

func (c *Clause) insertDocument(entity interface{}) (bson.M, error) {
	return c.convTypeInput(entity, c.option.CreateAtField)
}
