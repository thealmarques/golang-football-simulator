package schemas

import "github.com/lib/pq"

type CreateTeamInput struct {
	Name    string         `json:"name" binding:"required"`
	Players pq.StringArray `json:"players" binding:"required"`
}
