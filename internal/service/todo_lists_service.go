package service

import (
	"github.com/fojnk/todo-app3/internal/models"
	"github.com/fojnk/todo-app3/internal/repository"
)

type ListsService struct {
	repo repository.TodoLists
}

func NewListsService(repo repository.TodoLists) *ListsService {
	return &ListsService{
		repo: repo,
	}
}

func (l *ListsService) Create(userId int, list models.TodoList) (int, error) {
	return l.repo.Create(userId, list)
}

func (l *ListsService) GetAll(userId int) ([]models.TodoList, error) {
	return l.repo.GetAll(userId)
}

func (l *ListsService) GetById(userId, id int) (models.TodoList, error) {
	return l.repo.GetById(userId, id)
}

func (l *ListsService) Delete(userId, id int) error {
	return l.repo.Delete(userId, id)
}

func (l *ListsService) Update(userId, id int, updatedList models.TodoList) error {
	return l.repo.Update(userId, id, updatedList)
}
