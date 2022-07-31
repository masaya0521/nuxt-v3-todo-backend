package port

import (
	"context"

	"github.com/masaya0521/go-clean-architecture-sample/entity"
)

type TodoInputPort interface {
	GetTodoByID(ctx context.Context, todoID string)
	GetTodo(ctx context.Context)
}

type TodoOutputPort interface {
	PostRender(*[]entity.Todo)
	Render(*entity.Todo)
	RenderError(error)
}

type TodoRepository interface {
	GetTodoByID(ctx context.Context, todoID string) (*entity.Todo, error)
	GetTodo(ctx context.Context)(*[]entity.Todo, error)
}