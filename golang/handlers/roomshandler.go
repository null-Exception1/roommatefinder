package handlers

import (
	"encoding/json"
	"fmt"
	"golang/db"
	"golang/globals"
	"golang/structs"
	"net/http"
)

func Rooms(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()

	blockno := q.Get("block")

	sqlquery := fmt.Sprintf("SELECT * FROM people WHERE blockno='%s';", blockno)

	fmt.Println("hey", string(sqlquery))
	rows := db.Query(string(sqlquery), globals.Globaldb)
	fmt.Println("rows", rows)
	rooms := make(map[string]*structs.Room, 0)
	for _, row := range rows {
		if _, ok := rooms[row.Roomno]; ok {
			rooms[row.Roomno].People = append(rooms[row.Roomno].People, &structs.Person{Admnno: row.Admnno, Name: row.Name, Social: row.Social, Socialtype: row.Socialtype, Roomno: row.Roomno, Blockno: row.Blockno, Created_at: row.Created_at})
		} else {
			rooms[row.Roomno] = &structs.Room{People: make([]*structs.Person, 0)}
			rooms[row.Roomno].People = append(rooms[row.Roomno].People, &structs.Person{Admnno: row.Admnno, Name: row.Name, Social: row.Social, Socialtype: row.Socialtype, Roomno: row.Roomno, Blockno: row.Blockno, Created_at: row.Created_at})

		}
	}

	str, _ := json.Marshal(rooms)

	fmt.Fprintf(w, "%s", string(str))

}
