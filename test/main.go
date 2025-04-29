package test

import (
	"context"
	"testing"

	"github.com/nghialthanh/morn-go"
	"github.com/nghialthanh/morn-go/logger"
	"github.com/nghialthanh/morn-go/option"
)

// setupTestDB creates a new test database and returns the instance.
// WARNING:
// - The current test functions do not fully cover all the features of the library.
func setupTestDB(t *testing.T) *morn.Instance {
	t.Helper()

	logger := logger.NewFmtLogger()
	url := ""
	ins, err := morn.SetupMongoByURI(url, &option.MornOption{
		IsGenID:       true,
		DefaultNumber: 100000,
		Logger:        logger,

		CreateAtField: "created_at",
		UpdateAtField: "updated_at",
	})
	if err != nil {
		logger.Error(err.Error())
		return nil
	}

	//set database
	ins.SetDB("Cluster0")

	return ins
}

// cleanupTestDB drops the test database and disconnects the client.
func cleanupTestDB(t *testing.T, dao *morn.Dao, ins *morn.Instance) {
	t.Helper()

	if _, err := dao.Clause().MDeleteMany(); err != nil {
		t.Errorf("failed to delete test data: %v", err)
	}

	if err := ins.GetClient().Disconnect(context.TODO()); err != nil {
		t.Errorf("failed to disconnect MongoDB client: %v", err)
	}
}
