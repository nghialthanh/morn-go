package morn

import (
	"context"
	"errors"

	"github.com/nghialthanh/morn-go/gen"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// GenerateNewKey ensures the counter document for a collection.
// A second call for the same collection leaves the existing counter in place.
func (i *Instance) GenerateNewKey(key string) error {
	if i.genDao == nil {
		return errors.New("generator dao not found")
	}

	_, err := i.genDao.Ctx(context.TODO()).MCreateOne(bson.M{
		"_id":   key,
		"value": i.optField.DefaultNumber,
	})
	return err
}

// GenIDForDao hands out the current counter for this Dao's collection.
// The stored counter becomes the next value. The returned id is the value from before that increment.
func (d *Dao) GenIDForDao() (int64, error) {
	if d.genDao == nil {
		return 0, errors.New("generator dao not found")
	}
	generatorDao := *d.genDao
	res := generatorDao.collection.FindOneAndUpdate(context.TODO(), bson.M{
		"_id": d.colName,
	}, bson.M{
		"$inc": bson.M{"value": 1},
	})
	if res == nil || res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return 0, errors.New("no document found")
		}
		return 0, res.Err()
	}

	var generator gen.Generator
	err := res.Decode(&generator)
	if err != nil {
		return 0, err
	}

	return generator.Value, nil
}
