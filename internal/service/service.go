package service

import (
	"github.com/fojnk/todo-app3/internal/models"
	"github.com/fojnk/todo-app3/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoLists interface {
	Create(userId int, list models.TodoList) (int, error)
	GetAll(userId int) ([]models.TodoList, error)
	GetById(userId, id int) (models.TodoList, error)
	Delete(userId int, id int) error
	Update(userId, id int, updatedList models.TodoList) error
}

type TodoItems interface {
	Create(list_id int, item models.TodoItem) (int, error)
	GetAll(user_id, list_id int) ([]models.TodoItem, error)
	GetItemById(user_id, item_id int) (models.TodoItem, error)
	Delete(user_id, item_id int) error
	Update(user_id, item_id int, updatedItem models.TodoItem) error
}

type Service struct {
	Authorization
	TodoLists
	TodoItems
}

func NewService(repos *repository.Respository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoLists:     NewListsService(repos.TodoLists),
		TodoItems:     NewItemService(repos.TodoItems),
	}
}
