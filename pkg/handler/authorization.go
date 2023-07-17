package handler

import (
	"MyApi/pkg/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input models.UserInputFields

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid json", fmt.Sprintf("error while binding json: %s", err.Error()))
		return
	}
	id, err := h.Services.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "corrupted user", fmt.Sprintf("corrupted creating user: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":     id,
		"status": "success",
	})
}

func (h *Handler) SignIn(c *gin.Context) {
	var input models.UserInputCreds
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid json", err.Error())
		return
	}

	token, err := h.Services.Authorization.SignIn(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "wrong credentials", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"authorization": token,
		"status":        "success",
	})
}
