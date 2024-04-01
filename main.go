package main

import (
	"context"
	"go-clean-architecture-be/app/entity"
	"go-clean-architecture-be/app/handler"
	"go-clean-architecture-be/app/repository"
	db "go-clean-architecture-be/db/sqlc"
	"os"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main(){
	r := gin.Default()
	r.Use(cors.Default())

	conn,_ := pgxpool.New(context.Background(),os.Getenv("DATABASE_URL"))
	store := db.NewStore(conn)

	repoUser := repository.NewUserRepositry(&store)
	entityUser := entity.NewUserEntity(repoUser)

	api := r.Group("/api/v1")

	handler.NewUsersHandler(api, entityUser)

	r.Run()
}