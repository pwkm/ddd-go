// Package services holds all the services that connects repositories into a business flow
package order

import (
	"context"
	"log"

	"github.com/google/uuid"

	"github.com/pwkm/tavern/domain/customer"
	"github.com/pwkm/tavern/domain/customer/memory"
	"github.com/pwkm/tavern/domain/customer/mongo"
	"github.com/pwkm/tavern/domain/product"
	prodmemory "github.com/pwkm/tavern/domain/product/memory"
)

// OrderConfiguration is an alias for a function that will take in a pointer to an
// OrderService and modify it
type OrderConfiguration func(os *OrderService) error

// OrderService is a implementation of the orderService
type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}

// ----------------------- CREATE ORDER ----------------------------
// CreateOrder will chaintogether all repositories to create a order for a customer
// will return the collected price of all Products
func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	// Get the customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	// Get each Product, Ouchie, We need a ProductRepository
	var products []product.Product
	var price float64
	for _, id := range productIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		price += p.GetPrice()
	}

	// All Products exists in store, now we can create the order
	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))

	return price, nil
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

// ----------------- WithMemoryProductRepository -------------------
// WithMemoryProductRepository adds a in memory product repo and adds all input products
func WithMemoryProductRepository(products []product.Product) OrderConfiguration {
	return func(os *OrderService) error {
		// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
		pr := prodmemory.New()

		// Add Items to repo
		for _, p := range products {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

// ----------------- WithMongoCustomerRepository -------------------
func WithMongoCustomerRepository(connectionString string) OrderConfiguration {

	return func(os *OrderService) error {
		// Create the mongo repo, if we needed parameters, such as connection strings they could be inputted here
		cr, err := mongo.New(context.Background(), connectionString)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}
}

// AddCustomer will add a new customer and return the customerID
func (o *OrderService) AddCustomer(name string) (uuid.UUID, error) {
	c, err := customer.NewCustomer(name)
	if err != nil {
		return uuid.Nil, err
	}
	// Add to Repo
	err = o.customers.Add(c)
	if err != nil {
		return uuid.Nil, err
	}

	return c.GetID(), nil
}
