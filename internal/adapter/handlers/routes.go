package handlers

import "github.com/gin-gonic/gin"

type Router struct {
	CustomerCreatorHandler gin.HandlerFunc
	CustomerGetterHandler  gin.HandlerFunc
	OrderCreatorHandler    gin.HandlerFunc
	ProductCreatorHandler  gin.HandlerFunc
	ProductDeleteHandler   gin.HandlerFunc
	ProductUpdateHandler   gin.HandlerFunc
	ProductGetterHandler   gin.HandlerFunc
}

func (r *Router) Register(app *gin.Engine) {
	delivery := app.Group("/v1/delivery")
	{
		//user routes
		delivery.POST("/customers", r.CustomerCreatorHandler)
		delivery.GET("/customers", r.CustomerGetterHandler)

		//product routes
		delivery.POST("/products", r.ProductCreatorHandler)
		delivery.PUT("/products", r.ProductUpdateHandler)
		delivery.DELETE("/products", r.ProductDeleteHandler)
		delivery.GET("/products", r.ProductGetterHandler)

		// order routes
		delivery.POST("/orders", r.OrderCreatorHandler)
	}
}
