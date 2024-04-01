package domain

import "context"


type (USER_ID int64; EMAIL string;NAME string;PASSWORD string)

type User struct{
	ID			USER_ID  `json:"id"`
	Name 		NAME 	 `json:"name"`
	Email 		EMAIL 	 `json:"email"`
	Password	PASSWORD `json:"password"`
}

type UserEntity interface{
	LoginByEmail(ctx context.Context, email string,password string) (error)
}

type UserRepository interface{
	LoginByEmail(ctx context.Context, email string, password string) (USER_ID,error)
}