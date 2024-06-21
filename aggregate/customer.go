// Package aggregates holds aggregates that combines many entities into a full object
// An important rule in DDD aggregates is that they should only have one entity act as a root entity.
// What this means is that the reference of the root entity is also used to reference the aggregate.
// For our customer aggregate, this means that the Person ID is the unique identifier.
//
// Notice that all fields in the struct begins with lower case letters, this is a way in Go to make an object inaccessible from outside of
// the package the struct is defined in.
// This is done because an Aggregate should not allow direct access to the data.
// Neither does the struct define any tags for how the data is formatted such as json.
// I set all the entities as pointers, this is because an entity can change state and I want that to reflect across all instances of the runtime that has access to it.
// The value objects are held as nonpointers though since they cannot change state.

package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/pwkm/ddd-go/entity"
	"github.com/pwkm/ddd-go/valueobject"
)

var (
	// ErrInvalidPerson is returned when the person is not valid in the NewCustome factory
	ErrInvalidPerson = errors.New("a customer has to have an valid person")
)

// Customer is a aggregate that combines all entities needed to represent a customer
type Customer struct {
	// person is the root entity of a customer
	// which means the person.ID is the main identifier for this aggregate
	person *entity.Person
	// a customer can hold many products
	products []*entity.Item
	// a customer can perform many transactions
	transactions []valueobject.Transaction
}

// NewCustomer is a factory to create a new Customer aggregate
// he factory pattern is a design pattern that is used to encapsulate complex logic
// in functions that creates the wanted instance, without the caller knowing anything about the implementation details.
// It will validate that the name is not empty
func NewCustomer(name string) (Customer, error) {
	// Validate that the Name is not empty
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	// Create a new person and generate ID
	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}
	// Create a customer object and initialize all the values to avoid nil pointer exceptions
	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}

// GetID returns the customers root entity ID
func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

// SetID sets the root ID
func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

// SetName changes the name of the Customer
func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}

// SetName changes the name of the Customer
func (c *Customer) GetName() string {
	return c.person.Name
}
