package MyApi

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Server(r *gin.Engine) error {

	if err := r.Run(viper.GetString("port")); err != nil {
		return err
	}
	return nil
}
