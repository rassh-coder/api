package handler

import (
	"MyApi/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateCategory(c *gin.Context) {
	var input models.CategoryInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid json", err.Error())
		return
	}
	id, err := h.Services.CreateCategory(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "corrupted creating", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":     id,
		"status": "success",
	})
}

func (h *Handler) GetAllCategories(c *gin.Context) {
	data, err := h.Services.GetAllCategories()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "db error", err.Error())
		return
	}

	responseWithData(c, http.StatusOK, dataResponse{
		Data:   data,
		Status: "success",
	})
}

func (h *Handler) UpdateCategory(c *gin.Context) {
	var input models.CategoryInput

	id, err := strconv.Atoi(c.Param("id"))

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid json", err.Error())
		return
	}

	category, err := h.Services.UpdateCategory(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "category is not found", err.Error())
		return
	}

	responseWithData(c, http.StatusOK, dataResponse{
		Data:   category,
		Status: "success",
	})
}

func (h *Handler) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param", err.Error())
		return
	}

	err = h.Services.DeleteCategory(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "category is not found", err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}

func (h *Handler) GetCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param", err.Error())
		return
	}

	category, err := h.Services.GetCategory(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "category is not found", err.Error())
		return
	}

	responseWithData(c, http.StatusOK, dataResponse{
		Data:   category,
		Status: "success",
	})
}
