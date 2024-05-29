package price_service

import (
	"context"
	"errors"
	"github.com/sergey4qb/ltp/internal/core/domain/valueobject"
	"github.com/sergey4qb/ltp/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockTickerClient struct {
	mock.Mock
}

func (m *MockTickerClient) GetTickerInfo(ctx context.Context, pair string) (*dto.TickerResponse, error) {
	args := m.Called(ctx, pair)
	return args.Get(0).(*dto.TickerResponse), args.Error(1)
}

func TestGetLastTradedPrice(t *testing.T) {
	mockClient := &MockTickerClient{}
	service := NewService(mockClient)
	ctx := context.Background()
	exp := "45000.0"
	pair := &valueobject.Pair{BaseCurrency: "BTC", QuoteCurrency: "USD"}
	krakenStandard, err := pair.ToStringKraken()
	assert.NoError(t, err)

	mockTickerInfo := &dto.TickerResponse{
		Error:  nil,
		Result: map[string]dto.Pair{krakenStandard: {LastTradeClosed: [2]string{exp}}},
	}
	mockClient.On(
		"GetTickerInfo",
		ctx,
		"BTCUSD",
	).Return(mockTickerInfo, nil)

	price, err := service.getLastTradedPrice(ctx, pair)
	assert.NoError(t, err)
	assert.Equal(t, exp, price)

	mockClient.AssertExpectations(t)
}

func TestGetPricesForPairs(t *testing.T) {
	mockClient := &MockTickerClient{}
	service := NewService(mockClient)
	ctx := context.Background()
	BTCUSDExp, BTCEURExp := "45000.0", "42.000"
	pairs := []*valueobject.Pair{
		{BaseCurrency: "BTC", QuoteCurrency: "USD"},
		{BaseCurrency: "BTC", QuoteCurrency: "EUR"},
	}
	krakenStandardBTCUSD, err := pairs[0].ToStringKraken()
	assert.NoError(t, err)
	krakenStandardBTCEUR, err := pairs[1].ToStringKraken()
	assert.NoError(t, err)
	mockTickerInfoBTCUSD, mockTickerInfoBTCEUR := &dto.TickerResponse{
		Error:  nil,
		Result: map[string]dto.Pair{krakenStandardBTCUSD: {LastTradeClosed: [2]string{BTCUSDExp}}},
	}, &dto.TickerResponse{
		Error:  nil,
		Result: map[string]dto.Pair{krakenStandardBTCEUR: {LastTradeClosed: [2]string{BTCEURExp}}},
	}
	mockClient.On("GetTickerInfo", ctx, "BTCUSD").Return(mockTickerInfoBTCUSD, nil)
	mockClient.On("GetTickerInfo", ctx, "BTCEUR").Return(mockTickerInfoBTCEUR, nil)

	results, err := service.GetPricesForPairs(ctx, pairs)
	assert.NoError(t, err)
	assert.Len(t, results, 2)

	expectedResults := []valueobject.LTP{
		{Pair: "BTC/USD", Amount: BTCUSDExp},
		{Pair: "BTC/EUR", Amount: BTCEURExp},
	}
	assert.ElementsMatch(t, expectedResults, results)

	mockClient.AssertExpectations(t)
}

func TestGetPricesForPairs_Error(t *testing.T) {
	mockClient := new(MockTickerClient)
	service := NewService(mockClient)
	ctx := context.Background()
	BTCEURExp := "42.000"
	e := errors.New("some error")
	pairs := []*valueobject.Pair{
		{BaseCurrency: "BTC", QuoteCurrency: "USD"},
		{BaseCurrency: "BTC", QuoteCurrency: "EUR"},
	}
	krakenStandardBTCEUR, err := pairs[1].ToStringKraken()
	assert.NoError(t, err)

	mockClient.On("GetTickerInfo", ctx, "BTCUSD").Return(
		&dto.TickerResponse{},
		e,
	)
	mockClient.On("GetTickerInfo", ctx, "BTCEUR").Return(&dto.TickerResponse{
		Error:  nil,
		Result: map[string]dto.Pair{krakenStandardBTCEUR: {LastTradeClosed: [2]string{BTCEURExp}}}},
		nil,
	)
	results, err := service.GetPricesForPairs(ctx, pairs)
	assert.Error(t, err)
	assert.Nil(t, results)
	assert.Equal(t, e, err)
	mockClient.AssertExpectations(t)
}
