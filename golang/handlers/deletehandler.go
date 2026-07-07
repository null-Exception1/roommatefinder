package handlers

import (
	"fmt"
	"golang/globals"
	"net/http"
)

func DeleteHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/text")

	cookie, err := req.Cookie("sess_id")
	if err != nil {
		http.Error(w, "no session", http.StatusUnauthorized)
		return
	}

	// find session
	var admn_hash string
	err = globals.Globaldb.QueryRow("SELECT admnno FROM sessions WHERE id=$1", cookie.Value).Scan(&admn_hash)
	if err != nil {

		http.Error(w, "invalid session", http.StatusUnauthorized)
		return
	}

	// delete user + session
	_, _ = globals.Globaldb.Exec("DELETE FROM people WHERE admn_hash=$1", admn_hash)
	_, _ = globals.Globaldb.Exec("DELETE FROM sessions WHERE id=$1", cookie.Value)

	fmt.Fprint(w, "deleted")

}
