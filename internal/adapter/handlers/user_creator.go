package handlers

import (
	"context"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/domain"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type UserCreatorHandler struct {
	userCreatorUseCase ports.UserCreator
}

func NewUserCreatorHandler(userCreatorUseCase ports.UserCreator) *UserCreatorHandler {
	return &UserCreatorHandler{
		userCreatorUseCase: userCreatorUseCase,
	}
}

func (uh *UserCreatorHandler) Handle(c *gin.Context) {
	var user domain.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	ctx := context.Background()

	if err := uh.userCreatorUseCase.Create(ctx, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, user)
}
