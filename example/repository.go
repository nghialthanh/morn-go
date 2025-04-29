package main

import (
	"context"

	"github.com/nghialthanh/morn-go"
	"github.com/nghialthanh/morn-go/option"
)

type UserRepository struct {
	userDao *morn.Dao
}

func NewUserRepository(ins *morn.Instance) *UserRepository {
	return &UserRepository{userDao: InitUserModel(ins)}
}

func (r *UserRepository) CreateUser(user *User) (interface{}, error) {
	id, err := r.userDao.Clause().MCreateOne(user)
	return id, err
}

func (r *UserRepository) GetUser() ([]User, error) {

	err := r.userDao.Clause().MUpdateMany(map[string]interface{}{"email": "up@example.com"})
	if err != nil {
		return nil, err
	}
	result := &[]User{}
	err = r.userDao.Clause().
		Limit(1).
		Sort("email:asc").
		Where(map[string]interface{}{"email": "up@example.com"}).
		MFindMany(result)
	return *result, err
}

func (r *UserRepository) UpdateUser(id string) (*User, error) {
	user := &User{
		Username: "test",
		Email:    "test@test.com",
	}
	err := r.userDao.Clause().Option(option.QueryOption{
		Upsert: &[]bool{true}[0],
	}).Where(map[string]interface{}{"_id": id}).MUpdateOne(user)
	return user, err
}

func (r *UserRepository) TransactionUser() (interface{}, error) {

	err := r.userDao.Session(context.Background(), func(ctx context.Context) error {
		user := &User{
			Username: "test",
			Email:    "test@test.com",
		}
		_, err := r.userDao.Clause().MCreateOne(user)
		if err != nil {
			return err
		}
		user2 := &User{
			Username: "test2",
			Email:    "test2@test.com",
		}
		_, err = r.userDao.Clause().MCreateOne(user2)
		if err != nil {
			return err
		}
		return nil
	}, nil)

	return nil, err
}
