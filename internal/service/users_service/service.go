package users_service

import (
	"context"

	"github.com/joshuatheokurniawansiregar/eljena/internal/models/users"
)

type UsersRepository interface {
	GetUser(ctx context.Context, email string) (*users.UserModel, error)
	CreateUser(ctx context.Context, userModel users.UserModel)error
}

type Service struct{
	UsersRepository UsersRepository
}

func NewService(usersRepository UsersRepository)*Service{
	return &Service{
		UsersRepository: usersRepository,
	}
}