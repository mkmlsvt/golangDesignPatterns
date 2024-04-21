package main

import "strategy/interfaces"

type NumberIncreaser struct {
	Number    *int
	Increaser interfaces.IncreaseApproach
}

func (n *NumberIncreaser) IncreaseNumber() {
	n.Increaser.Increase(n.Number)
}
func (n *NumberIncreaser) SetIncreaseApproach(increaseApproach interfaces.IncreaseApproach) {
	n.Increaser = increaseApproach
}
func (n *NumberIncreaser) GetNumber() *int {
	return n.Number
}
