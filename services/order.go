// Package services holds all the services that connects repositories into a business flow
package services

import (
	"github.com/google/uuid"
	"github.com/pwkm/ddd-go/domain/customer"
	"github.com/pwkm/ddd-go/domain/customer/memory"
)

// OrderConfiguration is an alias for a function that will take in a pointer to an
// OrderService and modify it
type OrderConfiguration func(os *OrderService) error

// OrderService is a implementation of the orderService
type OrderService struct {
	customers customer.CustomerRepository
}

// CreateOrder will chaintogether all repositories to create a order for a customer
func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) error {
	// Get the customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return err
	}

	// Get each Product, Ouchie, We need a ProductRepository

	return nil
}

// NewOrderService takes a variable amount of OrderConfiguration functions and returns a new OrderService
// Each OrderConfiguration will be called in the order they are passed in

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	// Create the OrderService
	os := &OrderService{}

	//Apply all configurations passed in
	for _, cfg := range cfgs {
		// pass the services into the configuration function
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

// WithCustomerRepository applies a givin customer repository to the Orderservice
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// return a function that matches the OrderConfiguration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

// WithMemoryCustomerRepository applies a memory customer repository to the OrderService
func WithMemoryCustomerRepository() OrderConfiguration {
	// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
	cr := memory.New()
	return WithCustomerRepository(cr)
}
