package main

import "fmt"

func main() {
	factory := GetFactory("Turkey")

	sportsCarTr := factory.GetSportsCar()

	factory2 := GetFactory("Italy")
	famCarFer := factory2.GetFamilyCar()

	fmt.Printf("sports tog: %v", sportsCarTr)
	fmt.Printf("family ferrari: %v", famCarFer)
}
