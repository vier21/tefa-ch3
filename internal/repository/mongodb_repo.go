package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/vier21/tefa-ch3/config"
	"github.com/vier21/tefa-ch3/db"
	"github.com/vier21/tefa-ch3/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongodbRepositoryInterface interface {
	InsertUser(ctx context.Context, user model.User) (model.User, error)
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
