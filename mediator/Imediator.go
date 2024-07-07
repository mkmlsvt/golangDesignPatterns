package main

type Mediator interface {
	CanAirplaneLand(airplane *AirPlane) bool
	NotifyControlTowerLeaving()
}
