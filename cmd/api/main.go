package main

import (
	"log"
	"net/http"

	_ "github.com/Pos-Tech-Challenge-48/delivery-api/cmd/api/docs"
	"github.com/Pos-Tech-Challenge-48/delivery-api/config"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/controllers"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/controllers/customercreatorhandler"
	customergetterhandler "github.com/Pos-Tech-Challenge-48/delivery-api/internal/controllers/customergetterandler"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/controllers/loginhandler"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/controllers/ordercreatorhandler"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/controllers/ordergetterhandler"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/controllers/orderupdatehandler"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/controllers/paymentcreatorhandler"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/controllers/paymentwebhookhandler"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/controllers/productcreatorhandler"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/controllers/productdeletehandler"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/controllers/productgetterhandler"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/controllers/productupdatehandler"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/external/api"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/external/db"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/external/repositories"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/usecases/customercreator"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/usecases/customergetter"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/usecases/login"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/usecases/ordercreator"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/usecases/ordergetter"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/usecases/orderupdater"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/usecases/paymentcreator"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/usecases/paymentwebhook"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/usecases/productcreator"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/usecases/productdelete"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/usecases/productgetter"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/usecases/productupdate"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	orderRepository := repositories.NewOrderRepository(db)

	paymentGateway := api.NewPaymentGateway("https://api.mercadopago.com/", "create_qr_code", "fake-api-key")

	authorizerAPI := api.NewAuthorizer()

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
	productGetterHandler := productgetterhandler.NewProductGetterHandler(productGetter)

	productUpdate := productupdate.NewProductUpdate(productRepository)
	productUpdateHandler := productupdatehandler.NewProductUpdateHandler(productUpdate)

	orderCreator := ordercreator.NewOrderCreator(orderRepository, productRepository)
	orderCreatorHandler := ordercreatorhandler.NewOrderCreatorHandler(orderCreator)

	orderGetter := ordergetter.NewOrderGetter(orderRepository, productRepository)
	orderGetterHandler := ordergetterhandler.NewOrderGetterHandler(orderGetter)

	paymentCreator := paymentcreator.NewPaymentCreator(orderRepository, paymentGateway)
	paymentCreatorHander := paymentcreatorhandler.NewPaymentCreatorHandler(paymentCreator)

	paymentWebhook := paymentwebhook.NewPaymentWebhook(orderRepository)
	paymentWebhookHander := paymentwebhookhandler.NewPaymentWebhookHandler(paymentWebhook)

	orderUpdater := orderupdater.NewOrderUpdater(orderRepository)
	orderUpdaterHandler := orderupdatehandler.NewOrderUpdaterHandler(orderUpdater)

	loginUseCase := login.NewLoginUseCase(customerRepository, authorizerAPI)
	loginHandler := loginhandler.NewLoginHandler(loginUseCase)

	app := gin.Default()

	router := controllers.Router{
		CustomerCreatorHandler: customerCreatorHandler.Handle,
		CustomerGetterHandler:  CustomerGetterHandler.Handle,
		OrderCreatorHandler:    orderCreatorHandler.Handle,
		OrderGetterHandler:     orderGetterHandler.Handle,
		ProductCreatorHandler:  productCreatorHandler.Handle,
		ProductDeleteHandler:   productDeleteHandler.Handle,
		ProductUpdateHandler:   productUpdateHandler.Handle,
		ProductGetterHandler:   productGetterHandler.Handle,
		PaymentCreatorHandler:  paymentCreatorHander.Handle,
		PaymentWebhookHandler:  paymentWebhookHander.Handle,
		OrderUpdaterHandler:    orderUpdaterHandler.Handle,
		LoginHandler:           loginHandler.Handle,
	}

	router.Register(app)

	app.GET("/healthcheck", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	app.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.Run(":8080")
}
