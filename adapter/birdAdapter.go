package main

import "adapter/model"

type BirdAdapter struct {
	Eagle model.Eagle
}

func (l BirdAdapter) RunOnGround() int {
	return l.Eagle.FlyOnAir()
}
