package main

import (
	"log"

	"github.com/ffelipelimao/Pos-Tech-Challenge-48/delivery-api/config"
	"github.com/ffelipelimao/Pos-Tech-Challenge-48/delivery-api/internal/adapter/db"
	"github.com/ffelipelimao/Pos-Tech-Challenge-48/delivery-api/internal/adapter/handlers"
	"github.com/ffelipelimao/Pos-Tech-Challenge-48/delivery-api/internal/adapter/handlers/customercreatorhandler"
	"github.com/ffelipelimao/Pos-Tech-Challenge-48/delivery-api/internal/adapter/handlers/customergethandler"
	"github.com/ffelipelimao/Pos-Tech-Challenge-48/delivery-api/internal/adapter/repositories"
	"github.com/ffelipelimao/Pos-Tech-Challenge-48/delivery-api/internal/core/usecases/customercreator"
	"github.com/ffelipelimao/Pos-Tech-Challenge-48/delivery-api/internal/core/usecases/customergetdocument"
	"github.com/gin-gonic/gin"
)

func main() {

	config, err := config.LoadConfig()

	if err != nil {
		log.Panic("error to load config", err)
	}

	db := db.NewDatabase(config)
	defer db.Close()
	customerRepository := repositories.NewCustomerRepository(db)

	customerCreator := customercreator.NewCustomerCreator(customerRepository)
	customerCreatorHandler := customercreatorhandler.NewCustomerCreatorHandler(customerCreator)

	customerGetter := customergetdocument.NewCustomerGetter(customerRepository)
	CustomerGetterHandler := customergethandler.NewCustomerGetterHandler(customerGetter)

	app := gin.Default()

	router := handlers.Router{
		CustomerCreatorHandler: customerCreatorHandler.Handle,
		CustomerGetterHandler:  CustomerGetterHandler.Handle,
	}
	router.Register(app)

	app.Run(":8080")

}
