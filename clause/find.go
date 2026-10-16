package clause

import "go.mongodb.org/mongo-driver/v2/mongo"

// FindOne finds a single document in the collection
// With condition is a map[string]interface{} or bson.M take from Where method
// Warning:
// - entity must be a pointer to a struct
func (c *Clause) MFindOne(entity interface{}) error {
	res := c.collection.FindOne(c.ctx, c.condition, c.findOneOptions())
	if res == nil || res.Err() != nil {
		return res.Err()
	}

	return c.convResultToObj(entity, res)
}

func (c *Clause) FindOne() (*mongo.SingleResult, error) {
	res := c.collection.FindOne(c.ctx, c.condition, c.findOneOptions())
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
	res, err := c.collection.Find(c.ctx, c.condition, c.findOptions())
	if err != nil {
		return err
	}

	return c.convResultToObj(entity, res)
}

func (c *Clause) FindMany() (*mongo.Cursor, error) {
	return c.collection.Find(c.ctx, c.condition, c.findOptions())
}
