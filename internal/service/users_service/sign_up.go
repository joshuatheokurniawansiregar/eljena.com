package users_service

import (
	"context"
	"errors"
	"time"

	"github.com/joshuatheokurniawansiregar/eljena/internal/models/users"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) SignUp(ctx context.Context, signUpRequest users.SignUpRequest) error {
	// fmt.Println(os.Getenv("PORT"))
	
	userModel, err := s.UsersRepository.GetUser(ctx, signUpRequest.Email)
	if err != nil{
		return errors.New("GeUserError: "+err.Error())
	}

	if userModel != nil{
		return errors.New("email has been used")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpRequest.Password), bcrypt.DefaultCost)
	if err!= nil{
		return err
	}

	var now time.Time = time.Now()

	var newUser users.UserModel = users.UserModel{
		Email: signUpRequest.Email,
		Username: signUpRequest.Username,
		Password: string(hashedPassword),
		CreatedAt: now,
	}

	if err:= s.UsersRepository.CreateUser(ctx, newUser); err!=nil{
		return errors.New("CreateUserError: "+err.Error())
	}
	return nil
}