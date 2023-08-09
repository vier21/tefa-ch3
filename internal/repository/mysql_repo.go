package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/vier21/tefa-ch3/db"
	"github.com/vier21/tefa-ch3/internal/model"
)

type MysqlRepositoryInterface interface {
	InsertUser(ctx context.Context, user model.User) (model.User, error)
	GetUserByID(ctx context.Context, userID string) (model.User, error)
}

type mySqlRepository struct {
	db *sqlx.DB
}

func NewMysqlRepository() *mySqlRepository {
	return &mySqlRepository{
		db: db.DB,
	}
}

func (m *mySqlRepository) InsertUser(ctx context.Context, user model.User) (model.User, error) {

	sqlstr := "INSERT INTO user (id, name, address, email) values (?, ?, ?, ?)"
	user.UserID = uuid.NewString()

	_, err := m.db.ExecContext(ctx, sqlstr, user.UserID, user.Name, user.Address, user.Email)

	if err != nil {
		return model.User{}, err
	}

	result := model.User{
		UserID:  user.UserID,
		Name:    user.Name,
		Address: user.Name,
		Email:   user.Email,
	}

	return result, nil
}

func (m *mySqlRepository) GetUserByID(ctx context.Context, userID string) (model.User, error) {
	var user model.User
	sqlstr := "SELECT id, name, address, email FROM user WHERE id = ?"

	err := m.db.GetContext(ctx, &user, sqlstr, userID)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
