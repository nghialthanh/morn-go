package test

import (
	"context"
	"errors"
	"testing"
)

func TestSession(t *testing.T) {
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

	userID3, err := userDao.GenIDForDao()
	if err != nil {
		t.Errorf("Failed to generate user ID: %v", err)
	}

	tests := []struct {
		name    string
		fn      func(ctx context.Context) error
		wantErr bool
	}{
		{
			name: "Successful transaction",
			fn: func(ctx context.Context) error {
				// Create two users in a transaction
				user1 := &User{Username: "user1", Email: "user1@example.com", UserID: userID1}
				user2 := &User{Username: "user2", Email: "user2@example.com", UserID: userID2}

				_, err := userDao.Ctx(ctx).MCreateOne(user1)
				if err != nil {
					return err
				}

				_, err = userDao.Ctx(ctx).MCreateOne(user2)
				if err != nil {
					return err
				}

				return nil
			},
			wantErr: false,
		},
		{
			name: "Failed transaction with rollback",
			fn: func(ctx context.Context) error {
				// Create first user
				user1 := &User{Username: "user3", Email: "user3@example.com", UserID: userID3}
				_, err := userDao.Ctx(ctx).MCreateOne(user1)
				if err != nil {
					return err
				}

				// Try to create second user with same ID (should fail)
				user2 := &User{Username: "user4", Email: "user4@example.com", UserID: userID3}
				_, err = userDao.Ctx(ctx).MCreateOne(user2)
				if err != nil {
					return err
				}

				return nil
			},
			wantErr: true,
		},
		{
			name: "Transaction with error",
			fn: func(ctx context.Context) error {
				return errors.New("custom error")
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := userDao.Session(context.Background(), tt.fn, nil)

			if (err != nil) != tt.wantErr {
				t.Errorf("Session() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Verify the transaction was committed
				result := &[]User{}
				err := userDao.Clause().MFindMany(result)
				if err != nil {
					t.Errorf("Failed to verify transaction: %v", err)
					return
				}

				// Check if both users were created
				if len(*result) != 2 {
					t.Errorf("Expected 2 users, got %d", len(*result))
				}
			} else {
				// Verify the transaction was rolled back
				result := &[]User{}
				err := userDao.Clause().MFindMany(result)
				if err != nil {
					t.Errorf("Failed to verify rollback: %v", err)
					return
				}

				// Check if no users were created
				if len(*result) != 2 {
					t.Errorf("Expected 2 users after rollback, got %d", len(*result))
				}
			}
		})
	}
}
