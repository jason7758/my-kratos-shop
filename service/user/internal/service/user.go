package service

import (
	v1 "user/api/user/v1"
	"user/internal/biz"
)

type UserService struct {
	v1.UnimplementedUserServer

	uc *biz.UserUsecase
}

func NewUserService() *UserService {
	return &UserService{}
}
