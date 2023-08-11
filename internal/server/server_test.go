package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vier21/tefa-ch3/db"
	"github.com/vier21/tefa-ch3/internal/model"
	"github.com/vier21/tefa-ch3/internal/repository"
	"github.com/vier21/tefa-ch3/internal/usecase"
)

func TestGetUserMongoHandler(t *testing.T) {

	db.InitMongoDB()
	db.InitMysqlDB()

	mysqlRepo := repository.NewMysqlRepository()
	mongoRepo := repository.NewMongoRepository()
	usecase := usecase.NewUserUsecase(mysqlRepo, mongoRepo)

	server := NewServer(usecase)

	req, err := http.NewRequest("GET", "http://localhost:3001/42307263-2b56-45d5-9f86-db0fcb93958b/user/mongo", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r := server.NewRouter()
	r.HandleFunc("/{id}/user/mongo", server.GetUserMongoHandler)

	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "status code should be %d but got %d", http.StatusOK, rr.Code)
}

func TestGetUserMysqlHandler(t *testing.T) {
	db.InitMongoDB()
	db.InitMysqlDB()

	mysqlRepo := repository.NewMysqlRepository()
	mongoRepo := repository.NewMongoRepository()
	usecase := usecase.NewUserUsecase(mysqlRepo, mongoRepo)

	server := NewServer(usecase)

	req, err := http.NewRequest("GET", "http://localhost:3001/14ba3470-e99c-4ae5-aeba-0d875258dbf5/user/mysql", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r := server.NewRouter()
	r.HandleFunc("/{id}/user/mysql", server.GetUserMysqlHandler)

	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "status code should be %d but got %d", http.StatusOK, rr.Code)
}

func TestGetUserByAccountIDHandler(t *testing.T) {
	db.InitMongoDB()
	db.InitMysqlDB()

	mysqlRepo := repository.NewMysqlRepository()
	mongoRepo := repository.NewMongoRepository()
	usecase := usecase.NewUserUsecase(mysqlRepo, mongoRepo)

	server := NewServer(usecase)

	req, err := http.NewRequest("GET", "http://localhost:3001/90cb8ccb-8c52-42fe-9e6a-b03d9933b6e5/account", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r := server.NewRouter()
	r.HandleFunc("/{accountID}/account", server.GetUserByAccountIDHandler)

	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "status code should be %d but got %d", http.StatusOK, rr.Code)
}

func TestRegisterUserHandler(t *testing.T) {

	db.InitMongoDB()
	db.InitMysqlDB()

	mysqlRepo := repository.NewMysqlRepository()
	mongoRepo := repository.NewMongoRepository()
	usecase := usecase.NewUserUsecase(mysqlRepo, mongoRepo)

	server := NewServer(usecase)
	usr := model.User{
		Name:    "sdasd",
		Address: "sdasd",
		Email:   "sdasd",
	}

	jsonBytes, err := json.Marshal(usr)
	if err != nil {
		panic(err)
	}

	// Create a bytes.Buffer from the JSON bytes
	jsonBuffer := bytes.NewBuffer(jsonBytes)

	req, err := http.NewRequest("POST", "http://localhost:3001/user", jsonBuffer)
	if err != nil {
		panic(err)
	}


	rr := httptest.NewRecorder()
	r := server.NewRouter()
	r.Post("/user", server.RegisterUserHandler)

	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "status code should be %d but got %d", http.StatusOK, rr.Code)
}

func TestRegisterAccountHandler(t *testing.T) {

	db.InitMongoDB()
	db.InitMysqlDB()

	mysqlRepo := repository.NewMysqlRepository()
	mongoRepo := repository.NewMongoRepository()
	usecase := usecase.NewUserUsecase(mysqlRepo, mongoRepo)

	server := NewServer(usecase)
	usr := model.Account{
		MsisdnCustomer: "xxxxx",
		UserID: "bf53c3d8-aa04-4064-a1e1-f3c2c4072edc",
	}

	jsonBytes, err := json.Marshal(usr)
	if err != nil {
		panic(err)
	}

	// Create a bytes.Buffer from the JSON bytes
	jsonBuffer := bytes.NewBuffer(jsonBytes)

	req, err := http.NewRequest("POST", "http://localhost:3001/user", jsonBuffer)
	if err != nil {
		panic(err)
	}


	rr := httptest.NewRecorder()
	r := server.NewRouter()
	r.Post("/user", server.RegisterUserHandler)

	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "status code should be %d but got %d", http.StatusOK, rr.Code)
}
