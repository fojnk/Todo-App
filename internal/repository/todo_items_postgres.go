package repository

import (
	"fmt"

	"github.com/fojnk/todo-app3/internal/models"
	"github.com/jmoiron/sqlx"
)

type ItemsPostgres struct {
	db *sqlx.DB
}

func NewItemsPostgres(db *sqlx.DB) *ItemsPostgres {
	return &ItemsPostgres{
		db: db,
	}
}

func (i *ItemsPostgres) Create(list_id int, item models.TodoItem) (int, error) {
	var id int
	tr, err := i.db.Begin()
	if err != nil {
		return 0, err
	}

	query1 := fmt.Sprintf("INSERT INTO %s (title, description, done) values ($1, $2, $3) RETURNING id", todoItemsTable)
	row := tr.QueryRow(query1, item.Title, item.Description, item.Done)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	query2 := fmt.Sprintf("INSERT INTO %s (item_id, list_id) values ($1, $2)", listsItemsTable)
	_, err = tr.Exec(query2, id, list_id)
	if err != nil {
		tr.Rollback()
		return 0, err
	}

	return id, tr.Commit()
}

func (i *ItemsPostgres) GetAll(user_id, list_id int) ([]models.TodoItem, error) {
	var items []models.TodoItem
	query := fmt.Sprintf(`SELECT il.id, il.title, il.description, il.done FROM %s il 
						INNER JOIN %s li on li.item_id = il.id
						INNER JOIN %s ul on li.list_id = ul.list_id
						WHERE ul.user_id = $1 AND li.list_id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)

	err := i.db.Select(&items, query, user_id, list_id)
	return items, err
}

func (i *ItemsPostgres) GetItemById(user_id, item_id int) (models.TodoItem, error) {
	var input models.TodoItem
	query := fmt.Sprintf(`SELECT il.id, il.title, il.description, il.done FROM %s il 
						INNER JOIN %s li on li.item_id = il.id
						INNER JOIN %s ul on li.list_id = ul.list_id
						WHERE ul.user_id = $1 AND il.id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)
	err := i.db.Get(&input, query, user_id, item_id)
	return input, err
}

func (i *ItemsPostgres) Delete(user_id, item_id int) error {
	query := fmt.Sprintf("DELETE FROM %s ti USING %s ul, %s li WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2",
		todoItemsTable, usersListsTable, listsItemsTable)

	_, err := i.db.Exec(query, user_id, item_id)
	return err
}

func (i *ItemsPostgres) Update(user_id, item_id int, updatedItem models.TodoItem) error {
	query := fmt.Sprintf(`UPDATE %s ti SET title = $1, description = $2, done = $3 FROM %s ul, %s li 
							WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $4 AND ti.id = $5`,
		todoItemsTable, usersListsTable, listsItemsTable)
	_, err := i.db.Exec(query, updatedItem.Title, updatedItem.Description, updatedItem.Done, user_id, item_id)
	return err
}
