package controller

import (
	"fmt"
	"net/http"
	"structapp/model"
	"structapp/repository"

	"github.com/labstack/echo"
)

func Test(c echo.Context) error {
	return c.String(http.StatusOK, "test")
}

func GetAllTodos(c echo.Context) error {
	var todo []model.Todo
	err := repository.GetAllTodos(&todo)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	} else {
		return c.JSON(http.StatusOK, todo)
	}
}

func CreateTodo(c echo.Context) error {
	todo := new(model.Todo)
	if err := c.Bind(todo); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err := repository.CreateTodo(todo)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusNotFound, err)
	} else {
		return c.JSON(http.StatusCreated, todo)
	}

}

func UpdateTodo(c echo.Context) error {
	id := c.Param("id")
	var todo model.Todo
	foundErr := repository.GetTodoById(&todo, id)
	// Todo bulunmaz ise 404 döndürülecek..
	if foundErr != nil {
		return c.JSON(http.StatusNotFound, foundErr.Error())
	}

	repository.UpdateTodo(id)
	return c.String(http.StatusOK, "id= "+id+" todo successfully updated")
}

func DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	var todo model.Todo
	foundErr := repository.GetTodoById(&todo, id)
	// Todo bulunmaz ise 404 döndürülecek..
	if foundErr != nil {
		return c.JSON(http.StatusNotFound, foundErr.Error())
	}

	repository.DeleteTodo(&todo, id)
	return c.String(http.StatusOK, "id= "+id+" todo successfully deleted")
}
