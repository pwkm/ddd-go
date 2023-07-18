package product

import "github.com/google/uuid"

type ProductRepository interface {
	Get(uuid.UUID) aggregate.
}