// Package entities holds all the entities that are shared across all subdomains
// An entity is a struct with a unique identifier to reference it, which has states that can change.
// And by this, the names of the variables are in Upper Case.

package entity

import (
	"github.com/google/uuid"
)

// Person is a entity that represents a person in all Domains
type Person struct {
	// ID is the identifier of the Entity, the ID is shared for all sub domains
	ID uuid.UUID
	// Name is the name of the person
	Name string
	// Age is the age of the person
	Age int
}
