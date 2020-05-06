package mock

import (
	"database/sql"
	"fmt"
)

func CleanDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
	DROP TABLE IF EXISTS teams, players, matches`

	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database cleaned with success.")
}
