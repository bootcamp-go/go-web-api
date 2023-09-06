package handlers

import (
	"app/internal/products/storage"
	"app/pkg/web/response"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

// NewControllerProducts creates a new ControllerProducts
func NewControllerProducts(st storage.StorageProduct) *ControllerProducts {
	return &ControllerProducts{st: st}
}

// ControllerProducts is the struct that returns the handlers for products
type ControllerProducts struct {
	st storage.StorageProduct
}

// GetAll returns all products
type ProductHandlerGetAll struct {
	Id		  	int			`json:"id"`
	Name	  	string		`json:"name"`
	Quantity  	int			`json:"quantity"`
	CodeValue 	string		`json:"code_value"`
	IsPublished bool		`json:"is_published"`
	Expiration  time.Time	`json:"expiration"`
	Price 		float64		`json:"price"`
}
type ResponseGetAllProducts struct {
	Message string					`json:"message"`
	Data    []*ProductHandlerGetAll `json:"data"`
	Error	bool					`json:"error"`
}
func (c *ControllerProducts) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		ps, err := c.st.GetAll()
		if err != nil {
			code := http.StatusInternalServerError
			body := &ResponseGetAllProducts{Message: "Internal Server Error", Error: true}

			response.JSON(w, code, body)
			return
		}

		// response
		code := http.StatusOK
		body := &ResponseGetAllProducts{Message: "OK", Data: make([]*ProductHandlerGetAll, len(ps)), Error: false}
		for i, p := range ps {
			body.Data[i] = &ProductHandlerGetAll{
				Id:          p.Id,
				Name:        p.Name,
				Quantity:    p.Quantity,
				CodeValue:   p.CodeValue,
				IsPublished: p.IsPublished,
				Expiration:  p.Expiration,
				Price:       p.Price,
			}
		}

		response.JSON(w, code, body)
	}
}

// GetByID returns a product by its id
type ProductHandlerGetByID struct {
	Id		  	int			`json:"id"`
	Name	  	string		`json:"name"`
	Quantity  	int			`json:"quantity"`
	CodeValue 	string		`json:"code_value"`
	IsPublished bool		`json:"is_published"`
	Expiration  time.Time	`json:"expiration"`
	Price 		float64		`json:"price"`
}
type ResponseGetProductByID struct {
	Message string					`json:"message"`
	Data    *ProductHandlerGetByID  `json:"data"`
	Error	bool					`json:"error"`
}
func (c *ControllerProducts) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			code := http.StatusBadRequest
			body := &ResponseGetProductByID{Message: "Bad Request", Error: true}

			response.JSON(w, code, body)
			return
		}

		// process
		p, err := c.st.GetByID(id)
		if err != nil {
			var code int; var body *ResponseGetProductByID
			switch {
			case errors.Is(err, storage.ErrStorageProductNotFound):
				code = http.StatusNotFound
				body = &ResponseGetProductByID{Message: "Not Found", Error: true}
			default:
				code = http.StatusInternalServerError
				body = &ResponseGetProductByID{Message: "Internal Server Error", Error: true}
			}

			response.JSON(w, code, body)
			return
		}

		// response
		code := http.StatusOK
		body := &ResponseGetProductByID{Message: "OK", Data: &ProductHandlerGetByID{
			Id:          p.Id,
			Name:        p.Name,
			Quantity:    p.Quantity,
			CodeValue:   p.CodeValue,
			IsPublished: p.IsPublished,
			Expiration:  p.Expiration,
			Price:       p.Price,
		}, Error: false}

		response.JSON(w, code, body)
	}
}

// Search returns all products that match the given search criteria
type ProductQueryHandlerSearch struct {
	Id		  	int			`json:"id"`
}
type ProductHandlerSearch struct {
	Id		  	int			`json:"id"`
	Name	  	string		`json:"name"`
	Quantity  	int			`json:"quantity"`
	CodeValue 	string		`json:"code_value"`
	IsPublished bool		`json:"is_published"`
	Expiration  time.Time	`json:"expiration"`
	Price 		float64		`json:"price"`
}
type ResponseSearchProducts struct {
	Message string					`json:"message"`
	Data    []*ProductHandlerSearch `json:"data"`
	Error	bool					`json:"error"`
}
func (c *ControllerProducts) Search() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var query ProductQueryHandlerSearch
		query.Id, _ = strconv.Atoi(r.URL.Query().Get("id"))

		// process
		q := &storage.Query{Id: query.Id}
		ps, err := c.st.Search(q)
		if err != nil {
			code := http.StatusInternalServerError
			body := &ResponseSearchProducts{Message: "Internal Server Error", Error: true}

			response.JSON(w, code, body)
			return
		}

		// response
		code := http.StatusOK
		body := &ResponseSearchProducts{Message: "OK", Data: make([]*ProductHandlerSearch, len(ps)), Error: false}
		for i, p := range ps {
			body.Data[i] = &ProductHandlerSearch{
				Id:          p.Id,
				Name:        p.Name,
				Quantity:    p.Quantity,
				CodeValue:   p.CodeValue,
				IsPublished: p.IsPublished,
				Expiration:  p.Expiration,
				Price:       p.Price,
			}
		}

		response.JSON(w, code, body)
	}
}