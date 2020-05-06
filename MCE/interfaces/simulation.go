package interfaces

import "github.com/lib/pq"

type Simulation struct {
	ID      uint           `json:"id"`
	MatchId uint           `json:"match_id"`
	Result  string         `json:"result"`
	Events  pq.StringArray `json:"events"`
}
