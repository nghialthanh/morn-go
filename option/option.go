package option

import (
	"github.com/nghialthanh/morn-go/logger"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readconcern"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"go.mongodb.org/mongo-driver/v2/mongo/writeconcern"
)

type MornOption struct {
	otpMongo *options.ClientOptions
	// generator config
	// Create a new table named 'generator' to manage and control the incremental ID sequence for other collections in MongoDB
	// The table will be created in the database when the setDatabase function of the instance is executed
	IsGenID       bool
	DefaultNumber int64

	// logger config
	Logger logger.ILogger

	// field config
	CreateAtField string
	UpdateAtField string
}

type SessionOption struct {
	ReadConcern    *readconcern.ReadConcern
	ReadPreference *readpref.ReadPref
	WriteConcern   *writeconcern.WriteConcern
}

func (o *SessionOption) ToTransactionOptions() *options.TransactionOptionsBuilder {
	opts := options.Transaction()
	if o.ReadConcern != nil {
		opts = opts.SetReadConcern(o.ReadConcern)
	}
	if o.ReadPreference != nil {
		opts = opts.SetReadPreference(o.ReadPreference)
	}
	if o.WriteConcern != nil {
		opts = opts.SetWriteConcern(o.WriteConcern)
	}
	return opts
}
