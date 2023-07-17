package handler

import (
	"MyApi/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{Services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	categories := router.Group("/categories", h.adminIdentity)
	{
		categories.POST("/", h.CreateCategory)
		categories.PUT(":id/", h.UpdateCategory)
		categories.DELETE(":id/", h.DeleteCategory)
		categories.GET("/", h.GetAllCategories)
		categories.GET(":id", h.GetCategory)
	}

	products := router.Group("/products")
	{
		products.POST("/", h.CreateProduct)
		products.GET("/", nil)
		products.GET(":id", nil)
	}

	comments := router.Group("/comments")
	{
		comments.GET("/", nil)
		comments.POST("/", nil)
	}

	tags := router.Group("/tags")
	{
		tags.POST("/", nil)
		tags.GET("/", nil)
	}
	return router
}
