package tavern

import (
	"testing"

	"github.com/google/uuid"
	"github.com/pwkm/tavern/domain/customer"
	"github.com/pwkm/tavern/services/order"
)

func Test_Tavern(t *testing.T) {
	// Create OrderService

	products := order.init_products(t)

	os, err := order.NewOrderService(
		// WithMemoryCustomerRepository(),
		order.WithMongoCustomerRepository("mongodb://localhost:2717"),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	cust, err := customer.NewCustomer("Percy")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	// Execute Order
	err = tavern.Order(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}

}
