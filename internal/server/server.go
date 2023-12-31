package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/vier21/tefa-ch3/internal/model"
	"github.com/vier21/tefa-ch3/internal/usecase"
)

type ApiServer struct {
	Services usecase.UserInterface
	Router   *chi.Mux
	Server   *http.Server
}

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

var (
	ErrFetchResp       = "fail to fetch responses"
	ErrMethodNotAllow  = "method not allowed"
	ErrReqBodyNotValid = "request body not valid"
)

func NewServer(usersvc usecase.UserInterface) *ApiServer {
	mux := chi.NewRouter()

	return &ApiServer{
		Services: usersvc,
		Router:   mux,
		Server: &http.Server{
			Addr:         ":3001",
			Handler:      mux,
			IdleTimeout:  120 * time.Second,
			WriteTimeout: 1 * time.Second,
			ReadTimeout:  1 * time.Second,
		},
	}
}

func (a *ApiServer) NewRouter() *chi.Mux {
	return a.Router
}

func (a *ApiServer) Run() {
	r := a.NewRouter()

	r.Post("/user", a.RegisterUserHandler)
	r.Get("/{id}/user/mysql", a.GetUserMysqlHandler)
	r.Get("/{id}/user/mongo", a.GetUserMongoHandler)
	r.Post("/account", a.RegisterAccountHandler)
	r.Get("/{accountID}/account", a.GetUserByAccountIDHandler)

	go func() {
		log.Printf("Server start on localhost%s \n", ":3001")
		err := a.Server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server Error: %s \n", err)
		}
	}()

	a.GracefullShutdown()

}

func (a *ApiServer) GracefullShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := a.Server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Server stopped gracefully")
}

func (s *ApiServer) GetUserMongoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	user, err := s.Services.GetUserDataMongo(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	httpcode := strconv.Itoa(http.StatusOK)
	status := fmt.Sprintf("Success (%s)", httpcode)

	res := Response{
		Status: status,
		Data:   user,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (s *ApiServer) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var req model.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, ErrReqBodyNotValid, http.StatusBadRequest)
		return
	}

	reg, err := s.Services.RegisterUser(r.Context(), req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpcode := strconv.Itoa(http.StatusOK)
	status := fmt.Sprintf("Success (%s)", httpcode)

	res := Response{
		Status: status,
		Data:   reg,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "fail to fetch responses", http.StatusInternalServerError)
		return
	}

}

func (a *ApiServer) GetUserMysqlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	userID := chi.URLParam(r, "id") // Get userID from URL parameter

	user, err := a.Services.GetUserByID(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpcode := strconv.Itoa(http.StatusOK)
	status := fmt.Sprintf("Success (%s)", httpcode)

	res := Response{
		Status: status,
		Data:   user,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, ErrFetchResp, http.StatusInternalServerError)
		return
	}
}

func (a *ApiServer) RegisterAccountHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var req model.Account
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, ErrReqBodyNotValid, http.StatusBadRequest)
		return
	}

	account, err := a.Services.RegisterAccount(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpcode := strconv.Itoa(http.StatusOK)
	status := fmt.Sprintf("Success (%s)", httpcode)

	res := Response{
		Status: status,
		Data:   account,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "fail to fetch responses", http.StatusInternalServerError)
		return
	}

}

func (s *ApiServer) GetUserByAccountIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	accountIDStr := chi.URLParam(r, "accountID")
	accountID := accountIDStr // Sesuaikan tipe data accountID sesuai perubahan di model.go

	user, err := s.Services.GetUserByAccountID(r.Context(), accountID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpcode := strconv.Itoa(http.StatusOK)
	status := fmt.Sprintf("Success (%s)", httpcode)

	res := Response{
		Status: status,
		Data:   user,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, ErrFetchResp, http.StatusInternalServerError)
		return
	}
}
