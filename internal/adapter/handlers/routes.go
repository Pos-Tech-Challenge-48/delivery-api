package handlers

import "github.com/gin-gonic/gin"

type Router struct {
	CustomerCreatorHandler gin.HandlerFunc
}

func (r *Router) Register(app *gin.Engine) {
	app.POST("/v1/customer", r.CustomerCreatorHandler)
}
