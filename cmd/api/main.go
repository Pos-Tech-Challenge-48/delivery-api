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
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/adapter/handlers/ordergetterhandler"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/adapter/handlers/productcreatorhandler"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/adapter/handlers/productdeletehandler"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/adapter/handlers/productgetterhandler"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/adapter/handlers/productupdatehandler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/adapter/repositories"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/usecases/customercreator"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/usecases/customergetter"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/usecases/ordercreator"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/usecases/ordergetter"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/usecases/productcreator"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/usecases/productdelete"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/usecases/productgetter"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/usecases/productupdate"
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

	productCreator := productcreator.NewProductCreator(productRepository)
	productCreatorHandler := productcreatorhandler.NewProductCreatorHandler(productCreator)

	productDelete := productdelete.NewProductDelete(productRepository)
	productDeleteHandler := productdeletehandler.NewProductDeleteHandler(productDelete)

	productGetter := productgetter.NewProductGetter(productRepository)
	productGetterHandler := productgetterhandler.NewProductGetterHandler((productGetter))

	productUpdate := productupdate.NewProductUpdate(productRepository)
	productUpdateHandler := productupdatehandler.NewProductUpdateHandler(productUpdate)

	orderRepository := repositories.NewOrderRepository(db)
	orderCreator := ordercreator.NewOrderCreator(orderRepository, productRepository)
	orderCreatorHandler := ordercreatorhandler.NewOrderCreatorHandler(orderCreator)

	orderGetter := ordergetter.NewOrderGetter(orderRepository, productRepository)
	orderGetterHandler := ordergetterhandler.NewOrderGetterHandler(orderGetter)

	app := gin.Default()

	router := handlers.Router{
		CustomerCreatorHandler: customerCreatorHandler.Handle,
		CustomerGetterHandler:  CustomerGetterHandler.Handle,
		OrderCreatorHandler:    orderCreatorHandler.Handle,
		OrderGetterHandler:     orderGetterHandler.Handle,
		ProductCreatorHandler:  productCreatorHandler.Handle,
		ProductDeleteHandler:   productDeleteHandler.Handle,
		ProductUpdateHandler:   productUpdateHandler.Handle,
		ProductGetterHandler:   productGetterHandler.Handle,
	}

	router.Register(app)

	app.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.Run(":8080")

}
