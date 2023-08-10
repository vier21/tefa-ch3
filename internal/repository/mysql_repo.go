package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/vier21/tefa-ch3/db"
	"github.com/vier21/tefa-ch3/internal/model"
)

type MysqlRepositoryInterface interface {
	InsertUser(ctx context.Context, user model.User) (model.User, error)
	GetUserByID(ctx context.Context, userID string) (model.User, error)
	InsertAccount(ctx context.Context, account model.Account) (model.Account, error)
	GetUserByAccountID(ctx context.Context, accountID string) (model.User, error)
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

func (m *mySqlRepository) InsertAccount(ctx context.Context, account model.Account) (model.Account, error) {
	var accounts []model.Account
	rows, err := m.db.QueryContext(ctx, "SELECT * FROM account WHERE user_id=?", account.UserID)
	if err != nil {
		return model.Account{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var account model.Account
		err := rows.Scan(&account.AccountID, &account.MsisdnCustomer, &account.UserID)
		if err != nil {
			return model.Account{}, err
		}
		accounts = append(accounts, account)
	}

	if err := rows.Err(); err != nil {
		return model.Account{}, err
	}

	if len(accounts) >= 3 {
		return model.Account{}, errors.New("MSISDN Limit Reached")
	}

	sqlstr := "INSERT INTO account (id, msisdn_customer, user_id) VALUES (?, ?, ?)"
	account.AccountID = uuid.NewString()

	_, err = m.db.ExecContext(ctx, sqlstr, account.AccountID, account.MsisdnCustomer, account.UserID)

	if err != nil {
		return model.Account{}, err
	}

	return account, nil

}

func (m *mySqlRepository) GetUserByAccountID(ctx context.Context, accountID string) (model.User, error) {
	sqlstr := "SELECT id, name, address, email FROM user WHERE account_id = ?"

	var user model.User
	err := m.db.GetContext(ctx, &user, sqlstr, accountID)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
