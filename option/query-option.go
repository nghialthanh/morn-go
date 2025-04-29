package option

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type QueryOption struct {
	AllowPartialResults *bool              // Find, FindOne
	Collation           *options.Collation // Find, FindOne, FindOneAndUpdate, DeleteOne, DeleteMany, Count, UpdateOne, UpdateMany, Aggregate, CreateIndex
	Comment             interface{}        // Find, FindOne, FindOneAndUpdate, DeleteOne, DeleteMany, Count, UpdateOne, UpdateMany, InsertOne, InsertMany, Aggregate,
	Hint                interface{}        // Find, FindOne, FindOneAndUpdate, DeleteOne, DeleteMany, Count, UpdateOne, UpdateMany, Aggregate,
	Max                 interface{}        // Find, FindOne, CreateIndex(float64)
	MaxAwaitTime        *time.Duration     // Find, Aggregate,
	Min                 interface{}        // Find, FindOne, CreateIndex(float64)
	OplogReplay         *bool              // Find, FindOne,
	Projection          interface{}        // Find, FindOne, FindOneAndUpdate,
	ReturnKey           *bool              // Find, FindOne,
	ShowRecordID        *bool              // Find, FindOne,
	Skip                *int64             // Find, FindOne, Count,
	Sort                interface{}        // Find, FindOne, UpdateOne, FindOneAndUpdate,

	AllowDiskUse    *bool               // Find, Aggregate,
	BatchSize       *int32              // Find, Aggregate,
	CursorType      *options.CursorType // Find,
	Let             interface{}         // Find, DeleteOne, DeleteMany, UpdateOne, UpdateMany, Aggregate, FindOneAndUpdate,
	Limit           *int64              // Find, Count,
	NoCursorTimeout *bool               // Find,

	ArrayFilters             []interface{} // UpdateOne, UpdateMany, FindOneAndUpdate,
	BypassDocumentValidation *bool         // UpdateOne, UpdateMany, InsertOne, InsertMany, Aggregate, FindOneAndUpdate,
	Upsert                   *bool         // UpdateOne, UpdateMany, FindOneAndUpdate,

	Ordered *bool // InsertMany,

	Custom bson.M // Aggregate,

	ReturnDocument *options.ReturnDocument // FindOneAndUpdate,

	// CreateIndex
	ExpireAfterSeconds      *int32
	Name                    *string
	Sparse                  *bool
	StorageEngine           interface{}
	Unique                  *bool
	Version                 *int32
	DefaultLanguage         *string
	LanguageOverride        *string
	TextVersion             *int32
	Weights                 interface{}
	SphereVersion           *int32
	Bits                    *int32
	BucketSize              *int32
	PartialFilterExpression interface{}
	WildcardProjection      interface{}
	Hidden                  *bool
}

func (q *QueryOption) ToFindOne() *options.FindOneOptionsBuilder {
	opts := options.FindOne()
	if q.AllowPartialResults != nil {
		opts = opts.SetAllowPartialResults(*q.AllowPartialResults)
	}
	if q.Collation != nil {
		opts = opts.SetCollation(q.Collation)
	}
	if q.Comment != nil {
		opts = opts.SetComment(q.Comment)
	}
	if q.Hint != nil {
		opts = opts.SetHint(q.Hint)
	}
	if q.Max != nil {
		opts = opts.SetMax(q.Max)
	}
	if q.Min != nil {
		opts = opts.SetMin(q.Min)
	}
	if q.OplogReplay != nil {
		opts = opts.SetOplogReplay(*q.OplogReplay)
	}
	if q.Projection != nil {
		opts = opts.SetProjection(q.Projection)
	}
	if q.ReturnKey != nil {
		opts = opts.SetReturnKey(*q.ReturnKey)
	}
	if q.ShowRecordID != nil {
		opts = opts.SetShowRecordID(*q.ShowRecordID)
	}
	if q.Skip != nil {
		opts = opts.SetSkip(*q.Skip)
	}
	if q.Sort != nil {
		opts = opts.SetSort(q.Sort)
	}
	return opts
}

