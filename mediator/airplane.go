package main

import (
	"fmt"
)

type AirPlane struct {
	Name  string
	Tower Mediator
}

func NewAirPlane(mediator Mediator, name string) *AirPlane {
	return &AirPlane{Tower: mediator, Name: name}
}

func (p *AirPlane) LandTheAirport() {
	if p.Tower.CanAirplaneLand(p) {
		fmt.Println("plane is landed on airport.. name: " + p.Name)
		return
	}
	fmt.Printf("plane is cant land on airport name:%s . Waiting for other planes..\n", p.Name)
}

func (p *AirPlane) LeaveTheAirport() {
	fmt.Printf("the plane leave the airport. name:%s\n", p.Name)
	p.Tower.NotifyControlTowerLeaving()
}
