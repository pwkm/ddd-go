package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/pwkm/ddd-go/entity"
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
	item     *entity.Item
	price    float64
	quantity int
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() *entity.Item {
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
		item: &entity.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 0,
	}, nil
}
