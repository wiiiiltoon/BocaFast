package models

import "github.com/google/uuid"

/* Estructura de producto
* ID -> Identificador interno
* Name -> Nombre
* Description -> Descripcion
* Category -> Categoria (bocadillo, bebida...)
* Price -> Precio
* Stock -> Disponibilidad
 */
type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float32 `json:"price"`
	Stock       bool    `json:"stock"`
	Preparation int     `json:"preparation"`
}

func NewProduct(name string, description string, category string, price float32, stock bool, preparation int) *Product {
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Category:    category,
		Price:       price,
		Stock:       stock,
		Preparation: preparation,
	}
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetPrice() float32 {
	return p.Price
}

func (p *Product) GetDescription() string {
	return p.Description
}

func (p *Product) IsAvailable() bool {
	return p.Stock
}

func (p *Product) GetPreparationTime() int {
	return p.Preparation
}
