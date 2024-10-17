package entity

import "errors"

type OrderRepositoryInterface interface {
	Save(order Order) error
	GetTotal() (int, error)
}

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price, tax float64) (Order, error) {
	order := Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	if err := order.IsValid(); err != nil {
		return Order{}, err
	}
	order.FinalPrice = order.Price + order.Tax
	return order, nil
}

func (o *Order) CalculateFinalPrice() error {
	if err := o.IsValid(); err != nil {
		return err
	}
	o.FinalPrice = o.Price + o.Tax
	return nil
}

func (order Order) IsValid() error {
	if order.ID == "" {
		return errors.New("the order ID is required")
	}
	if order.Price == 0 {
		return errors.New("the order price is required") //error strings should not be capitalized (ST1005)go-staticcheck
	}
	if order.Tax == 0 {
		return errors.New("the order tax is required") //error strings should not be capitalized (ST1005)go-staticcheck
	}
	return nil
}
