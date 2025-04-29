package test

import (
	"time"

	"github.com/nghialthanh/morn-go"
	"github.com/nghialthanh/morn-go/option"
	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	UserCollection = "users"
)

type User struct {
	ID        *bson.ObjectID `bson:"_id,omitempty"`
	CreatedAt *time.Time     `bson:"created_at"`
	UpdatedAt *time.Time     `bson:"updated_at"`
	Username  string         `bson:"username"`
	UserID    int64          `bson:"user_id"`
	Password  string         `bson:"password"`
	Email     string         `bson:"email"`
	Point     int64          `bson:"point"`
}

func InitUserModel(ins *morn.Instance) *morn.Dao {
	dao := morn.NewDao(UserCollection, User{}, ins, nil)

	enumTrue := true
	dao.Clause().Option(option.QueryOption{
		Unique: &enumTrue,
	}).CreateIndex("user_id:1")

	//composite index
	dao.Clause().CreateIndex("username:1", "email:1")
	return dao
}
