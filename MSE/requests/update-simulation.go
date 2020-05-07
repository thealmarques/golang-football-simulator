package requests

import (
	"MSE/models"
	"bytes"
	"encoding/json"
	"net/http"
)

func UpdateSimulation(simulation models.Simulation) {
	values := map[string]interface{}{"id": simulation.ID, "match_id": simulation.MatchId,
		"result": simulation.Result, "events": simulation.Events}

	jsonValue, _ := json.Marshal(values)

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPut, "http://0.0.0.0:9000/simulation", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}

	// set the request header Content-Type for json
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if err != nil {
		panic(err)
	}
}
