package test

import (
	"testing"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// ... existing code ...

func TestUpdateOne(t *testing.T) {
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
		updater   map[string]interface{}
		wantErr   bool
	}{
		{
			name:      "Update existing user",
			condition: map[string]interface{}{"user_id": userID},
			updater:   map[string]interface{}{"email": "updated@example.com"},
			wantErr:   false,
		},
		{
			name:      "Update non-existent user",
			condition: map[string]interface{}{"user_id": 999999},
			updater:   map[string]interface{}{"email": "updated@example.com"},
			wantErr:   false,
		},
		{
			name:      "Update with invalid condition",
			condition: "invalid",
			updater:   map[string]interface{}{"email": "updated@example.com"},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := userDao.Clause().Where(tt.condition).MUpdateOne(tt.updater)
			if (err != nil) != tt.wantErr {
				t.Errorf("MUpdateOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Verify the update
				result := &User{}
				err := userDao.Clause().Where(tt.condition).MFindOne(result)
				if err != nil {
					t.Errorf("Failed to verify update: %v", err)
					return
				}

				if result.Email != tt.updater["email"] {
					t.Errorf("MUpdateOne() field email = %v, want %v", result.Email, tt.updater["email"])
				}
			}
		})
	}
}

func TestUpdateMany(t *testing.T) {
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
		updater   map[string]interface{}
		wantErr   bool
	}{
		{
			name:      "Update all users",
			condition: bson.M{},
			updater:   map[string]interface{}{"email": "updated@example.com"},
			wantErr:   false,
		},
		{
			name:      "Update users with condition",
			condition: map[string]interface{}{"username": "user1"},
			updater:   map[string]interface{}{"email": "updated1@example.com"},
			wantErr:   false,
		},
		{
			name:      "Update with invalid condition",
			condition: "invalid",
			updater:   map[string]interface{}{"email": "updated@example.com"},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := userDao.Clause().Where(tt.condition).MUpdateMany(tt.updater)

			if (err != nil) != tt.wantErr {
				t.Errorf("MUpdateMany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Verify the updates
				result := &[]User{}
				err := userDao.Clause().Where(tt.condition).MFindMany(result)
				if err != nil {
					t.Errorf("Failed to verify updates: %v", err)
					return
				}

				for _, user := range *result {

					if user.Email != tt.updater["email"] {
						t.Errorf("MUpdateMany() field email = %v, want %v", user.Email, tt.updater["email"])
					}
				}
			}
		})
	}
}

func TestIncreaseValue(t *testing.T) {
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
		Point:    100,
	}

	_, err = userDao.Clause().MCreateOne(user)
	if err != nil {
		t.Errorf("Failed to create test user: %v", err)
	}

	tests := []struct {
		name        string
		condition   interface{}
		field       string
		upsert      bool
		wantErr     bool
		expectPoint int64
	}{
		{
			name:        "Increase user_id",
			condition:   map[string]interface{}{"user_id": userID},
			field:       "point:5",
			upsert:      false,
			wantErr:     false,
			expectPoint: 105,
		},
		{
			name:        "Decrease user_id",
			condition:   map[string]interface{}{"user_id": userID},
			field:       "point:-2",
			upsert:      false,
			wantErr:     false,
			expectPoint: 103,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := userDao.Clause().Where(tt.condition).MIncreaseValue(nil, tt.field, tt.upsert)
			if (err != nil) != tt.wantErr {
				t.Errorf("MIncreaseValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				result := &User{}
				err := userDao.Clause().Where(tt.condition).MFindOne(result)
				if err != nil {
					t.Errorf("Failed to verify update: %v", err)
					return
				}
				if result.Point != tt.expectPoint {
					t.Errorf("MIncreaseValue() field point = %v, want %v", result.Point, tt.expectPoint)
				}
			}
		})
	}
}

// func TestFindOneAndUpdate(t *testing.T) {
// 	ins := setupTestDB(t)

// 	userDao := InitUserModel(ins)
// 	defer cleanupTestDB(t, userDao, ins)

// 	// Create test data
// 	userID, err := userDao.GenIDForDao()
// 	if err != nil {
// 		t.Errorf("Failed to generate user ID: %v", err)
// 	}

// 	user := &User{
// 		Username: "testuser",
// 		Email:    "test@example.com",
// 		UserID:   userID,
// 	}

// 	_, err = userDao.Clause().MCreateOne(user)
// 	if err != nil {
// 		t.Errorf("Failed to create test user: %v", err)
// 	}

// 	tests := []struct {
// 		name      string
// 		condition interface{}
// 		updater   interface{}
// 		want      *User
// 		wantErr   bool
// 	}{
// 		{
// 			name:      "Update and return existing user",
// 			condition: map[string]interface{}{"user_id": userID},
// 			updater:   map[string]interface{}{"email": "updated@example.com"},
// 			want:      &User{Username: "testuser", Email: "updated@example.com", UserID: userID},
// 			wantErr:   false,
// 		},
// 		{
// 			name:      "Update non-existent user",
// 			condition: map[string]interface{}{"user_id": 999999},
// 			updater:   map[string]interface{}{"email": "updated@example.com"},
// 			want:      nil,
// 			wantErr:   true,
// 		},
// 		{
// 			name:      "Update with invalid condition",
// 			condition: "invalid",
// 			updater:   map[string]interface{}{"email": "updated@example.com"},
// 			want:      nil,
// 			wantErr:   true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			result := &User{}
// 			err := userDao.Clause().Where(tt.condition).MFindOneAndUpdate(tt.updater, result)

// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("MFindOneAndUpdate() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}

// 			if !tt.wantErr {
// 				if result.Email != tt.want.Email {
// 					t.Errorf("MFindOneAndUpdate() = %+v, want %+v", result, tt.want)
// 				}
// 				resultAfter := &User{}
// 				err := userDao.Clause().Where(tt.condition).MFindOne(resultAfter)
// 				if err != nil {
// 					t.Errorf("Failed to verify update: %v", err)
// 					return
// 				}
// 				if resultAfter.Username != tt.want.Username || resultAfter.Email != tt.want.Email || resultAfter.UserID != tt.want.UserID {
// 					t.Errorf("MFindOneAndUpdate() = %+v, want %+v", result, tt.want)
// 				}
// 			}
// 		})
// 	}
// }
