package main

import (
	"database/sql"
	"github.com/betawulan/sahamrakyat/delivery"
	"log"

	"github.com/labstack/echo"
	"github.com/spf13/viper"

	"github.com/betawulan/sahamrakyat/repository"
	"github.com/betawulan/sahamrakyat/service"
)

func main() {
	viper.AutomaticEnv()
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("failed running because file .env")
	}

	dsn := viper.GetString("mysql_dsn")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("can't connect database")
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	e := echo.New()
	delivery.RegisterUserRoute(userService, e)

	e.Logger.Fatal(e.Start(":9090"))
}