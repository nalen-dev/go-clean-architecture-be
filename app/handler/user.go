package handler

import (
	"go-clean-architecture-be/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersHandler struct{
	UserEntity domain.UserEntity
}

func NewUsersHandler(r *gin.RouterGroup, us domain.UserEntity){
	handler := &UsersHandler{
		UserEntity: us,
	}
	r.POST("/users", handler.LoginByEmail)

}


func (a *UsersHandler ) LoginByEmail(c *gin.Context){
	
	err := a.UserEntity.LoginByEmail(c.Request.Context(),"","")

	if err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"message":"Login Failed"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"Login Success"})
} 
