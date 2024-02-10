package models

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type UsersLists struct {
	Id     int
	UserId int
	ListId int
}

type ListsItems struct {
	Id     int
	ItemId int
	ListId int
}
