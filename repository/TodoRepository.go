package repository

import (
	"fmt"
	"structapp/config"
	"structapp/model"
)

// type TodoRepository interface {
// 	GetAllTodos() []model.Todo
// }

func GetAllTodos(todo *[]model.Todo) (err error) {
	if err = config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateTodo(todo *model.Todo) (err error) {
	if err = config.DB.Create(todo).Error; err != nil {
		fmt.Println("hata")
		return err
	}
	return nil
}

func UpdateTodo(id string) (err error) {
	config.DB.Model(&model.Todo{}).Where("id = ?", id).Update("statu", "1")
	return nil
}

func DeleteTodo(id string) (err error) {
	config.DB.Where("id = ?", id).Delete(&model.Todo{})
	return nil
}
