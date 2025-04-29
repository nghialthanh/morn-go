package clause

import (
	"context"

	"github.com/nghialthanh/morn-go/logger"
	"github.com/nghialthanh/morn-go/option"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Clause struct {
	// dao layer
	collection *mongo.Collection
	logger     logger.ILogger
	template   interface{}
	ctx        context.Context
	option     option.MornOption

	// clause layer
	condition interface{}
	offset    int
	limit     int
	sort      bson.M
	opts      *option.QueryOption
}

func NewClause(
	collection *mongo.Collection,
	logger logger.ILogger,
	template interface{},
	option option.MornOption,
	ctx context.Context,
) *Clause {
	return &Clause{
		collection: collection,
		logger:     logger,
		template:   template,
		option:     option,
		ctx:        ctx,
		condition:  bson.M{},
	}
}

// --------------------------------- PUBLIC METHODS ---------------------------------//
// Where set the condition for the query
// With condition is a map[string]interface{} or bson.M
func (c *Clause) Where(condition interface{}) *Clause {
	c.condition = condition
	return c
}

func (c *Clause) Limit(limit int) *Clause {
	c.limit = limit
	return c
}

func (c *Clause) Skip(offset int) *Clause {
	c.offset = offset
	return c
}

// Page set the skip and limit for the query
// This function is used briefly to replace the above 2 functions.
func (c *Clause) Page(page int, limit int) *Clause {
	c.offset = page
	c.limit = limit
	return c
}

// Sort sort the documents in the collection
// With sort is a string in the format of field:direction
// Value of direction is "asc" or "desc"
// Example: "name:asc" or "name:desc"
func (c *Clause) Sort(sort string) *Clause {
	if sort == "" {
		return c
	}
	sortFields, err := c.convSorted(sort)
	if err != nil {
		c.logger.Errorf("error convert sort: %v", err)
		return c
	}
	c.sort = sortFields
	return c
}

func (c *Clause) Option(opts option.QueryOption) *Clause {
	c.opts = &opts
	return c
}
