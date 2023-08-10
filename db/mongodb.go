package db

import (
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vier21/tefa-ch3/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoCLI *mongo.Client

func InitMongoDB() (err error) {

	MongoCLI, err = mongo.Connect(context.Background(), options.Client().ApplyURI(config.GetConfig().MongoDBURL).SetMaxPoolSize(50))
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}

	if err := MongoCLI.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	return
}

func Disconnect() {
	if err := MongoCLI.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}
