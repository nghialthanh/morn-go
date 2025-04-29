package morn

import (
	"context"
	"errors"

	"github.com/nghialthanh/morn-go/gen"
	"github.com/nghialthanh/morn-go/logger"
	"github.com/nghialthanh/morn-go/option"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Instance struct {
	db     *mongo.Database
	client *mongo.Client

	dao      *Dao
	optField *option.MornOption
	logger   logger.ILogger
	genDao   *Dao
}

// SetupMongo with default options
// If you want to use custom options, you can config inside url or manual setup by SetupMongoByClient method
func SetupMongoByURI(uri string, opts *option.MornOption) (*Instance, error) {
	if uri == "" {
		return nil, errors.New("uri is required")
	}

	ins := Instance{
		optField: opts,
		logger:   opts.Logger,
	}

	if opts.Logger == nil {
		ins.logger = logger.NewFmtLogger()
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	client, err := mongo.Connect(options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI))
	if err != nil {
		ins.logger.Error("Failed to connect to MongoDB", err)
		return nil, err
	}

	ins.client = client

	return &ins, nil
}

// SetupManual with custom options
func (i *Instance) SetupMongoByClient(client *mongo.Client, opts *option.MornOption) *Instance {
	ins := Instance{
		optField: opts,
		logger:   opts.Logger,
	}
	ins.client = client
	return &ins
}

func (i *Instance) SetDB(db string) *Instance {
	i.db = i.client.Database(db)
	if i.optField.IsGenID {
		i.genDao = &Dao{
			colName:    "generator",
			template:   gen.Generator{},
			collection: i.GetDB().Collection("generator"),
			option:     i.GetOptsField(),
			logger:     i.GetLogger(),
		}
	}
	return i
}

func (i *Instance) Disconnect() error {
	return i.client.Disconnect(context.TODO())
}

func (i *Instance) GetDB() *mongo.Database {
	return i.db
}

func (i *Instance) GetOptsField() option.MornOption {
	return *i.optField
}

func (i *Instance) GetLogger() logger.ILogger {
	return i.logger
}

func (i *Instance) GetClient() *mongo.Client {
	return i.client
}

func (i *Instance) GetDao() *Dao {
	return i.genDao
}

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