func (q *QueryOption) ToFind() *options.FindOptionsBuilder {
	opts := options.Find()

	if q.AllowPartialResults != nil {
		opts = opts.SetAllowPartialResults(*q.AllowPartialResults)
	}
	if q.Collation != nil {
		opts = opts.SetCollation(q.Collation)
	}
	if q.Comment != nil {
		opts = opts.SetComment(q.Comment)
	}
	if q.Hint != nil {
		opts = opts.SetHint(q.Hint)
	}
	if q.Max != nil {
		opts = opts.SetMax(q.Max)
	}
	if q.MaxAwaitTime != nil {
		opts = opts.SetMaxAwaitTime(*q.MaxAwaitTime)
	}
	if q.Min != nil {
		opts = opts.SetMin(q.Min)
	}
	if q.OplogReplay != nil {
		opts = opts.SetOplogReplay(*q.OplogReplay)
	}
	if q.Projection != nil {
		opts = opts.SetProjection(q.Projection)
	}
	if q.ReturnKey != nil {
		opts = opts.SetReturnKey(*q.ReturnKey)
	}
	if q.ShowRecordID != nil {
		opts = opts.SetShowRecordID(*q.ShowRecordID)
	}
	if q.Skip != nil {
		opts = opts.SetSkip(*q.Skip)
	}
	if q.Sort != nil {
		opts = opts.SetSort(q.Sort)
	}
	if q.AllowDiskUse != nil {
		opts = opts.SetAllowDiskUse(*q.AllowDiskUse)
	}
	if q.BatchSize != nil {
		opts = opts.SetBatchSize(*q.BatchSize)
	}
	if q.CursorType != nil {
		opts = opts.SetCursorType(*q.CursorType)
	}
	if q.Let != nil {
		opts = opts.SetLet(q.Let)
	}
	if q.Limit != nil {
		opts = opts.SetLimit(*q.Limit)
	}
	if q.NoCursorTimeout != nil {
		opts = opts.SetNoCursorTimeout(*q.NoCursorTimeout)
	}
	return opts
}

func (q *QueryOption) ToFindOneAndUpdate() *options.FindOneAndUpdateOptionsBuilder {
	opts := options.FindOneAndUpdate()

	if q.Collation != nil {
		opts = opts.SetCollation(q.Collation)
	}
	if q.Comment != nil {
		opts = opts.SetComment(q.Comment)
	}
	if q.Hint != nil {
		opts = opts.SetHint(q.Hint)
	}
	if q.Projection != nil {
		opts = opts.SetProjection(q.Projection)
	}
	if q.Sort != nil {
		opts = opts.SetSort(q.Sort)
	}
	if q.Let != nil {
		opts = opts.SetLet(q.Let)
	}
	if q.ArrayFilters != nil {
		opts = opts.SetArrayFilters(q.ArrayFilters)
	}
	if q.BypassDocumentValidation != nil {
		opts = opts.SetBypassDocumentValidation(*q.BypassDocumentValidation)
	}
	if q.Upsert != nil {
		opts = opts.SetUpsert(*q.Upsert)
	}
	if q.ReturnDocument != nil {
		opts = opts.SetReturnDocument(*q.ReturnDocument)
	}
	return opts
}

func (q *QueryOption) ToCount() *options.CountOptionsBuilder {
	opts := options.Count()
	if q.Collation != nil {
		opts = opts.SetCollation(q.Collation)
	}
	if q.Comment != nil {
		opts = opts.SetComment(q.Comment)
	}
	if q.Hint != nil {
		opts = opts.SetHint(q.Hint)
	}
	if q.Limit != nil {
		opts = opts.SetLimit(*q.Limit)
	}
	if q.Skip != nil {
		opts = opts.SetSkip(*q.Skip)
	}
	return opts
}

func (q *QueryOption) ToDeleteOne() *options.DeleteOneOptionsBuilder {
	opts := options.DeleteOne()
	if q.Collation != nil {
		opts = opts.SetCollation(q.Collation)
	}
	if q.Comment != nil {
		opts = opts.SetComment(q.Comment)
	}
	if q.Hint != nil {
		opts = opts.SetHint(q.Hint)
	}
	if q.Let != nil {
		opts = opts.SetLet(q.Let)
	}
	return opts
}

func (q *QueryOption) ToDeleteMany() *options.DeleteManyOptionsBuilder {
	opts := options.DeleteMany()
	if q.Collation != nil {
		opts = opts.SetCollation(q.Collation)
	}
	if q.Comment != nil {
		opts = opts.SetComment(q.Comment)
	}
	if q.Hint != nil {
		opts = opts.SetHint(q.Hint)
	}
	if q.Let != nil {
		opts = opts.SetLet(q.Let)
	}
	return opts
}

func (q *QueryOption) ToUpdateOne() *options.UpdateOneOptionsBuilder {
	opts := options.UpdateOne()
	if q.Collation != nil {
		opts = opts.SetCollation(q.Collation)
	}
	if q.Comment != nil {
		opts = opts.SetComment(q.Comment)
	}
	if q.Hint != nil {
		opts = opts.SetHint(q.Hint)
	}
	if q.Sort != nil {
		opts = opts.SetSort(q.Sort)
	}
	if q.Let != nil {
		opts = opts.SetLet(q.Let)
	}
	if q.ArrayFilters != nil {
		opts = opts.SetArrayFilters(q.ArrayFilters)
	}
	if q.BypassDocumentValidation != nil {
		opts = opts.SetBypassDocumentValidation(*q.BypassDocumentValidation)
	}
	if q.Upsert != nil {
		opts = opts.SetUpsert(*q.Upsert)
	}
	return opts
}

