package gateway

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/masaya0521/go-clean-architecture-sample/entity"
	"github.com/masaya0521/go-clean-architecture-sample/usecase/port"
)

type TodoRepository struct {
	conn *sql.DB
}

func NewTodoRepository(conn *sql.DB) port.TodoRepository{
	return &TodoRepository{
		conn: conn,
	}
}

func (u *TodoRepository) GetTodoByID(ctx context.Context, todoID string) (*entity.Todo, error) {
	conn := u.GetDBConn()
	row := conn.QueryRowContext(ctx, "SELECT * FROM `todo` WHERE todo_id=?", todoID)
	todo := entity.Todo{}
	err := row.Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Content, &todo.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Todo Not Found. TodoID = %s", todoID)
		}
		log.Println(err)
		return nil, errors.New("Internal Server Error. adapter/gateway/GetTodoByID")
	}
	return &todo, nil
}

func (u *TodoRepository) GetTodo(ctx context.Context)(*[]entity.Todo, error) {
	todos := []entity.Todo{}
	conn := u.GetDBConn()
	rows, err := conn.Query("SELECT * FROM `todo`")
	if err != nil {
		return nil, fmt.Errorf("Todo Empty")
	}
	for rows.Next() {
		todo := entity.Todo{}
		err = rows.Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Content, &todo.Status)
		if err != nil {
			return nil, fmt.Errorf("Scan failed")
		}
		todos = append(todos, todo)
	}
	return &todos, nil
}

// GetDBConn はconnectionを取得します．
func (u *TodoRepository) GetDBConn() *sql.DB {
	return u.conn
}