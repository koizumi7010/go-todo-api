package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/koizumi7010/go-todo-api/controller"
	"github.com/koizumi7010/go-todo-api/infrastructure/mysql"
	"github.com/koizumi7010/go-todo-api/usecase"
	"gorm.io/gorm"
)

func main() {
	db, err := mysql.NewDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	r := setupRouter(db)
	r.Run()
}

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	repository := mysql.NewTodo(db)
	usecase := usecase.NewTodo(repository)
	controller := controller.NewTodo(usecase)

	todo := r.Group("/todo")
	{
		todo.POST("/", controller.Create)
		todo.DELETE("/:id", controller.Delete)
		todo.PUT("/:id", controller.Update)
		todo.GET("/:id", controller.Get)
		todo.GET("/", controller.GetAll)
	}

	return r
}
