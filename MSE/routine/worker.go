package routine

import (
	"MSE/interfaces"
	"MSE/models"
	"MSE/requests"
	"MSE/routine/method"
	"MSE/schemas"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func Run(teams schemas.CreateSimulation, simulation models.Simulation, matchId uint) {
	socketClient := NewSocketIO("http://localhost:9090/socketio/")
	home_factor := 10
	home_power := getOverallPower(teams.HomeTeam) + home_factor
	away_power := getOverallPower(teams.AwayTeam)

	// Generate the number of events in one match
	events := generateEvents(home_power, away_power)
	method.New(mapPowerToPercentage(home_power, away_power))

	// Pick the minutes corresponding to those events in the match (60 minutes)
	minutes := []int{}
	for i := 0; i < events; i++ {
		rand.Seed(time.Now().UnixNano() + int64(i))
		minutes = append(minutes, rand.Intn(1+92)+1)
	}
	sort.Sort(sort.IntSlice(minutes))

	home_score := 0
	away_score := 0
	for i := 1; i <= 93; i++ {
		nextTime := time.Now().Truncate(time.Minute)
		nextTime = nextTime.Add(time.Minute)
		time.Sleep(time.Until(nextTime))
		found := sort.SearchInts(minutes, i)
		if found < len(minutes) && minutes[found] == i {
			strEvent := ""
			if method.Generate() == 0 {
				// Home team event {implement player rules here}
				if rand.Float32() <= 0.5 {
					// Away team defense
					strEvent = fmt.Sprint(teams.AwayTeam.Name + " has some solid defense !")
					fmt.Println(strEvent)
				} else {
					// Home team scores
					home_score += 1
					strEvent = fmt.Sprint(teams.HomeTeam.Name + " scores!")
					fmt.Println(strEvent)
				}
			} else {
				// Away team event {implement player rules here}
				if rand.Float32() <= 0.5 {
					// Home team defense
					strEvent = fmt.Sprint(teams.HomeTeam.Name + " has some solid defense !")
					fmt.Println(strEvent)
				} else {
					// Away team scores
					away_score += 1
					strEvent = fmt.Sprint(teams.AwayTeam.Name + " scores!")
					fmt.Println(strEvent)
				}
			}
			simulation.Events = append(simulation.Events, strEvent)
			simulation.Result = fmt.Sprintf("Result %d' %s %d - %d %s\n", i, teams.HomeTeam.Name, home_score, away_score, teams.AwayTeam.Name)
			requests.UpdateSimulation(simulation)
			socketClient.Emit(simulation.Result, matchId)
			fmt.Println(simulation.Result)
		}
	}
	simulation.Result = fmt.Sprintf("%s %d - %d %s", teams.HomeTeam.Name, home_score, away_score, teams.AwayTeam.Name)
	requests.UpdateSimulation(simulation)
	fmt.Println(simulation.Result)
}

func generateEvents(home int, away int) int {
	events := int(math.Abs(float64(away - home)))
	if events > 20 {
		rand.Seed(time.Now().UnixNano())
		events = rand.Intn(7) + 8
	} else if events < 6 {
		rand.Seed(time.Now().UnixNano())
		events = rand.Intn(6) + 5
	}
	return events
}

func getOverallPower(team interfaces.Team) int {
	power := 0
	for _, player := range team.Players {
		power += player.Attack + player.Defense + player.GoalKeeping + player.Creativity
	}
	return power
}

func mapPowerToPercentage(home int, away int) []float32 {
	return []float32{float32(home) / float32(home+away), float32(away) / float32(home+away)}
}
