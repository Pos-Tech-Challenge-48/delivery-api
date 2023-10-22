package main

import (
	"fmt"

	"github.com/Pos-Tech-Challenge-48/delivery-api/config"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/adapter/db"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/adapter/handlers"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/adapter/repositories"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/usecases"
	"github.com/gin-gonic/gin"
)

func main() {

	config, err := config.LoadConfig()

	if err != nil {
		fmt.Println("error loading configs")
	}

	db := db.NewDatabase(config)
	userRepository := repositories.NewUserRepository(db)
	userCreator := usecases.NewUserCreator(userRepository)
	userCreatorHandler := handlers.NewUserCreatorHandler(userCreator)

	app := gin.Default()

	router := handlers.Router{
		UserCreatorHandler: userCreatorHandler.Handle,
	}
	router.Register(app)

	app.Run(":8080")

}
