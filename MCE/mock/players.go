package mock

import (
	"MCE/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func InsertPlayers() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO players (name, age, goal_keeping, attack, defense, creativity)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	players := readPlayers()
	for i := 0; i < len(players); i++ {
		id := 0

		err = db.QueryRow(sqlStatement, players[i].Name, players[i].Age, players[i].GoalKeeping,
			players[i].Attack, players[i].Defense, players[i].Creativity).Scan(&id)
		if err != nil {
			panic(err)
		}

		fmt.Println("Added player with ID:", id)
	}
}

func readPlayers() []models.Player {
	jsonFile, err := os.Open("mock/testdata/players.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var players Players

	json.Unmarshal(byteValue, &players)

	return players.Players
}
