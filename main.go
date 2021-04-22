package main

import (
	"fmt"
	"structapp/config"
	"structapp/controller"
	"structapp/model"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("test")

	dsn := config.BuildDBConfig()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	config.DB = db
	if err != nil {
		fmt.Println("No Connected")
	}
	config.DB.AutoMigrate(&model.Todo{})

	e := echo.New()

	e.GET("/test", controller.Test)
	// Get All Todos
	e.GET("/todo", controller.GetAllTodos)
	//Save Todo
	e.POST("/todo", controller.CreateTodo)
	//Update Todo
	e.GET("/todo/update/:id", controller.UpdateTodo)
	//Delete Todo
	e.DELETE("/todo/:id", controller.DeleteTodo)

	e.Logger.Fatal(e.Start(":8080"))
}
