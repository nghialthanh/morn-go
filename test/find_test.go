package test

import (
	"testing"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestFindOne(t *testing.T) {
	ins := setupTestDB(t)

	userDao := InitUserModel(ins)
	defer cleanupTestDB(t, userDao, ins)

	// Create test data
	userID, err := userDao.GenIDForDao()
	if err != nil {
		t.Errorf("Failed to generate user ID: %v", err)
	}

	user := &User{
		Username: "testuser",
		Email:    "test@example.com",
		UserID:   userID,
	}

	_, err = userDao.Clause().MCreateOne(user)
	if err != nil {
		t.Errorf("Failed to create test user: %v", err)
	}

	tests := []struct {
		name      string
		condition interface{}
		want      *User
		wantErr   bool
	}{
		{
			name:      "Find existing user by ID",
			condition: map[string]interface{}{"user_id": userID},
			want:      user,
			wantErr:   false,
		},
		{
			name:      "Find existing user by username",
			condition: map[string]interface{}{"username": user.Username},
			want:      user,
			wantErr:   false,
		},
		{
			name:      "Find non-existent user",
			condition: map[string]interface{}{"_id": 999999},
			want:      nil,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := &User{}
			err := userDao.Clause().Where(tt.condition).MFindOne(result)

			if (err != nil) != tt.wantErr {
				t.Errorf("MFindOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if result.Username != tt.want.Username || result.Email != tt.want.Email || result.UserID != tt.want.UserID {
					t.Errorf("MFindOne() = %+v, want %+v", result, tt.want)
				}
			}
		})
	}
}

func TestFindMany(t *testing.T) {
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

	users := &[]User{
		{Username: "user1", Email: "user1@example.com", UserID: userID1},
		{Username: "user2", Email: "user2@example.com", UserID: userID2},
	}

	_, err = userDao.Clause().MCreateMany(users)
	if err != nil {
		t.Errorf("Failed to create test users: %v", err)
	}

	tests := []struct {
		name      string
		condition interface{}
		limit     int
		offset    int
		sort      string
		want      []User
		wantErr   bool
	}{
		{
			name:      "Find all users",
			condition: bson.M{},
			limit:     0,
			offset:    0,
			sort:      "",
			want:      *users,
			wantErr:   false,
		},
		{
			name:      "Find users with limit",
			condition: bson.M{},
			limit:     1,
			offset:    0,
			sort:      "",
			want:      []User{(*users)[0]},
			wantErr:   false,
		},
		{
			name:      "Find users with offset",
			condition: bson.M{},
			limit:     0,
			offset:    1,
			sort:      "",
			want:      []User{(*users)[1]},
			wantErr:   false,
		},
		{
			name:      "Find users with sort",
			condition: bson.M{},
			limit:     0,
			offset:    0,
			sort:      "username:desc",
			want:      []User{(*users)[1], (*users)[0]},
			wantErr:   false,
		},
		{
			name:      "Find users with condition",
			condition: map[string]interface{}{"username": "user1"},
			limit:     0,
			offset:    0,
			sort:      "",
			want:      []User{(*users)[0]},
			wantErr:   false,
		},
		{
			name:      "Find with invalid condition",
			condition: "invalid",
			limit:     0,
			offset:    0,
			sort:      "",
			want:      nil,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := &[]User{}
			clause := userDao.Clause().Where(tt.condition)

			if tt.limit > 0 {
				clause = clause.Limit(tt.limit)
			}
			if tt.offset > 0 {
				clause = clause.Skip(tt.offset)
			}
			if tt.sort != "" {
				clause = clause.Sort(tt.sort)
			}

			err := clause.MFindMany(result)

			if (err != nil) != tt.wantErr {
				t.Errorf("MFindMany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if len(*result) != len(tt.want) {
					t.Errorf("MFindMany() got %d results, want %d", len(*result), len(tt.want))
					return
				}

				for i, user := range *result {
					if user.Username != tt.want[i].Username || user.Email != tt.want[i].Email || user.UserID != tt.want[i].UserID {
						t.Errorf("MFindMany() result[%d] = %+v, want %+v", i, user, tt.want[i])
					}
				}
			}
		})
	}
}
