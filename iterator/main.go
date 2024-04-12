package main

import (
	"bufio"
	"fmt"
	"iterator/concretes"
	"iterator/model"
	"os"
	"strings"
)

func main() {
	city1 := &model.City{
		Name: "istanbul",
		Code: 34,
	}
	city2 := &model.City{
		Name: "canakkale",
		Code: 17,
	}
	city3 := &model.City{
		Name: "ankara",
		Code: 6,
	}
	city4 := &model.City{
		Name: "antalya",
		Code: 7,
	}
	city5 := &model.City{
		Name: "rize",
		Code: 53,
	}
	cityCollection := concretes.CityCollection{Cities: []*model.City{city1, city2, city3, city4, city5}}

	cityIterator := cityCollection.CreateIterator()
	fmt.Println("-------- guess the next city --------")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("enter the text")
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, "\n")

		city := cityIterator.GetNext()
		if city == nil {
			fmt.Println("YOU LOST GAME..")
			break
		}
		if text == city.Name {
			fmt.Println("You guessed the city correctly. YOU WİN")
			fmt.Printf("your point is ==> %d", 6-cityIterator.GetCursor())
			break
		} else {
			fmt.Println("You wrong try again")
		}
	}

}
