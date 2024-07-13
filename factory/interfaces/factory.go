package interfaces

import "factory/model"

type Factory interface {
	GetSportsCar() model.Vehicle
	GetFamilyCar() model.Vehicle
	GetTruck() model.Vehicle
}
