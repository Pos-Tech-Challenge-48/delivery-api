package handlers

import "github.com/gin-gonic/gin"

type Router struct {
	UserCreatorHandler gin.HandlerFunc
}

func (r *Router) Register(app *gin.Engine) {
	app.POST("/v1/user", r.UserCreatorHandler)
}
