package repository

import (
	"github.com/fojnk/todo-app3/internal/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
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

type Respository struct {
	Authorization
	TodoLists
	TodoItems
}

func NewRepository(db *sqlx.DB) *Respository {
	return &Respository{
		Authorization: NewAuthPostgres(db),
		TodoLists:     NewListsPostgres(db),
		TodoItems:     NewItemsPostgres(db),
	}
}
