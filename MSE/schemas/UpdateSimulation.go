package schemas

import "github.com/lib/pq"

type UpdateSimulation struct {
	ID      uint           `json:"id" binding:"required"`
	MatchId uint           `json:"match_id" binding:"required"`
	Result  string         `json:"result" binding:"required"`
	Events  pq.StringArray `json:"events" binding:"required"`
}
