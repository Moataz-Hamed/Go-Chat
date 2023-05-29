package user

import (
	"context"
)

type User struct {
	ID       int64  `json:"id" db:"id"`
	UserName string `json:"username" db:"username"`
	Email    string
	Password string
}

type LoginUserReq struct {
	Email    string
	Password string
}

type LoginUserRes struct {
	accessToken string
	ID          string `json:"id" db:"id"`
	UserName    string `json:"username" db:"username"`
}

type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type Service interface {
	CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error)
	Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error)
}

type CreateUserReq struct {
	UserName string `json:"username" db:"username"`
	Email    string
	Password string
}

type CreateUserRes struct {
	ID       string `json:"id" db:"id"`
	UserName string `json:"username" db:"username"`
	Email    string
}
