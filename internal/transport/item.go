package transport

import (
	"net/http"
	"strconv"

	"github.com/fojnk/todo-app3/internal/models"
	"github.com/gin-gonic/gin"
)

type AllItems struct {
	Data []models.TodoItem `json:"items"`
}

// @Summary Get all items
// @Security ApiKeyAuth
// @Tags items
// @Description Get all items
// @ID get-all-items
// @Accept  json
// @Produce  json
// @Param list_id  path int true "List ID"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} transort_error
// @Failure 500 {object} transort_error
// @Failure default {object} transort_error
// @Router /api/lists/{list_id}/ [get]
func (h *Handler) getAllItems(c *gin.Context) {
	user_id, ok := c.Get(UserId)
	if !ok {
		NewTransportErrorResponse(c, http.StatusBadRequest, "You are not authorized!!!")
		return
	}

	list_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewTransportErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	items, err := h.services.TodoItems.GetAll(user_id.(int), list_id)
	if err != nil {
		NewTransportErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &AllItems{
		Data: items,
	})
}

// @Summary Create item
// @Security ApiKeyAuth
// @Tags items
// @Description Create item
// @ID create-item
// @Accept  json
// @Produce  json
// @Param input body models.TodoItem true "item info"
// @Param list_id  path int true "List ID"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} transort_error
// @Failure 500 {object} transort_error
// @Failure default {object} transort_error
// @Router /api/lists/{list_id}/ [post]
func (h *Handler) createItem(c *gin.Context) {
	list_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewTransportErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input models.TodoItem
	if err := c.BindJSON(&input); err != nil {
		NewTransportErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	item_id, err := h.services.TodoItems.Create(list_id, input)
	if err != nil {
		NewTransportErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"item_id": item_id,
	})
}

// @Summary Get item by Id
// @Security ApiKeyAuth
// @Tags items
// @Description Get item by Id
// @ID get-item-by-id
// @Accept  json
// @Produce  json
// @Param list_id  path int true "List ID"
// @Param item_id  path int true "Item ID"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} transort_error
// @Failure 500 {object} transort_error
// @Failure default {object} transort_error
// @Router /api/lists/{list_id}/{item_id} [get]
func (h *Handler) getItemById(c *gin.Context) {
	item_id, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		NewTransportErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user_id, ok := c.Get(UserId)
	if !ok {
		NewTransportErrorResponse(c, http.StatusBadRequest, "You are not authorized!!!")
		return
	}

	item, err := h.services.TodoItems.GetItemById(user_id.(int), item_id)
	if err != nil {
		NewTransportErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

// @Summary Delete item by Id
// @Security ApiKeyAuth
// @Tags items
// @Description Delete item by Id
// @ID delete-item-by-id
// @Accept  json
// @Produce  json
// @Param list_id  path int true "List ID"
// @Param item_id  path int true "Item ID"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} transort_error
// @Failure 500 {object} transort_error
// @Failure default {object} transort_error
// @Router /api/lists/{list_id}/{item_id} [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	item_id, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		NewTransportErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user_id, ok := c.Get(UserId)
	if !ok {
		NewTransportErrorResponse(c, http.StatusBadRequest, "You are not authorized!!!")
		return
	}

	if err := h.services.TodoItems.Delete(user_id.(int), item_id); err != nil {
		NewTransportErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "i deleted it :D")
}

// @Summary Update item
// @Security ApiKeyAuth
// @Tags items
// @Description Update item by Id
// @ID update-item
// @Accept  json
// @Produce  json
// @Param list_id  path int true "List ID"
// @Param item_id  path int true "Item ID"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} transort_error
// @Failure 500 {object} transort_error
// @Failure default {object} transort_error
// @Router /api/lists/{list_id}/{item_id} [put]
func (h *Handler) updateItem(c *gin.Context) {
	item_id, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		NewTransportErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user_id, ok := c.Get(UserId)
	if !ok {
		NewTransportErrorResponse(c, http.StatusBadRequest, "You are not authorized!!!")
		return
	}

	var input models.TodoItem
	if err := c.BindJSON(&input); err != nil {
		NewTransportErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoItems.Update(user_id.(int), item_id, input); err != nil {
		NewTransportErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "i updated it :D")
}
