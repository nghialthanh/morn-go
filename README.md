# MORN-GO

One challenge I consistently face when starting a new project using the MongoDB and Golang tech stack is that the official MongoDB driver can be quite difficult to work with. Each project often has its own unique configuration, which can be confusing and inconsistent for developers.

This library was created to help address that issue. It's a lightweight and intuitive ORM for MongoDB and Golang, called MORM-Go (Mongo ORM for Go).

## Feature Overview

* Full-Featured ORM
* Find, FindOne, FindOneAndUpdate, DeleteOne, DeleteMany,...
* Easy to manage table indexes right in the source code
* Integrated ascending GenID method
* Support self-conversion from mongo object to a certain struct
* Session, Save Point, RollbackTo to Saved Point
* Context, Prepared Statement Mode, DryRun Mode
* Logger
* Every feature comes with tests
* Developer Friendly

## Getting Started

```sh
go get github.com/nghialthanh/morn-go
```
The latest version of MORM-Go supports the last four major Go releases and is fully compatible with MongoDB 8.0 as well as the official mongo-driver v2.


## Usage examples

```go
package main

import (
	"fmt"

	"github.com/nghialthanh/morn-go"
	"github.com/nghialthanh/morn-go/logger"
	"github.com/nghialthanh/morn-go/option"
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

func main() {
	fmt.Println("Hello, World!")

	logger := logger.NewFmtLogger()
	url := ""
	ins, err := morn.SetupMongoByURI(url, &option.MornOption{
		IsGenID:       true,
		DefaultNumber: 100000,
		Logger:        logger,
	})
	if err != nil {
		logger.Error(err.Error())
		return
	}

	defer func() {
		if err = ins.Disconnect(); err != nil {
			logger.Error("Failed to disconnect from MongoDB", err)
			panic(err)
		}
	}()

	//set database
	ins.SetDB("Cluster0")

	//create dao user
	userDao := InitUserModel(ins)

	//create user
	user := &User{
		Username: "test",
		Email:    "test@test.com",
	}

	user.UserID, err = userRepository.userDao.GenIDForDao()
	if err != nil {
		logger.Error(err.Error())
	}
    id, err := r.userDao.Clause().MCreateOne(user)
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("User ID: %v", id)

}

```
You can explore how to organize your code and use the provided functions by referring to the examples in the example folder or the test cases in the test folder.

## Contribute

**Use issues for everything**

- For a small change, just send a PR.
- For bigger changes open an issue for discussion before sending a PR.
- PR should have:
  - Test case
  - Documentation
  - Example (If it makes sense)
- You can also contribute by:
  - Reporting issues
  - Suggesting new features or enhancements
  - Improve/fix documentation

## Reference

MongoDB         [MongoDB 8.0](https://www.mongodb.com/docs/manual/release-notes/8.0)
Mongo-Driver    [Mongo-Driver 2.0](https://www.mongodb.com/docs/drivers/go/v2.0/)
Golang          [Golang](https://go.dev/)

## License

© Nghia, 2025

Released under the [MIT License](https://github.com/go-gorm/gorm/blob/master/LICENSE)