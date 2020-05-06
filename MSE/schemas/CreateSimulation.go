package schemas

import "MSE/interfaces"

type CreateSimulation struct {
	MatchId  uint            `json:"match_id" binding:"required"`
	HomeTeam interfaces.Team `json:"home_team" binding:"required"`
	AwayTeam interfaces.Team `json:"away_team" binding:"required"`
}
