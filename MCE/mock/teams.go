package mock

import (
	"MSS/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func InsertTeams() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
		INSERT INTO teams (name, players)
		VALUES ($1, $2)
		RETURNING id`

	teams := readTeams()
	for i := 0; i < len(teams); i++ {
		id := 0

		err = db.QueryRow(sqlStatement, teams[i].Name, teams[i].Players).Scan(&id)
		if err != nil {
			panic(err)
		}

		fmt.Println("Added team with ID:", id)
	}
}

func readTeams() []models.Teams {
	jsonFile, err := os.Open("mock/testdata/teams.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var teams Teams

	json.Unmarshal(byteValue, &teams)

	return teams.Teams
}
