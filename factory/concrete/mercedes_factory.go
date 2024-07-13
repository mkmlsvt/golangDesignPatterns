package concrete

import "factory/model"

type MercedesFactory struct {
}

func NewMercedesFactory() *MercedesFactory {
	return &MercedesFactory{}
}
func (t MercedesFactory) GetSportsCar() model.Vehicle {
	return model.Vehicle{
		Power:      500,
		Type:       "slsamg v2",
		DoorsCount: 2,
	}
}

func (t MercedesFactory) GetFamilyCar() model.Vehicle {

	return model.Vehicle{
		Power:      300,
		Type:       "benz",
		DoorsCount: 4,
	}
}

func (t MercedesFactory) GetTruck() model.Vehicle {

	return model.Vehicle{
		Power:      1000,
		Type:       "merco truck",
		DoorsCount: 2,
	}
}
