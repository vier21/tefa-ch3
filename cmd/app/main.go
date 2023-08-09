package main

import (
	"github.com/vier21/tefa-ch3/db"
	"github.com/vier21/tefa-ch3/internal/repository"
	"github.com/vier21/tefa-ch3/internal/server"
	"github.com/vier21/tefa-ch3/internal/usecase"
)

func main() {
	db.InitMongoDB()
	db.InitMysqlDB()

	mysqlRepo := repository.NewMysqlRepository()
	mongoRepo := repository.NewMongoRepository()
	usecase := usecase.NewUserUsecase(mysqlRepo, mongoRepo)

	server := server.NewServer(usecase)
	server.Run()
}
