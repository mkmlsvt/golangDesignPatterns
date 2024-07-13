package concrete

import "factory/model"

type FerrariFactory struct {
}

func NewFerrariFactory() *FerrariFactory {
	return &FerrariFactory{}
}

func (t FerrariFactory) GetSportsCar() model.Vehicle {
	return model.Vehicle{
		Power:      1000,
		Type:       "ferro v.325",
		DoorsCount: 2,
	}
}

func (t FerrariFactory) GetFamilyCar() model.Vehicle {
	return model.Vehicle{
		Power:      900,
		Type:       "ferfamily v.20",
		DoorsCount: 4,
	}
}

func (t FerrariFactory) GetTruck() model.Vehicle {
	return model.Vehicle{
		Power:      2000,
		Type:       "ferTruck",
		DoorsCount: 2,
	}
}
