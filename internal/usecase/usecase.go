package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/vier21/tefa-ch3/internal/model"
	"github.com/vier21/tefa-ch3/internal/repository"
)

type UserInterface interface {
	RegisterUser(ctx context.Context, user model.User) (Result, error)
	GetUserByID(ctx context.Context, userID string) (model.User, error)
	RegisterAccount(ctx context.Context, account model.Account) (model.Account, error)
	GetUserByAccountID(ctx context.Context, accountID string) (model.User, error)
	GetUserDataMongo(ctx context.Context, id string) (model.User, error)
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

func (u *userUsecase) GetUserDataMongo(ctx context.Context, id string) (model.User, error) {
	if id == "" {
		return model.User{}, fmt.Errorf("error id not specified")
	}
	user, err := u.userMongoRepository.GetUser(ctx, id)

	if err != nil {
		log.Println("error retrieving user: %s", err.Error())
		return model.User{}, err
	}

	return user, nil
}

func (u *userUsecase) GetUserByID(ctx context.Context, userID string) (model.User, error) {
	user, err := u.userMysqlRepository.GetUserByID(ctx, userID)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userUsecase) RegisterAccount(ctx context.Context, account model.Account) (model.Account, error) {
	account, err := u.userMysqlRepository.InsertAccount(ctx, account)
	if err != nil {
		return model.Account{}, err
	}

	return account, nil
}

func (u *userUsecase) GetUserByAccountID(ctx context.Context, accountID string) (model.User, error) {
	user, err := u.userMysqlRepository.GetUserByAccountID(ctx, accountID)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
