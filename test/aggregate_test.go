package test

import (
	"testing"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestMAggregate(t *testing.T) {
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
		t.Errorf("Failed to create test users: %v", err)
	}

	tests := []struct {
		name      string
		pipeline  []bson.M
		entity    interface{}
		wantErr   bool
		checkFunc func(t *testing.T, result interface{})
	}{
		{
			name: "Basic aggregation - count documents",
			pipeline: []bson.M{
				{
					"$count": "total",
				},
			},
			entity: &struct {
				Total int `bson:"total"`
			}{},
			wantErr: false,
			checkFunc: func(t *testing.T, result interface{}) {
				count := result.(*struct {
					Total int `bson:"total"`
				})
				if count.Total != 2 {
					t.Errorf("Expected 2 documents, got %d", count.Total)
				}
			},
		},
		{
			name: "Aggregation with match stage",
			pipeline: []bson.M{
				{
					"$match": bson.M{
						"point": 100,
					},
				},
			},
			entity:  &[]User{},
			wantErr: false,
			checkFunc: func(t *testing.T, result interface{}) {
				users := result.(*[]User)
				if len(*users) != 1 {
					t.Errorf("Expected 1 user, got %d", len(*users))
				}
				if (*users)[0].Point != 100 {
					t.Errorf("Expected point 100, got %d", (*users)[0].Point)
				}
			},
		},
		{
			name: "Complex aggregation with group and sort",
			pipeline: []bson.M{
				{
					"$group": bson.M{
						"_id": nil,
						"totalPoints": bson.M{
							"$sum": "$point",
						},
						"avgPoints": bson.M{
							"$avg": "$point",
						},
					},
				},
			},
			entity: &struct {
				TotalPoints int     `bson:"totalPoints"`
				AvgPoints   float64 `bson:"avgPoints"`
			}{},
			wantErr: false,
			checkFunc: func(t *testing.T, result interface{}) {
				stats := result.(*struct {
					TotalPoints int     `bson:"totalPoints"`
					AvgPoints   float64 `bson:"avgPoints"`
				})
				if stats.TotalPoints != 300 {
					t.Errorf("Expected total points 300, got %d", stats.TotalPoints)
				}
				if stats.AvgPoints != 150 {
					t.Errorf("Expected average points 150, got %f", stats.AvgPoints)
				}
			},
		},
		{
			name: "Invalid pipeline stage",
			pipeline: []bson.M{
				{
					"$invalidStage": bson.M{},
				},
			},
			entity:  &[]User{},
			wantErr: true,
			checkFunc: func(t *testing.T, result interface{}) {
				// No check needed for error case
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := userDao.Clause().MAggregate(tt.entity, tt.pipeline)
			if (err != nil) != tt.wantErr {
				t.Errorf("MAggregate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				tt.checkFunc(t, tt.entity)
			}
		})
	}
}
