package schemas

type CreatePlayerInput struct {
	Name        string `json:"name" binding:"required"`
	Age         int    `json:"age" binding:"required"`
	GoalKeeping int    `json:"goalkeeping" binding:"required"`
	Attack      int    `json:"attack" binding:"required"`
	Defense     int    `json:"defense" binding:"required"`
	Creativity  int    `json:"creativity" binding:"required"`
}
