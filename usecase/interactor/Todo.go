package interactor

import (
	"context"

	"github.com/masaya0521/go-clean-architecture-sample/usecase/port"
)

type Todo struct {
	OutputPort port.TodoOutputPort
	TodoRepo   port.TodoRepository
}

func NewTodoInputPort(outputPort port.TodoOutputPort, userRepository port.TodoRepository) port.TodoInputPort {
	return &Todo {
		OutputPort: outputPort,
		TodoRepo: userRepository,
	} 
}

func (u *Todo) GetTodoByID (ctx context.Context, todoID string){
	user, err  := u.TodoRepo.GetTodoByID(ctx, todoID)
	if err != nil {
		u.OutputPort.RenderError(err)
		return 
	}
	u.OutputPort.Render(user)
}

func (u *Todo) GetTodo(ctx context.Context){
	user, err := u.TodoRepo.GetTodo(ctx)
	if err != nil {
		u.OutputPort.RenderError(err)
		return
	}
	u.OutputPort.PostRender(user)
}