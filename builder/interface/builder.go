package _interface

import (
	"builder/concrete"
	"builder/model"
)

type IBuilder interface {
	SetTeamMemberCount()
	SetGoalKeeper()
	SetCoach()
	GetTeam() model.Team
}

func getBuilder(builderType string) IBuilder {
	if builderType == "basket" {
		return concrete.NewBasketballTeamBuilder()
	}

	if builderType == "football" {
		return concrete.NewFootballTeamBuilder()
	}
	return nil
}
