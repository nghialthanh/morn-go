package clause

import "go.mongodb.org/mongo-driver/v2/mongo/options"

// Count counts the number of documents in the collection
// With condition is a map[string]interface{} or bson.M take from Where method
// Offset and limit will apply to the result
// Warning:
// - Operation will run EstimatedDocumentCount if condition is nil
func (c *Clause) MCount() (int64, error) {
	var opts *options.CountOptionsBuilder = options.Count()
	if c.opts != nil {
		opts = c.opts.ToCount()
	}

	var res int64
	var err error
	if c.condition == nil {
		res, err = c.collection.EstimatedDocumentCount(c.ctx)
	} else {
		if c.offset > 0 {
			opts = options.Count().SetSkip(int64(c.offset))
		}
		if c.limit > 0 {
			opts = options.Count().SetLimit(int64(c.limit))
		}
		res, err = c.collection.CountDocuments(c.ctx, c.condition, opts)
	}

	if err != nil {
		return 0, err
	}
	return res, nil
}
