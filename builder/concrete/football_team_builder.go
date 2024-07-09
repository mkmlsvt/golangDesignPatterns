package concrete

import "builder/model"

type FootballTeamBuilder struct {
	teamMemberCount int
	goalKeeper      bool
	coach           string
}

func NewFootballTeamBuilder() *FootballTeamBuilder {
	return &FootballTeamBuilder{}
}

func (f *FootballTeamBuilder) SetTeamMemberCount() {
	f.teamMemberCount = 11
}
func (f *FootballTeamBuilder) SetGoalKeeper() {
	f.goalKeeper = true
}
func (f *FootballTeamBuilder) SetCoach() {
	f.coach = "mourinho"
}
func (f *FootballTeamBuilder) GetTeam() model.Team {
	return model.Team{
		TeamMemberCount: f.teamMemberCount,
		GoalKeeper:      f.goalKeeper,
		Coach:           f.coach,
	}
}
