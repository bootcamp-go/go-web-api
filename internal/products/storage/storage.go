package storage

import (
	"errors"
	"time"
)

// Product is the struct that defines a product
type Product struct {
	Id		  	int
	Name	  	string
	Quantity  	int
	CodeValue 	string
	IsPublished bool
	Expiration  time.Time
	Price 		float64
}

// Storage is an interface that wraps the basic methods for a storage
type Storage interface {
	// Read reads all products from the storage
	Read() (p []*Product, err error)

	// Write writes all products to the storage
	Write(p []*Product) (err error)
}

var (
	// ErrStorageInternal is the error returned when an internal error happens in the storage
	ErrStorageInternal = errors.New("internal storage error")
)


// Query is the struct that defines a query for a product
type Query struct {
	Id		  	int
}

// StorageProduct is the interface that wraps the basic methods for a product storage
type StorageProduct interface {
	// GetAll returns all products from the storage
	GetAll() (p []*Product, err error)

	// GetByID returns a product from the storage by its id
	GetByID(id int) (p *Product, err error)

	// Search returns all products from the storage that match the given search criteria
	Search(query *Query) (p []*Product, err error)
}

var (
	// ErrStorageProductInternal is the error returned when an internal error happens in the storage
	ErrStorageProductInternal = errors.New("internal storage error")

	// ErrStorageProductNotFound is the error returned when a product is not found
	ErrStorageProductNotFound = errors.New("product not found")
)