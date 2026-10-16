package clause

// Count counts the number of documents in the collection
// With condition is a map[string]interface{} or bson.M take from Where method
// Offset and limit will apply to the result
// Warning:
// - Operation will run EstimatedDocumentCount if condition is nil
func (c *Clause) MCount() (int64, error) {
	var res int64
	var err error
	if c.condition == nil {
		res, err = c.collection.EstimatedDocumentCount(c.ctx)
	} else {
		res, err = c.collection.CountDocuments(c.ctx, c.condition, c.countOptions())
	}

	if err != nil {
		return 0, err
	}
	return res, nil
}
