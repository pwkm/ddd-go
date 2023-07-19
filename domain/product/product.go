package product

import (
	"errors"

	"github.com/google/uuid"
	"github.com/pwkm/tavern"
)

var (
	ErrMissingValues = errors.New("missing values")
)

// -----------------------------------------------------
//
//	AGGREGATE PRODUCT
//
// -----------------------------------------------------
// Product is a aggregate that combines item with a price and quantity
type Product struct {
	item     *tavern.Item
	price    float64
	quantity int
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() *tavern.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}

// ------ Factory NewProduct ------------------------
// NewProduct will create a new product
// will return error if name of description is empty
func NewProduct(name string, description string, price float64) (Product, error) {

	if name == "" || description == "" {
		return Product{}, ErrMissingValues
	}

	return Product{
		item: &tavern.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 0,
	}, nil
}
