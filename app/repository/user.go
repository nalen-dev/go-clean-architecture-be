package repository

import (
	"context"
	db "go-clean-architecture-be/db/sqlc"
	"go-clean-architecture-be/domain"
)

type UserRepository struct{
	Conn *db.SQLStore
}

func NewUserRepositry(Conn *db.SQLStore)domain.UserRepository{
	return &UserRepository{
		Conn,
	}
}

func (m *UserRepository) LoginByEmail(c context.Context, email string, password string) (domain.USER_ID, error){
	//some application business logic
	user,err :=	m.Conn.Queries.UserLoginByEmail(c,email)
	
	if(err != nil){
		return 0,err
	}

	return domain.USER_ID(user.ID),nil
}