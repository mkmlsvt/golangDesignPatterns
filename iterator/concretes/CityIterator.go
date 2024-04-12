package concretes

import "iterator/model"

type CityIterator struct {
	Cursor int
	Cities []*model.City
}

func (ci *CityIterator) GetNext() *model.City {
	if ci.HasNext() {
		city := ci.Cities[ci.Cursor]
		ci.Cursor++
		return city
	} else {
		return nil
	}
}

func (ci *CityIterator) HasNext() bool {
	if ci.Cursor == len(ci.Cities) {
		return false
	}
	return true
}

func (ci *CityIterator) GetCursor() int {
	return ci.Cursor
}
