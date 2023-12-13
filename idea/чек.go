package main

import (
	"fmt"
)

type House struct {
	Family     []FamilyMember // Список членов семьи
	Appliances []Appliance    // Список бытовых приборов
	Furniture  []Furniture    // Список мебели
}

type FamilyMember struct {
	Name string // Имя члена семьи
	Age  int    // Возраст члена семьи
}

type Appliance struct {
	Name  string // Название бытового прибора
	Size  string // Размер бытового прибора
	Power int    // Мощность бытового прибора
}

type Furniture struct {
	Name  string // Название предмета мебели
	Size  string // Размер предмета мебели
	Color string // Цвет предмета мебели
}

func main() {
	familyMembers := []FamilyMember{
		{Name: "Галя", Age: 35},
		{Name: "Люся", Age: 30},
		{Name: "Толя", Age: 10},
		{Name: "Валя", Age: 5},
		{Name: "Тоня", Age: 65},
	}

	appliances := []Appliance{
		{Name: "Духовка", Size: "Громоздкая", Power: 2000},
		{Name: "Холодильник", Size: "Терпимо большой", Power: 500},
		{Name: "Посудомоечная машина", Size: "Комфортно средняя", Power: 1000},
		{Name: "Телевизор", Size: "Малюсенький", Power: 400},
		{Name: "Микроволновая печь", Size: "Малюсенькая", Power: 800},
	}

	furniture := []Furniture{
		{Name: "Диван", Size: "Здоровенный", Color: "Серый"},
		{Name: "Обеденный стол", Size: "Скромных размеров", Color: "Коричневый"},
		{Name: "Кровать", Size: "Огромная", Color: "Белая"},
		{Name: "Шкаф", Size: "Двухметровый", Color: "Черный"},
		{Name: "Письменный стол", Size: "Небольшой", Color: "Коричневый"},
	}

	house := House{
		Family:     familyMembers,
		Appliances: appliances,
		Furniture:  furniture,
	}

	house.printHouse()
}

func (h House) printHouse() {
	fmt.Println("Family Members:")
	for _, member := range h.Family {
		fmt.Println("Name:", member.Name)
		fmt.Println("Age:", member.Age)
		fmt.Println()
	}

	fmt.Println("Appliances:")
	for _, appliance := range h.Appliances {
		fmt.Println("Name:", appliance.Name)
		fmt.Println("Size:", appliance.Size)
		fmt.Println("Power:", appliance.Power)
		fmt.Println()
	}

	fmt.Println("Furniture:")
	for _, furniture := range h.Furniture {
		fmt.Println("Name:", furniture.Name)
		fmt.Println("Size:", furniture.Size)
		fmt.Println("Color:", furniture.Color)
		fmt.Println()
	}
}
