package store

// STEP TWO

type Boat struct {
	*Product
	Capacity  int
	Motorized bool
}

// Create new boat using the NewpProduct method to declare the *Product value in boat and then insert rest of values
func NewBoat(name string, price float64, capacity int, motorized bool) *Boat {
	return &Boat{
		NewProduct(name, "Watersports", price),
		capacity,
		motorized,
	}
}
