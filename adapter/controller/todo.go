package controller

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/masaya0521/go-clean-architecture-sample/usecase/port"
)

type Todo struct {
	OutputFactory func(w http.ResponseWriter) port.TodoOutputPort
	InputFactory func(o port.TodoOutputPort, u port.TodoRepository) port.TodoInputPort
	RepoFactory func(c *sql.DB) port.TodoRepository
	Conn *sql.DB
}

func (u *Todo) GetTodoByID (w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	todoID := strings.TrimPrefix(r.URL.Path, "/todo/")
	outputPort := u.OutputFactory(w)
    repository := u.RepoFactory(u.Conn)
	inputPort := u.InputFactory(outputPort, repository)
	inputPort.GetTodoByID(ctx, todoID)
}

func (u *Todo) GetTodo (w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	outputPort := u.OutputFactory(w)
    repository := u.RepoFactory(u.Conn)
	inputPort := u.InputFactory(outputPort, repository)
	inputPort.GetTodo(ctx)
}