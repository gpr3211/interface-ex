package main

import "github.com/gpr3211/interface-ex/store"
import "fmt"

func main() {

	//// STEP ONE

	fmt.Printf("------------------\n Base Product type \n \n")
	// Introducue first type

	kayak := store.NewProduct("Kayak", "Watersports", 275)
	lifejacket := &store.Product{Name: "Lifejacket", Category: "Watersports"}

	// NOTICE LIFEJACKET IS inserted directly into the products struct WITHOUT A PRICE
	// KEY DETAIL ON WHY IT DOESNT PRINT LATER WHILE RANGING PRODUCTS

	for _, p := range []*store.Product{kayak, lifejacket} {
		fmt.Println("Name", p.Name, "Category", p.Category, "Price", p.Price(0.2)) // Notice here p.Price is  the method Price() which adds a 20% tax to the price
	}
	// STEP TWO
	// Introduce second type

	fmt.Printf("------------------- \n Boat Type \n")
	boats := []*store.Boat{
		store.NewBoat("Kayak", 275, 1, false),
		store.NewBoat("Canoe", 400, 3, false),
		store.NewBoat("Tender", 650.25, 1, true),
	}

	for _, b := range boats {
		fmt.Println("Name", b.Name, "Capacity", b.Capacity, "Price: ", b.Price(0.2))
	}
	fmt.Printf("--------------\n  Interface \n")

	// STEP THREE

	// Time for Interfaces
	// Having an Interface that uses the common Price method allows us to easily define structures with multiple types of structs
	//	which share common methods

	products := map[string]store.ItemForSale{
		"Kayak": store.NewBoat("Kayak", 279, 1, false),
		"Ball":  store.NewProduct("Hockey Stick", "Hockey", 129.99),
	}

	for key, p := range products {
		fmt.Println("Key", key, "Price :", p.Price(0.2))
	}

	// STEP FOUR
	//
	// In order to acces the Product fields in products we need to separate the different types in our interface
	// That can be done by first getting the type with p.(type) and then using a switch case to differentiate our output
	//

	for key, p := range products {
		switch item := p.(type) {
		case *store.Product:
			fmt.Println("Name: ", item.Name, "Category: ", item.Category, "Price: ", item.Price(0.2))
		case *store.Boat:
			fmt.Println("Name: ", item.Name, "Category: ", item.Category, "Capacity: ", item.Capacity, "Motorized: ", item.Motorized)

		default:
			fmt.Println("Key: ", key, "Price: ", p.Price(0.2))
		}
	}

	// Step 5
	fmt.Printf("---------------\n\nDescribable interface\n")
	for _, p := range products {
		switch item := p.(type) {
		case store.Describable:
			fmt.Println("Name", item.GetName(), "Category", item.GetCategory(), "Price: ", item.Price(0.2))
		}
	}

}
