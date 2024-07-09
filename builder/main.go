package main

import (
	"builder/concrete"
	"fmt"
)

func main() {
	director := NewDirector(concrete.NewFootballTeamBuilder())

	footballTeam := director.BuildTeam()

	director.SetBuilder(concrete.NewBasketballTeamBuilder())

	basketballTeam := director.BuildTeam()

	fmt.Printf("%v", footballTeam)
	fmt.Printf("%v", basketballTeam)

}
