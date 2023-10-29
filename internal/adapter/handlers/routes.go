package handlers

import "github.com/gin-gonic/gin"

type Router struct {
	CustomerCreatorHandler gin.HandlerFunc
	CustomerGetterHandler  gin.HandlerFunc
	ProductAddHandler      gin.HandlerFunc
	ProductInstanceHandler gin.HandlerFunc
}

func (r *Router) Register(app *gin.Engine) {
	delivery := app.Group("/v1/delivery")
	{
		//user routes
		delivery.POST("/customer", r.CustomerCreatorHandler)
		delivery.GET("/customer", r.CustomerGetterHandler)
		//product routes
		product := delivery.Group("/product")
		{
			product.POST("/add", r.ProductAddHandler)
			product.POST("/update", r.ProductInstanceHandler)
			product.POST("/delete", r.ProductInstanceHandler)
		}
		delivery.GET("/products", r.ProductInstanceHandler)
		// order routes
	}
}
