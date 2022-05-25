package main

import (
	"fmt"
	"task8_8_2/auto"
)

// Вывод на экран характеристик автомобилей,
// с указанием габаритных размеров в заданных единицах.
func printCharacterics(auto auto.Auto, units auto.UnitType) {
	length := auto.Dimensions().Length().Get(units)
	width := auto.Dimensions().Width().Get(units)
	height := auto.Dimensions().Height().Get(units)

	fmt.Println("Brand:\t\t", auto.Brand())
	fmt.Println("Model:\t\t", auto.Model())
	fmt.Println("Length:\t\t", length, units)
	fmt.Println("Width:\t\t", width, units)
	fmt.Println("Height:\t\t", height, units)
	fmt.Println("MaxSpeed:\t", auto.MaxSpeed())
	fmt.Println("EnginePower:\t", auto.EnginePower())
	fmt.Println()
}

func main() {
	bmw := auto.NewAuto("BMW", "X5 30i 3.0 AT", 485.4, 193.3, 173.9, auto.CM, 210, 272)
	mercedes := auto.NewAuto("Mersedes-Benz", "A 35 AMG DCT 4MATIC", 454.9, 179.6, 144.6, auto.CM, 250, 306)
	dodge := auto.NewAuto("Dodge", "Dodge Viper (8.3 MT SRT-10 Roadster)", 445.8, 167.5, 120.9, auto.CM, 314, 507)

	//Информация о характеристиках автомобилей, с габаритными размерами в сантиметрах.
	printCharacterics(bmw, auto.CM)
	printCharacterics(mercedes, auto.CM)
	printCharacterics(dodge, auto.CM)

	//Информация о характеристиках автомобилей, с габаритными размерами в дюймах.
	printCharacterics(bmw, auto.Inch)
	printCharacterics(mercedes, auto.Inch)
	printCharacterics(dodge, auto.Inch)
}
