package main

import (
	_interface "builder/interface"
	"builder/model"
)

type Director struct {
	builder _interface.IBuilder
}

func NewDirector(b _interface.IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b _interface.IBuilder) {
	d.builder = b
}

func (d *Director) BuildTeam() model.Team {
	d.builder.SetTeamMemberCount()
	d.builder.SetGoalKeeper()
	d.builder.SetCoach()
	return d.builder.GetTeam()
}
