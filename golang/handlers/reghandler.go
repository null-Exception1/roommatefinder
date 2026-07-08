package handlers

import (
	"fmt"
	"golang/db"
	"golang/globals"
	"golang/structs"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func RegistrationHandler(w http.ResponseWriter, req *http.Request) {

	logrus.WithFields(logrus.Fields{
		"package":  "handlers",
		"endpoint": "/registration",
		"method":   req.Method,
		"remote":   req.RemoteAddr,
	}).Info("requested /registration")

	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("SITE_URL"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/text")

	query := req.URL.Query()

	admnno := query.Get("admn_hash")
	admn_hash := globals.SecureHash(admnno, os.Getenv("PEPPER"))

	name := query.Get("name")
	social := query.Get("social")
	socialtype := query.Get("socialtype")
	blockno := query.Get("blockno")
	roomno := query.Get("roomno")
	created_at := query.Get("created_at")

	p := structs.Person{Admnno: admn_hash, Name: name, Social: social, Socialtype: socialtype, Roomno: roomno, Blockno: blockno, Created_at: created_at}

	logrus.WithFields(logrus.Fields{
		"package":  "handlers",
		"endpoint": "/registration",
		"method":   req.Method,
		"remote":   req.RemoteAddr,
	}).Debug("creating person struct")

	err := db.Insert(p, globals.Globaldb)

	if err {

		logrus.WithFields(logrus.Fields{
			"package":  "handlers",
			"endpoint": "/registration",
			"error":    err,
			"method":   req.Method,
			"remote":   req.RemoteAddr,
		}).Error("error occured in inserting a new row")

		fmt.Fprintf(w, "err")
	} else {

		token, err := RandomToken(16)

		if err != nil {
			logrus.WithError(err).Error("failed to generate session token")
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		// flow for adding new session
		_, err = globals.Globaldb.Exec("INSERT INTO sessions (id, admnno, expires_at) VALUES ($1, $2, NOW() + interval '1 day');", token, admn_hash)

		if err != nil {
			logrus.WithError(err).Error("failed to insert new session")
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

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
			"endpoint": "/registration",
			"status":   http.StatusOK,
			"token":    token,
			"method":   req.Method,
			"remote":   req.RemoteAddr,
		}).Info("assigned new session")

		fmt.Fprintf(w, "done")

	}
}
