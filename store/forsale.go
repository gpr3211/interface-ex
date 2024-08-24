package store

// STEP THREE

// We create our first interface. This will pick off only the items that have an already set price field and therefore
type ItemForSale interface {
	Price(taxRate float64) float64
}
