package requests

import (
	"MSS/interfaces"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Data interfaces.Simulation `json:"data"`
}

func PostSimulation(home interfaces.Team, away interfaces.Team, matchId uint) interfaces.Simulation {
	values := map[string]interface{}{"match_id": matchId, "home_team": home, "away_team": away}

	jsonValue, _ := json.Marshal(values)

	resp, err := http.Post("http://localhost:9083/simulation", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var simulation Response
	err = json.Unmarshal(body, &simulation)

	if err != nil {
		panic(err)
	}

	return simulation.Data
}
