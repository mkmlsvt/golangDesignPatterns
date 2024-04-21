package main

import (
	"fmt"
	"math/rand"
	"strategy/concreates"
)

func main() {
	randomNumber := rand.Intn(100)
	twoByTwoIncreaser := concreates.IncreaseTwoByTwo{}
	oneByOneIncreaser := concreates.IncreaseOneByOne{}
	increaser := NumberIncreaser{
		Number: &randomNumber,
	}
	if randomNumber%2 == 0 {
		increaser.SetIncreaseApproach(twoByTwoIncreaser)
	} else {
		increaser.SetIncreaseApproach(oneByOneIncreaser)
	}
	fmt.Printf("the number is: %d \n", randomNumber)
	increaser.IncreaseNumber()
	increaser.IncreaseNumber()
	fmt.Printf("the new number is: %d \n", randomNumber)
}
