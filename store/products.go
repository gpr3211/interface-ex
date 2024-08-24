package store

// STEP ONE

type Product struct {
	Name     string
	Category string
	price    float64
}

// Method on the Product struct that calculates the price + tax rate(input)
func (p *Product) Price(taxRate float64) float64 {
	return p.price + (p.price * taxRate)
}

// Product Constructor method
// Inputs name, category(strings), price(float64)
func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}
