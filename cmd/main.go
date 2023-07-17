package main

import (
	apiPkg "MyApi"
	"MyApi/pkg/handler"
	"MyApi/pkg/repository"
	"MyApi/pkg/service"
	"github.com/joho/godotenv"
	"github.com/siruspen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

type User struct {
	Id           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	UserName     string `json:"user_name"`
	PasswordHash string `json:"password_hash"`
}

func main() {
	// Logrus setting formatt
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Viper init config
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("error on reading configs: %s", err.Error())
	}

	// Init .env conf

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error with loading env file: %s", err.Error())
	}

	// DB connection
	db, err := repository.NewMySqlDB(repository.Config{
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.name"),
	})

	if err != nil {
		logrus.Fatalf("error with connection db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	defer db.Close()
	// Start server
	if err := apiPkg.Server(handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error with start server: %s", err.Error())
	}

}
