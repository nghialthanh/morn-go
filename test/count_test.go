package test

import (
	"testing"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestMCount(t *testing.T) {
	ins := setupTestDB(t)

	userDao := InitUserModel(ins)
	defer cleanupTestDB(t, userDao, ins)

	// Create test data
	userID1, err := userDao.GenIDForDao()
	if err != nil {
		t.Errorf("Failed to generate user ID: %v", err)
	}

	userID2, err := userDao.GenIDForDao()
	if err != nil {
		t.Errorf("Failed to generate user ID: %v", err)
	}

	users := []*User{
		{
			Username: "user1",
			Email:    "user1@example.com",
			UserID:   userID1,
			Point:    100,
		},
		{
			Username: "user2",
			Email:    "user2@example.com",
			UserID:   userID2,
			Point:    200,
		},
	}

	_, err = userDao.Clause().MCreateMany(users)
	if err != nil {
		t.Errorf("Failed to create test user: %v", err)
	}

	tests := []struct {
		name      string
		condition interface{}
		limit     int
		offset    int
		want      int64
		wantErr   bool
	}{
		{
			name:      "Count all documents",
			condition: nil,
			limit:     0,
			offset:    0,
			want:      2,
			wantErr:   false,
		},
		{
			name:      "Count with condition",
			condition: bson.M{"point": 100},
			limit:     0,
			offset:    0,
			want:      1,
			wantErr:   false,
		},
		{
			name:      "Count with limit",
			condition: nil,
			limit:     1,
			offset:    0,
			want:      1,
			wantErr:   false,
		},
		{
			name:      "Count with offset",
			condition: nil,
			limit:     0,
			offset:    1,
			want:      1,
			wantErr:   false,
		},
		{
			name:      "Count with limit and offset",
			condition: nil,
			limit:     1,
			offset:    1,
			want:      1,
			wantErr:   false,
		},
		{
			name:      "Count with non-existent condition",
			condition: bson.M{"point": 999},
			limit:     0,
			offset:    0,
			want:      0,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clause := userDao.Clause()
			if tt.condition != nil {
				clause = clause.Where(tt.condition)
			}
			if tt.limit > 0 {
				clause = clause.Limit(tt.limit)
			}
			if tt.offset > 0 {
				clause = clause.Skip(tt.offset)
			}

			got, err := clause.MCount()
			if (err != nil) != tt.wantErr {
				t.Errorf("MCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
