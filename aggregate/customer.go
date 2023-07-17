package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/pwkm/ddd-go/entity"
	"github.com/pwkm/ddd-go/valueobject"
)

type Customer struct {
	// Person is the root entity of a customer
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction
}

var (
	ErrInvalidPerson = errors.New("a customer has to have an valid person")
)

// Factory Pattern NewCustomer on top of the Customer aggregate
// ----------------------------------------------------------

func NewCustomer(name string) (Customer, error) {
	// validate that the name is not empty
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		ID:   uuid.New(),
		Name: name,
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}
