package clause

import (
	"context"
	"reflect"
	"testing"

	"github.com/nghialthanh/morn-go/logger"
	"github.com/nghialthanh/morn-go/option"
	"go.mongodb.org/mongo-driver/v2/bson"
	mongoopts "go.mongodb.org/mongo-driver/v2/mongo/options"
)

func TestCountKeepsQueryOptionWhenSkipIsSet(t *testing.T) {
	comment := "trace"
	queryLimit := int64(9)
	c := NewClause(nil, logger.NewFmtLogger(), struct{}{}, option.MornOption{}, context.Background())
	c.Option(option.QueryOption{Comment: comment, Limit: &queryLimit}).Skip(2)

	got := countOptionsOf(c.countOptions())
	if got.Comment != "trace" {
		t.Fatalf("comment = %v, want trace", got.Comment)
	}
	if got.Skip == nil || *got.Skip != 2 {
		t.Fatalf("skip = %v, want 2", got.Skip)
	}
	if got.Limit == nil || *got.Limit != 9 {
		t.Fatalf("limit = %v, want 9", got.Limit)
	}
}

func TestClauseLimitSkipAndSortOverrideQueryOption(t *testing.T) {
	comment := "trace"
	queryLimit := int64(9)
	c := NewClause(nil, logger.NewFmtLogger(), struct{}{}, option.MornOption{}, context.Background())
	c.Option(option.QueryOption{
		Comment: comment,
		Limit:   &queryLimit,
		Sort:    bson.M{"age": 1},
	}).Limit(3).Skip(2).Sort("name:desc")

	got := findOptionsOf(c.findOptions())
	if commentValue(got.Comment) != "trace" {
		t.Fatalf("comment = %v, want trace", got.Comment)
	}
	if got.Limit == nil || *got.Limit != 3 {
		t.Fatalf("limit = %v, want 3", got.Limit)
	}
	if got.Skip == nil || *got.Skip != 2 {
		t.Fatalf("skip = %v, want 2", got.Skip)
	}
	if !reflect.DeepEqual(got.Sort, bson.M{"name": -1}) {
		t.Fatalf("sort = %#v, want name desc", got.Sort)
	}
}

func TestQueryOptionSortRemainsWhenClauseSortIsUnset(t *testing.T) {
	c := NewClause(nil, logger.NewFmtLogger(), struct{}{}, option.MornOption{}, context.Background())
	c.Option(option.QueryOption{Sort: bson.M{"age": 1}})

	got := findOptionsOf(c.findOptions())
	if !reflect.DeepEqual(got.Sort, bson.M{"age": 1}) {
		t.Fatalf("sort = %#v, want age ascending from the query option", got.Sort)
	}
}

func commentValue(v interface{}) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	ptr, ok := v.(*interface{})
	if !ok || ptr == nil {
		return ""
	}
	s, _ := (*ptr).(string)
	return s
}

func countOptionsOf(b *mongoopts.CountOptionsBuilder) mongoopts.CountOptions {
	var opts mongoopts.CountOptions
	for _, set := range b.List() {
		if err := set(&opts); err != nil {
			panic(err)
		}
	}
	return opts
}

func findOptionsOf(b *mongoopts.FindOptionsBuilder) mongoopts.FindOptions {
	var opts mongoopts.FindOptions
	for _, set := range b.List() {
		if err := set(&opts); err != nil {
			panic(err)
		}
	}
	return opts
}
