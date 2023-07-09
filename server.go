package MyApi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Server(port string) error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "pong",
		})
	})

	if err := r.Run(port); err != nil {
		return err
	}
	return nil
}
