package transport

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationKey = "Authorization"
	UserId           = "user_id"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationKey)

	splitParts := strings.Split(header, " ")
	if len(splitParts) != 2 {
		NewTransportErrorResponse(c, http.StatusBadRequest, "bad token (user doesn't authorized)")
		return
	}

	id, err := h.services.Authorization.ParseToken(splitParts[1])
	if err != nil {
		NewTransportErrorResponse(c, http.StatusBadRequest, "parse token failed")
		return
	}

	c.Set("user_id", id)
}
