package schemas

type CreateMatchInput struct {
	HomeId uint `json:"home_id" binding:"required"`
	AwayId uint `json:"away_id" binding:"required"`
}
