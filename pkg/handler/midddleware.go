package handler

import (
	"MyApi/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
)

func (h *Handler) adminIdentity(c *gin.Context) {
	header := c.GetHeader(authHeader)
	params := strings.Split(header, " ")
	if len(params) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header", "empty auth header")
		return
	}
	if params[1] == "" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid token", "invalid token")
		return
	}
	isValid := service.ParseSUToken(params[1])
	if !isValid {
		newErrorResponse(c, http.StatusUnauthorized, "no permissions", "invalid admin token")
		return
	}
}
