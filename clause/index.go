package clause

import (
	"strconv"

	"github.com/nghialthanh/morn-go/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (c *Clause) CreateIndex(index ...string) error {
	indexList := bson.D{}
	for _, index := range index {
		key, value, err := utils.ConvKeyValue(index)
		if err != nil {
			return err
		}
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		indexList = append(indexList, bson.E{Key: key, Value: intValue})
	}

	_, err := c.collection.Indexes().CreateOne(c.ctx, mongo.IndexModel{
		Keys:    indexList,
		Options: c.indexOptions(),
	})
	if err != nil {
		c.logger.Error("Failed to create index", err)
	}
	return err
}
