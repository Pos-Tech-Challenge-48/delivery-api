package ordersorterhandler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/controllers/ordersorterhandler"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/mocks/ordersortermock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_OrderSorter_Handler(t *testing.T) {

	fakeSortedOrders := []entities.Order{
		{ID: "1", CustomerID: "uuid1", Status: "Pronto"},
		{ID: "2", CustomerID: "uuid2", Status: "Em Preparação"},
	}

	tests := []struct {
		name               string
		expectedStatusCode int
		executeMock        func(m *ordersortermock.MockOrderSorter)
	}{
		{
			name:               "Should retrieve sorted orders successfully",
			expectedStatusCode: http.StatusOK,
			executeMock: func(m *ordersortermock.MockOrderSorter) {
				m.EXPECT().GetAllSortedByStatus(gomock.Any()).Return(fakeSortedOrders, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			orderSorterUseCaseMock := ordersortermock.NewMockOrderSorter(ctrl)
			tt.executeMock(orderSorterUseCaseMock)

			appFake := gin.Default()
			endpoint := "/v1/orders/sorted"
			orderSorterHandler := ordersorterhandler.NewOrderSorterHandler(orderSorterUseCaseMock)

			appFake.GET(endpoint, orderSorterHandler.Handle)
			req, _ := http.NewRequest("GET", endpoint, nil)
			w := httptest.NewRecorder()
			appFake.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}
