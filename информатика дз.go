package main

import "fmt"

type Furniture struct {
	name     string
	material string
	color    string
	price    float64
	quantity int
}

type Appliance struct {
	name       string
	brand      string
	powerUsage float64
	price      float64
	quantity   int
}

type FamilyMember struct {
	name   string
	age    int
	gender string
	role   string
}

type House struct {
	furniture     []Furniture
	appliances    []Appliance
	familyMembers []FamilyMember
}

func main() {
	house := House{
		furniture: []Furniture{
			Furniture{name: "Sofa", material: "Leather", color: "Black", price: 799.99, quantity: 1},
			Furniture{name: "Bed", material: "Wood", color: "Brown", price: 1299.99, quantity: 2},
			Furniture{name: "Table", material: "Glass", color: "Clear", price: 299.99, quantity: 1},
			Furniture{name: "Chairs", material: "Metal", color: "Silver", price: 149.99, quantity: 6},
			Furniture{name: "Shelves", material: "Particle board", color: "White", price: 99.99, quantity: 3},
		},
		appliances: []Appliance{
			Appliance{name: "Refrigerator", brand: "Samsung", powerUsage: 60, price: 1099.99, quantity: 1},
			Appliance{name: "Washing Machine", brand: "LG", powerUsage: 120, price: 699.99, quantity: 1},
			Appliance{name: "Microwave", brand: "Panasonic", powerUsage: 100, price: 199.99, quantity: 1},
			Appliance{name: "TV", brand: "Sony", powerUsage: 75, price: 799.99, quantity: 1},
			Appliance{name: "Vacuum Cleaner", brand: "Dyson", powerUsage: 1400, price: 299.99, quantity: 1},
		},
		familyMembers: []FamilyMember{
			FamilyMember{name: "Толя", age: 35, gender: "Male", role: "Father"},
			FamilyMember{name: "Валя", age: 32, gender: "Female", role: "Mother"},
			FamilyMember{name: "Ваня", age: 8, gender: "Male", role: "Son"},
			FamilyMember{name: "Ира", age: 5, gender: "Female", role: "Daughter"},
			FamilyMember{name: "Таня", age: 65, gender: "Female", role: "Grandmother"},
		},
	}

	fmt.Println(house)
}
