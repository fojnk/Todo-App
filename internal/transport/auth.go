package transport

import (
	"net/http"

	"github.com/fojnk/todo-app3/internal/models"
	"github.com/gin-gonic/gin"
)

type InputAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body models.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404,500,default {object} transort_error
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		NewTransportErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewTransportErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

// @Summary SignIn
// @Tags auth
// @Description login to account
// @ID login-to-account
// @Accept  json
// @Produce  json
// @Param input body InputAuth true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404,500,default {object} transort_error
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input InputAuth

	if err := c.BindJSON(&input); err != nil {
		NewTransportErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		NewTransportErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}
