package main

import (
	apiPkg "MyApi"
	"github.com/siruspen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// Logrus setting formatt
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Viper init config
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("error on reading configs: %s", err.Error())
	}

	// Start server
	if err := apiPkg.Server(viper.GetString("port")); err != nil {
		logrus.Errorf("error with start server: %s", err.Error())
	}

}
