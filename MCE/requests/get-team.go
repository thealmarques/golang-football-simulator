package requests

import (
	"MCE/interfaces"
	"MCE/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TeamResponse struct {
	Data models.Teams `json:"data"`
}

type PlayerResponse struct {
	Data models.Player `json:"data"`
}

func GetTeam(id uint) interfaces.Team {
	var team interfaces.Team
	url := fmt.Sprintf("%s%d", "http://0.0.0.0:9000/teams/", id)
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	// read the payload, in this case team information
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	var jsonResponse TeamResponse
	err = json.Unmarshal(body, &jsonResponse)

	if err != nil {
		panic(err)
	}

	team.ID = jsonResponse.Data.ID
	team.Name = jsonResponse.Data.Name

	for _, playerId := range jsonResponse.Data.Players {
		team.Players = append(team.Players, GetPlayer(playerId))
	}

	return team
}

func GetPlayer(id string) models.Player {
	url := fmt.Sprintf("%s%s", "http://0.0.0.0:9000/player/", id)
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	// read the payload, in this case team information
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	var jsonResponse PlayerResponse
	err = json.Unmarshal(body, &jsonResponse)

	if err != nil {
		panic(err)
	}

	return jsonResponse.Data
}
