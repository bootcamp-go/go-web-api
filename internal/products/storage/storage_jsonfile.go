package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// NewStorageJsonFile creates a new StorageJsonFile
func NewStorageJsonFile(file string) *StorageJsonFile {
	return &StorageJsonFile{file: file}
}

// StorageJsonFile is the struct that defines a storage for products in a JSON file
type StorageJsonFile struct {
	// file is the path to the JSON file
	file string
}

type ProductJson struct {
	Id		  	int			`json:"id"`
	Name	  	string		`json:"name"`
	Quantity  	int			`json:"quantity"`
	CodeValue 	string		`json:"code_value"`
	IsPublished bool		`json:"is_published"`
	Expiration  time.Time	`json:"expiration"`
	Price 		float64		`json:"price"`
}

// Read reads all products from the storage
func (s *StorageJsonFile) Read() (p []*Product, err error) {
	// open file
	f, err := os.Open(s.file)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageInternal, err)
		return
	}
	defer f.Close()

	// decode JSON
	var productsJson []*ProductJson
	err = json.NewDecoder(f).Decode(&productsJson)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageInternal, err)
		return
	}

	// serialization
	p = make([]*Product, len(productsJson))
	for i, productJson := range productsJson {
		p[i] = &Product{
			Id: productJson.Id,
			Name: productJson.Name,
			Quantity: productJson.Quantity,
			CodeValue: productJson.CodeValue,
			IsPublished: productJson.IsPublished,
			Expiration: productJson.Expiration,
			Price: productJson.Price,
		}
	}

	return
}

// Write writes all products to the storage
func (s *StorageJsonFile) Write(p []*Product) (err error) {
	// open file
	f, err := os.Create(s.file)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageInternal, err)
		return
	}
	defer f.Close()

	// deserialization
	productsJson := make([]*ProductJson, len(p))
	for i, product := range p {
		productsJson[i] = &ProductJson{
			Id: product.Id,
			Name: product.Name,
			Quantity: product.Quantity,
			CodeValue: product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration: product.Expiration,
			Price: product.Price,
		}
	}

	// encode JSON
	err = json.NewEncoder(f).Encode(productsJson)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageInternal, err)
		return
	}

	return
}