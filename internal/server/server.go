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
	"github.com/vier21/tefa-ch3/config"
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
			Addr:         config.GetConfig().ServerPort,
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
	r.Get("/user/accountID", a.GetUserByAccountIDHandler)


	go func() {
		log.Printf("Server start on localhost%s \n", config.GetConfig().ServerPort)
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

func (s *ApiServer) GetUserByAccountIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	accountIDStr := chi.URLParam(r, "accountID")
	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		http.Error(w, "invalid accountID", http.StatusBadRequest)
		return
	}

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

