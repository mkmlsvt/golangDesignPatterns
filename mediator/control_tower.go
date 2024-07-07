package main

type ControlTower struct {
	airplanesQueue []*AirPlane
	isAirportFree  bool
}

func NewControlTower() Mediator {
	return &ControlTower{isAirportFree: true}
}

func (c *ControlTower) CanAirplaneLand(airplane *AirPlane) bool {
	if c.isAirportFree {
		c.isAirportFree = false
		return true
	}
	c.airplanesQueue = append(c.airplanesQueue, airplane)
	return false
}

func (c *ControlTower) NotifyControlTowerLeaving() {
	if !c.isAirportFree {
		c.isAirportFree = true
	}
	if len(c.airplanesQueue) > 0 {
		firstPlaneInQueue := c.airplanesQueue[0]
		c.airplanesQueue = c.airplanesQueue[1:]
		firstPlaneInQueue.LandTheAirport()
	}
}