func (q *QueryOption) ToUpdateMany() *options.UpdateManyOptionsBuilder {
	opts := options.UpdateMany()
	if q.Collation != nil {
		opts = opts.SetCollation(q.Collation)
	}
	if q.Comment != nil {
		opts = opts.SetComment(q.Comment)
	}
	if q.Hint != nil {
		opts = opts.SetHint(q.Hint)
	}
	if q.Let != nil {
		opts = opts.SetLet(q.Let)
	}
	if q.ArrayFilters != nil {
		opts = opts.SetArrayFilters(q.ArrayFilters)
	}
	if q.BypassDocumentValidation != nil {
		opts = opts.SetBypassDocumentValidation(*q.BypassDocumentValidation)
	}
	if q.Upsert != nil {
		opts = opts.SetUpsert(*q.Upsert)
	}
	return opts
}

func (q *QueryOption) ToInsertOne() *options.InsertOneOptionsBuilder {
	opts := options.InsertOne()
	if q.BypassDocumentValidation != nil {
		opts = opts.SetBypassDocumentValidation(*q.BypassDocumentValidation)
	}
	if q.Comment != nil {
		opts = opts.SetComment(q.Comment)
	}
	return opts
}

func (q *QueryOption) ToInsertMany() *options.InsertManyOptionsBuilder {
	opts := options.InsertMany()
	if q.BypassDocumentValidation != nil {
		opts = opts.SetBypassDocumentValidation(*q.BypassDocumentValidation)
	}
	if q.Comment != nil {
		opts = opts.SetComment(q.Comment)
	}
	if q.Ordered != nil {
		opts = opts.SetOrdered(*q.Ordered)
	}
	return opts
}

func (q *QueryOption) ToAggregate() *options.AggregateOptionsBuilder {
	opts := options.Aggregate()
	if q.Collation != nil {
		opts = opts.SetCollation(q.Collation)
	}
	if q.Comment != nil {
		opts = opts.SetComment(q.Comment)
	}
	if q.Hint != nil {
		opts = opts.SetHint(q.Hint)
	}
	if q.MaxAwaitTime != nil {
		opts = opts.SetMaxAwaitTime(*q.MaxAwaitTime)
	}
	if q.AllowDiskUse != nil {
		opts = opts.SetAllowDiskUse(*q.AllowDiskUse)
	}
	if q.BatchSize != nil {
		opts = opts.SetBatchSize(*q.BatchSize)
	}
	if q.Let != nil {
		opts = opts.SetLet(q.Let)
	}
	if q.BypassDocumentValidation != nil {
		opts = opts.SetBypassDocumentValidation(*q.BypassDocumentValidation)
	}
	if q.Custom != nil {
		opts = opts.SetCustom(q.Custom)
	}
	return opts
}

func (q *QueryOption) ToCreateIndex() *options.IndexOptionsBuilder {
	opts := options.Index()
	if q.Collation != nil {
		opts = opts.SetCollation(q.Collation)
	}
	if q.Max != nil && q.Max.(float64) != 0 {
		opts = opts.SetMax(q.Max.(float64))
	}
	if q.Min != nil && q.Min.(float64) != 0 {
		opts = opts.SetMin(q.Min.(float64))
	}
	if q.ExpireAfterSeconds != nil {
		opts = opts.SetExpireAfterSeconds(*q.ExpireAfterSeconds)
	}
	if q.Name != nil {
		opts = opts.SetName(*q.Name)
	}
	if q.Sparse != nil {
		opts = opts.SetSparse(*q.Sparse)
	}
	if q.StorageEngine != nil {
		opts = opts.SetStorageEngine(q.StorageEngine)
	}
	if q.Unique != nil {
		opts = opts.SetUnique(*q.Unique)
	}
	if q.Version != nil {
		opts = opts.SetVersion(*q.Version)
	}
	if q.DefaultLanguage != nil {
		opts = opts.SetDefaultLanguage(*q.DefaultLanguage)
	}
	if q.LanguageOverride != nil {
		opts = opts.SetLanguageOverride(*q.LanguageOverride)
	}
	if q.TextVersion != nil {
		opts = opts.SetTextVersion(*q.TextVersion)
	}
	if q.Weights != nil {
		opts = opts.SetWeights(q.Weights)
	}
	if q.SphereVersion != nil {
		opts = opts.SetSphereVersion(*q.SphereVersion)
	}
	if q.Bits != nil {
		opts = opts.SetBits(*q.Bits)
	}
	if q.BucketSize != nil {
		opts = opts.SetBucketSize(*q.BucketSize)
	}
	if q.PartialFilterExpression != nil {
		opts = opts.SetPartialFilterExpression(q.PartialFilterExpression)
	}
	if q.WildcardProjection != nil {
		opts = opts.SetWildcardProjection(q.WildcardProjection)
	}
	if q.Hidden != nil {
		opts = opts.SetHidden(*q.Hidden)
	}

	return opts
}
