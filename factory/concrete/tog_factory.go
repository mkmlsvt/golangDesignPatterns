package concrete

import "factory/model"

type TogFactory struct {
}

func NewTogFactory() *TogFactory {
	return &TogFactory{}
}

func (t TogFactory) GetSportsCar() model.Vehicle {

	return model.Vehicle{
		Power:      20000,
		Type:       "wolf",
		DoorsCount: 2,
	}
}

func (t TogFactory) GetFamilyCar() model.Vehicle {
	return model.Vehicle{
		Power:      1500,
		Type:       "classic",
		DoorsCount: 4,
	}
}

func (t TogFactory) GetTruck() model.Vehicle {
	return model.Vehicle{
		Power:      3000,
		Type:       "uzay gemisi",
		DoorsCount: 2,
	}
}
