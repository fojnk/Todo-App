package service

import (
	"github.com/fojnk/todo-app3/internal/models"
	"github.com/fojnk/todo-app3/internal/repository"
)

type ItemService struct {
	repo repository.TodoItems
}

func NewItemService(repo repository.TodoItems) *ItemService {
	return &ItemService{
		repo: repo,
	}
}

func (i *ItemService) Create(list_id int, item models.TodoItem) (int, error) {
	return i.repo.Create(list_id, item)
}

func (i *ItemService) GetAll(user_id, list_id int) ([]models.TodoItem, error) {
	return i.repo.GetAll(user_id, list_id)
}

func (i *ItemService) GetItemById(user_id, item_id int) (models.TodoItem, error) {
	return i.repo.GetItemById(user_id, item_id)
}

func (i *ItemService) Delete(user_id, item_id int) error {
	return i.repo.Delete(user_id, item_id)
}

func (i *ItemService) Update(user_id, item_id int, updatedItem models.TodoItem) error {
	return i.repo.Update(user_id, item_id, updatedItem)
}
