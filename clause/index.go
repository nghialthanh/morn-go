package clause

import (
	"strconv"

	"github.com/nghialthanh/morn-go/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (c *Clause) CreateIndex(index ...string) error {
	var opts *options.IndexOptionsBuilder = options.Index()
	if c.opts != nil {
		opts = c.opts.ToCreateIndex()
	}

	indexList := bson.D{}
	for _, index := range index {
		key, value, err := utils.ConvKeyValue(index)
		intValue, _ := strconv.Atoi(value)
		if err != nil {
			return err
		}
		indexList = append(indexList, bson.E{Key: key, Value: intValue})
	}

	_, err := c.collection.Indexes().CreateOne(c.ctx, mongo.IndexModel{
		Keys:    indexList,
		Options: opts,
	})
	if err != nil {
		c.logger.Error("Failed to create index", err)
	}
	return err
}
