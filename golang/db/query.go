package db

import (
	"database/sql"
	"golang/structs"
)

func Query(query string, db *sql.DB) []structs.Person {
	rows, _ := db.Query(query)
	results := make([]structs.Person, 0)

	for rows.Next() {
		var admnno string
		var name string
		var social string
		var socialtype string
		var roomno string
		var blockno string
		var created_at string

		err := rows.Scan(&admnno, &name, &social, &socialtype, &roomno, &blockno, &created_at)

		p := structs.Person{Admnno: admnno, Name: name, Social: social, Socialtype: socialtype, Roomno: roomno, Blockno: blockno, Created_at: created_at}

		results = append(results, p)

		if err != nil {
			panic(err)
		}

		//fmt.Println(admnno, name, social, socialtype, roomno, blockno, created_at)

	}
	return results
}
