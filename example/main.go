package main

import (
	"fmt"

	"github.com/nghialthanh/morn-go"
	"github.com/nghialthanh/morn-go/logger"
	"github.com/nghialthanh/morn-go/option"
)

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

	//create user repository
	userRepository := NewUserRepository(ins)

	//create user
	user := &User{
		Username: "test",
		Email:    "test@test.com",
	}

	user.UserID, err = userRepository.userDao.GenIDForDao()
	if err != nil {
		logger.Error(err.Error())
	}

	//get user by id
	users, err := userRepository.GetUser()
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("User: %v", users)

}
