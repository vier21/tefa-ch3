package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/vier21/tefa-ch3/config"
	"github.com/vier21/tefa-ch3/db"
	"github.com/vier21/tefa-ch3/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongodbRepositoryInterface interface {
	InsertUser(ctx context.Context, user model.User) (model.User, error)
	GetUser(ctx context.Context, userid string) (model.User, error)
}

type MongoRepository struct {
	db         *mongo.Client
	collection string
}

func NewMongoRepository() *MongoRepository {
	return &MongoRepository{
		db:         db.MongoCLI,
		collection: "user",
	}
}

func (m *MongoRepository) InsertUser(ctx context.Context, user model.User) (model.User, error) {
	coll := m.db.Database(config.GetConfig().UserDBName).Collection(m.collection)
	id := uuid.NewString()
	user.UserID = id

	doc, err := coll.InsertOne(ctx, user)

	if err != nil {
		return model.User{}, err
	}

	if doc.InsertedID.(string) != id {
		return model.User{}, fmt.Errorf("document not contain expected id %s", id)
	}

	return user, nil

}

func (m *MongoRepository) GetUser(ctx context.Context, userid string) (model.User, error) {
	coll := m.db.Database(config.GetConfig().UserDBName).Collection(m.collection)
	filter := bson.M{
		"_id": userid,
	}
	doc := coll.FindOne(ctx, filter)

	if doc.Err() != nil {
		log.Println(doc.Err().Error())

	}

	var user model.User

	if err := doc.Decode(&user); err != nil {
		log.Println(err.Error())
		return model.User{}, err
	}

	return user, nil

}
