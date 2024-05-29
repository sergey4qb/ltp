package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sergey4qb/ltp/internal/core/domain/valueobject"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockPriceService struct {
	mock.Mock
}

func (m *MockPriceService) GetPricesForPairs(ctx context.Context, pairs []*valueobject.Pair) ([]valueobject.LTP, error) {
	args := m.Called(ctx, pairs)
	return args.Get(0).([]valueobject.LTP), args.Error(1)
}

func TestGetPricesForPairs(t *testing.T) {
	mockService := &MockPriceService{}
	controller := NewPriceController(mockService)

	expectedPairs := []*valueobject.Pair{
		valueobject.NewPair("BTC", "USD"),
		valueobject.NewPair("BTC", "CHF"),
		valueobject.NewPair("BTC", "EUR"),
	}
	expectedValues := []valueobject.LTP{{
		Pair:   expectedPairs[0].ToStringDivider(),
		Amount: "50000.0",
	}, {
		Pair:   expectedPairs[1].ToStringDivider(),
		Amount: "45000.0",
	}, {
		Pair:   expectedPairs[2].ToStringDivider(),
		Amount: "47000.0",
	}}
	mockService.On(
		"GetPricesForPairs",
		mock.Anything,
		expectedPairs,
	).Return(expectedValues, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/api/v1/ltp", nil)

	controller.GetPricesForPairs(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{
    "ltp": [
        {
            "pair": "`+expectedValues[0].Pair+`",
            "amount": "`+expectedValues[0].Amount+`"
        },
        {
            "pair": "`+expectedValues[1].Pair+`",
            "amount": "`+expectedValues[1].Amount+`"
        },
        {
            "pair": "`+expectedValues[2].Pair+`",
            "amount": "`+expectedValues[2].Amount+`"
        }
    ]
}`, w.Body.String())

}

func TestGetPricesForPairs_Error(t *testing.T) {
	mockService := &MockPriceService{}
	controller := NewPriceController(mockService)

	expectedPairs := []*valueobject.Pair{
		valueobject.NewPair("BTC", "USD"),
		valueobject.NewPair("BTC", "CHF"),
		valueobject.NewPair("BTC", "EUR"),
	}
	expectedError := fmt.Errorf("some error occurred")
	mockService.On(
		"GetPricesForPairs",
		mock.Anything,
		expectedPairs,
	).Return([]valueobject.LTP{}, expectedError)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/api/v1/ltp", nil)

	controller.GetPricesForPairs(c)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.JSONEq(t, `{"error":"some error occurred"}`, w.Body.String())
}
