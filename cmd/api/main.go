package main

import (
	"log"

	_ "github.com/Pos-Tech-Challenge-48/delivery-api/cmd/api/docs"
	"github.com/Pos-Tech-Challenge-48/delivery-api/config"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/adapter/db"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/adapter/handlers"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/adapter/handlers/customercreatorhandler"
	customergetterhandler "github.com/Pos-Tech-Challenge-48/delivery-api/internal/adapter/handlers/customergetterandler"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/adapter/handlers/ordercreatorhandler"
	producthandler "github.com/Pos-Tech-Challenge-48/delivery-api/internal/adapter/handlers/product"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/adapter/repositories"
	usecase "github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/usecases"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/usecases/customercreator"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/usecases/customergetter"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/usecases/ordercreator"
	"github.com/gin-gonic/gin"
)

// @title Delivery API
// @version 1.0
// @description Aplicativo que gerencia atividades de um serviço de pedidos em um restaurante. Desde a base de clientes, catálogo de produtos, pedidos e fila de preparo
// @host localhost:8080
// @BasePath /v1/delivery
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

	customerGetter := customergetter.NewCustomerGetter(customerRepository)
	CustomerGetterHandler := customergetterhandler.NewCustomerGetterHandler(customerGetter)

	productRepository := repositories.NewProductRepository(db)
	productService := usecase.NewProductService(productRepository)
	productHandler := producthandler.NewProductHandler(productService)

	orderRepository := repositories.NewOrderRepository(db)
	orderCreator := ordercreator.NewOrderCreator(orderRepository, productRepository)
	orderCreatorHandler := ordercreatorhandler.NewOrderCreatorHandler(orderCreator)

	app := gin.Default()

	router := handlers.Router{
		CustomerCreatorHandler: customerCreatorHandler.Handle,
		CustomerGetterHandler:  CustomerGetterHandler.Handle,
		ProductHandler:         productHandler.Handle,
		OrderCreatorHandler:    orderCreatorHandler.Handle,
	}
	router.Register(app)

	app.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.Run(":8080")

}
