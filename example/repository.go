package main

import (
	"github.com/nghialthanh/morn-go"
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
		Sort("").
		MFindMany(result)
	return *result, err
}
