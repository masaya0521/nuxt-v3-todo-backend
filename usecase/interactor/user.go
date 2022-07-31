package interactor

import (
	"context"

	"github.com/masaya0521/go-clean-architecture-sample/usecase/port"
)

type User struct {
	OutputPort port.UserOutputPort
	UserRepo   port.UserRepository
}

func NewUserInputPort(outputPort port.UserOutputPort, userRepository port.UserRepository) port.UserInputPort {
	return &User {
		OutputPort: outputPort,
		UserRepo: userRepository,
	} 
}

func (u *User) GetUserByID (ctx context.Context, userID string){
	user, err  := u.UserRepo.GetUserByID(ctx, userID)
	if err != nil {
		u.OutputPort.RenderError(err)
		return 
	}
	u.OutputPort.Render(user)
}

func (u *User) GetUser(ctx context.Context){
	user, err := u.UserRepo.GetUser(ctx)
	if err != nil {
		u.OutputPort.RenderError(err)
		return
	}
	u.OutputPort.PostRender(user)
}