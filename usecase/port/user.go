package port

import (
	"context"

	"github.com/masaya0521/go-clean-architecture-sample/entity"
)

type UserInputPort interface {
	GetUserByID(ctx context.Context, userID string)
	GetUser(ctx context.Context)
}

type UserOutputPort interface {
	PostRender(*[]entity.User)
	Render(*entity.User)
	RenderError(error)
}

type UserRepository interface {
	GetUserByID(ctx context.Context, userID string) (*entity.User, error)
	GetUser(ctx context.Context)(*[]entity.User, error)
}