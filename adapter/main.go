package main

import (
	"adapter/interfaces"
	"adapter/model"
	"fmt"
)

func main() {
	lion := model.Lion{}
	cheetah := model.Cheetah{}
	eagle := model.Eagle{}
	birdadapter := BirdAdapter{Eagle: eagle}
	lionMoveCount := MoveAnimals(lion)
	cheetahMoveCount := MoveAnimals(cheetah)
	eagleMoveCount := MoveAnimals(birdadapter)

	fmt.Printf("lion, cheetah, eagle moves ---->  %d, %d, %d", lionMoveCount, cheetahMoveCount, eagleMoveCount)
}

func MoveAnimals(animal interfaces.MoveAnimal) int {
	return animal.RunOnGround()
}
