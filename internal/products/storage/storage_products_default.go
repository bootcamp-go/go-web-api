package storage

import "fmt"

// NewStorageProductDefault creates a new StorageProductDefault
func NewStorageProductDefault(st Storage) *StorageProductDefault {
	return &StorageProductDefault{st: st}
}

// StorageProductDefault is the struct that defines a default storage for products
type StorageProductDefault struct {
	// st is the storage for products
	st Storage
}

// GetAll returns all products from the storage
func (s *StorageProductDefault) GetAll() (p []*Product, err error) {
	p, err = s.st.Read()
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageProductInternal, err)
		return
	}

	return
}

// GetByID returns a product from the storage by its id
func (s *StorageProductDefault) GetByID(id int) (p *Product, err error) {
	// get all products
	products, err := s.st.Read()
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageProductInternal, err)
		return
	}

	// search product by id
	var found bool
	for _, product := range products {
		if product.Id == id {
			p = product
			found = true
			break
		}
	}
	// check if product was found
	if !found {
		err = fmt.Errorf("%w. %v", ErrStorageProductNotFound, id)
		return
	}

	return
}

// Search returns all products from the storage that match the given search criteria
func (s *StorageProductDefault) Search(query *Query) (p []*Product, err error) {
	// get all products
	products, err := s.st.Read()
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageProductInternal, err)
		return
	}

	// search products by query
	for _, product := range products {
		// check if query is set
		if query != nil && query.Id != 0 {
			// check if product matches query
			if product.Id != query.Id {
				continue
			}
		}

		p = append(p, product)
	}

	return
}