package morn

import (
	"context"
	"errors"

	"github.com/nghialthanh/morn-go/clause"
	"github.com/nghialthanh/morn-go/gen"
	"github.com/nghialthanh/morn-go/logger"
	"github.com/nghialthanh/morn-go/option"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Dao struct {
	collection *mongo.Collection
	client     *mongo.Client
	colName    string
	template   interface{}
	logger     logger.ILogger

	option option.MornOption
	genDao *Dao
}

func NewDao(colName string, template interface{}, ins *Instance, opt *option.MornOption) *Dao {
	optionDao := ins.GetOptsField()
	if opt != nil {
		optionDao = *opt
	}
	if optionDao.IsGenID {
		ins.GetLogger().Infof("Generate ID for collection %s", colName)
		ins.GenerateNewKey(colName)
	}
	return &Dao{
		colName:    colName,
		template:   template,
		collection: ins.GetDB().Collection(colName),
		option:     optionDao,
		logger:     ins.GetLogger(),
		genDao:     ins.GetDao(),
		client:     ins.GetClient(),
	}
}

func (d *Dao) Clause() *clause.Clause {
	return clause.NewClause(
		d.collection,
		d.logger,
		d.template,
		d.option,
		context.TODO(),
	)
}

func (d *Dao) Ctx(ctx context.Context) *clause.Clause {
	if ctx == nil {
		ctx = context.TODO()
	}
	if d.collection == nil {
		d.logger.Error("Collection not connected")
		return nil
	}
	return clause.NewClause(
		d.collection,
		d.logger,
		d.template,
		d.option,
		ctx,
	)
}

// ----------------------- Get/Set --------------------------//
func (d *Dao) Col() *mongo.Collection {
	return d.collection
}

func (d *Dao) GetOptionField() *option.MornOption {
	return &d.option
}

// func (d *Dao) SetOptionField(opt *option.MornOption) {
// 	currOpt := d.option
// 	if opt.CreateAtField != "" {
// 		currOpt.CreateAtField = opt.CreateAtField
// 	}
// 	if opt.UpdateAtField != "" {
// 		currOpt.UpdateAtField = opt.UpdateAtField
// 	}
// 	if opt.IsGenID != nil {
// 		currOpt.IsGenID = opt.IsGenID
// 	}
// 	if opt.DefaultNumber != nil {
// 		currOpt.DefaultNumber = opt.DefaultNumber
// 	}
// 	d.option = currOpt
// }

// GenIDForDao is used to generate a new ID for the collection
// This function will update the value of the document with the _id is the collection name
// and then return the new value by plus 1
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

// Session is a function that starts a session and executes a function with the session context
func (d *Dao) Session(ctx context.Context, f func(ctx context.Context) error, opt *option.SessionOption) error {
	session, err := d.client.StartSession()
	if err != nil {
		d.logger.Error("Failed to start session", err)
		return err
	}
	defer session.EndSession(ctx)

	txnOptions := options.Transaction()
	if opt != nil {
		txnOptions = opt.ToTransactionOptions()
	}

	err = mongo.WithSession(ctx, session, func(ctxSession context.Context) error {
		err := session.StartTransaction(txnOptions)
		if err != nil {
			return err
		}

		err = f(ctxSession)
		if err != nil {
			session.AbortTransaction(ctxSession)
			return err
		}
		// Commit the transaction
		return session.CommitTransaction(ctxSession)
	})

	return err
}
