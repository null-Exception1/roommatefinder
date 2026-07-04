package handlers

import (
	"database/sql"
	"fmt"
	"golang/globals"
	"net/http"
	"time"
)

func Verify(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/text")

	if req.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	cookie, err := req.Cookie("sess_id")

	fmt.Println(cookie)
	if err != nil {
		fmt.Fprint(w, "invalid")
		return
	}

	var expires time.Time
	err = globals.Globaldb.QueryRow("SELECT expires_at FROM sessions WHERE id=$1", cookie.Value).
		Scan(&expires)

	if err == sql.ErrNoRows {
		fmt.Fprint(w, "invalid")
		return
	}
	if time.Now().After(expires) {
		fmt.Fprint(w, "expired")
	} else {
		fmt.Fprint(w, "valid")
	}

}
