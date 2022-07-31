package presenter

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/masaya0521/go-clean-architecture-sample/entity"
	"github.com/masaya0521/go-clean-architecture-sample/usecase/port"
)

type Todo struct {
	w http.ResponseWriter
}

func NewTodoOutputPort(w http.ResponseWriter) port.TodoOutputPort{
	return &Todo{
		w: w,
	}
}

// usecase.TodoOutputPortを実装している
// Render はNameを出力します．
func (u *Todo) Render(todo *entity.Todo) {
	u.w.WriteHeader(http.StatusOK)
	// httpでentity.Todo.Nameを出力
	fmt.Fprint(u.w, todo.Title)
}

// RenderError はErrorを出力します．
func (u *Todo) RenderError(err error) {
	u.w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(u.w, err)
}

func (u *Todo) PostRender(todos *[]entity.Todo) {
	u.w.WriteHeader(http.StatusOK)
	// httpでentity.Todo.Nameを出力
	json, err := json.Marshal(&todos)
	if err != nil {
		fmt.Println("json error")
	}
	fmt.Fprint(u.w,string(json))
}