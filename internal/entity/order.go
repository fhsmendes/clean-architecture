package entity

import (
	"errors"
	"fmt"
)

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	err := order.IsValid()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) IsValid() error {
	if o.ID == "" {
		return errors.New("invalid id")
	}
	if o.Price <= 0 {
		return errors.New("invalid price")
	}
	if o.Tax <= 0 {
		return errors.New("invalid tax")
	}
	return nil
}

func (o *Order) CalculateFinalPrice() error {
	fmt.Println("Calculating final price...")

	o.FinalPrice = o.Price + o.Tax
	err := o.IsValid()
	if err != nil {
		fmt.Println("Error calculating final price:", err)
		return err
	}

	fmt.Println("Final price calculated successfully:", o.FinalPrice)
	return nil
}
