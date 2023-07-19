// Package memory is a in-memory implementation of the Product repository
package memory

import (
	"sync"

	"github.com/google/uuid"

	"github.com/pwkm/ddd-go/domain/product"
)

// MemoryProductRepository fulfills the ProductRepository interface
type MemoryProductRepository struct {
	products map[uuid.UUID]product.Product
	sync.Mutex
}

// ------------------- GETALL ---------------------
// GetAll returns all products as a slice
// Yes, it never returns an error, but
// A database implementation could return an error for instance
func (mpr *MemoryProductRepository) GetAll() ([]product.Product, error) {
	var products []product.Product

	for _, product := range mpr.products {
		products = append(products, product)
	}
	return products, nil
}

// -------------------FIND product by ID -------------
func (mpr *MemoryProductRepository) GetByID(id uuid.UUID) (product.Product, error) {
	if product, ok := mpr.products[id]; ok {
		return product, nil
	}
	return product.Product{}, product.ErrProductNotFound
}

// -------- ADD a product to the repository  ---------
func (mpr *MemoryProductRepository) Add(newprod product.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[newprod.GetID()]; ok {
		return product.ErrProductAlreadyExist
	}

	mpr.products[newprod.GetID()] = newprod
	return nil
}

// ------------------ UPDATE A PRODUCT --------------
func (mpr *MemoryProductRepository) Update(upprod product.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[upprod.GetID()]; !ok {
		return product.ErrProductNotFound

	}

	mpr.products[upprod.GetID()] = upprod
	return nil
}

// ------------------ DELETE A PRODUCT --------------------------

func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return product.ErrProductNotFound
	}
	delete(mpr.products, id)
	return nil
}

// -------------- FACTORY NEW -------------------
// New is a factory function to generate a new repository of customers
func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]product.Product),
	}
}
