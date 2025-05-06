package users_repository

import (
	"context"
	"database/sql"

	"github.com/joshuatheokurniawansiregar/eljena/internal/models/users"
)

func (r *Repository) GetUser(ctx context.Context, email string)(*users.UserModel,  error){
	var query string = `SELECT id, email, username, password, created_at FROM users WHERE email=$1`
	// fmt.Println(email)
	var row *sql.Row = r.DB.QueryRowContext(ctx, query, email)
	var userModel users.UserModel

	err:= row.Scan(&userModel.ID, &userModel.Email, &userModel.Username, &userModel.Password, &userModel.CreatedAt)
	if err == sql.ErrNoRows{
		return nil, nil
	}else if err != nil{
		return nil, err
	}
	return &userModel, nil
}

func (r *Repository) CreateUser(ctx context.Context, userModel users.UserModel)error{
	var query string = "INSERT INTO users(email, username, password, created_at) VALUES($1, $2, $3, $4)"
	var _, err = r.DB.ExecContext(ctx, query, userModel.Email, userModel.Username, userModel.Password, userModel.CreatedAt)
	if err != nil{
		return err
	}
	return nil
}