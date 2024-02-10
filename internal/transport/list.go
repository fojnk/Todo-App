package transport

import (
	"net/http"
	"strconv"

	"github.com/fojnk/todo-app3/internal/models"
	"github.com/gin-gonic/gin"
)

type getAllListsResponses struct {
	Data []models.TodoList `json:"data"`
}

// @Summary Get all lists
// @Security ApiKeyAuth
// @Tags lists
// @Description Get all lists
// @ID get-all-lists
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} transort_error
// @Failure 500 {object} transort_error
// @Failure default {object} transort_error
// @Router /api/lists/ [get]
func (h *Handler) getAllLists(c *gin.Context) {
	id, ok := c.Get(UserId)
	if !ok {
		NewTransportErrorResponse(c, http.StatusBadRequest, "You are not authorized!!!")
		return
	}

	lists, err := h.services.TodoLists.GetAll(id.(int))
	if err != nil {
		NewTransportErrorResponse(c, http.StatusInternalServerError, "some problems while getting lists")
		return
	}

	c.JSON(http.StatusOK, &getAllListsResponses{
		Data: lists,
	})
}

// @Summary Create todo list
// @Security ApiKeyAuth
// @Tags lists
// @Description create todo list
// @ID create-list
// @Accept  json
// @Produce  json
// @Param input body models.TodoList true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} transort_error
// @Failure 500 {object} transort_error
// @Failure default {object} transort_error
// @Router /api/lists/ [post]
func (h *Handler) createList(c *gin.Context) {
	var input models.TodoList

	userId, ok := c.Get(UserId)
	if !ok {
		NewTransportErrorResponse(c, http.StatusBadRequest, "You are not authorized!!!")
		return
	}

	if err := c.BindJSON(&input); err != nil {
		NewTransportErrorResponse(c, http.StatusBadRequest, "Bad request")
		return
	}

	id, err := h.services.TodoLists.Create(userId.(int), input)
	if err != nil {
		NewTransportErrorResponse(c, http.StatusInternalServerError, "some problems")
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get list by Id
// @Security ApiKeyAuth
// @Tags lists
// @Description Get todo list by Id
// @ID get-list-by-id
// @Accept  json
// @Produce  json
// @Param id  path int true "List ID"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} transort_error
// @Failure 500 {object} transort_error
// @Failure default {object} transort_error
// @Router /api/lists/{id} [get]
func (h *Handler) getListById(c *gin.Context) {
	userId, ok := c.Get(UserId)
	if !ok {
		NewTransportErrorResponse(c, http.StatusBadRequest, "You not authorized!!!")
		return
	}

	list_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewTransportErrorResponse(c, http.StatusBadRequest, "bad list id")
		return
	}

	list, err := h.services.TodoLists.GetById(userId.(int), list_id)
	if err != nil {
		NewTransportErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

// @Summary Update list
// @Security ApiKeyAuth
// @Tags lists
// @Description Update list by Id
// @ID update-list
// @Accept  json
// @Produce  json
// @Param id  path int true "List ID"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} transort_error
// @Failure 500 {object} transort_error
// @Failure default {object} transort_error
// @Router /api/lists/{id} [put]
func (h *Handler) updateList(c *gin.Context) {
	userId, ok := c.Get(UserId)
	if !ok {
		NewTransportErrorResponse(c, http.StatusBadRequest, "You are not authorized!!!")
		return
	}

	list_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewTransportErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input models.TodoList
	if err := c.BindJSON(&input); err != nil {
		NewTransportErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoLists.Update(userId.(int), list_id, input); err != nil {
		NewTransportErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "i updated it")
}

// @Summary Delete list
// @Security ApiKeyAuth
// @Tags lists
// @Description Delete list by Id
// @ID delete-list
// @Accept  json
// @Produce  json
// @Param id  path int true "List ID"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} transort_error
// @Failure 500 {object} transort_error
// @Failure default {object} transort_error
// @Router /api/lists/{id} [delete]
func (h *Handler) deleteList(c *gin.Context) {
	userId, ok := c.Get(UserId)
	if !ok {
		NewTransportErrorResponse(c, http.StatusBadRequest, "You are not authorized!!!")
		return
	}

	list_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewTransportErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoLists.Delete(userId.(int), list_id); err != nil {
		NewTransportErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "i deleted it :D")
}
