package main

func main() {

	controlTower := NewControlTower()

	airplane1 := NewAirPlane(controlTower, "1.ucak")
	airplane2 := NewAirPlane(controlTower, "2.ucak")
	airplane3 := NewAirPlane(controlTower, "3.ucak")

	airplane1.LandTheAirport()
	airplane2.LandTheAirport()
	airplane3.LandTheAirport()
	airplane1.LeaveTheAirport()
	airplane2.LeaveTheAirport()
}
