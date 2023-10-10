package main

import (
	"fmt"
)

func fuelGauge(fuel int) {
	//prints below string when called
	//to call function with the parameter use "fuelGauge(22)"
	fmt.Printf("You have %v gallons of fuel left!", fuel)
}

func calculateFuel(planet string) int {
	var fuel int
	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	default:
		fmt.Println("Enter a real planet to visit")
	}
	return (fuel)
}

func greetPlanet(planet string) {
	fmt.Printf("Welcome to %v\n", planet)
}

func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {

	var fuelRemaining, fuelCost int
	fuelRemaining = fuel

	fuelCost = calculateFuel(planet)

	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}
	return fuelRemaining

}

func main() {

	var fuel int
	fuel = 1000000
	planetChoice := "Venus"

	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)
}
