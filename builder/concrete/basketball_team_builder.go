package concrete

import "builder/model"

type BasketballTeamBuilder struct {
	teamMemberCount int
	goalKeeper      bool
	coach           string
}

func NewBasketballTeamBuilder() *BasketballTeamBuilder {
	return &BasketballTeamBuilder{}
}

func (f *BasketballTeamBuilder) SetTeamMemberCount() {
	f.teamMemberCount = 5
}
func (f *BasketballTeamBuilder) SetGoalKeeper() {
	f.goalKeeper = false
}
func (f *BasketballTeamBuilder) SetCoach() {
	f.coach = "obradovic"
}
func (f *BasketballTeamBuilder) GetTeam() model.Team {
	return model.Team{
		TeamMemberCount: f.teamMemberCount,
		GoalKeeper:      f.goalKeeper,
		Coach:           f.coach,
	}
}
