package handlers

import (
	"fmt"
	"golang/db"
	"golang/globals"
	"golang/structs"
	"net/http"
)

func RegistrationHandler(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()

	admnno := query.Get("admnno")
	name := query.Get("name")
	social := query.Get("social")
	socialtype := query.Get("socialtype")
	blockno := query.Get("blockno")
	roomno := query.Get("roomno")
	created_at := query.Get("created_at")

	p := structs.Person{Admnno: admnno, Name: name, Social: social, Socialtype: socialtype, Roomno: roomno, Blockno: blockno, Created_at: created_at}
	err := db.Insert(p, globals.Globaldb)
	if err {
		fmt.Fprintf(w, "err")
	} else {
		fmt.Fprintf(w, "done")
	}

}
