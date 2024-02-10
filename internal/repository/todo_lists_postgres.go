package repository

import (
	"fmt"

	"github.com/fojnk/todo-app3/internal/models"
	"github.com/jmoiron/sqlx"
)

type ListsPostgres struct {
	db *sqlx.DB
}

func NewListsPostgres(db *sqlx.DB) *ListsPostgres {
	return &ListsPostgres{db: db}
}

func (l *ListsPostgres) Create(userId int, list models.TodoList) (int, error) {
	var id int
	tr, err := l.db.Begin()
	if err != nil {
		return 0, err
	}

	query1 := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", todoListsTable)
	row := tr.QueryRow(query1, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	query2 := fmt.Sprintf("INSERT INTO %s (user_id, list_id) values ($1, $2)", usersListsTable)
	_, err = tr.Exec(query2, userId, id)
	if err != nil {
		tr.Rollback()
		return 0, err
	}

	return id, tr.Commit()
}

func (l *ListsPostgres) GetAll(userId int) ([]models.TodoList, error) {
	var lists []models.TodoList
	query := fmt.Sprintf("SELECT lt.id, lt.title, lt.description FROM %s lt INNER JOIN %s ul on ul.list_id = lt.id WHERE ul.user_id = $1", todoListsTable, usersListsTable)

	err := l.db.Select(&lists, query, userId)
	if err != nil {
		return nil, err
	}
	return lists, nil
}

func (l *ListsPostgres) GetById(userId, id int) (models.TodoList, error) {
	var list models.TodoList
	query := fmt.Sprintf("SELECT lt.id, lt.title, lt.description FROM %s lt INNER JOIN %s ul on ul.list_id = lt.id WHERE ul.user_id = $1 AND lt.id = $2",
		todoListsTable, usersListsTable)

	err := l.db.Get(&list, query, userId, id)
	return list, err
}

func (l *ListsPostgres) Delete(userId int, id int) error {
	query := fmt.Sprintf("DELETE FROM %s lt USING %s ul WHERE ul.list_id = lt.id AND ul.user_id = $1 AND lt.id = $2",
		todoListsTable, usersListsTable)
	_, err := l.db.Exec(query, userId, id)
	return err
}

func (l *ListsPostgres) Update(userId, id int, updatedList models.TodoList) error {
	query := fmt.Sprintf("UPDATE %s lt SET title = $1, description = $2 FROM %s ul WHERE lt.id = ul.list_id AND ul.user_id = $3 AND lt.id = $4",
		todoListsTable, usersListsTable)

	_, err := l.db.Exec(query, updatedList.Title, updatedList.Description, userId, id)
	return err
}
