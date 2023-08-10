package usecase

import (
	"context"

	"github.com/vier21/tefa-ch3/internal/model"
	"github.com/vier21/tefa-ch3/internal/repository"
)

type UserInterface interface {
	RegisterUser(ctx context.Context, user model.User) (Result, error)
	GetUserByAccountID(ctx context.Context, accountID string) (model.User, error)
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

func (u *userUsecase) GetUserByAccountID(ctx context.Context, accountID int) (Result, error) {
	userMysql, err := u.userMysqlRepository.GetUserByAccountID(ctx, accountID)
	if err != nil {
		// Jika tidak ditemukan di MySQL, coba mencari di MongoDB
		userMongo, err := u.userMongoRepository.GetUserByAccountID(ctx, accountID)
		if err != nil {
			return Result{}, err
		}
		return Result{
			UserMysql: model.User{},
			UserMongo: userMongo,
		}, nil
	}
	return Result{
		UserMysql: userMysql,
		UserMongo: model.User{},
	}, nil
}

