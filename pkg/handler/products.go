package handler

import (
	"MyApi/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateProduct(c *gin.Context) {
	var input models.ProductInputFields

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid json", err.Error())
		return
	}

	id, err := h.Services.CreateProduct(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "corrupted creating product", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":     id,
		"status": "success",
	})
}
