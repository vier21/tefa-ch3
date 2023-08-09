package usecase

import (
	"context"

	"github.com/vier21/tefa-ch3/internal/model"
	"github.com/vier21/tefa-ch3/internal/repository"
)

type UserInterface interface {
	RegisterUser(ctx context.Context, user model.User) (Result, error)
}

type Result struct {
	UserMongo model.User `json:"userMongo"`
	UserMysql model.User `json:"userMysql"`
}

type userUsecase struct {
	userMysqlRepository repository.MysqlRepositoryInterface
	userMongoRepository repository.MongodbRepositoryInterface
}

func NewUserUsecase(mysql repository.MysqlRepositoryInterface, mongodb repository.MongodbRepositoryInterface) *userUsecase {
	return &userUsecase{
		userMysqlRepository: mysql,
		userMongoRepository: mongodb,
	}
}

func (u *userUsecase) RegisterUser(ctx context.Context, user model.User) (Result, error) {
	insMysql, err := u.userMysqlRepository.InsertUser(ctx, user)
	if err != nil {
		return Result{}, err
	}

	insMongo, err := u.userMongoRepository.InsertUser(ctx, user)
	if err != nil {
		return Result{}, err
	}

	return Result{
		UserMysql: insMysql,
		UserMongo: insMongo,
	}, nil
}
