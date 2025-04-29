package clause

import "go.mongodb.org/mongo-driver/v2/mongo/options"

// Delete all documents mapping with condition in the collection
// With condition is a map[string]interface{} or bson.M take from Where method
// Warning:
// - Operation will delete all documents if condition is nil
func (c *Clause) MDelete() error {
	var opts *options.DeleteOneOptionsBuilder = options.DeleteOne()
	if c.opts != nil {
		opts = c.opts.ToDeleteOne()
	}

	res, err := c.collection.DeleteOne(c.ctx, c.condition, opts)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		c.logger.Warn("No document deleted")
	}

	return nil
}

func (c *Clause) MDeleteMany() (int64, error) {
	var opts *options.DeleteManyOptionsBuilder = options.DeleteMany()
	if c.opts != nil {
		opts = c.opts.ToDeleteMany()
	}

	res, err := c.collection.DeleteMany(c.ctx, c.condition, opts)
	if err != nil {
		return 0, err
	}

	if res.DeletedCount == 0 {
		c.logger.Warn("No document deleted")
	}

	return res.DeletedCount, nil
}
