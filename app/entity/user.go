package entity

import (
	"context"
	"go-clean-architecture-be/domain"
)

type UserEntity struct {
	userRepo domain.UserRepository
}

func NewUserEntity (a domain.UserRepository) domain.UserEntity{
	return &UserEntity{userRepo: a}
}

func (a *UserEntity) LoginByEmail(c context.Context, email string, password string)error{
	//some business logic
	_,err := a.userRepo.LoginByEmail(c,email,password)

	if(err != nil){
		return err
	}

	return nil
}