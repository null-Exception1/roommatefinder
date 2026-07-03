package db

import (
	"database/sql"
	"fmt"
	"golang/structs"
)

func Insert(person structs.Person, db *sql.DB) bool {

	_, err := db.Exec(`INSERT INTO people VALUES ($1 , $2, $3, $4, $5, $6)`, person.Admnno, person.Name, person.Social, person.Socialtype, person.Roomno, person.Blockno)

	if err != nil {
		fmt.Println("[ERR] db.Insert failed:", err)
		return true
	} else {
		return false
	}

}
