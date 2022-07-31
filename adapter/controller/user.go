package controller

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/masaya0521/go-clean-architecture-sample/usecase/port"
)

type User struct {
	OutputFactory func(w http.ResponseWriter) port.UserOutputPort
	InputFactory func(o port.UserOutputPort, u port.UserRepository) port.UserInputPort
	RepoFactory func(c *sql.DB) port.UserRepository
	Conn *sql.DB
}

func (u *User) GetUserByID (w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := strings.TrimPrefix(r.URL.Path, "/user/")
	outputPort := u.OutputFactory(w)
    repository := u.RepoFactory(u.Conn)
	inputPort := u.InputFactory(outputPort, repository)
	inputPort.GetUserByID(ctx, userID)
}

func (u *User) GetUser (w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	outputPort := u.OutputFactory(w)
    repository := u.RepoFactory(u.Conn)
	inputPort := u.InputFactory(outputPort, repository)
	inputPort.GetUser(ctx)
}