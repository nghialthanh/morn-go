package clause

// Delete all documents mapping with condition in the collection
// With condition is a map[string]interface{} or bson.M take from Where method
// Warning:
// - Operation will delete all documents if condition is nil
func (c *Clause) MDelete() error {
	res, err := c.collection.DeleteOne(c.ctx, c.condition, c.deleteOneOptions())
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		c.logger.Warn("No document deleted")
	}

	return nil
}

func (c *Clause) MDeleteMany() (int64, error) {
	res, err := c.collection.DeleteMany(c.ctx, c.condition, c.deleteManyOptions())
	if err != nil {
		return 0, err
	}

	if res.DeletedCount == 0 {
		c.logger.Warn("No document deleted")
	}

	return res.DeletedCount, nil
}
