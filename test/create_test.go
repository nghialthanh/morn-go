package test

import (
	"testing"
)

func TestCreate(t *testing.T) {
	ins := setupTestDB(t)

	userDao := InitUserModel(ins)

	defer cleanupTestDB(t, userDao, ins)
	userID, err := userDao.GenIDForDao()
	if err != nil {
		t.Errorf("Failed to generate user ID: %v", err)
	}

	tests := []struct {
		name    string
		user    *User
		wantErr bool
	}{
		{
			name:    "Valid user",
			user:    &User{Username: "John Doe", Email: "john@example.com", UserID: userID},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := userDao.Clause().MCreateOne(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// Verify the user was inserted
				result := &User{}
				err = userDao.Clause().Where(map[string]interface{}{"_id": id}).MFindOne(result)
				if err != nil {
					t.Errorf("Failed to find inserted user: %v", err)
				}
				if result.Username != tt.user.Username || result.Email != tt.user.Email || result.UserID != tt.user.UserID {
					t.Errorf("Inserted user does not match: got %+v, want %+v", result, tt.user)
				}
			}
		})
	}
}

func TestCreateMany(t *testing.T) {
	ins := setupTestDB(t)

	userDao := InitUserModel(ins)
	defer cleanupTestDB(t, userDao, ins)
	userID1, err := userDao.GenIDForDao()
	if err != nil {
		t.Errorf("Failed to generate user ID: %v", err)
	}

	userID2, err := userDao.GenIDForDao()
	if err != nil {
		t.Errorf("Failed to generate user ID: %v", err)
	}

	tests := []struct {
		name    string
		user    *[]User
		wantErr bool
	}{
		{
			name: "Valid user",
			user: &[]User{
				User{Username: "John Doe", Email: "john@example.com", UserID: userID1},
				User{Username: "Jane Doe", Email: "jane@example.com", UserID: userID2},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := userDao.Clause().MCreateMany(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {

				// filter := bson.M{"_id": bson.M{"$in": id}} use bson
				filter := map[string]interface{}{
					"_id": map[string]interface{}{"$in": id},
				}
				result := &[]User{}
				err = userDao.Clause().Where(filter).MFindMany(result)
				if err != nil {
					t.Errorf("Failed to find inserted user: %v", err)
				}
				if len(*result) != len(*tt.user) {
					t.Errorf("Inserted user does not match: got %+v, want %+v", result, tt.user)
				}

				mapResult := make(map[string]User)
				for _, user := range *result {
					mapResult[user.Username] = user
				}
				for _, user := range *tt.user {
					if mapResult[user.Username].Username != user.Username || mapResult[user.Username].Email != user.Email || mapResult[user.Username].UserID != user.UserID {
						t.Errorf("Inserted user does not match: got %+v, want %+v", mapResult[user.Username], user)
					}
				}
			}
		})
	}
}
