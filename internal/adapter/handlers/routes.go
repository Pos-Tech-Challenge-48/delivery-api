package handlers

import "github.com/gin-gonic/gin"

type Router struct {
	CustomerCreatorHandler gin.HandlerFunc
	CustomerGetterHandler  gin.HandlerFunc
	ProductHandler         gin.HandlerFunc
}

func (r *Router) Register(app *gin.Engine) {
	delivery := app.Group("/v1/delivery")
	{
		//user routes
		delivery.POST("/customer", r.CustomerCreatorHandler)
		delivery.GET("/customer", r.CustomerGetterHandler)

		//product routes
		delivery.POST("/products", r.ProductHandler)
		delivery.PUT("/products", r.ProductHandler)
		delivery.DELETE("/products", r.ProductHandler)
		delivery.GET("/products", r.ProductHandler)
		// order routes
	}
}
