package main

import (
	"github.com/ffelipelimao/hex-arch-api/internal/adapter/db"
	"github.com/ffelipelimao/hex-arch-api/internal/adapter/handlers"
	"github.com/ffelipelimao/hex-arch-api/internal/adapter/repositories"
	"github.com/ffelipelimao/hex-arch-api/internal/core/usecases"
	"github.com/gin-gonic/gin"
)

func main() {

	db := db.NewDatabase()
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
