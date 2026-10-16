package clause_test

import (
	"context"
	"testing"

	"github.com/nghialthanh/morn-go/clause"
	"github.com/nghialthanh/morn-go/logger"
	"github.com/nghialthanh/morn-go/option"
)

func TestCreateIndexRejectsANonNumericDirection(t *testing.T) {
	c := clause.NewClause(nil, logger.NewFmtLogger(), struct{ Name string }{}, option.MornOption{}, context.Background())

	err := c.CreateIndex("user_id:abc")
	if err == nil {
		t.Fatal("index direction must be a number")
	}
}
