package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/siruspen/logrus"
)

type errorResponse struct {
	Message       string `json:"message"`
	ServerMessage string `json:"-"`
	Status        string `json:"status"`
}

type dataResponse struct {
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message, serverMessage string) {
	logrus.Error(serverMessage)
	c.AbortWithStatusJSON(statusCode, errorResponse{
		message,
		serverMessage,
		"error",
	})
}

func responseWithData(c *gin.Context, statusCode int, messageWithData dataResponse) {
	c.JSON(statusCode, messageWithData)
}
