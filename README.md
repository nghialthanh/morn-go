# morn-go

morn-go is a small MongoDB library for Go. It sits on the official driver and gives you a collection handle, struct mapping, indexes, a numeric id counter, and transactions without setting those up from scratch in every project.

## Why it exists

The official driver can do almost anything, so each project ends up with its own way of connecting, naming options, and decoding documents. morn-go keeps that setup in one place: one Dao per collection, one Clause per query.

## What you can do

- Find, create, update, delete, count, and aggregate
- Decode a result into your struct, or take the driver value and decode it yourself
- Declare indexes beside the collection
- Hand out an ascending id with `GenIDForDao`
- Run a callback inside a transaction
- Pass a context, and plug in a logger

An index string looks like `"user_id:1"`. A sort string looks like `"age:asc"` or `"age:desc"`.

## Install

```sh
go get github.com/nghialthanh/morn-go
```

The module requires Go 1.23 and uses [mongo-driver v2](https://www.mongodb.com/docs/drivers/go/v2.0/).

## Two kinds of methods

`FindOne`, `FindMany`, `UpdateMany`, and `Aggregate` return the driver's own values, such as `*mongo.Cursor` or `*mongo.SingleResult`.

Methods with an `M` prefix write the result into the value you pass. `MFindOne` fills one struct. `MFindMany` fills a slice. `MCreateOne` inserts one document and returns the inserted id.

The `M` methods use reflection. On a hot path, call the method that returns the driver value and decode it yourself.

## Usage

`NewDao` binds a collection to the struct its documents use. `Clause` starts a query on `context.TODO()`. `Ctx` starts a query on the context you pass.

Pass a pointer to that struct into `MFindOne` and `MCreateOne`. Pass a pointer to a slice into `MFindMany`.

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/nghialthanh/morn-go"
	"github.com/nghialthanh/morn-go/logger"
	"github.com/nghialthanh/morn-go/option"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID       *bson.ObjectID `bson:"_id,omitempty"`
	Created  *time.Time     `bson:"created_at"`
	Username string         `bson:"username"`
	UserID   int64          `bson:"user_id"`
	Email    string         `bson:"email"`
}

func main() {
	log := logger.NewFmtLogger()
	ins, err := morn.SetupMongoByURI("mongodb://localhost:27017", &option.MornOption{
		IsGenID:       true,
		DefaultNumber: 100000,
		Logger:        log,
	})
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer ins.Disconnect()

	ins.SetDB("app")
	users := morn.NewDao("users", User{}, ins, nil)

	unique := true
	err = users.Clause().Option(option.QueryOption{Unique: &unique}).CreateIndex("user_id:1")
	if err != nil {
		log.Error(err.Error())
		return
	}

	ctx := context.Background()
	user := &User{Username: "ada", Email: "ada@example.com"}
	user.UserID, err = users.GenIDForDao()
	if err != nil {
		log.Error(err.Error())
		return
	}

	id, err := users.Ctx(ctx).MCreateOne(user)
	if err != nil {
		log.Error(err.Error())
		return
	}
	fmt.Println(id)

	var found User
	err = users.Ctx(ctx).Where(bson.M{"user_id": user.UserID}).MFindOne(&found)
	if err != nil {
		log.Error(err.Error())
		return
	}

	var page []User
	err = users.Ctx(ctx).
		Where(bson.M{"user_id": bson.M{"$gt": 0}}).
		Sort("user_id:asc").
		Limit(20).
		Skip(0).
		MFindMany(&page)
	if err != nil {
		log.Error(err.Error())
		return
	}
}
```

`GenIDForDao` returns the counter's current value, then stores the next one. Set that value on the document yourself before you insert. With `IsGenID` set, `NewDao` creates the counter document for the collection on first use.

Inside `Session`, build each query with `Ctx` and the context the callback receives. A query started with `Clause` stays on `context.TODO()`.

```go
err = users.Session(ctx, func(tx context.Context) error {
	_, err := users.Ctx(tx).MCreateOne(user)
	return err
}, nil)
```

`Limit`, `Skip`, and `Sort` on the clause replace those same fields when you also set them on `option.QueryOption`. A hint, a collation, and the other query options still apply.

The [example](example) folder shows the same calls inside a repository. The [test](test) folder runs each operation against a live database.

## Contributing

Open an issue before a large change. A small fix can go straight to a pull request. Include a test, and update this file or the example when callers need to see the change.

## Reference

- [MongoDB Go driver v2](https://www.mongodb.com/docs/drivers/go/v2.0/)
- [Go](https://go.dev/)

## License

© Nghia, 2025

MIT. See [LICENSE](LICENSE).
