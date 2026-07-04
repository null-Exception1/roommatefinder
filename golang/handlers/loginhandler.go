package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang/db"
	"golang/globals"
	"net/http"
)

func RandomToken(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func Login(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/text")

	if req.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	q := req.URL.Query()
	// check if it exist
	str := "SELECT * FROM people WHERE admn_hash='" + q.Get("admn_hash") + "' AND name='" + q.Get("name") + "'"

	results := db.Query(str, globals.Globaldb)

	if len(results) == 1 {
		token, _ := RandomToken(16)

		// flow for adding new session
		globals.Globaldb.Exec("INSERT INTO sessions (id, admnno, expires_at) VALUES ($1, $2, NOW() + interval '1 day');", token, q.Get("admn_hash"))

		http.SetCookie(w, &http.Cookie{
			Name:     "sess_id",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,                  // must be false on plain http://localhost
			SameSite: http.SameSiteNoneMode, // <-- allow cross-port
		})

		//fmt.Println(token)

		fmt.Fprintf(w, "%s", token)

	} else {
		fmt.Fprintf(w, "not found")
	}

}
