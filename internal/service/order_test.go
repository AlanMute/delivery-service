package service

import (
	"errors"
	"testing"

	"github.com/KrizzMU/delivery-service/internal/core"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) Add(ord core.Order) error {
	args := m.Called(ord)
	return args.Error(0)
}

func (m *MockOrderRepository) GetAll() []core.Order {
	args := m.Called()
	return args.Get(0).([]core.Order)
}

type MockCache struct {
	mock.Mock
}

func (m *MockCache) Add(key string, data interface{}) {
	m.Called(key, data)
}

func (m *MockCache) Get(key string) (interface{}, bool) {
	args := m.Called(key)
	return args.Get(0), args.Bool(1)
}

func TestOrder_Add(t *testing.T) {
	tests := []struct {
		name             string
		inputOrder       core.Order
		mockRepoError    error
		expectedError    error
		expectedCacheKey string
	}{
		{
			name: "OK",
			inputOrder: core.Order{
				OrderUID:    "asd12as",
				TrackNumber: "testtrack",
				Locale:      "EN",
			},
			mockRepoError:    nil,
			expectedError:    nil,
			expectedCacheKey: "asd12as",
		},
		{
			name: "Error",
			inputOrder: core.Order{
				OrderUID:    "asd12as",
				TrackNumber: "testtrack",
				Locale:      "EN",
			},
			mockRepoError:    errors.New("Something wrong"),
			expectedError:    errors.New("Something wrong"),
			expectedCacheKey: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			orderRepo := new(MockOrderRepository)
			orderCache := new(MockCache)

			orderService := NewOrderService(orderRepo, orderCache)

			orderRepo.On("Add", test.inputOrder).Return(test.mockRepoError)

			if test.mockRepoError == nil {
				orderCache.On("Add", test.expectedCacheKey, test.inputOrder)
			}

			err := orderService.Create(test.inputOrder)

			assert.Equal(t, test.expectedError, err)

			orderRepo.AssertExpectations(t)
			orderCache.AssertExpectations(t)

		})
	}
}

func TestOrder_Get(t *testing.T) {
	tests := []struct {
		name          string
		inputKey      string
		cacheVal      interface{}
		cacheFound    bool
		expectedOrder core.Order
		expectedError error
	}{
		{
			name:          "OK",
			inputKey:      "testID",
			cacheVal:      core.Order{OrderUID: "testUID"},
			cacheFound:    true,
			expectedOrder: core.Order{OrderUID: "testUID"},
			expectedError: nil,
		},
		{
			name:          "Error Not Found",
			inputKey:      "testID",
			cacheVal:      core.Order{},
			cacheFound:    false,
			expectedOrder: core.Order{},
			expectedError: errors.New("Order with id: testID does not exist!"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			orderCache := new(MockCache)
			orderService := &OrderService{c: orderCache}

			orderCache.On("Get", test.inputKey).Return(test.cacheVal, test.cacheFound)

			order, err := orderService.Get(test.inputKey)

			assert.Equal(t, err, test.expectedError)
			assert.Equal(t, test.expectedOrder, order)

			orderCache.AssertExpectations(t)

		})
	}
}
