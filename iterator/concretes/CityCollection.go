package concretes

import (
	"iterator/interfaces"
	"iterator/model"
)

type CityCollection struct {
	Cities []*model.City
}

func (cc CityCollection) CreateIterator() interfaces.Iterator {
	return &CityIterator{
		Cursor: 0,
		Cities: cc.Cities,
	}
}
