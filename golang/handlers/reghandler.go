package handlers

import (
	"fmt"
	"golang/db"
	"golang/globals"
	"golang/structs"
	"net/http"
)

func RegistrationHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("[REG HANDLER]")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/text")

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

		token, _ := RandomToken(16)

		// flow for adding new session
		globals.Globaldb.Exec("INSERT INTO sessions (id, admnno, expires_at) VALUES ($1, $2, NOW() + interval '1 day');", token, admnno)

		http.SetCookie(w, &http.Cookie{
			Name:     "sess_id",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			Secure:   false, // true if HTTPS
			SameSite: http.SameSiteLaxMode,
		})

		fmt.Fprintf(w, "done")
	}
}
