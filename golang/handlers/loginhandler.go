package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang/db"
	"golang/globals"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
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

	logrus.WithFields(logrus.Fields{
		"package":  "handlers",
		"endpoint": "/login",
		"method":   req.Method,
		"remote":   req.RemoteAddr,
	}).Info("requested /login by ", req.RemoteAddr)

	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("SITE_URL"))
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
	admnno := q.Get("admn_hash")
	admn_hash := globals.SecureHash(admnno, os.Getenv("PEPPER"))

	str := "SELECT * FROM people WHERE admn_hash='" + admn_hash + "' AND name='" + q.Get("name") + "'"

	logrus.WithFields(logrus.Fields{
		"package":  "handlers",
		"endpoint": "/login",
		"query":    str,
		"method":   req.Method,
		"remote":   req.RemoteAddr,
	}).Debug("building query ", str)

	results := db.Query(str, globals.Globaldb)

	logrus.WithFields(logrus.Fields{
		"package":  "handlers",
		"endpoint": "/login",
		"results":  len(results),
		"method":   req.Method,
		"remote":   req.RemoteAddr,
	}).Debug("fetch query")

	if len(results) == 1 {
		token, _ := RandomToken(16)

		// flow for adding new session
		globals.Globaldb.Exec("INSERT INTO sessions (id, admnno, expires_at) VALUES ($1, $2, NOW() + interval '1 day');", token, admn_hash)

		http.SetCookie(w, &http.Cookie{
			Name:     "sess_id",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,                  // MUST be true for cross-site cookie delivery over HTTPS
			SameSite: http.SameSiteNoneMode, // MUST be None to allow the Vercel -> Render jump
		})

		logrus.WithFields(logrus.Fields{
			"package":  "handlers",
			"endpoint": "/login",
			"method":   req.Method,
			"remote":   req.RemoteAddr,
		}).Info("insert new session (login succeeded) ", token)

		fmt.Fprintf(w, "%s", token)

	} else {

		logrus.WithFields(logrus.Fields{
			"package":  "handlers",
			"endpoint": "/login",
			"method":   req.Method,
			"remote":   req.RemoteAddr,
		}).Warn("insert new session failed (login failed) ")

		logrus.WithFields(logrus.Fields{
			"endpoint": "/login",
			"status":   http.StatusOK,
			"user":     q.Get("name"),
		}).Info("Response sent")
		fmt.Fprintf(w, "not found")
	}

}
