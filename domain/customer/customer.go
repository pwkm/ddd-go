package customer

import (
	"errors"

	"github.com/google/uuid"

	"github.com/pwkm/tavern"
	"github.com/pwkm/tavern/valueobject"
)

// ---    ERROR Messages    ---
var (
	ErrInvalidPerson = errors.New("a customer has to have an valid person")
)

// -----------------------------------------------------
//
//	AGGREGATE CUSTOMER
//
// -----------------------------------------------------
type Customer struct {
	// Person is the root entity of a customer
	person       *tavern.Person
	products     []*tavern.Item
	transactions []valueobject.Transaction
}

// GetID returns the customers root entity ID
func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

// SetID sets the root ID
func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}
	c.person.ID = id
}

// SetName changes the name of the Customer
func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}
	c.person.Name = name
}

// SetName changes the name of the Customer
func (c *Customer) GetName() string {
	return c.person.Name
}

// ----------------------------------------------------------
//                 Factory NewCustomer
// ----------------------------------------------------------

func NewCustomer(name string) (Customer, error) {
	// validate that the name is not empty
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &tavern.Person{
		ID:   uuid.New(),
		Name: name,
	}

	return Customer{
		person:       person,
		products:     make([]*tavern.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}
