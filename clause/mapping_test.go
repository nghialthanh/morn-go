package clause_test

import (
	"context"
	"testing"

	"github.com/nghialthanh/morn-go/clause"
	"github.com/nghialthanh/morn-go/logger"
	"github.com/nghialthanh/morn-go/option"
)

func TestMCreateOneRejectsADocumentThatIsNotTheTemplate(t *testing.T) {
	type User struct{ Name string }
	type Order struct{ Name string }

	c := clause.NewClause(nil, logger.NewFmtLogger(), User{}, option.MornOption{}, context.Background())
	_, err := c.MCreateOne(&Order{Name: "a"})
	if err == nil || err.Error() != "struct of entity is not match with template" {
		t.Fatalf("err = %v, want template mismatch", err)
	}
}
