package entity_test

import (
	"testing"

	"github.com/devfullcycle/pfa-go/internal/order/entity"
	"github.com/stretchr/testify/assert"
)

func TestGivenAndEmpty_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	// Given When
	order := entity.Order{}

	// Then
	assert.Error(t, order.IsValid(), "The order is not valid")
}

func TestGivenAndEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	// Given When
	order := entity.Order{
		ID: "123",
	}

	// Then
	assert.Error(t, order.IsValid(), "The order is not valid")
}

func TestGivenAndEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	// Given When
	order := entity.Order{
		ID:    "123",
		Price: 10,
	}

	// Then
	assert.Error(t, order.IsValid(), "The order is not valid")
}

func TestGivenValidParams_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	// Given When
	order, err := entity.NewOrder("123", 10, 1)

	// Then
	assert.NoError(t, err, "The order is valid")
	assert.Equal(t, "123", order.ID, "The order ID is invalid")
	assert.Equal(t, 10.0, order.Price, "The order price is invalid")
	assert.Equal(t, 1.0, order.Tax, "The order tax is invalid")
}

func TestGivenValidParams_WhenCreateANewOrder_ThenShouldCalculateFinlPrice(t *testing.T) {
	// Given When
	order, err := entity.NewOrder("123", 10, 2)

	// Then
	assert.NoError(t, err, "The order is valid")
	err = order.CalculateFinalPrice()
	assert.NoError(t, err, "The order is valid")
	assert.Equal(t, 12.0, order.FinalPrice, "The order final price is invalid")
}
