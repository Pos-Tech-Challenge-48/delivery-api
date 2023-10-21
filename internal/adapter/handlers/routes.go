package handlers

import "github.com/gin-gonic/gin"

type Router struct {
	CustomerCreatorHandler gin.HandlerFunc
	CustomerGetterHandler  gin.HandlerFunc
}

func (r *Router) Register(app *gin.Engine) {
	app.POST("/v1/customer", r.CustomerCreatorHandler)
	app.GET("/v1/customer", r.CustomerGetterHandler)
}
