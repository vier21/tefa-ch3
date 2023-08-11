package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vier21/tefa-ch3/internal/model"
)

type MockMysqlRepository struct{}
type MockMongoRepository struct{}

func (m *MockMongoRepository) InsertUser(ctx context.Context, user model.User) (model.User, error) {
	return user, nil
}

func (m *MockMongoRepository) GetUserByAccountID(ctx context.Context, accountID string) (model.User, error) {
	return model.User{}, nil
}

func (m *MockMongoRepository) GetUser(ctx context.Context, userID string) (model.User, error) {
	mockUser := model.User{
		UserID:  userID,
		Name:    "Mock User",
		Address: "Mock Address",
		Email:   "mock@example.com",
	}

	return mockUser, nil
}

func (m *MockMysqlRepository) InsertUser(ctx context.Context, user model.User) (model.User, error) {
	return user, nil
}

func (m *MockMysqlRepository) GetUserByID(ctx context.Context, userID string) (model.User, error) {
	return model.User{}, nil
}

func (m *MockMysqlRepository) InsertAccount(ctx context.Context, account model.Account) (model.Account, error) {
	return account, nil
}

func (m *MockMysqlRepository) GetUserByAccountID(ctx context.Context, accountID string) (model.User, error) {
	return model.User{}, nil
}

func TestUserUsecase_RegisterUser(t *testing.T) {
	mysqlRepo := &MockMysqlRepository{}
	mongoRepo := &MockMongoRepository{}
	usecase := NewUserUsecase(mysqlRepo, mongoRepo)

	user := model.User{
		Name:    "John Doe",
		Address: "123 Main St",
		Email:   "john@example.com",
	}

	result, err := usecase.RegisterUser(context.Background(), user)
	assert.NoError(t, err)
	assert.Equal(t, user.Name, result.UserMysql.Name)
	assert.Equal(t, user.Name, result.UserMongo.Name)
}

func TestUserUsecase_GetUserByID(t *testing.T) {
	mysqlRepo := &MockMysqlRepository{}
	mongoRepo := &MockMongoRepository{}
	usecase := NewUserUsecase(mysqlRepo, mongoRepo)

	userID := "someUserID"
	user, err := usecase.GetUserByID(context.Background(), userID)
	assert.NoError(t, err)
	assert.Equal(t, model.User{}, user)
}

func TestUserUsecase_GetUserByIDMongo(t *testing.T) {
	mysqlRepo := &MockMysqlRepository{}
	mongoRepo := &MockMongoRepository{}
	usecase := NewUserUsecase(mysqlRepo, mongoRepo)

	userID := "someUserID"
	user, err := usecase.GetUserByID(context.Background(), userID)
	assert.NoError(t, err)
	assert.Equal(t, model.User{}, user)
}

func TestUserUsecase_RegisterAccount(t *testing.T) {
	mysqlRepo := &MockMysqlRepository{}
	mongoRepo := &MockMongoRepository{}
	usecase := NewUserUsecase(mysqlRepo, mongoRepo)

	account := model.Account{
		MsisdnCustomer: "1234567890",
		UserID:         "someUserID",
	}

	result, err := usecase.RegisterAccount(context.Background(), account)
	assert.NoError(t, err)
	assert.Equal(t, account.MsisdnCustomer, result.MsisdnCustomer)
}

func TestUserUsecase_GetUserByAccountID(t *testing.T) {
	mysqlRepo := &MockMysqlRepository{}
	mongoRepo := &MockMongoRepository{}
	usecase := NewUserUsecase(mysqlRepo, mongoRepo)

	accountID := "someAccountID"
	user, err := usecase.GetUserByAccountID(context.Background(), accountID)
	assert.NoError(t, err)
	assert.Equal(t, model.User{}, user)
}
