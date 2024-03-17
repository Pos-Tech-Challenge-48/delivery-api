package loginhandler

import (
	"context"
	"errors"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-api/internal/interfaces/usecases/login"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginUseCase interfaces.LoginUseCase
}

func NewLoginHandler(LoginUseCase interfaces.LoginUseCase) *LoginHandler {
	return &LoginHandler{
		LoginUseCase: LoginUseCase,
	}
}

func (h *LoginHandler) Handle(c *gin.Context) {
	var login entities.Login

	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx := context.Background()

	customer, err := h.LoginUseCase.SignUp(ctx, &login)
	if err != nil {
		if errors.Is(err, entities.ErrInvalidCustomer) {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		if errors.Is(err, entities.ErrCustomerEmptyEmail) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if errors.Is(err, entities.ErrCustomerInvalidDocument) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if errors.Is(err, entities.ErrCustomerInvalidEmail) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}
